// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isConfigVar_UnusedScope is an autogenerated mock type for the isConfigVar_UnusedScope type
type isConfigVar_UnusedScope struct {
	mock.Mock
}

// isConfigVar_UnusedScope provides a mock function with given fields:
func (_m *isConfigVar_UnusedScope) isConfigVar_UnusedScope() {
	_m.Called()
}

type mockConstructorTestingTnewIsConfigVar_UnusedScope interface {
	mock.TestingT
	Cleanup(func())
}

// newIsConfigVar_UnusedScope creates a new instance of isConfigVar_UnusedScope. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsConfigVar_UnusedScope(t mockConstructorTestingTnewIsConfigVar_UnusedScope) *isConfigVar_UnusedScope {
	mock := &isConfigVar_UnusedScope{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
