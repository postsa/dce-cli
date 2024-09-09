// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAuthReader is a Reader for the GetAuth structure.
type GetAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /auth] GetAuth", response, response.Code())
	}
}

// NewGetAuthOK creates a GetAuthOK with default headers values
func NewGetAuthOK() *GetAuthOK {
	return &GetAuthOK{}
}

/*
GetAuthOK describes a response with status code 200, with default header values.

GetAuthOK get auth o k
*/
type GetAuthOK struct {
	AccessControlAllowHeaders string
	AccessControlAllowMethods string
	AccessControlAllowOrigin  string

	Payload interface{}
}

// IsSuccess returns true when this get auth o k response has a 2xx status code
func (o *GetAuthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get auth o k response has a 3xx status code
func (o *GetAuthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get auth o k response has a 4xx status code
func (o *GetAuthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth o k response has a 5xx status code
func (o *GetAuthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get auth o k response a status code equal to that given
func (o *GetAuthOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get auth o k response
func (o *GetAuthOK) Code() int {
	return 200
}

func (o *GetAuthOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth][%d] getAuthOK %s", 200, payload)
}

func (o *GetAuthOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth][%d] getAuthOK %s", 200, payload)
}

func (o *GetAuthOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Access-Control-Allow-Headers
	hdrAccessControlAllowHeaders := response.GetHeader("Access-Control-Allow-Headers")

	if hdrAccessControlAllowHeaders != "" {
		o.AccessControlAllowHeaders = hdrAccessControlAllowHeaders
	}

	// hydrates response header Access-Control-Allow-Methods
	hdrAccessControlAllowMethods := response.GetHeader("Access-Control-Allow-Methods")

	if hdrAccessControlAllowMethods != "" {
		o.AccessControlAllowMethods = hdrAccessControlAllowMethods
	}

	// hydrates response header Access-Control-Allow-Origin
	hdrAccessControlAllowOrigin := response.GetHeader("Access-Control-Allow-Origin")

	if hdrAccessControlAllowOrigin != "" {
		o.AccessControlAllowOrigin = hdrAccessControlAllowOrigin
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
