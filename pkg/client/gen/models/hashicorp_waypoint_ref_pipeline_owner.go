// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointRefPipelineOwner hashicorp waypoint ref pipeline owner
//
// swagger:model hashicorp.waypoint.Ref.PipelineOwner
type HashicorpWaypointRefPipelineOwner struct {

	// the name of the defined pipeline config
	PipelineName string `json:"pipeline_name,omitempty"`

	// the project this pipeline is associated with
	Project *HashicorpWaypointRefProject `json:"project,omitempty"`
}

// Validate validates this hashicorp waypoint ref pipeline owner
func (m *HashicorpWaypointRefPipelineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointRefPipelineOwner) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint ref pipeline owner based on the context it is used
func (m *HashicorpWaypointRefPipelineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointRefPipelineOwner) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {
		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRefPipelineOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRefPipelineOwner) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRefPipelineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
