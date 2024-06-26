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

// WaypointListReleases2Reader is a Reader for the WaypointListReleases2 structure.
type WaypointListReleases2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListReleases2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListReleases2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListReleases2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListReleases2OK creates a WaypointListReleases2OK with default headers values
func NewWaypointListReleases2OK() *WaypointListReleases2OK {
	return &WaypointListReleases2OK{}
}

/*
WaypointListReleases2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListReleases2OK struct {
	Payload *models.HashicorpWaypointListReleasesResponse
}

// IsSuccess returns true when this waypoint list releases2 o k response has a 2xx status code
func (o *WaypointListReleases2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list releases2 o k response has a 3xx status code
func (o *WaypointListReleases2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list releases2 o k response has a 4xx status code
func (o *WaypointListReleases2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list releases2 o k response has a 5xx status code
func (o *WaypointListReleases2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list releases2 o k response a status code equal to that given
func (o *WaypointListReleases2OK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointListReleases2OK) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/releases][%d] waypointListReleases2OK  %+v", 200, o.Payload)
}

func (o *WaypointListReleases2OK) String() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/releases][%d] waypointListReleases2OK  %+v", 200, o.Payload)
}

func (o *WaypointListReleases2OK) GetPayload() *models.HashicorpWaypointListReleasesResponse {
	return o.Payload
}

func (o *WaypointListReleases2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListReleasesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListReleases2Default creates a WaypointListReleases2Default with default headers values
func NewWaypointListReleases2Default(code int) *WaypointListReleases2Default {
	return &WaypointListReleases2Default{
		_statusCode: code,
	}
}

/*
WaypointListReleases2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListReleases2Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list releases2 default response
func (o *WaypointListReleases2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint list releases2 default response has a 2xx status code
func (o *WaypointListReleases2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list releases2 default response has a 3xx status code
func (o *WaypointListReleases2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list releases2 default response has a 4xx status code
func (o *WaypointListReleases2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list releases2 default response has a 5xx status code
func (o *WaypointListReleases2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list releases2 default response a status code equal to that given
func (o *WaypointListReleases2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointListReleases2Default) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/releases][%d] Waypoint_ListReleases2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListReleases2Default) String() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/releases][%d] Waypoint_ListReleases2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListReleases2Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListReleases2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
