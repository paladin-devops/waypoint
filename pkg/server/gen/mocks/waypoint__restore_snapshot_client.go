// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gen "github.com/hashicorp/waypoint/pkg/server/gen"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// Waypoint_RestoreSnapshotClient is an autogenerated mock type for the Waypoint_RestoreSnapshotClient type
type Waypoint_RestoreSnapshotClient struct {
	mock.Mock
}

// CloseAndRecv provides a mock function with given fields:
func (_m *Waypoint_RestoreSnapshotClient) CloseAndRecv() (*emptypb.Empty, error) {
	ret := _m.Called()

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func() *emptypb.Empty); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseSend provides a mock function with given fields:
func (_m *Waypoint_RestoreSnapshotClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *Waypoint_RestoreSnapshotClient) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Header provides a mock function with given fields:
func (_m *Waypoint_RestoreSnapshotClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecvMsg provides a mock function with given fields: m
func (_m *Waypoint_RestoreSnapshotClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *Waypoint_RestoreSnapshotClient) Send(_a0 *gen.RestoreSnapshotRequest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gen.RestoreSnapshotRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *Waypoint_RestoreSnapshotClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trailer provides a mock function with given fields:
func (_m *Waypoint_RestoreSnapshotClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

type mockConstructorTestingTNewWaypoint_RestoreSnapshotClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewWaypoint_RestoreSnapshotClient creates a new instance of Waypoint_RestoreSnapshotClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWaypoint_RestoreSnapshotClient(t mockConstructorTestingTNewWaypoint_RestoreSnapshotClient) *Waypoint_RestoreSnapshotClient {
	mock := &Waypoint_RestoreSnapshotClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
