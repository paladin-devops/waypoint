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

// WaypointExpediteStatusReport3Reader is a Reader for the WaypointExpediteStatusReport3 structure.
type WaypointExpediteStatusReport3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointExpediteStatusReport3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointExpediteStatusReport3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointExpediteStatusReport3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointExpediteStatusReport3OK creates a WaypointExpediteStatusReport3OK with default headers values
func NewWaypointExpediteStatusReport3OK() *WaypointExpediteStatusReport3OK {
	return &WaypointExpediteStatusReport3OK{}
}

/*
WaypointExpediteStatusReport3OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointExpediteStatusReport3OK struct {
	Payload *models.HashicorpWaypointExpediteStatusReportResponse
}

// IsSuccess returns true when this waypoint expedite status report3 o k response has a 2xx status code
func (o *WaypointExpediteStatusReport3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint expedite status report3 o k response has a 3xx status code
func (o *WaypointExpediteStatusReport3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint expedite status report3 o k response has a 4xx status code
func (o *WaypointExpediteStatusReport3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint expedite status report3 o k response has a 5xx status code
func (o *WaypointExpediteStatusReport3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint expedite status report3 o k response a status code equal to that given
func (o *WaypointExpediteStatusReport3OK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointExpediteStatusReport3OK) Error() string {
	return fmt.Sprintf("[PUT /project/{deployment.sequence.application.project}/application/{deployment.sequence.application.application}/deployment/{deployment.sequence.number}/status-report][%d] waypointExpediteStatusReport3OK  %+v", 200, o.Payload)
}

func (o *WaypointExpediteStatusReport3OK) String() string {
	return fmt.Sprintf("[PUT /project/{deployment.sequence.application.project}/application/{deployment.sequence.application.application}/deployment/{deployment.sequence.number}/status-report][%d] waypointExpediteStatusReport3OK  %+v", 200, o.Payload)
}

func (o *WaypointExpediteStatusReport3OK) GetPayload() *models.HashicorpWaypointExpediteStatusReportResponse {
	return o.Payload
}

func (o *WaypointExpediteStatusReport3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointExpediteStatusReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointExpediteStatusReport3Default creates a WaypointExpediteStatusReport3Default with default headers values
func NewWaypointExpediteStatusReport3Default(code int) *WaypointExpediteStatusReport3Default {
	return &WaypointExpediteStatusReport3Default{
		_statusCode: code,
	}
}

/*
WaypointExpediteStatusReport3Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointExpediteStatusReport3Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint expedite status report3 default response
func (o *WaypointExpediteStatusReport3Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint expedite status report3 default response has a 2xx status code
func (o *WaypointExpediteStatusReport3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint expedite status report3 default response has a 3xx status code
func (o *WaypointExpediteStatusReport3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint expedite status report3 default response has a 4xx status code
func (o *WaypointExpediteStatusReport3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint expedite status report3 default response has a 5xx status code
func (o *WaypointExpediteStatusReport3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint expedite status report3 default response a status code equal to that given
func (o *WaypointExpediteStatusReport3Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointExpediteStatusReport3Default) Error() string {
	return fmt.Sprintf("[PUT /project/{deployment.sequence.application.project}/application/{deployment.sequence.application.application}/deployment/{deployment.sequence.number}/status-report][%d] Waypoint_ExpediteStatusReport3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointExpediteStatusReport3Default) String() string {
	return fmt.Sprintf("[PUT /project/{deployment.sequence.application.project}/application/{deployment.sequence.application.application}/deployment/{deployment.sequence.number}/status-report][%d] Waypoint_ExpediteStatusReport3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointExpediteStatusReport3Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointExpediteStatusReport3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
