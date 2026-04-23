// Package consul contains components for syncing app configuration with Consul.
package consul

// Options are the SDK options to use for instantiation for this plugin.
var Options = []sdk.Option{
	sdk.WithComponents(&ConfigSourcer{}),
}
