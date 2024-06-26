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

// HashicorpWaypointServerConfig ServerConfig is the configuration for the server that can be read and
// set online. This differs from the configuration used to start the server
// since some settings can only be set via the file vs. the API.
//
// swagger:model hashicorp.waypoint.ServerConfig
type HashicorpWaypointServerConfig struct {

	// The addresses that are advertised for entrypoints. These define how
	// applications reach back to the server. Currently you may only set
	// EXACTLY ONE address. In the future, we'll support multiple advertise
	// addrs and more controls over which are advertised when.
	AdvertiseAddrs []*HashicorpWaypointServerConfigAdvertiseAddr `json:"advertise_addrs"`

	// Cookie is a unique cookie for this server. This can be sent in metadata
	// as a semi-secret mechanism to just ensure you're talking to the correct
	// cluster. The cookie does not enable access to data directly. Some API
	// endpoints (such as RunnerToken) require it. This is auto-generated on
	// startup and cannot currently be changed. Any attempts to change this
	// value are ignored.
	Cookie string `json:"cookie,omitempty"`

	// The platform that the server is currently installed to. This is set
	// through the CLI flag '-platform' on installation.
	Platform string `json:"platform,omitempty"`
}

// Validate validates this hashicorp waypoint server config
func (m *HashicorpWaypointServerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvertiseAddrs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointServerConfig) validateAdvertiseAddrs(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvertiseAddrs) { // not required
		return nil
	}

	for i := 0; i < len(m.AdvertiseAddrs); i++ {
		if swag.IsZero(m.AdvertiseAddrs[i]) { // not required
			continue
		}

		if m.AdvertiseAddrs[i] != nil {
			if err := m.AdvertiseAddrs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advertise_addrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advertise_addrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp waypoint server config based on the context it is used
func (m *HashicorpWaypointServerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvertiseAddrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointServerConfig) contextValidateAdvertiseAddrs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdvertiseAddrs); i++ {

		if m.AdvertiseAddrs[i] != nil {
			if err := m.AdvertiseAddrs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advertise_addrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advertise_addrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointServerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointServerConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointServerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
