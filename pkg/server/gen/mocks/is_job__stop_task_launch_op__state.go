// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isJob_StopTaskLaunchOp_State is an autogenerated mock type for the isJob_StopTaskLaunchOp_State type
type isJob_StopTaskLaunchOp_State struct {
	mock.Mock
}

// isJob_StopTaskLaunchOp_State provides a mock function with given fields:
func (_m *isJob_StopTaskLaunchOp_State) isJob_StopTaskLaunchOp_State() {
	_m.Called()
}

type mockConstructorTestingTnewIsJob_StopTaskLaunchOp_State interface {
	mock.TestingT
	Cleanup(func())
}

// newIsJob_StopTaskLaunchOp_State creates a new instance of isJob_StopTaskLaunchOp_State. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsJob_StopTaskLaunchOp_State(t mockConstructorTestingTnewIsJob_StopTaskLaunchOp_State) *isJob_StopTaskLaunchOp_State {
	mock := &isJob_StopTaskLaunchOp_State{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
