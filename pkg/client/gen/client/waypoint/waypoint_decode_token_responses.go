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

// WaypointDecodeTokenReader is a Reader for the WaypointDecodeToken structure.
type WaypointDecodeTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointDecodeTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointDecodeTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointDecodeTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointDecodeTokenOK creates a WaypointDecodeTokenOK with default headers values
func NewWaypointDecodeTokenOK() *WaypointDecodeTokenOK {
	return &WaypointDecodeTokenOK{}
}

/*
WaypointDecodeTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointDecodeTokenOK struct {
	Payload *models.HashicorpWaypointDecodeTokenResponse
}

// IsSuccess returns true when this waypoint decode token o k response has a 2xx status code
func (o *WaypointDecodeTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint decode token o k response has a 3xx status code
func (o *WaypointDecodeTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint decode token o k response has a 4xx status code
func (o *WaypointDecodeTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint decode token o k response has a 5xx status code
func (o *WaypointDecodeTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint decode token o k response a status code equal to that given
func (o *WaypointDecodeTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointDecodeTokenOK) Error() string {
	return fmt.Sprintf("[POST /token/decode][%d] waypointDecodeTokenOK  %+v", 200, o.Payload)
}

func (o *WaypointDecodeTokenOK) String() string {
	return fmt.Sprintf("[POST /token/decode][%d] waypointDecodeTokenOK  %+v", 200, o.Payload)
}

func (o *WaypointDecodeTokenOK) GetPayload() *models.HashicorpWaypointDecodeTokenResponse {
	return o.Payload
}

func (o *WaypointDecodeTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointDecodeTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointDecodeTokenDefault creates a WaypointDecodeTokenDefault with default headers values
func NewWaypointDecodeTokenDefault(code int) *WaypointDecodeTokenDefault {
	return &WaypointDecodeTokenDefault{
		_statusCode: code,
	}
}

/*
WaypointDecodeTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointDecodeTokenDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint decode token default response
func (o *WaypointDecodeTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint decode token default response has a 2xx status code
func (o *WaypointDecodeTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint decode token default response has a 3xx status code
func (o *WaypointDecodeTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint decode token default response has a 4xx status code
func (o *WaypointDecodeTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint decode token default response has a 5xx status code
func (o *WaypointDecodeTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint decode token default response a status code equal to that given
func (o *WaypointDecodeTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointDecodeTokenDefault) Error() string {
	return fmt.Sprintf("[POST /token/decode][%d] Waypoint_DecodeToken default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointDecodeTokenDefault) String() string {
	return fmt.Sprintf("[POST /token/decode][%d] Waypoint_DecodeToken default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointDecodeTokenDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointDecodeTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
