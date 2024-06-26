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

// WaypointListWorkspaces3Reader is a Reader for the WaypointListWorkspaces3 structure.
type WaypointListWorkspaces3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListWorkspaces3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListWorkspaces3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListWorkspaces3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListWorkspaces3OK creates a WaypointListWorkspaces3OK with default headers values
func NewWaypointListWorkspaces3OK() *WaypointListWorkspaces3OK {
	return &WaypointListWorkspaces3OK{}
}

/*
WaypointListWorkspaces3OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListWorkspaces3OK struct {
	Payload *models.HashicorpWaypointListWorkspacesResponse
}

// IsSuccess returns true when this waypoint list workspaces3 o k response has a 2xx status code
func (o *WaypointListWorkspaces3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list workspaces3 o k response has a 3xx status code
func (o *WaypointListWorkspaces3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list workspaces3 o k response has a 4xx status code
func (o *WaypointListWorkspaces3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list workspaces3 o k response has a 5xx status code
func (o *WaypointListWorkspaces3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list workspaces3 o k response a status code equal to that given
func (o *WaypointListWorkspaces3OK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointListWorkspaces3OK) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspaces][%d] waypointListWorkspaces3OK  %+v", 200, o.Payload)
}

func (o *WaypointListWorkspaces3OK) String() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspaces][%d] waypointListWorkspaces3OK  %+v", 200, o.Payload)
}

func (o *WaypointListWorkspaces3OK) GetPayload() *models.HashicorpWaypointListWorkspacesResponse {
	return o.Payload
}

func (o *WaypointListWorkspaces3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListWorkspacesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListWorkspaces3Default creates a WaypointListWorkspaces3Default with default headers values
func NewWaypointListWorkspaces3Default(code int) *WaypointListWorkspaces3Default {
	return &WaypointListWorkspaces3Default{
		_statusCode: code,
	}
}

/*
WaypointListWorkspaces3Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListWorkspaces3Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list workspaces3 default response
func (o *WaypointListWorkspaces3Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint list workspaces3 default response has a 2xx status code
func (o *WaypointListWorkspaces3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list workspaces3 default response has a 3xx status code
func (o *WaypointListWorkspaces3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list workspaces3 default response has a 4xx status code
func (o *WaypointListWorkspaces3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list workspaces3 default response has a 5xx status code
func (o *WaypointListWorkspaces3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list workspaces3 default response a status code equal to that given
func (o *WaypointListWorkspaces3Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointListWorkspaces3Default) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspaces][%d] Waypoint_ListWorkspaces3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListWorkspaces3Default) String() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspaces][%d] Waypoint_ListWorkspaces3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListWorkspaces3Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListWorkspaces3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
