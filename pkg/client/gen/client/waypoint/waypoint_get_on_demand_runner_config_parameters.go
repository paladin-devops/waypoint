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

// NewWaypointGetOnDemandRunnerConfigParams creates a new WaypointGetOnDemandRunnerConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetOnDemandRunnerConfigParams() *WaypointGetOnDemandRunnerConfigParams {
	return &WaypointGetOnDemandRunnerConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetOnDemandRunnerConfigParamsWithTimeout creates a new WaypointGetOnDemandRunnerConfigParams object
// with the ability to set a timeout on a request.
func NewWaypointGetOnDemandRunnerConfigParamsWithTimeout(timeout time.Duration) *WaypointGetOnDemandRunnerConfigParams {
	return &WaypointGetOnDemandRunnerConfigParams{
		timeout: timeout,
	}
}

// NewWaypointGetOnDemandRunnerConfigParamsWithContext creates a new WaypointGetOnDemandRunnerConfigParams object
// with the ability to set a context for a request.
func NewWaypointGetOnDemandRunnerConfigParamsWithContext(ctx context.Context) *WaypointGetOnDemandRunnerConfigParams {
	return &WaypointGetOnDemandRunnerConfigParams{
		Context: ctx,
	}
}

// NewWaypointGetOnDemandRunnerConfigParamsWithHTTPClient creates a new WaypointGetOnDemandRunnerConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetOnDemandRunnerConfigParamsWithHTTPClient(client *http.Client) *WaypointGetOnDemandRunnerConfigParams {
	return &WaypointGetOnDemandRunnerConfigParams{
		HTTPClient: client,
	}
}

/*
WaypointGetOnDemandRunnerConfigParams contains all the parameters to send to the API endpoint

	for the waypoint get on demand runner config operation.

	Typically these are written to a http.Request.
*/
type WaypointGetOnDemandRunnerConfigParams struct {

	// ConfigID.
	ConfigID string

	// ConfigName.
	ConfigName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get on demand runner config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetOnDemandRunnerConfigParams) WithDefaults() *WaypointGetOnDemandRunnerConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get on demand runner config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetOnDemandRunnerConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) WithTimeout(timeout time.Duration) *WaypointGetOnDemandRunnerConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) WithContext(ctx context.Context) *WaypointGetOnDemandRunnerConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) WithHTTPClient(client *http.Client) *WaypointGetOnDemandRunnerConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigID adds the configID to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) WithConfigID(configID string) *WaypointGetOnDemandRunnerConfigParams {
	o.SetConfigID(configID)
	return o
}

// SetConfigID adds the configId to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) SetConfigID(configID string) {
	o.ConfigID = configID
}

// WithConfigName adds the configName to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) WithConfigName(configName *string) *WaypointGetOnDemandRunnerConfigParams {
	o.SetConfigName(configName)
	return o
}

// SetConfigName adds the configName to the waypoint get on demand runner config params
func (o *WaypointGetOnDemandRunnerConfigParams) SetConfigName(configName *string) {
	o.ConfigName = configName
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetOnDemandRunnerConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param config.id
	if err := r.SetPathParam("config.id", o.ConfigID); err != nil {
		return err
	}

	if o.ConfigName != nil {

		// query param config.name
		var qrConfigName string

		if o.ConfigName != nil {
			qrConfigName = *o.ConfigName
		}
		qConfigName := qrConfigName
		if qConfigName != "" {

			if err := r.SetQueryParam("config.name", qConfigName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
