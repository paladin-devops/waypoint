// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isTrigger_Operation is an autogenerated mock type for the isTrigger_Operation type
type isTrigger_Operation struct {
	mock.Mock
}

// isTrigger_Operation provides a mock function with given fields:
func (_m *isTrigger_Operation) isTrigger_Operation() {
	_m.Called()
}

type mockConstructorTestingTnewIsTrigger_Operation interface {
	mock.TestingT
	Cleanup(func())
}

// newIsTrigger_Operation creates a new instance of isTrigger_Operation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsTrigger_Operation(t mockConstructorTestingTnewIsTrigger_Operation) *isTrigger_Operation {
	mock := &isTrigger_Operation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
