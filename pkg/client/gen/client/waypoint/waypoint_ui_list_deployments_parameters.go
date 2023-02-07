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
	"github.com/go-openapi/swag"
)

// NewWaypointUIListDeploymentsParams creates a new WaypointUIListDeploymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUIListDeploymentsParams() *WaypointUIListDeploymentsParams {
	return &WaypointUIListDeploymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUIListDeploymentsParamsWithTimeout creates a new WaypointUIListDeploymentsParams object
// with the ability to set a timeout on a request.
func NewWaypointUIListDeploymentsParamsWithTimeout(timeout time.Duration) *WaypointUIListDeploymentsParams {
	return &WaypointUIListDeploymentsParams{
		timeout: timeout,
	}
}

// NewWaypointUIListDeploymentsParamsWithContext creates a new WaypointUIListDeploymentsParams object
// with the ability to set a context for a request.
func NewWaypointUIListDeploymentsParamsWithContext(ctx context.Context) *WaypointUIListDeploymentsParams {
	return &WaypointUIListDeploymentsParams{
		Context: ctx,
	}
}

// NewWaypointUIListDeploymentsParamsWithHTTPClient creates a new WaypointUIListDeploymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUIListDeploymentsParamsWithHTTPClient(client *http.Client) *WaypointUIListDeploymentsParams {
	return &WaypointUIListDeploymentsParams{
		HTTPClient: client,
	}
}

/*
WaypointUIListDeploymentsParams contains all the parameters to send to the API endpoint

	for the waypoint UI list deployments operation.

	Typically these are written to a http.Request.
*/
type WaypointUIListDeploymentsParams struct {

	// ApplicationApplication.
	ApplicationApplication *string

	// ApplicationProject.
	ApplicationProject *string

	// OrderDesc.
	OrderDesc *bool

	/* OrderLimit.

	   Limit the number of results.

	   Format: int64
	*/
	OrderLimit *int64

	/* OrderOrder.

	   Order for the results.

	   Default: "UNSET"
	*/
	OrderOrder *string

	/* PhysicalState.

	     The physical state to filter for. If this is zero or unset then no
	filtering on physical state will be done.

	     Default: "UNKNOWN"
	*/
	PhysicalState *string

	// WorkspaceWorkspace.
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint UI list deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIListDeploymentsParams) WithDefaults() *WaypointUIListDeploymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint UI list deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIListDeploymentsParams) SetDefaults() {
	var (
		orderOrderDefault = string("UNSET")

		physicalStateDefault = string("UNKNOWN")
	)

	val := WaypointUIListDeploymentsParams{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithTimeout(timeout time.Duration) *WaypointUIListDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithContext(ctx context.Context) *WaypointUIListDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithHTTPClient(client *http.Client) *WaypointUIListDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithApplicationApplication(applicationApplication *string) *WaypointUIListDeploymentsParams {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetApplicationApplication(applicationApplication *string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithApplicationProject(applicationProject *string) *WaypointUIListDeploymentsParams {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetApplicationProject(applicationProject *string) {
	o.ApplicationProject = applicationProject
}

// WithOrderDesc adds the orderDesc to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithOrderDesc(orderDesc *bool) *WaypointUIListDeploymentsParams {
	o.SetOrderDesc(orderDesc)
	return o
}

// SetOrderDesc adds the orderDesc to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetOrderDesc(orderDesc *bool) {
	o.OrderDesc = orderDesc
}

// WithOrderLimit adds the orderLimit to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithOrderLimit(orderLimit *int64) *WaypointUIListDeploymentsParams {
	o.SetOrderLimit(orderLimit)
	return o
}

// SetOrderLimit adds the orderLimit to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetOrderLimit(orderLimit *int64) {
	o.OrderLimit = orderLimit
}

// WithOrderOrder adds the orderOrder to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithOrderOrder(orderOrder *string) *WaypointUIListDeploymentsParams {
	o.SetOrderOrder(orderOrder)
	return o
}

// SetOrderOrder adds the orderOrder to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetOrderOrder(orderOrder *string) {
	o.OrderOrder = orderOrder
}

// WithPhysicalState adds the physicalState to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithPhysicalState(physicalState *string) *WaypointUIListDeploymentsParams {
	o.SetPhysicalState(physicalState)
	return o
}

// SetPhysicalState adds the physicalState to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetPhysicalState(physicalState *string) {
	o.PhysicalState = physicalState
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointUIListDeploymentsParams {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list deployments params
func (o *WaypointUIListDeploymentsParams) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUIListDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationApplication != nil {

		// query param application.application
		var qrApplicationApplication string

		if o.ApplicationApplication != nil {
			qrApplicationApplication = *o.ApplicationApplication
		}
		qApplicationApplication := qrApplicationApplication
		if qApplicationApplication != "" {

			if err := r.SetQueryParam("application.application", qApplicationApplication); err != nil {
				return err
			}
		}
	}

	if o.ApplicationProject != nil {

		// query param application.project
		var qrApplicationProject string

		if o.ApplicationProject != nil {
			qrApplicationProject = *o.ApplicationProject
		}
		qApplicationProject := qrApplicationProject
		if qApplicationProject != "" {

			if err := r.SetQueryParam("application.project", qApplicationProject); err != nil {
				return err
			}
		}
	}

	if o.OrderDesc != nil {

		// query param order.desc
		var qrOrderDesc bool

		if o.OrderDesc != nil {
			qrOrderDesc = *o.OrderDesc
		}
		qOrderDesc := swag.FormatBool(qrOrderDesc)
		if qOrderDesc != "" {

			if err := r.SetQueryParam("order.desc", qOrderDesc); err != nil {
				return err
			}
		}
	}

	if o.OrderLimit != nil {

		// query param order.limit
		var qrOrderLimit int64

		if o.OrderLimit != nil {
			qrOrderLimit = *o.OrderLimit
		}
		qOrderLimit := swag.FormatInt64(qrOrderLimit)
		if qOrderLimit != "" {

			if err := r.SetQueryParam("order.limit", qOrderLimit); err != nil {
				return err
			}
		}
	}

	if o.OrderOrder != nil {

		// query param order.order
		var qrOrderOrder string

		if o.OrderOrder != nil {
			qrOrderOrder = *o.OrderOrder
		}
		qOrderOrder := qrOrderOrder
		if qOrderOrder != "" {

			if err := r.SetQueryParam("order.order", qOrderOrder); err != nil {
				return err
			}
		}
	}

	if o.PhysicalState != nil {

		// query param physical_state
		var qrPhysicalState string

		if o.PhysicalState != nil {
			qrPhysicalState = *o.PhysicalState
		}
		qPhysicalState := qrPhysicalState
		if qPhysicalState != "" {

			if err := r.SetQueryParam("physical_state", qPhysicalState); err != nil {
				return err
			}
		}
	}

	// path param workspace.workspace
	if err := r.SetPathParam("workspace.workspace", o.WorkspaceWorkspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
