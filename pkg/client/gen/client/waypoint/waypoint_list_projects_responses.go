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

// WaypointListProjectsReader is a Reader for the WaypointListProjects structure.
type WaypointListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListProjectsOK creates a WaypointListProjectsOK with default headers values
func NewWaypointListProjectsOK() *WaypointListProjectsOK {
	return &WaypointListProjectsOK{}
}

/*
WaypointListProjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListProjectsOK struct {
	Payload *models.HashicorpWaypointListProjectsResponse
}

// IsSuccess returns true when this waypoint list projects o k response has a 2xx status code
func (o *WaypointListProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list projects o k response has a 3xx status code
func (o *WaypointListProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list projects o k response has a 4xx status code
func (o *WaypointListProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list projects o k response has a 5xx status code
func (o *WaypointListProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list projects o k response a status code equal to that given
func (o *WaypointListProjectsOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /projects][%d] waypointListProjectsOK  %+v", 200, o.Payload)
}

func (o *WaypointListProjectsOK) String() string {
	return fmt.Sprintf("[GET /projects][%d] waypointListProjectsOK  %+v", 200, o.Payload)
}

func (o *WaypointListProjectsOK) GetPayload() *models.HashicorpWaypointListProjectsResponse {
	return o.Payload
}

func (o *WaypointListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListProjectsDefault creates a WaypointListProjectsDefault with default headers values
func NewWaypointListProjectsDefault(code int) *WaypointListProjectsDefault {
	return &WaypointListProjectsDefault{
		_statusCode: code,
	}
}

/*
WaypointListProjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListProjectsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list projects default response
func (o *WaypointListProjectsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint list projects default response has a 2xx status code
func (o *WaypointListProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list projects default response has a 3xx status code
func (o *WaypointListProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list projects default response has a 4xx status code
func (o *WaypointListProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list projects default response has a 5xx status code
func (o *WaypointListProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list projects default response a status code equal to that given
func (o *WaypointListProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointListProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /projects][%d] Waypoint_ListProjects default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListProjectsDefault) String() string {
	return fmt.Sprintf("[GET /projects][%d] Waypoint_ListProjects default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListProjectsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
