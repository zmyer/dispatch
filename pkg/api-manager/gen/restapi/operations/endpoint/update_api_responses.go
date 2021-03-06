///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api-manager/gen/models"
)

// UpdateAPIOKCode is the HTTP code returned for type UpdateAPIOK
const UpdateAPIOKCode int = 200

/*UpdateAPIOK Successful update

swagger:response updateApiOK
*/
type UpdateAPIOK struct {

	/*
	  In: Body
	*/
	Payload *models.API `json:"body,omitempty"`
}

// NewUpdateAPIOK creates UpdateAPIOK with default headers values
func NewUpdateAPIOK() *UpdateAPIOK {
	return &UpdateAPIOK{}
}

// WithPayload adds the payload to the update Api o k response
func (o *UpdateAPIOK) WithPayload(payload *models.API) *UpdateAPIOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Api o k response
func (o *UpdateAPIOK) SetPayload(payload *models.API) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAPIOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAPIBadRequestCode is the HTTP code returned for type UpdateAPIBadRequest
const UpdateAPIBadRequestCode int = 400

/*UpdateAPIBadRequest Invalid input

swagger:response updateApiBadRequest
*/
type UpdateAPIBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateAPIBadRequest creates UpdateAPIBadRequest with default headers values
func NewUpdateAPIBadRequest() *UpdateAPIBadRequest {
	return &UpdateAPIBadRequest{}
}

// WithPayload adds the payload to the update Api bad request response
func (o *UpdateAPIBadRequest) WithPayload(payload *models.Error) *UpdateAPIBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Api bad request response
func (o *UpdateAPIBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAPIBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAPINotFoundCode is the HTTP code returned for type UpdateAPINotFound
const UpdateAPINotFoundCode int = 404

/*UpdateAPINotFound API not found

swagger:response updateApiNotFound
*/
type UpdateAPINotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateAPINotFound creates UpdateAPINotFound with default headers values
func NewUpdateAPINotFound() *UpdateAPINotFound {
	return &UpdateAPINotFound{}
}

// WithPayload adds the payload to the update Api not found response
func (o *UpdateAPINotFound) WithPayload(payload *models.Error) *UpdateAPINotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Api not found response
func (o *UpdateAPINotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAPINotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAPIInternalServerErrorCode is the HTTP code returned for type UpdateAPIInternalServerError
const UpdateAPIInternalServerErrorCode int = 500

/*UpdateAPIInternalServerError Internal error

swagger:response updateApiInternalServerError
*/
type UpdateAPIInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateAPIInternalServerError creates UpdateAPIInternalServerError with default headers values
func NewUpdateAPIInternalServerError() *UpdateAPIInternalServerError {
	return &UpdateAPIInternalServerError{}
}

// WithPayload adds the payload to the update Api internal server error response
func (o *UpdateAPIInternalServerError) WithPayload(payload *models.Error) *UpdateAPIInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Api internal server error response
func (o *UpdateAPIInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAPIInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
