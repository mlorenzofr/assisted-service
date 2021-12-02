// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// RegenerateInfraEnvSigningKeyNoContentCode is the HTTP code returned for type RegenerateInfraEnvSigningKeyNoContent
const RegenerateInfraEnvSigningKeyNoContentCode int = 204

/*RegenerateInfraEnvSigningKeyNoContent Success.

swagger:response regenerateInfraEnvSigningKeyNoContent
*/
type RegenerateInfraEnvSigningKeyNoContent struct {
}

// NewRegenerateInfraEnvSigningKeyNoContent creates RegenerateInfraEnvSigningKeyNoContent with default headers values
func NewRegenerateInfraEnvSigningKeyNoContent() *RegenerateInfraEnvSigningKeyNoContent {

	return &RegenerateInfraEnvSigningKeyNoContent{}
}

// WriteResponse to the client
func (o *RegenerateInfraEnvSigningKeyNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// RegenerateInfraEnvSigningKeyUnauthorizedCode is the HTTP code returned for type RegenerateInfraEnvSigningKeyUnauthorized
const RegenerateInfraEnvSigningKeyUnauthorizedCode int = 401

/*RegenerateInfraEnvSigningKeyUnauthorized Unauthorized.

swagger:response regenerateInfraEnvSigningKeyUnauthorized
*/
type RegenerateInfraEnvSigningKeyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewRegenerateInfraEnvSigningKeyUnauthorized creates RegenerateInfraEnvSigningKeyUnauthorized with default headers values
func NewRegenerateInfraEnvSigningKeyUnauthorized() *RegenerateInfraEnvSigningKeyUnauthorized {

	return &RegenerateInfraEnvSigningKeyUnauthorized{}
}

// WithPayload adds the payload to the regenerate infra env signing key unauthorized response
func (o *RegenerateInfraEnvSigningKeyUnauthorized) WithPayload(payload *models.InfraError) *RegenerateInfraEnvSigningKeyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the regenerate infra env signing key unauthorized response
func (o *RegenerateInfraEnvSigningKeyUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegenerateInfraEnvSigningKeyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegenerateInfraEnvSigningKeyForbiddenCode is the HTTP code returned for type RegenerateInfraEnvSigningKeyForbidden
const RegenerateInfraEnvSigningKeyForbiddenCode int = 403

/*RegenerateInfraEnvSigningKeyForbidden Forbidden.

swagger:response regenerateInfraEnvSigningKeyForbidden
*/
type RegenerateInfraEnvSigningKeyForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewRegenerateInfraEnvSigningKeyForbidden creates RegenerateInfraEnvSigningKeyForbidden with default headers values
func NewRegenerateInfraEnvSigningKeyForbidden() *RegenerateInfraEnvSigningKeyForbidden {

	return &RegenerateInfraEnvSigningKeyForbidden{}
}

// WithPayload adds the payload to the regenerate infra env signing key forbidden response
func (o *RegenerateInfraEnvSigningKeyForbidden) WithPayload(payload *models.InfraError) *RegenerateInfraEnvSigningKeyForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the regenerate infra env signing key forbidden response
func (o *RegenerateInfraEnvSigningKeyForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegenerateInfraEnvSigningKeyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegenerateInfraEnvSigningKeyNotFoundCode is the HTTP code returned for type RegenerateInfraEnvSigningKeyNotFound
const RegenerateInfraEnvSigningKeyNotFoundCode int = 404

/*RegenerateInfraEnvSigningKeyNotFound Error.

swagger:response regenerateInfraEnvSigningKeyNotFound
*/
type RegenerateInfraEnvSigningKeyNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegenerateInfraEnvSigningKeyNotFound creates RegenerateInfraEnvSigningKeyNotFound with default headers values
func NewRegenerateInfraEnvSigningKeyNotFound() *RegenerateInfraEnvSigningKeyNotFound {

	return &RegenerateInfraEnvSigningKeyNotFound{}
}

// WithPayload adds the payload to the regenerate infra env signing key not found response
func (o *RegenerateInfraEnvSigningKeyNotFound) WithPayload(payload *models.Error) *RegenerateInfraEnvSigningKeyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the regenerate infra env signing key not found response
func (o *RegenerateInfraEnvSigningKeyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegenerateInfraEnvSigningKeyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegenerateInfraEnvSigningKeyMethodNotAllowedCode is the HTTP code returned for type RegenerateInfraEnvSigningKeyMethodNotAllowed
const RegenerateInfraEnvSigningKeyMethodNotAllowedCode int = 405

/*RegenerateInfraEnvSigningKeyMethodNotAllowed Method Not Allowed.

swagger:response regenerateInfraEnvSigningKeyMethodNotAllowed
*/
type RegenerateInfraEnvSigningKeyMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegenerateInfraEnvSigningKeyMethodNotAllowed creates RegenerateInfraEnvSigningKeyMethodNotAllowed with default headers values
func NewRegenerateInfraEnvSigningKeyMethodNotAllowed() *RegenerateInfraEnvSigningKeyMethodNotAllowed {

	return &RegenerateInfraEnvSigningKeyMethodNotAllowed{}
}

// WithPayload adds the payload to the regenerate infra env signing key method not allowed response
func (o *RegenerateInfraEnvSigningKeyMethodNotAllowed) WithPayload(payload *models.Error) *RegenerateInfraEnvSigningKeyMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the regenerate infra env signing key method not allowed response
func (o *RegenerateInfraEnvSigningKeyMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegenerateInfraEnvSigningKeyMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegenerateInfraEnvSigningKeyInternalServerErrorCode is the HTTP code returned for type RegenerateInfraEnvSigningKeyInternalServerError
const RegenerateInfraEnvSigningKeyInternalServerErrorCode int = 500

/*RegenerateInfraEnvSigningKeyInternalServerError Error.

swagger:response regenerateInfraEnvSigningKeyInternalServerError
*/
type RegenerateInfraEnvSigningKeyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegenerateInfraEnvSigningKeyInternalServerError creates RegenerateInfraEnvSigningKeyInternalServerError with default headers values
func NewRegenerateInfraEnvSigningKeyInternalServerError() *RegenerateInfraEnvSigningKeyInternalServerError {

	return &RegenerateInfraEnvSigningKeyInternalServerError{}
}

// WithPayload adds the payload to the regenerate infra env signing key internal server error response
func (o *RegenerateInfraEnvSigningKeyInternalServerError) WithPayload(payload *models.Error) *RegenerateInfraEnvSigningKeyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the regenerate infra env signing key internal server error response
func (o *RegenerateInfraEnvSigningKeyInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegenerateInfraEnvSigningKeyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}