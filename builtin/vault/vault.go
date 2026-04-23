// Package vault contains components for syncing secrets with Vault.
package vault

import sdk "github.com/paladin-devops/waypoint-plugin-sdk"

// Options are the SDK options to use for instantiation for this plugin.
var Options = []sdk.Option{
	sdk.WithComponents(&ConfigSourcer{}),
}
