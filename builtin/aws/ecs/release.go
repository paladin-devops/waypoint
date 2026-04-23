package ecs

import "github.com/paladin-devops/waypoint-plugin-sdk/component"

func (r *Release) URL() string { return r.Url }

var _ component.Release = (*Release)(nil)
