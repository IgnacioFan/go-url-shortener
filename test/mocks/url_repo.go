// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/url/impl.go

// Package mock_url is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUrlRepository is a mock of UrlRepository interface.
type MockUrlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUrlRepositoryMockRecorder
}

// MockUrlRepositoryMockRecorder is the mock recorder for MockUrlRepository.
type MockUrlRepositoryMockRecorder struct {
	mock *MockUrlRepository
}

// NewMockUrlRepository creates a new mock instance.
func NewMockUrlRepository(ctrl *gomock.Controller) *MockUrlRepository {
	mock := &MockUrlRepository{ctrl: ctrl}
	mock.recorder = &MockUrlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlRepository) EXPECT() *MockUrlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUrlRepository) Create(longURL, shortUrl string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", longURL, shortUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUrlRepositoryMockRecorder) Create(longURL, shortUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUrlRepository)(nil).Create), longURL, shortUrl)
}

// FindBy mocks base method.
func (m *MockUrlRepository) FindBy(shortUrl string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBy", shortUrl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBy indicates an expected call of FindBy.
func (mr *MockUrlRepositoryMockRecorder) FindBy(shortUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBy", reflect.TypeOf((*MockUrlRepository)(nil).FindBy), shortUrl)
}
