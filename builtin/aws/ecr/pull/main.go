package ecrpull

import (
	sdk "github.com/paladin-devops/waypoint-plugin-sdk"
)

// Options are the SDK options to use for instantiation.
var Options = []sdk.Option{
	sdk.WithComponents(&Builder{}),
}
