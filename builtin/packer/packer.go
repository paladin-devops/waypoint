package packer

import (
	sdk "github.com/paladin-devops/waypoint-plugin-sdk"
)

var Options = []sdk.Option{
	sdk.WithComponents(&ConfigSourcer{}),
}
