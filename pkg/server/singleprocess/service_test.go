package singleprocess

import (
	"context"

	"github.com/paladin-devops/waypoint/internal/server/boltdbstate"
	pb "github.com/paladin-devops/waypoint/pkg/server/gen"
)

func testServiceImpl(impl pb.WaypointServer) *Service {
	return impl.(*Service)
}

func testStateInmem(impl pb.WaypointServer) *boltdbstate.State {
	return testServiceImpl(impl).state(context.Background()).(*boltdbstate.State)
}
