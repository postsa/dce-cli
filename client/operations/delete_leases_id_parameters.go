// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteLeasesIDParams creates a new DeleteLeasesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLeasesIDParams() *DeleteLeasesIDParams {
	return &DeleteLeasesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLeasesIDParamsWithTimeout creates a new DeleteLeasesIDParams object
// with the ability to set a timeout on a request.
func NewDeleteLeasesIDParamsWithTimeout(timeout time.Duration) *DeleteLeasesIDParams {
	return &DeleteLeasesIDParams{
		timeout: timeout,
	}
}

// NewDeleteLeasesIDParamsWithContext creates a new DeleteLeasesIDParams object
// with the ability to set a context for a request.
func NewDeleteLeasesIDParamsWithContext(ctx context.Context) *DeleteLeasesIDParams {
	return &DeleteLeasesIDParams{
		Context: ctx,
	}
}

// NewDeleteLeasesIDParamsWithHTTPClient creates a new DeleteLeasesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLeasesIDParamsWithHTTPClient(client *http.Client) *DeleteLeasesIDParams {
	return &DeleteLeasesIDParams{
		HTTPClient: client,
	}
}

/*
DeleteLeasesIDParams contains all the parameters to send to the API endpoint

	for the delete leases ID operation.

	Typically these are written to a http.Request.
*/
type DeleteLeasesIDParams struct {

	/* ID.

	   The ID of the lease to be deleted.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete leases ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLeasesIDParams) WithDefaults() *DeleteLeasesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete leases ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLeasesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete leases ID params
func (o *DeleteLeasesIDParams) WithTimeout(timeout time.Duration) *DeleteLeasesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete leases ID params
func (o *DeleteLeasesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete leases ID params
func (o *DeleteLeasesIDParams) WithContext(ctx context.Context) *DeleteLeasesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete leases ID params
func (o *DeleteLeasesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete leases ID params
func (o *DeleteLeasesIDParams) WithHTTPClient(client *http.Client) *DeleteLeasesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete leases ID params
func (o *DeleteLeasesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete leases ID params
func (o *DeleteLeasesIDParams) WithID(id string) *DeleteLeasesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete leases ID params
func (o *DeleteLeasesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLeasesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
