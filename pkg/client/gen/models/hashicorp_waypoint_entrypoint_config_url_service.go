// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointEntrypointConfigURLService hashicorp waypoint entrypoint config URL service
//
// swagger:model hashicorp.waypoint.EntrypointConfig.URLService
type HashicorpWaypointEntrypointConfigURLService struct {

	// address to the control server and the token for auth
	ControlAddr string `json:"control_addr,omitempty"`

	// labels to register this instance under
	Labels string `json:"labels,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this hashicorp waypoint entrypoint config URL service
func (m *HashicorpWaypointEntrypointConfigURLService) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint entrypoint config URL service based on context it is used
func (m *HashicorpWaypointEntrypointConfigURLService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointEntrypointConfigURLService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointEntrypointConfigURLService) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointEntrypointConfigURLService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
