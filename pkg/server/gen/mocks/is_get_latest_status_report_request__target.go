// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isGetLatestStatusReportRequest_Target is an autogenerated mock type for the isGetLatestStatusReportRequest_Target type
type isGetLatestStatusReportRequest_Target struct {
	mock.Mock
}

// isGetLatestStatusReportRequest_Target provides a mock function with given fields:
func (_m *isGetLatestStatusReportRequest_Target) isGetLatestStatusReportRequest_Target() {
	_m.Called()
}

type mockConstructorTestingTnewIsGetLatestStatusReportRequest_Target interface {
	mock.TestingT
	Cleanup(func())
}

// newIsGetLatestStatusReportRequest_Target creates a new instance of isGetLatestStatusReportRequest_Target. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsGetLatestStatusReportRequest_Target(t mockConstructorTestingTnewIsGetLatestStatusReportRequest_Target) *isGetLatestStatusReportRequest_Target {
	mock := &isGetLatestStatusReportRequest_Target{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
