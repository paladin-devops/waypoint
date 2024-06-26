// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewWaypointListRunnersParams creates a new WaypointListRunnersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListRunnersParams() *WaypointListRunnersParams {
	return &WaypointListRunnersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListRunnersParamsWithTimeout creates a new WaypointListRunnersParams object
// with the ability to set a timeout on a request.
func NewWaypointListRunnersParamsWithTimeout(timeout time.Duration) *WaypointListRunnersParams {
	return &WaypointListRunnersParams{
		timeout: timeout,
	}
}

// NewWaypointListRunnersParamsWithContext creates a new WaypointListRunnersParams object
// with the ability to set a context for a request.
func NewWaypointListRunnersParamsWithContext(ctx context.Context) *WaypointListRunnersParams {
	return &WaypointListRunnersParams{
		Context: ctx,
	}
}

// NewWaypointListRunnersParamsWithHTTPClient creates a new WaypointListRunnersParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListRunnersParamsWithHTTPClient(client *http.Client) *WaypointListRunnersParams {
	return &WaypointListRunnersParams{
		HTTPClient: client,
	}
}

/*
WaypointListRunnersParams contains all the parameters to send to the API endpoint

	for the waypoint list runners operation.

	Typically these are written to a http.Request.
*/
type WaypointListRunnersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list runners params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListRunnersParams) WithDefaults() *WaypointListRunnersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list runners params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListRunnersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list runners params
func (o *WaypointListRunnersParams) WithTimeout(timeout time.Duration) *WaypointListRunnersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list runners params
func (o *WaypointListRunnersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list runners params
func (o *WaypointListRunnersParams) WithContext(ctx context.Context) *WaypointListRunnersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list runners params
func (o *WaypointListRunnersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list runners params
func (o *WaypointListRunnersParams) WithHTTPClient(client *http.Client) *WaypointListRunnersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list runners params
func (o *WaypointListRunnersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListRunnersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
