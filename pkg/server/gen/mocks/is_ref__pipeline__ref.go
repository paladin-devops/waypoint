// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isRef_Pipeline_Ref is an autogenerated mock type for the isRef_Pipeline_Ref type
type isRef_Pipeline_Ref struct {
	mock.Mock
}

// isRef_Pipeline_Ref provides a mock function with given fields:
func (_m *isRef_Pipeline_Ref) isRef_Pipeline_Ref() {
	_m.Called()
}

type mockConstructorTestingTnewIsRef_Pipeline_Ref interface {
	mock.TestingT
	Cleanup(func())
}

// newIsRef_Pipeline_Ref creates a new instance of isRef_Pipeline_Ref. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsRef_Pipeline_Ref(t mockConstructorTestingTnewIsRef_Pipeline_Ref) *isRef_Pipeline_Ref {
	mock := &isRef_Pipeline_Ref{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
