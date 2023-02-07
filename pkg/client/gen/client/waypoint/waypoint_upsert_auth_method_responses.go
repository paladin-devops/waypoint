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

// WaypointUpsertAuthMethodReader is a Reader for the WaypointUpsertAuthMethod structure.
type WaypointUpsertAuthMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUpsertAuthMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUpsertAuthMethodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUpsertAuthMethodDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUpsertAuthMethodOK creates a WaypointUpsertAuthMethodOK with default headers values
func NewWaypointUpsertAuthMethodOK() *WaypointUpsertAuthMethodOK {
	return &WaypointUpsertAuthMethodOK{}
}

/*
WaypointUpsertAuthMethodOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointUpsertAuthMethodOK struct {
	Payload *models.HashicorpWaypointUpsertAuthMethodResponse
}

// IsSuccess returns true when this waypoint upsert auth method o k response has a 2xx status code
func (o *WaypointUpsertAuthMethodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint upsert auth method o k response has a 3xx status code
func (o *WaypointUpsertAuthMethodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint upsert auth method o k response has a 4xx status code
func (o *WaypointUpsertAuthMethodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint upsert auth method o k response has a 5xx status code
func (o *WaypointUpsertAuthMethodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint upsert auth method o k response a status code equal to that given
func (o *WaypointUpsertAuthMethodOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointUpsertAuthMethodOK) Error() string {
	return fmt.Sprintf("[POST /auth-method][%d] waypointUpsertAuthMethodOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertAuthMethodOK) String() string {
	return fmt.Sprintf("[POST /auth-method][%d] waypointUpsertAuthMethodOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertAuthMethodOK) GetPayload() *models.HashicorpWaypointUpsertAuthMethodResponse {
	return o.Payload
}

func (o *WaypointUpsertAuthMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUpsertAuthMethodResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUpsertAuthMethodDefault creates a WaypointUpsertAuthMethodDefault with default headers values
func NewWaypointUpsertAuthMethodDefault(code int) *WaypointUpsertAuthMethodDefault {
	return &WaypointUpsertAuthMethodDefault{
		_statusCode: code,
	}
}

/*
WaypointUpsertAuthMethodDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointUpsertAuthMethodDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint upsert auth method default response
func (o *WaypointUpsertAuthMethodDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint upsert auth method default response has a 2xx status code
func (o *WaypointUpsertAuthMethodDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint upsert auth method default response has a 3xx status code
func (o *WaypointUpsertAuthMethodDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint upsert auth method default response has a 4xx status code
func (o *WaypointUpsertAuthMethodDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint upsert auth method default response has a 5xx status code
func (o *WaypointUpsertAuthMethodDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint upsert auth method default response a status code equal to that given
func (o *WaypointUpsertAuthMethodDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointUpsertAuthMethodDefault) Error() string {
	return fmt.Sprintf("[POST /auth-method][%d] Waypoint_UpsertAuthMethod default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertAuthMethodDefault) String() string {
	return fmt.Sprintf("[POST /auth-method][%d] Waypoint_UpsertAuthMethod default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertAuthMethodDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUpsertAuthMethodDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
