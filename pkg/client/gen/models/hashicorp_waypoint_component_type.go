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

// HashicorpWaypointComponentType Supported component types, the values here MUST match the enum values
// in the Go sdk/component package exactly. A test in internal/server
// validates this.
//
// swagger:model hashicorp.waypoint.Component.Type
type HashicorpWaypointComponentType string

func NewHashicorpWaypointComponentType(value HashicorpWaypointComponentType) *HashicorpWaypointComponentType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpWaypointComponentType.
func (m HashicorpWaypointComponentType) Pointer() *HashicorpWaypointComponentType {
	return &m
}

const (

	// HashicorpWaypointComponentTypeUNKNOWN captures enum value "UNKNOWN"
	HashicorpWaypointComponentTypeUNKNOWN HashicorpWaypointComponentType = "UNKNOWN"

	// HashicorpWaypointComponentTypeBUILDER captures enum value "BUILDER"
	HashicorpWaypointComponentTypeBUILDER HashicorpWaypointComponentType = "BUILDER"

	// HashicorpWaypointComponentTypeREGISTRY captures enum value "REGISTRY"
	HashicorpWaypointComponentTypeREGISTRY HashicorpWaypointComponentType = "REGISTRY"

	// HashicorpWaypointComponentTypePLATFORM captures enum value "PLATFORM"
	HashicorpWaypointComponentTypePLATFORM HashicorpWaypointComponentType = "PLATFORM"

	// HashicorpWaypointComponentTypeRELEASEMANAGER captures enum value "RELEASEMANAGER"
	HashicorpWaypointComponentTypeRELEASEMANAGER HashicorpWaypointComponentType = "RELEASEMANAGER"
)

// for schema
var hashicorpWaypointComponentTypeEnum []interface{}

func init() {
	var res []HashicorpWaypointComponentType
	if err := json.Unmarshal([]byte(`["UNKNOWN","BUILDER","REGISTRY","PLATFORM","RELEASEMANAGER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointComponentTypeEnum = append(hashicorpWaypointComponentTypeEnum, v)
	}
}

func (m HashicorpWaypointComponentType) validateHashicorpWaypointComponentTypeEnum(path, location string, value HashicorpWaypointComponentType) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointComponentTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint component type
func (m HashicorpWaypointComponentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointComponentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp waypoint component type based on context it is used
func (m HashicorpWaypointComponentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
