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

// HashicorpWaypointExecStreamRequest hashicorp waypoint exec stream request
//
// swagger:model hashicorp.waypoint.ExecStreamRequest
type HashicorpWaypointExecStreamRequest struct {

	// input
	Input *HashicorpWaypointExecStreamRequestInput `json:"input,omitempty"`

	// input_eof should be sent as an event when the input stream is
	// closed. After this, no more Input messages can be sent. Any Input
	// messages sent will be ignored. This will send an EOF on the remote
	// end as well to close stdin for the exec process.
	InputEOF interface{} `json:"input_eof,omitempty"`

	// start
	Start *HashicorpWaypointExecStreamRequestStart `json:"start,omitempty"`

	// winch
	Winch *HashicorpWaypointExecStreamRequestWindowSize `json:"winch,omitempty"`
}

// Validate validates this hashicorp waypoint exec stream request
func (m *HashicorpWaypointExecStreamRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWinch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointExecStreamRequest) validateInput(formats strfmt.Registry) error {
	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointExecStreamRequest) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if m.Start != nil {
		if err := m.Start.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("start")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("start")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointExecStreamRequest) validateWinch(formats strfmt.Registry) error {
	if swag.IsZero(m.Winch) { // not required
		return nil
	}

	if m.Winch != nil {
		if err := m.Winch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("winch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("winch")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint exec stream request based on the context it is used
func (m *HashicorpWaypointExecStreamRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWinch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointExecStreamRequest) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	if m.Input != nil {
		if err := m.Input.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointExecStreamRequest) contextValidateStart(ctx context.Context, formats strfmt.Registry) error {

	if m.Start != nil {
		if err := m.Start.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("start")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("start")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointExecStreamRequest) contextValidateWinch(ctx context.Context, formats strfmt.Registry) error {

	if m.Winch != nil {
		if err := m.Winch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("winch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("winch")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointExecStreamRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointExecStreamRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointExecStreamRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
