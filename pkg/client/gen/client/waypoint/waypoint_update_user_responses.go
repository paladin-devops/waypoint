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

// WaypointUpdateUserReader is a Reader for the WaypointUpdateUser structure.
type WaypointUpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUpdateUserOK creates a WaypointUpdateUserOK with default headers values
func NewWaypointUpdateUserOK() *WaypointUpdateUserOK {
	return &WaypointUpdateUserOK{}
}

/*
WaypointUpdateUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointUpdateUserOK struct {
	Payload *models.HashicorpWaypointUpdateUserResponse
}

// IsSuccess returns true when this waypoint update user o k response has a 2xx status code
func (o *WaypointUpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint update user o k response has a 3xx status code
func (o *WaypointUpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint update user o k response has a 4xx status code
func (o *WaypointUpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint update user o k response has a 5xx status code
func (o *WaypointUpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint update user o k response a status code equal to that given
func (o *WaypointUpdateUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointUpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /user/{user.id}][%d] waypointUpdateUserOK  %+v", 200, o.Payload)
}

func (o *WaypointUpdateUserOK) String() string {
	return fmt.Sprintf("[PUT /user/{user.id}][%d] waypointUpdateUserOK  %+v", 200, o.Payload)
}

func (o *WaypointUpdateUserOK) GetPayload() *models.HashicorpWaypointUpdateUserResponse {
	return o.Payload
}

func (o *WaypointUpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUpdateUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUpdateUserDefault creates a WaypointUpdateUserDefault with default headers values
func NewWaypointUpdateUserDefault(code int) *WaypointUpdateUserDefault {
	return &WaypointUpdateUserDefault{
		_statusCode: code,
	}
}

/*
WaypointUpdateUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointUpdateUserDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint update user default response
func (o *WaypointUpdateUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint update user default response has a 2xx status code
func (o *WaypointUpdateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint update user default response has a 3xx status code
func (o *WaypointUpdateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint update user default response has a 4xx status code
func (o *WaypointUpdateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint update user default response has a 5xx status code
func (o *WaypointUpdateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint update user default response a status code equal to that given
func (o *WaypointUpdateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointUpdateUserDefault) Error() string {
	return fmt.Sprintf("[PUT /user/{user.id}][%d] Waypoint_UpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpdateUserDefault) String() string {
	return fmt.Sprintf("[PUT /user/{user.id}][%d] Waypoint_UpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpdateUserDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
