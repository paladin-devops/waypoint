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

// HashicorpWaypointListOIDCAuthMethodsResponse hashicorp waypoint list o ID c auth methods response
//
// swagger:model hashicorp.waypoint.ListOIDCAuthMethodsResponse
type HashicorpWaypointListOIDCAuthMethodsResponse struct {

	// auth methods
	AuthMethods []*HashicorpWaypointOIDCAuthMethod `json:"auth_methods"`
}

// Validate validates this hashicorp waypoint list o ID c auth methods response
func (m *HashicorpWaypointListOIDCAuthMethodsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthMethods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListOIDCAuthMethodsResponse) validateAuthMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthMethods); i++ {
		if swag.IsZero(m.AuthMethods[i]) { // not required
			continue
		}

		if m.AuthMethods[i] != nil {
			if err := m.AuthMethods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auth_methods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auth_methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp waypoint list o ID c auth methods response based on the context it is used
func (m *HashicorpWaypointListOIDCAuthMethodsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthMethods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListOIDCAuthMethodsResponse) contextValidateAuthMethods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuthMethods); i++ {

		if m.AuthMethods[i] != nil {
			if err := m.AuthMethods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auth_methods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auth_methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointListOIDCAuthMethodsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointListOIDCAuthMethodsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointListOIDCAuthMethodsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
