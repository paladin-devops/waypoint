// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues hashicorp waypoint get job stream response terminal event named values
//
// swagger:model hashicorp.waypoint.GetJobStreamResponse.Terminal.Event.NamedValues
type HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues struct {

	// values
	Values []*HashicorpWaypointGetJobStreamResponseTerminalEventNamedValue `json:"values"`
}

// Validate validates this hashicorp waypoint get job stream response terminal event named values
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp waypoint get job stream response terminal event named values based on the context it is used
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Values); i++ {

		if m.Values[i] != nil {
			if err := m.Values[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetJobStreamResponseTerminalEventNamedValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
