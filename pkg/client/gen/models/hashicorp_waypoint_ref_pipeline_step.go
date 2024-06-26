// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointRefPipelineStep hashicorp waypoint ref pipeline step
//
// swagger:model hashicorp.waypoint.Ref.PipelineStep
type HashicorpWaypointRefPipelineStep struct {

	// ID of the pipeline
	PipelineID string `json:"pipeline_id,omitempty"`

	// Name of the pipeline
	PipelineName string `json:"pipeline_name,omitempty"`

	// Pipeline run sequence
	RunSequence string `json:"run_sequence,omitempty"`

	// Step name within the pipeline.
	Step string `json:"step,omitempty"`
}

// Validate validates this hashicorp waypoint ref pipeline step
func (m *HashicorpWaypointRefPipelineStep) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint ref pipeline step based on context it is used
func (m *HashicorpWaypointRefPipelineStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRefPipelineStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRefPipelineStep) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRefPipelineStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
