// Code generated by go-swagger; DO NOT EDIT.

package c_o_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOptionsUsageParams creates a new OptionsUsageParams object
// with the default values initialized.
func NewOptionsUsageParams() *OptionsUsageParams {

	return &OptionsUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOptionsUsageParamsWithTimeout creates a new OptionsUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOptionsUsageParamsWithTimeout(timeout time.Duration) *OptionsUsageParams {

	return &OptionsUsageParams{

		timeout: timeout,
	}
}

// NewOptionsUsageParamsWithContext creates a new OptionsUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewOptionsUsageParamsWithContext(ctx context.Context) *OptionsUsageParams {

	return &OptionsUsageParams{

		Context: ctx,
	}
}

// NewOptionsUsageParamsWithHTTPClient creates a new OptionsUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOptionsUsageParamsWithHTTPClient(client *http.Client) *OptionsUsageParams {

	return &OptionsUsageParams{
		HTTPClient: client,
	}
}

/*
OptionsUsageParams contains all the parameters to send to the API endpoint
for the options usage operation typically these are written to a http.Request
*/
type OptionsUsageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the options usage params
func (o *OptionsUsageParams) WithTimeout(timeout time.Duration) *OptionsUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the options usage params
func (o *OptionsUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the options usage params
func (o *OptionsUsageParams) WithContext(ctx context.Context) *OptionsUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the options usage params
func (o *OptionsUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the options usage params
func (o *OptionsUsageParams) WithHTTPClient(client *http.Client) *OptionsUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the options usage params
func (o *OptionsUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OptionsUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
