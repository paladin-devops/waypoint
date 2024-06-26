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

// WaypointListWorkspacesReader is a Reader for the WaypointListWorkspaces structure.
type WaypointListWorkspacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListWorkspacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListWorkspacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListWorkspacesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListWorkspacesOK creates a WaypointListWorkspacesOK with default headers values
func NewWaypointListWorkspacesOK() *WaypointListWorkspacesOK {
	return &WaypointListWorkspacesOK{}
}

/*
WaypointListWorkspacesOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListWorkspacesOK struct {
	Payload *models.HashicorpWaypointListWorkspacesResponse
}

// IsSuccess returns true when this waypoint list workspaces o k response has a 2xx status code
func (o *WaypointListWorkspacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list workspaces o k response has a 3xx status code
func (o *WaypointListWorkspacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list workspaces o k response has a 4xx status code
func (o *WaypointListWorkspacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list workspaces o k response has a 5xx status code
func (o *WaypointListWorkspacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list workspaces o k response a status code equal to that given
func (o *WaypointListWorkspacesOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointListWorkspacesOK) Error() string {
	return fmt.Sprintf("[GET /workspaces][%d] waypointListWorkspacesOK  %+v", 200, o.Payload)
}

func (o *WaypointListWorkspacesOK) String() string {
	return fmt.Sprintf("[GET /workspaces][%d] waypointListWorkspacesOK  %+v", 200, o.Payload)
}

func (o *WaypointListWorkspacesOK) GetPayload() *models.HashicorpWaypointListWorkspacesResponse {
	return o.Payload
}

func (o *WaypointListWorkspacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListWorkspacesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListWorkspacesDefault creates a WaypointListWorkspacesDefault with default headers values
func NewWaypointListWorkspacesDefault(code int) *WaypointListWorkspacesDefault {
	return &WaypointListWorkspacesDefault{
		_statusCode: code,
	}
}

/*
WaypointListWorkspacesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListWorkspacesDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list workspaces default response
func (o *WaypointListWorkspacesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint list workspaces default response has a 2xx status code
func (o *WaypointListWorkspacesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list workspaces default response has a 3xx status code
func (o *WaypointListWorkspacesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list workspaces default response has a 4xx status code
func (o *WaypointListWorkspacesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list workspaces default response has a 5xx status code
func (o *WaypointListWorkspacesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list workspaces default response a status code equal to that given
func (o *WaypointListWorkspacesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointListWorkspacesDefault) Error() string {
	return fmt.Sprintf("[GET /workspaces][%d] Waypoint_ListWorkspaces default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListWorkspacesDefault) String() string {
	return fmt.Sprintf("[GET /workspaces][%d] Waypoint_ListWorkspaces default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListWorkspacesDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListWorkspacesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
