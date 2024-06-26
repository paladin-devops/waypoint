// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointConfigVarDynamicVal DynamicVal is the configuration for dynamic values for configuration.
//
// swagger:model hashicorp.waypoint.ConfigVar.DynamicVal
type HashicorpWaypointConfigVarDynamicVal struct {

	// config is the configuration for the config source plugin that
	// defines how the value is read. For example, for a "Vault" config
	// source this may contain the path in the KV store to read the value.
	Config map[string]string `json:"config,omitempty"`

	// from is the config source plugin to use
	From string `json:"from,omitempty"`
}

// Validate validates this hashicorp waypoint config var dynamic val
func (m *HashicorpWaypointConfigVarDynamicVal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint config var dynamic val based on context it is used
func (m *HashicorpWaypointConfigVarDynamicVal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointConfigVarDynamicVal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointConfigVarDynamicVal) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointConfigVarDynamicVal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
