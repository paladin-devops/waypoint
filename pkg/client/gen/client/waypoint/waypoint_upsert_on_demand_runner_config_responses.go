// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint/pkg/client/gen/models"
)

// WaypointUpsertOnDemandRunnerConfigReader is a Reader for the WaypointUpsertOnDemandRunnerConfig structure.
type WaypointUpsertOnDemandRunnerConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUpsertOnDemandRunnerConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUpsertOnDemandRunnerConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUpsertOnDemandRunnerConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUpsertOnDemandRunnerConfigOK creates a WaypointUpsertOnDemandRunnerConfigOK with default headers values
func NewWaypointUpsertOnDemandRunnerConfigOK() *WaypointUpsertOnDemandRunnerConfigOK {
	return &WaypointUpsertOnDemandRunnerConfigOK{}
}

/*
WaypointUpsertOnDemandRunnerConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointUpsertOnDemandRunnerConfigOK struct {
	Payload *models.HashicorpWaypointUpsertOnDemandRunnerConfigResponse
}

// IsSuccess returns true when this waypoint upsert on demand runner config o k response has a 2xx status code
func (o *WaypointUpsertOnDemandRunnerConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint upsert on demand runner config o k response has a 3xx status code
func (o *WaypointUpsertOnDemandRunnerConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint upsert on demand runner config o k response has a 4xx status code
func (o *WaypointUpsertOnDemandRunnerConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint upsert on demand runner config o k response has a 5xx status code
func (o *WaypointUpsertOnDemandRunnerConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint upsert on demand runner config o k response a status code equal to that given
func (o *WaypointUpsertOnDemandRunnerConfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointUpsertOnDemandRunnerConfigOK) Error() string {
	return fmt.Sprintf("[POST /runner/on-demand][%d] waypointUpsertOnDemandRunnerConfigOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertOnDemandRunnerConfigOK) String() string {
	return fmt.Sprintf("[POST /runner/on-demand][%d] waypointUpsertOnDemandRunnerConfigOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertOnDemandRunnerConfigOK) GetPayload() *models.HashicorpWaypointUpsertOnDemandRunnerConfigResponse {
	return o.Payload
}

func (o *WaypointUpsertOnDemandRunnerConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUpsertOnDemandRunnerConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUpsertOnDemandRunnerConfigDefault creates a WaypointUpsertOnDemandRunnerConfigDefault with default headers values
func NewWaypointUpsertOnDemandRunnerConfigDefault(code int) *WaypointUpsertOnDemandRunnerConfigDefault {
	return &WaypointUpsertOnDemandRunnerConfigDefault{
		_statusCode: code,
	}
}

/*
WaypointUpsertOnDemandRunnerConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointUpsertOnDemandRunnerConfigDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint upsert on demand runner config default response
func (o *WaypointUpsertOnDemandRunnerConfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint upsert on demand runner config default response has a 2xx status code
func (o *WaypointUpsertOnDemandRunnerConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint upsert on demand runner config default response has a 3xx status code
func (o *WaypointUpsertOnDemandRunnerConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint upsert on demand runner config default response has a 4xx status code
func (o *WaypointUpsertOnDemandRunnerConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint upsert on demand runner config default response has a 5xx status code
func (o *WaypointUpsertOnDemandRunnerConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint upsert on demand runner config default response a status code equal to that given
func (o *WaypointUpsertOnDemandRunnerConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointUpsertOnDemandRunnerConfigDefault) Error() string {
	return fmt.Sprintf("[POST /runner/on-demand][%d] Waypoint_UpsertOnDemandRunnerConfig default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertOnDemandRunnerConfigDefault) String() string {
	return fmt.Sprintf("[POST /runner/on-demand][%d] Waypoint_UpsertOnDemandRunnerConfig default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertOnDemandRunnerConfigDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUpsertOnDemandRunnerConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
