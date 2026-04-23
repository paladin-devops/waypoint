package ptypes

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/paladin-devops/waypoint/internal/pkg/validationext"
	pb "github.com/paladin-devops/waypoint/pkg/server/gen"
)

// ValidateCreateHostnameRequest
func ValidateCreateHostnameRequest(v *pb.CreateHostnameRequest) error {
	return validationext.Error(validation.ValidateStruct(v,
		validation.Field(&v.Target, validation.Required),
	))
}
