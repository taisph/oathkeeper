// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new health API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for health API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IsInstanceAlive(params *IsInstanceAliveParams, opts ...ClientOption) (*IsInstanceAliveOK, error)

	IsInstanceReady(params *IsInstanceReadyParams, opts ...ClientOption) (*IsInstanceReadyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	IsInstanceAlive checks alive status

	This endpoint returns a 200 status code when the HTTP server is up running.

This status does currently not include checks whether the database connection is working.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
*/
func (a *Client) IsInstanceAlive(params *IsInstanceAliveParams, opts ...ClientOption) (*IsInstanceAliveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsInstanceAliveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "isInstanceAlive",
		Method:             "GET",
		PathPattern:        "/health/alive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsInstanceAliveReader{formats: a.formats},
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
	success, ok := result.(*IsInstanceAliveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for isInstanceAlive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	IsInstanceReady checks readiness status

	This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g.

the database) are responsive as well.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
*/
func (a *Client) IsInstanceReady(params *IsInstanceReadyParams, opts ...ClientOption) (*IsInstanceReadyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsInstanceReadyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "isInstanceReady",
		Method:             "GET",
		PathPattern:        "/health/ready",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsInstanceReadyReader{formats: a.formats},
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
	success, ok := result.(*IsInstanceReadyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for isInstanceReady: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}