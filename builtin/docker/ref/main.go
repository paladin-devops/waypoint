package dockerref

// Options are the SDK options to use for instantiation.
var Options = []sdk.Option{
	sdk.WithComponents(&Builder{}),
}
