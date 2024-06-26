// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointRefProject Project references a project.
//
// swagger:model hashicorp.waypoint.Ref.Project
type HashicorpWaypointRefProject struct {

	// project
	Project string `json:"project,omitempty"`
}

// Validate validates this hashicorp waypoint ref project
func (m *HashicorpWaypointRefProject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint ref project based on context it is used
func (m *HashicorpWaypointRefProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRefProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRefProject) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRefProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
