package ecr

import sdk "github.com/paladin-devops/waypoint-plugin-sdk"

//go:generate protoc -I ../../../.. --go_out=../../../.. --go-grpc_out=../../../.. waypoint/builtin/aws/ecr/plugin.proto

// Options are the SDK options to use for instantiation.
var Options = []sdk.Option{
	sdk.WithComponents(&Registry{}),
	sdk.WithMappers(ECRImageMapper),
}
