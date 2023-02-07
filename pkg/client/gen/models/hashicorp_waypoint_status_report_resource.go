// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpWaypointStatusReportResource A resource as observed in a platform
//
// swagger:model hashicorp.waypoint.StatusReport.Resource
type HashicorpWaypointStatusReportResource struct {

	// The high level category of the resource, used as a hint to the UI on how to display the resource.
	CategoryDisplayHint *HashicorpWaypointResourceCategoryDisplayHint `json:"category_display_hint,omitempty"`

	// platform-reported time of resource creation
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"created_time,omitempty"`

	// declared resource that this directly references.
	DeclaredResource *HashicorpWaypointRefDeclaredResource `json:"declared_resource,omitempty"`

	// deprecated in favor of the Health enum and message to match the plugin protos. Was never used.
	DeprecatedHealth *HashicorpWaypointStatusReportHealth `json:"deprecated_health,omitempty"`

	// the current health state for a single resource
	Health *HashicorpWaypointStatusReportResourceHealth `json:"health,omitempty"`

	// a simple human readable message detailing the Health state
	HealthMessage string `json:"health_message,omitempty"`

	// The id of the resource, according to the platform.
	ID string `json:"id,omitempty"`

	// Friendly name of the resource, if applicable
	Name string `json:"name,omitempty"`

	// Resources that created this resource.
	ParentResourceID string `json:"parent_resource_id,omitempty"`

	// The platform on which the resource exists.
	Platform string `json:"platform,omitempty"`

	// A link directly to the resource in the platform, if applicable.
	PlatformURL string `json:"platform_url,omitempty"`

	// any additional metadata about the resource, encoded as JSON
	StateJSON string `json:"state_json,omitempty"`

	// platform-specific name of the resource type. i.e. instance, pod, auto-scaling group, etc
	Type string `json:"type,omitempty"`
}

// Validate validates this hashicorp waypoint status report resource
func (m *HashicorpWaypointStatusReportResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryDisplayHint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeclaredResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeprecatedHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointStatusReportResource) validateCategoryDisplayHint(formats strfmt.Registry) error {
	if swag.IsZero(m.CategoryDisplayHint) { // not required
		return nil
	}

	if m.CategoryDisplayHint != nil {
		if err := m.CategoryDisplayHint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category_display_hint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category_display_hint")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointStatusReportResource) validateCreatedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("created_time", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpWaypointStatusReportResource) validateDeclaredResource(formats strfmt.Registry) error {
	if swag.IsZero(m.DeclaredResource) { // not required
		return nil
	}

	if m.DeclaredResource != nil {
		if err := m.DeclaredResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("declared_resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("declared_resource")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointStatusReportResource) validateDeprecatedHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.DeprecatedHealth) { // not required
		return nil
	}

	if m.DeprecatedHealth != nil {
		if err := m.DeprecatedHealth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deprecated_health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deprecated_health")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointStatusReportResource) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint status report resource based on the context it is used
func (m *HashicorpWaypointStatusReportResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategoryDisplayHint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeclaredResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeprecatedHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointStatusReportResource) contextValidateCategoryDisplayHint(ctx context.Context, formats strfmt.Registry) error {

	if m.CategoryDisplayHint != nil {
		if err := m.CategoryDisplayHint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category_display_hint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category_display_hint")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointStatusReportResource) contextValidateDeclaredResource(ctx context.Context, formats strfmt.Registry) error {

	if m.DeclaredResource != nil {
		if err := m.DeclaredResource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("declared_resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("declared_resource")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointStatusReportResource) contextValidateDeprecatedHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.DeprecatedHealth != nil {
		if err := m.DeprecatedHealth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deprecated_health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deprecated_health")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointStatusReportResource) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {
		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointStatusReportResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointStatusReportResource) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointStatusReportResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
