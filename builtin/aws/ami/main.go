package ami

//go:generate protoc -I ../../../.. --go_out=../../../.. --go-grpc_out=../../../.. waypoint/builtin/aws/ami/plugin.proto

// Options are the SDK options to use for instantiation.
var Options = []sdk.Option{
	sdk.WithComponents(&Builder{}),
}
