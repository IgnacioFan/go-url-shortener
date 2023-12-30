// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/zookeeper/impl.go

// Package mock_zookeeper is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockZookeeper is a mock of Zookeeper interface.
type MockZookeeper struct {
	ctrl     *gomock.Controller
	recorder *MockZookeeperMockRecorder
}

// MockZookeeperMockRecorder is the mock recorder for MockZookeeper.
type MockZookeeperMockRecorder struct {
	mock *MockZookeeper
}

// NewMockZookeeper creates a new mock instance.
func NewMockZookeeper(ctrl *gomock.Controller) *MockZookeeper {
	mock := &MockZookeeper{ctrl: ctrl}
	mock.recorder = &MockZookeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockZookeeper) EXPECT() *MockZookeeperMockRecorder {
	return m.recorder
}

// SetNewRange mocks base method.
func (m *MockZookeeper) SetNewRange() (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNewRange")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetNewRange indicates an expected call of SetNewRange.
func (mr *MockZookeeperMockRecorder) SetNewRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNewRange", reflect.TypeOf((*MockZookeeper)(nil).SetNewRange))
}
