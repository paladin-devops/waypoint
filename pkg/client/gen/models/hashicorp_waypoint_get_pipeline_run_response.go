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

// HashicorpWaypointGetPipelineRunResponse hashicorp waypoint get pipeline run response
//
// swagger:model hashicorp.waypoint.GetPipelineRunResponse
type HashicorpWaypointGetPipelineRunResponse struct {

	// A single pipeline run
	PipelineRun *HashicorpWaypointPipelineRun `json:"pipeline_run,omitempty"`
}

// Validate validates this hashicorp waypoint get pipeline run response
func (m *HashicorpWaypointGetPipelineRunResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePipelineRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetPipelineRunResponse) validatePipelineRun(formats strfmt.Registry) error {
	if swag.IsZero(m.PipelineRun) { // not required
		return nil
	}

	if m.PipelineRun != nil {
		if err := m.PipelineRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipeline_run")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint get pipeline run response based on the context it is used
func (m *HashicorpWaypointGetPipelineRunResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePipelineRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetPipelineRunResponse) contextValidatePipelineRun(ctx context.Context, formats strfmt.Registry) error {

	if m.PipelineRun != nil {
		if err := m.PipelineRun.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipeline_run")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetPipelineRunResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetPipelineRunResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetPipelineRunResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
