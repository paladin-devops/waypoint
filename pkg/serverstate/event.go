package serverstate

import (
	"time"

	"github.com/paladin-devops/waypoint/pkg/server/gen"
)

type Event struct {
	Application    *gen.Ref_Application
	Project        *gen.Ref_Project
	EventTimestamp time.Time
	EventType      string
	EventData      []byte
}
