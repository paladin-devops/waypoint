// Package ssm contains components for syncing configuration with AWS SSM.
package ssm

// Options are the SDK options to use for instantiation for this plugin.
var Options = []sdk.Option{
	sdk.WithComponents(&ConfigSourcer{}),
}
