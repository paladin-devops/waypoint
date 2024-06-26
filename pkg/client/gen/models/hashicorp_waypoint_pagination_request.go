// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointPaginationRequest https://github.com/hashicorp/cloud-api/blob/master/hashicorp/cloud/common/pagination.proto
// PaginationRequest are the parameters for a paginated list request.
//
// swagger:model hashicorp.waypoint.PaginationRequest
type HashicorpWaypointPaginationRequest struct {

	// Specifies a page token to use to retrieve the next page. Set this to the
	// `next_page_token` returned by previous list requests to get the next page of
	// results. If set, `previous_page_token` must not be set.
	NextPageToken string `json:"next_page_token,omitempty"`

	// The max number of results per page that should be returned. If the number
	// of available results is larger than `page_size`, a `next_page_token` is
	// returned which can be used to get the next page of results in subsequent
	// requests. A value of zero will cause `page_size` to be defaulted.
	PageSize int64 `json:"page_size,omitempty"`

	// Specifies a page token to use to retrieve the previous page. Set this to
	// the `previous_page_token` returned by previous list requests to get the
	// previous page of results. If set, `next_page_token` must not be set.
	PreviousPageToken string `json:"previous_page_token,omitempty"`
}

// Validate validates this hashicorp waypoint pagination request
func (m *HashicorpWaypointPaginationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint pagination request based on context it is used
func (m *HashicorpWaypointPaginationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointPaginationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointPaginationRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointPaginationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
