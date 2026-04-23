package singleprocess

import (
	"context"

	empty "google.golang.org/protobuf/types/known/emptypb"

	"github.com/paladin-devops/waypoint/pkg/protocolversion"
	pb "github.com/paladin-devops/waypoint/pkg/server/gen"
)

func (s *Service) GetVersionInfo(
	ctx context.Context,
	req *empty.Empty,
) (*pb.GetVersionInfoResponse, error) {
	return &pb.GetVersionInfoResponse{
		Info: protocolversion.Current(),
		ServerFeatures: &pb.ServerFeatures{
			Features: s.features,
		},
	}, nil
}
