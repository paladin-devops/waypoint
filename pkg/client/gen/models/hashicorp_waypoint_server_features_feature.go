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

// HashicorpWaypointServerFeaturesFeature - FEATURE_INLINE_KEEPALIVES: Advertises that this server is capable of receiving inline keepalive messages
//
// swagger:model hashicorp.waypoint.ServerFeatures.feature
type HashicorpWaypointServerFeaturesFeature string

func NewHashicorpWaypointServerFeaturesFeature(value HashicorpWaypointServerFeaturesFeature) *HashicorpWaypointServerFeaturesFeature {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpWaypointServerFeaturesFeature.
func (m HashicorpWaypointServerFeaturesFeature) Pointer() *HashicorpWaypointServerFeaturesFeature {
	return &m
}

const (

	// HashicorpWaypointServerFeaturesFeatureFEATUREUNSPECIFIED captures enum value "FEATURE_UNSPECIFIED"
	HashicorpWaypointServerFeaturesFeatureFEATUREUNSPECIFIED HashicorpWaypointServerFeaturesFeature = "FEATURE_UNSPECIFIED"

	// HashicorpWaypointServerFeaturesFeatureFEATUREINLINEKEEPALIVES captures enum value "FEATURE_INLINE_KEEPALIVES"
	HashicorpWaypointServerFeaturesFeatureFEATUREINLINEKEEPALIVES HashicorpWaypointServerFeaturesFeature = "FEATURE_INLINE_KEEPALIVES"
)

// for schema
var hashicorpWaypointServerFeaturesFeatureEnum []interface{}

func init() {
	var res []HashicorpWaypointServerFeaturesFeature
	if err := json.Unmarshal([]byte(`["FEATURE_UNSPECIFIED","FEATURE_INLINE_KEEPALIVES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointServerFeaturesFeatureEnum = append(hashicorpWaypointServerFeaturesFeatureEnum, v)
	}
}

func (m HashicorpWaypointServerFeaturesFeature) validateHashicorpWaypointServerFeaturesFeatureEnum(path, location string, value HashicorpWaypointServerFeaturesFeature) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointServerFeaturesFeatureEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint server features feature
func (m HashicorpWaypointServerFeaturesFeature) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointServerFeaturesFeatureEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp waypoint server features feature based on context it is used
func (m HashicorpWaypointServerFeaturesFeature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
