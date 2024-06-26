// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointJobGitSSH SSH private key auth
//
// swagger:model hashicorp.waypoint.Job.Git.SSH
type HashicorpWaypointJobGitSSH struct {

	// password is an optional password for decoding the private key.
	Password string `json:"password,omitempty"`

	// private_key_pem is a PEM-encoded private key.
	// Format: byte
	PrivateKeyPem strfmt.Base64 `json:"private_key_pem,omitempty"`

	// user is the SSH user to use when cloning. This will default to
	// "git" if not specified.
	User string `json:"user,omitempty"`
}

// Validate validates this hashicorp waypoint job git SSH
func (m *HashicorpWaypointJobGitSSH) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint job git SSH based on context it is used
func (m *HashicorpWaypointJobGitSSH) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobGitSSH) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobGitSSH) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobGitSSH
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
