// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ShortUrlUsecase is an autogenerated mock type for the ShortUrlUsecase type
type ShortUrlUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: url
func (_m *ShortUrlUsecase) Create(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: code
func (_m *ShortUrlUsecase) Delete(code string) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Redirect provides a mock function with given fields: url
func (_m *ShortUrlUsecase) Redirect(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewShortUrlUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewShortUrlUsecase creates a new instance of ShortUrlUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewShortUrlUsecase(t mockConstructorTestingTNewShortUrlUsecase) *ShortUrlUsecase {
	mock := &ShortUrlUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
