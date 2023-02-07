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

// NewWaypointUIGetProjectParams creates a new WaypointUIGetProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUIGetProjectParams() *WaypointUIGetProjectParams {
	return &WaypointUIGetProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUIGetProjectParamsWithTimeout creates a new WaypointUIGetProjectParams object
// with the ability to set a timeout on a request.
func NewWaypointUIGetProjectParamsWithTimeout(timeout time.Duration) *WaypointUIGetProjectParams {
	return &WaypointUIGetProjectParams{
		timeout: timeout,
	}
}

// NewWaypointUIGetProjectParamsWithContext creates a new WaypointUIGetProjectParams object
// with the ability to set a context for a request.
func NewWaypointUIGetProjectParamsWithContext(ctx context.Context) *WaypointUIGetProjectParams {
	return &WaypointUIGetProjectParams{
		Context: ctx,
	}
}

// NewWaypointUIGetProjectParamsWithHTTPClient creates a new WaypointUIGetProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUIGetProjectParamsWithHTTPClient(client *http.Client) *WaypointUIGetProjectParams {
	return &WaypointUIGetProjectParams{
		HTTPClient: client,
	}
}

/*
WaypointUIGetProjectParams contains all the parameters to send to the API endpoint

	for the waypoint UI get project operation.

	Typically these are written to a http.Request.
*/
type WaypointUIGetProjectParams struct {

	// ProjectProject.
	ProjectProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint UI get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIGetProjectParams) WithDefaults() *WaypointUIGetProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint UI get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIGetProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) WithTimeout(timeout time.Duration) *WaypointUIGetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) WithContext(ctx context.Context) *WaypointUIGetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) WithHTTPClient(client *http.Client) *WaypointUIGetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectProject adds the projectProject to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) WithProjectProject(projectProject string) *WaypointUIGetProjectParams {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint UI get project params
func (o *WaypointUIGetProjectParams) SetProjectProject(projectProject string) {
	o.ProjectProject = projectProject
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUIGetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project.project
	if err := r.SetPathParam("project.project", o.ProjectProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
