// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpWaypointExecStreamResponseOutputChannel hashicorp waypoint exec stream response output channel
//
// swagger:model hashicorp.waypoint.ExecStreamResponse.Output.Channel
type HashicorpWaypointExecStreamResponseOutputChannel string

func NewHashicorpWaypointExecStreamResponseOutputChannel(value HashicorpWaypointExecStreamResponseOutputChannel) *HashicorpWaypointExecStreamResponseOutputChannel {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpWaypointExecStreamResponseOutputChannel.
func (m HashicorpWaypointExecStreamResponseOutputChannel) Pointer() *HashicorpWaypointExecStreamResponseOutputChannel {
	return &m
}

const (

	// HashicorpWaypointExecStreamResponseOutputChannelUNKNOWN captures enum value "UNKNOWN"
	HashicorpWaypointExecStreamResponseOutputChannelUNKNOWN HashicorpWaypointExecStreamResponseOutputChannel = "UNKNOWN"

	// HashicorpWaypointExecStreamResponseOutputChannelSTDOUT captures enum value "STDOUT"
	HashicorpWaypointExecStreamResponseOutputChannelSTDOUT HashicorpWaypointExecStreamResponseOutputChannel = "STDOUT"

	// HashicorpWaypointExecStreamResponseOutputChannelSTDERR captures enum value "STDERR"
	HashicorpWaypointExecStreamResponseOutputChannelSTDERR HashicorpWaypointExecStreamResponseOutputChannel = "STDERR"
)

// for schema
var hashicorpWaypointExecStreamResponseOutputChannelEnum []interface{}

func init() {
	var res []HashicorpWaypointExecStreamResponseOutputChannel
	if err := json.Unmarshal([]byte(`["UNKNOWN","STDOUT","STDERR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointExecStreamResponseOutputChannelEnum = append(hashicorpWaypointExecStreamResponseOutputChannelEnum, v)
	}
}

func (m HashicorpWaypointExecStreamResponseOutputChannel) validateHashicorpWaypointExecStreamResponseOutputChannelEnum(path, location string, value HashicorpWaypointExecStreamResponseOutputChannel) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointExecStreamResponseOutputChannelEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint exec stream response output channel
func (m HashicorpWaypointExecStreamResponseOutputChannel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointExecStreamResponseOutputChannelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp waypoint exec stream response output channel based on context it is used
func (m HashicorpWaypointExecStreamResponseOutputChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
