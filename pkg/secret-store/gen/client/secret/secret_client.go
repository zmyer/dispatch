///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new secret API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for secret API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddSecret add secret API
*/
func (a *Client) AddSecret(params *AddSecretParams, authInfo runtime.ClientAuthInfoWriter) (*AddSecretCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSecret",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddSecretCreated), nil

}

/*
DeleteSecret delete secret API
*/
func (a *Client) DeleteSecret(params *DeleteSecretParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSecretNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSecret",
		Method:             "DELETE",
		PathPattern:        "/{secretName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSecretNoContent), nil

}

/*
GetSecret get secret API
*/
func (a *Client) GetSecret(params *GetSecretParams, authInfo runtime.ClientAuthInfoWriter) (*GetSecretOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSecret",
		Method:             "GET",
		PathPattern:        "/{secretName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSecretOK), nil

}

/*
GetSecrets get secrets API
*/
func (a *Client) GetSecrets(params *GetSecretsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSecretsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecretsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSecrets",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSecretsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSecretsOK), nil

}

/*
UpdateSecret update secret API
*/
func (a *Client) UpdateSecret(params *UpdateSecretParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSecretCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSecret",
		Method:             "PUT",
		PathPattern:        "/{secretName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSecretCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
