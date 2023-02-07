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

// WaypointListInstances3Reader is a Reader for the WaypointListInstances3 structure.
type WaypointListInstances3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListInstances3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListInstances3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListInstances3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListInstances3OK creates a WaypointListInstances3OK with default headers values
func NewWaypointListInstances3OK() *WaypointListInstances3OK {
	return &WaypointListInstances3OK{}
}

/*
WaypointListInstances3OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListInstances3OK struct {
	Payload *models.HashicorpWaypointListInstancesResponse
}

// IsSuccess returns true when this waypoint list instances3 o k response has a 2xx status code
func (o *WaypointListInstances3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list instances3 o k response has a 3xx status code
func (o *WaypointListInstances3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list instances3 o k response has a 4xx status code
func (o *WaypointListInstances3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list instances3 o k response has a 5xx status code
func (o *WaypointListInstances3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list instances3 o k response a status code equal to that given
func (o *WaypointListInstances3OK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointListInstances3OK) Error() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/workspace/{application.workspace.workspace}/instances][%d] waypointListInstances3OK  %+v", 200, o.Payload)
}

func (o *WaypointListInstances3OK) String() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/workspace/{application.workspace.workspace}/instances][%d] waypointListInstances3OK  %+v", 200, o.Payload)
}

func (o *WaypointListInstances3OK) GetPayload() *models.HashicorpWaypointListInstancesResponse {
	return o.Payload
}

func (o *WaypointListInstances3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListInstances3Default creates a WaypointListInstances3Default with default headers values
func NewWaypointListInstances3Default(code int) *WaypointListInstances3Default {
	return &WaypointListInstances3Default{
		_statusCode: code,
	}
}

/*
WaypointListInstances3Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListInstances3Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list instances3 default response
func (o *WaypointListInstances3Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint list instances3 default response has a 2xx status code
func (o *WaypointListInstances3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list instances3 default response has a 3xx status code
func (o *WaypointListInstances3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list instances3 default response has a 4xx status code
func (o *WaypointListInstances3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list instances3 default response has a 5xx status code
func (o *WaypointListInstances3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list instances3 default response a status code equal to that given
func (o *WaypointListInstances3Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointListInstances3Default) Error() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/workspace/{application.workspace.workspace}/instances][%d] Waypoint_ListInstances3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListInstances3Default) String() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/workspace/{application.workspace.workspace}/instances][%d] Waypoint_ListInstances3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListInstances3Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListInstances3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
