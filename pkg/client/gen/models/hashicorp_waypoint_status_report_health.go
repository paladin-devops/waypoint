// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointStatusReportHealth hashicorp waypoint status report health
//
// swagger:model hashicorp.waypoint.StatusReport.Health
type HashicorpWaypointStatusReportHealth struct {

	// deprecated id
	DeprecatedID string `json:"deprecated_id,omitempty"`

	// deprecated name
	DeprecatedName string `json:"deprecated_name,omitempty"`

	// the overall health message of the report for a resource
	HealthMessage string `json:"health_message,omitempty"`

	// the overall health of the report for a resource
	HealthStatus string `json:"health_status,omitempty"`
}

// Validate validates this hashicorp waypoint status report health
func (m *HashicorpWaypointStatusReportHealth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint status report health based on context it is used
func (m *HashicorpWaypointStatusReportHealth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointStatusReportHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointStatusReportHealth) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointStatusReportHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
