package swarm

import (
	"context"
	"errors"
	"sync"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

// TODO: change this text when we fix the bug
var errDialCanceled = errors.New("dial was aborted internally, likely due to https://git.io/Je2wW")

// DialFunc is the type of function expected by DialSync.
type DialWorkerFunc func(context.Context, peer.ID, <-chan DialRequest)

// NewDialSync constructs a new DialSync
func NewDialSync(worker DialWorkerFunc) *DialSync {
	return &DialSync{
		dials:      make(map[peer.ID]*activeDial),
		dialWorker: worker,
	}
}

// DialSync is a dial synchronization helper that ensures that at most one dial
// to any given peer is active at any given time.
type DialSync struct {
	dials      map[peer.ID]*activeDial
	dialsLk    sync.Mutex
	dialWorker DialWorkerFunc
}

type activeDial struct {
	id     peer.ID
	refCnt int

	ctx    context.Context
	cancel func()

	reqch chan DialRequest

	ds *DialSync
}

func (ad *activeDial) decref() {
	ad.ds.dialsLk.Lock()
	ad.refCnt--
	if ad.refCnt == 0 {
		ad.cancel()
		close(ad.reqch)
		delete(ad.ds.dials, ad.id)
	}
	ad.ds.dialsLk.Unlock()
}

func (ad *activeDial) dial(ctx context.Context, p peer.ID) (*Conn, error) {
	dialCtx := ad.ctx

	if forceDirect, reason := network.GetForceDirectDial(ctx); forceDirect {
		dialCtx = network.WithForceDirectDial(dialCtx, reason)
	}
	if simConnect, reason := network.GetSimultaneousConnect(ctx); simConnect {
		dialCtx = network.WithSimultaneousConnect(dialCtx, reason)
	}

	resch := make(chan DialResponse, 1)
	select {
	case ad.reqch <- DialRequest{Ctx: dialCtx, Resch: resch}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	select {
	case res := <-resch:
		return res.Conn, res.Err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (ds *DialSync) getActiveDial(p peer.ID) *activeDial {
	ds.dialsLk.Lock()
	defer ds.dialsLk.Unlock()

	actd, ok := ds.dials[p]
	if !ok {
		// This code intentionally uses the background context. Otherwise, if the first call
		// to Dial is canceled, subsequent dial calls will also be canceled.
		// XXX: this also breaks direct connection logic. We will need to pipe the
		// information through some other way.
		adctx, cancel := context.WithCancel(context.Background())
		actd = &activeDial{
			id:     p,
			ctx:    adctx,
			cancel: cancel,
			reqch:  make(chan DialRequest),
			ds:     ds,
		}
		ds.dials[p] = actd

		go ds.dialWorker(adctx, p, actd.reqch)
	}

	// increase ref count before dropping dialsLk
	actd.refCnt++

	return actd
}

// DialLock initiates a dial to the given peer if there are none in progress
// then waits for the dial to that peer to complete.
func (ds *DialSync) DialLock(ctx context.Context, p peer.ID) (*Conn, error) {
	ad := ds.getActiveDial(p)
	defer ad.decref()

	return ad.dial(ctx, p)
}
