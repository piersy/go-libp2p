// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libp2p/go-libp2p/core/network (interfaces: ResourceScopeSpan)

// Package mocknetwork is a generated GoMock package.
package mocknetwork

import (
	reflect "reflect"

	network "github.com/libp2p/go-libp2p/core/network"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceScopeSpan is a mock of ResourceScopeSpan interface.
type MockResourceScopeSpan struct {
	ctrl     *gomock.Controller
	recorder *MockResourceScopeSpanMockRecorder
}

// MockResourceScopeSpanMockRecorder is the mock recorder for MockResourceScopeSpan.
type MockResourceScopeSpanMockRecorder struct {
	mock *MockResourceScopeSpan
}

// NewMockResourceScopeSpan creates a new mock instance.
func NewMockResourceScopeSpan(ctrl *gomock.Controller) *MockResourceScopeSpan {
	mock := &MockResourceScopeSpan{ctrl: ctrl}
	mock.recorder = &MockResourceScopeSpanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceScopeSpan) EXPECT() *MockResourceScopeSpanMockRecorder {
	return m.recorder
}

// BeginSpan mocks base method.
func (m *MockResourceScopeSpan) BeginSpan() (network.ResourceScopeSpan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginSpan")
	ret0, _ := ret[0].(network.ResourceScopeSpan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginSpan indicates an expected call of BeginSpan.
func (mr *MockResourceScopeSpanMockRecorder) BeginSpan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginSpan", reflect.TypeOf((*MockResourceScopeSpan)(nil).BeginSpan))
}

// Done mocks base method.
func (m *MockResourceScopeSpan) Done() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Done")
}

// Done indicates an expected call of Done.
func (mr *MockResourceScopeSpanMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockResourceScopeSpan)(nil).Done))
}

// ReleaseMemory mocks base method.
func (m *MockResourceScopeSpan) ReleaseMemory(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseMemory", arg0)
}

// ReleaseMemory indicates an expected call of ReleaseMemory.
func (mr *MockResourceScopeSpanMockRecorder) ReleaseMemory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseMemory", reflect.TypeOf((*MockResourceScopeSpan)(nil).ReleaseMemory), arg0)
}

// ReserveMemory mocks base method.
func (m *MockResourceScopeSpan) ReserveMemory(arg0 int, arg1 byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveMemory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveMemory indicates an expected call of ReserveMemory.
func (mr *MockResourceScopeSpanMockRecorder) ReserveMemory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveMemory", reflect.TypeOf((*MockResourceScopeSpan)(nil).ReserveMemory), arg0, arg1)
}

// Stat mocks base method.
func (m *MockResourceScopeSpan) Stat() network.ScopeStat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat")
	ret0, _ := ret[0].(network.ScopeStat)
	return ret0
}

// Stat indicates an expected call of Stat.
func (mr *MockResourceScopeSpanMockRecorder) Stat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockResourceScopeSpan)(nil).Stat))
}
