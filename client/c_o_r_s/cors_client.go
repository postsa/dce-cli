// Code generated by go-swagger; DO NOT EDIT.

package c_o_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new c o r s API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new c o r s API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new c o r s API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for c o r s API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	OptionsAccounts(params *OptionsAccountsParams, opts ...ClientOption) (*OptionsAccountsOK, error)

	OptionsAccountsID(params *OptionsAccountsIDParams, opts ...ClientOption) (*OptionsAccountsIDOK, error)

	OptionsAuth(params *OptionsAuthParams, opts ...ClientOption) (*OptionsAuthOK, error)

	OptionsAuthFile(params *OptionsAuthFileParams, opts ...ClientOption) (*OptionsAuthFileOK, error)

	OptionsLeases(params *OptionsLeasesParams, opts ...ClientOption) (*OptionsLeasesOK, error)

	OptionsLeasesID(params *OptionsLeasesIDParams, opts ...ClientOption) (*OptionsLeasesIDOK, error)

	OptionsLeasesIDAuth(params *OptionsLeasesIDAuthParams, opts ...ClientOption) (*OptionsLeasesIDAuthOK, error)

	OptionsUsage(params *OptionsUsageParams, opts ...ClientOption) (*OptionsUsageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
OptionsAccounts cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsAccounts(params *OptionsAccountsParams, opts ...ClientOption) (*OptionsAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsAccounts",
		Method:             "OPTIONS",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsAccountsID cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsAccountsID(params *OptionsAccountsIDParams, opts ...ClientOption) (*OptionsAccountsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsAccountsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsAccountsID",
		Method:             "OPTIONS",
		PathPattern:        "/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsAccountsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsAccountsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsAccountsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsAuth cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsAuth(params *OptionsAuthParams, opts ...ClientOption) (*OptionsAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsAuth",
		Method:             "OPTIONS",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsAuth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsAuthFile cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsAuthFile(params *OptionsAuthFileParams, opts ...ClientOption) (*OptionsAuthFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsAuthFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsAuthFile",
		Method:             "OPTIONS",
		PathPattern:        "/auth/{file+}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsAuthFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsAuthFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsAuthFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsLeases cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsLeases(params *OptionsLeasesParams, opts ...ClientOption) (*OptionsLeasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsLeasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsLeases",
		Method:             "OPTIONS",
		PathPattern:        "/leases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsLeasesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsLeasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsLeases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsLeasesID cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsLeasesID(params *OptionsLeasesIDParams, opts ...ClientOption) (*OptionsLeasesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsLeasesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsLeasesID",
		Method:             "OPTIONS",
		PathPattern:        "/leases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsLeasesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsLeasesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsLeasesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsLeasesIDAuth cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsLeasesIDAuth(params *OptionsLeasesIDAuthParams, opts ...ClientOption) (*OptionsLeasesIDAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsLeasesIDAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsLeasesIDAuth",
		Method:             "OPTIONS",
		PathPattern:        "/leases/{id}/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsLeasesIDAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsLeasesIDAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsLeasesIDAuth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsUsage cs o r s support

Enable CORS by returning correct headers
*/
func (a *Client) OptionsUsage(params *OptionsUsageParams, opts ...ClientOption) (*OptionsUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsUsage",
		Method:             "OPTIONS",
		PathPattern:        "/usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsUsageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsUsage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
