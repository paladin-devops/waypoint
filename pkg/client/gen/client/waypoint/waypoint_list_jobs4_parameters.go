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

// NewWaypointListJobs4Params creates a new WaypointListJobs4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointListJobs4Params() *WaypointListJobs4Params {
	return &WaypointListJobs4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListJobs4ParamsWithTimeout creates a new WaypointListJobs4Params object
// with the ability to set a timeout on a request.
func NewWaypointListJobs4ParamsWithTimeout(timeout time.Duration) *WaypointListJobs4Params {
	return &WaypointListJobs4Params{
		timeout: timeout,
	}
}

// NewWaypointListJobs4ParamsWithContext creates a new WaypointListJobs4Params object
// with the ability to set a context for a request.
func NewWaypointListJobs4ParamsWithContext(ctx context.Context) *WaypointListJobs4Params {
	return &WaypointListJobs4Params{
		Context: ctx,
	}
}

// NewWaypointListJobs4ParamsWithHTTPClient creates a new WaypointListJobs4Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointListJobs4ParamsWithHTTPClient(client *http.Client) *WaypointListJobs4Params {
	return &WaypointListJobs4Params{
		HTTPClient: client,
	}
}

/*
WaypointListJobs4Params contains all the parameters to send to the API endpoint

	for the waypoint list jobs4 operation.

	Typically these are written to a http.Request.
*/
type WaypointListJobs4Params struct {

	// ApplicationApplication.
	ApplicationApplication *string

	// ApplicationProject.
	ApplicationProject *string

	// JobState.
	JobState []string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	/* PipelinePipelineID.

	   ID of the pipeline.
	*/
	PipelinePipelineID *string

	/* PipelinePipelineName.

	   Name of the pipeline.
	*/
	PipelinePipelineName *string

	/* PipelineRunSequence.

	   Pipeline run sequence.

	   Format: uint64
	*/
	PipelineRunSequence *string

	/* PipelineStep.

	   Step name within the pipeline.
	*/
	PipelineStep *string

	// ProjectProject.
	ProjectProject *string

	// TargetRunnerIDID.
	TargetRunnerIDID *string

	// WorkspaceWorkspace.
	WorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint list jobs4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListJobs4Params) WithDefaults() *WaypointListJobs4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint list jobs4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointListJobs4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithTimeout(timeout time.Duration) *WaypointListJobs4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithContext(ctx context.Context) *WaypointListJobs4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithHTTPClient(client *http.Client) *WaypointListJobs4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithApplicationApplication(applicationApplication *string) *WaypointListJobs4Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetApplicationApplication(applicationApplication *string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithApplicationProject(applicationProject *string) *WaypointListJobs4Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetApplicationProject(applicationProject *string) {
	o.ApplicationProject = applicationProject
}

// WithJobState adds the jobState to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithJobState(jobState []string) *WaypointListJobs4Params {
	o.SetJobState(jobState)
	return o
}

// SetJobState adds the jobState to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetJobState(jobState []string) {
	o.JobState = jobState
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithPaginationNextPageToken(paginationNextPageToken *string) *WaypointListJobs4Params {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithPaginationPageSize(paginationPageSize *int64) *WaypointListJobs4Params {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *WaypointListJobs4Params {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithPipelinePipelineID adds the pipelinePipelineID to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithPipelinePipelineID(pipelinePipelineID *string) *WaypointListJobs4Params {
	o.SetPipelinePipelineID(pipelinePipelineID)
	return o
}

// SetPipelinePipelineID adds the pipelinePipelineId to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetPipelinePipelineID(pipelinePipelineID *string) {
	o.PipelinePipelineID = pipelinePipelineID
}

// WithPipelinePipelineName adds the pipelinePipelineName to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithPipelinePipelineName(pipelinePipelineName *string) *WaypointListJobs4Params {
	o.SetPipelinePipelineName(pipelinePipelineName)
	return o
}

// SetPipelinePipelineName adds the pipelinePipelineName to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetPipelinePipelineName(pipelinePipelineName *string) {
	o.PipelinePipelineName = pipelinePipelineName
}

// WithPipelineRunSequence adds the pipelineRunSequence to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithPipelineRunSequence(pipelineRunSequence *string) *WaypointListJobs4Params {
	o.SetPipelineRunSequence(pipelineRunSequence)
	return o
}

// SetPipelineRunSequence adds the pipelineRunSequence to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetPipelineRunSequence(pipelineRunSequence *string) {
	o.PipelineRunSequence = pipelineRunSequence
}

// WithPipelineStep adds the pipelineStep to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithPipelineStep(pipelineStep *string) *WaypointListJobs4Params {
	o.SetPipelineStep(pipelineStep)
	return o
}

// SetPipelineStep adds the pipelineStep to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetPipelineStep(pipelineStep *string) {
	o.PipelineStep = pipelineStep
}

// WithProjectProject adds the projectProject to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithProjectProject(projectProject *string) *WaypointListJobs4Params {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetProjectProject(projectProject *string) {
	o.ProjectProject = projectProject
}

// WithTargetRunnerIDID adds the targetRunnerIDID to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithTargetRunnerIDID(targetRunnerIDID *string) *WaypointListJobs4Params {
	o.SetTargetRunnerIDID(targetRunnerIDID)
	return o
}

// SetTargetRunnerIDID adds the targetRunnerIdId to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetTargetRunnerIDID(targetRunnerIDID *string) {
	o.TargetRunnerIDID = targetRunnerIDID
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) WithWorkspaceWorkspace(workspaceWorkspace *string) *WaypointListJobs4Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list jobs4 params
func (o *WaypointListJobs4Params) SetWorkspaceWorkspace(workspaceWorkspace *string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListJobs4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.JobState != nil {

		// binding items for jobState
		joinedJobState := o.bindParamJobState(reg)

		// path array param jobState
		// SetPathParam does not support variadic arguments, since we used JoinByFormat
		// we can send the first item in the array as it's all the items of the previous
		// array joined together
		if len(joinedJobState) > 0 {
			if err := r.SetPathParam("jobState", joinedJobState[0]); err != nil {
				return err
			}
		}
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if o.PipelinePipelineID != nil {

		// query param pipeline.pipeline_id
		var qrPipelinePipelineID string

		if o.PipelinePipelineID != nil {
			qrPipelinePipelineID = *o.PipelinePipelineID
		}
		qPipelinePipelineID := qrPipelinePipelineID
		if qPipelinePipelineID != "" {

			if err := r.SetQueryParam("pipeline.pipeline_id", qPipelinePipelineID); err != nil {
				return err
			}
		}
	}

	if o.PipelinePipelineName != nil {

		// query param pipeline.pipeline_name
		var qrPipelinePipelineName string

		if o.PipelinePipelineName != nil {
			qrPipelinePipelineName = *o.PipelinePipelineName
		}
		qPipelinePipelineName := qrPipelinePipelineName
		if qPipelinePipelineName != "" {

			if err := r.SetQueryParam("pipeline.pipeline_name", qPipelinePipelineName); err != nil {
				return err
			}
		}
	}

	if o.PipelineRunSequence != nil {

		// query param pipeline.run_sequence
		var qrPipelineRunSequence string

		if o.PipelineRunSequence != nil {
			qrPipelineRunSequence = *o.PipelineRunSequence
		}
		qPipelineRunSequence := qrPipelineRunSequence
		if qPipelineRunSequence != "" {

			if err := r.SetQueryParam("pipeline.run_sequence", qPipelineRunSequence); err != nil {
				return err
			}
		}
	}

	if o.PipelineStep != nil {

		// query param pipeline.step
		var qrPipelineStep string

		if o.PipelineStep != nil {
			qrPipelineStep = *o.PipelineStep
		}
		qPipelineStep := qrPipelineStep
		if qPipelineStep != "" {

			if err := r.SetQueryParam("pipeline.step", qPipelineStep); err != nil {
				return err
			}
		}
	}

	if o.ProjectProject != nil {

		// query param project.project
		var qrProjectProject string

		if o.ProjectProject != nil {
			qrProjectProject = *o.ProjectProject
		}
		qProjectProject := qrProjectProject
		if qProjectProject != "" {

			if err := r.SetQueryParam("project.project", qProjectProject); err != nil {
				return err
			}
		}
	}

	if o.TargetRunnerIDID != nil {

		// query param targetRunner.id.id
		var qrTargetRunnerIDID string

		if o.TargetRunnerIDID != nil {
			qrTargetRunnerIDID = *o.TargetRunnerIDID
		}
		qTargetRunnerIDID := qrTargetRunnerIDID
		if qTargetRunnerIDID != "" {

			if err := r.SetQueryParam("targetRunner.id.id", qTargetRunnerIDID); err != nil {
				return err
			}
		}
	}

	if o.WorkspaceWorkspace != nil {

		// query param workspace.workspace
		var qrWorkspaceWorkspace string

		if o.WorkspaceWorkspace != nil {
			qrWorkspaceWorkspace = *o.WorkspaceWorkspace
		}
		qWorkspaceWorkspace := qrWorkspaceWorkspace
		if qWorkspaceWorkspace != "" {

			if err := r.SetQueryParam("workspace.workspace", qWorkspaceWorkspace); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWaypointListJobs4 binds the parameter jobState
func (o *WaypointListJobs4Params) bindParamJobState(formats strfmt.Registry) []string {
	jobStateIR := o.JobState

	var jobStateIC []string
	for _, jobStateIIR := range jobStateIR { // explode []string

		jobStateIIV := jobStateIIR // string as string
		jobStateIC = append(jobStateIC, jobStateIIV)
	}

	// items.CollectionFormat: "csv"
	jobStateIS := swag.JoinByFormat(jobStateIC, "csv")

	return jobStateIS
}
