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

// NewWaypointListOnDemandRunnerConfigsParams creates a new WaypointListOnDemandRunnerConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListOnDemandRunnerConfigsParams() *WaypointListOnDemandRunnerConfigsParams {
	return &WaypointListOnDemandRunnerConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListOnDemandRunnerConfigsParamsWithTimeout creates a new WaypointListOnDemandRunnerConfigsParams object
// with the ability to set a timeout on a request.
func NewWaypointListOnDemandRunnerConfigsParamsWithTimeout(timeout time.Duration) *WaypointListOnDemandRunnerConfigsParams {
	return &WaypointListOnDemandRunnerConfigsParams{
		timeout: timeout,
	}
}

// NewWaypointListOnDemandRunnerConfigsParamsWithContext creates a new WaypointListOnDemandRunnerConfigsParams object
// with the ability to set a context for a request.
func NewWaypointListOnDemandRunnerConfigsParamsWithContext(ctx context.Context) *WaypointListOnDemandRunnerConfigsParams {
	return &WaypointListOnDemandRunnerConfigsParams{
		Context: ctx,
	}
}

// NewWaypointListOnDemandRunnerConfigsParamsWithHTTPClient creates a new WaypointListOnDemandRunnerConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListOnDemandRunnerConfigsParamsWithHTTPClient(client *http.Client) *WaypointListOnDemandRunnerConfigsParams {
	return &WaypointListOnDemandRunnerConfigsParams{
		HTTPClient: client,
	}
}

/*
WaypointListOnDemandRunnerConfigsParams contains all the parameters to send to the API endpoint

	for the waypoint list on demand runner configs operation.

	Typically these are written to a http.Request.
*/
type WaypointListOnDemandRunnerConfigsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list on demand runner configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListOnDemandRunnerConfigsParams) WithDefaults() *WaypointListOnDemandRunnerConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list on demand runner configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListOnDemandRunnerConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list on demand runner configs params
func (o *WaypointListOnDemandRunnerConfigsParams) WithTimeout(timeout time.Duration) *WaypointListOnDemandRunnerConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list on demand runner configs params
func (o *WaypointListOnDemandRunnerConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list on demand runner configs params
func (o *WaypointListOnDemandRunnerConfigsParams) WithContext(ctx context.Context) *WaypointListOnDemandRunnerConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list on demand runner configs params
func (o *WaypointListOnDemandRunnerConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list on demand runner configs params
func (o *WaypointListOnDemandRunnerConfigsParams) WithHTTPClient(client *http.Client) *WaypointListOnDemandRunnerConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list on demand runner configs params
func (o *WaypointListOnDemandRunnerConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListOnDemandRunnerConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
