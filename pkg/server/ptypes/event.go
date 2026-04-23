package ptypes

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/paladin-devops/waypoint/internal/pkg/validationext"
	pb "github.com/paladin-devops/waypoint/pkg/server/gen"
)

// ValidateUIListEventsRequest
func ValidateUIListEventsRequest(v *pb.UI_ListEventsRequest) error {
	return validationext.Error(validation.ValidateStruct(v,
		validation.Field(&v.Application, validation.Required),
		validationext.StructField(&v.Pagination, func() []*validation.FieldRules {
			return ValidatePaginationRequestRules(v.Pagination)
		}),
	))
}
