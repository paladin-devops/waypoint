// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointCreateSnapshotResponse hashicorp waypoint create snapshot response
//
// swagger:model hashicorp.waypoint.CreateSnapshotResponse
type HashicorpWaypointCreateSnapshotResponse struct {

	// Chunk is a next chunk of data. You should continue to expect
	// data until an EOF is received on the stream.
	// Format: byte
	Chunk strfmt.Base64 `json:"chunk,omitempty"`

	// Open is sent as the opening message with information about the
	// snapshot. This is always sent first (before any data).
	Open HashicorpWaypointCreateSnapshotResponseOpen `json:"open,omitempty"`
}

// Validate validates this hashicorp waypoint create snapshot response
func (m *HashicorpWaypointCreateSnapshotResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint create snapshot response based on context it is used
func (m *HashicorpWaypointCreateSnapshotResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointCreateSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointCreateSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointCreateSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
