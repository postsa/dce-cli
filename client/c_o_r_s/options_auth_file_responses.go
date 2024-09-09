// Code generated by go-swagger; DO NOT EDIT.

package c_o_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OptionsAuthFileReader is a Reader for the OptionsAuthFile structure.
type OptionsAuthFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OptionsAuthFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOptionsAuthFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[OPTIONS /auth/{file+}] OptionsAuthFile", response, response.Code())
	}
}

// NewOptionsAuthFileOK creates a OptionsAuthFileOK with default headers values
func NewOptionsAuthFileOK() *OptionsAuthFileOK {
	return &OptionsAuthFileOK{}
}

/*
OptionsAuthFileOK describes a response with status code 200, with default header values.

Default response for CORS method
*/
type OptionsAuthFileOK struct {
	AccessControlAllowHeaders string
	AccessControlAllowMethods string
	AccessControlAllowOrigin  string
}

// IsSuccess returns true when this options auth file o k response has a 2xx status code
func (o *OptionsAuthFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this options auth file o k response has a 3xx status code
func (o *OptionsAuthFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this options auth file o k response has a 4xx status code
func (o *OptionsAuthFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this options auth file o k response has a 5xx status code
func (o *OptionsAuthFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this options auth file o k response a status code equal to that given
func (o *OptionsAuthFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the options auth file o k response
func (o *OptionsAuthFileOK) Code() int {
	return 200
}

func (o *OptionsAuthFileOK) Error() string {
	return fmt.Sprintf("[OPTIONS /auth/{file+}][%d] optionsAuthFileOK", 200)
}

func (o *OptionsAuthFileOK) String() string {
	return fmt.Sprintf("[OPTIONS /auth/{file+}][%d] optionsAuthFileOK", 200)
}

func (o *OptionsAuthFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}
