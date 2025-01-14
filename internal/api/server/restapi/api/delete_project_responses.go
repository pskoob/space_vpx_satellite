// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteProjectOKCode is the HTTP code returned for type DeleteProjectOK
const DeleteProjectOKCode int = 200

/*
DeleteProjectOK Delete Project Response

swagger:response deleteProjectOK
*/
type DeleteProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateProjectResponse `json:"body,omitempty"`
}

// NewDeleteProjectOK creates DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {

	return &DeleteProjectOK{}
}

// WithPayload adds the payload to the delete project o k response
func (o *DeleteProjectOK) WithPayload(payload *models.CreateProjectResponse) *DeleteProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project o k response
func (o *DeleteProjectOK) SetPayload(payload *models.CreateProjectResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectBadRequestCode is the HTTP code returned for type DeleteProjectBadRequest
const DeleteProjectBadRequestCode int = 400

/*
DeleteProjectBadRequest Bad request

swagger:response deleteProjectBadRequest
*/
type DeleteProjectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectBadRequest creates DeleteProjectBadRequest with default headers values
func NewDeleteProjectBadRequest() *DeleteProjectBadRequest {

	return &DeleteProjectBadRequest{}
}

// WithPayload adds the payload to the delete project bad request response
func (o *DeleteProjectBadRequest) WithPayload(payload *models.Error) *DeleteProjectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project bad request response
func (o *DeleteProjectBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectForbiddenCode is the HTTP code returned for type DeleteProjectForbidden
const DeleteProjectForbiddenCode int = 403

/*
DeleteProjectForbidden Forbidden

swagger:response deleteProjectForbidden
*/
type DeleteProjectForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectForbidden creates DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {

	return &DeleteProjectForbidden{}
}

// WithPayload adds the payload to the delete project forbidden response
func (o *DeleteProjectForbidden) WithPayload(payload *models.Error) *DeleteProjectForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project forbidden response
func (o *DeleteProjectForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectUnprocessableEntityCode is the HTTP code returned for type DeleteProjectUnprocessableEntity
const DeleteProjectUnprocessableEntityCode int = 422

/*
DeleteProjectUnprocessableEntity Unprocessable Entity

swagger:response deleteProjectUnprocessableEntity
*/
type DeleteProjectUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectUnprocessableEntity creates DeleteProjectUnprocessableEntity with default headers values
func NewDeleteProjectUnprocessableEntity() *DeleteProjectUnprocessableEntity {

	return &DeleteProjectUnprocessableEntity{}
}

// WithPayload adds the payload to the delete project unprocessable entity response
func (o *DeleteProjectUnprocessableEntity) WithPayload(payload *models.Error) *DeleteProjectUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project unprocessable entity response
func (o *DeleteProjectUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectInternalServerErrorCode is the HTTP code returned for type DeleteProjectInternalServerError
const DeleteProjectInternalServerErrorCode int = 500

/*
DeleteProjectInternalServerError Internal server error

swagger:response deleteProjectInternalServerError
*/
type DeleteProjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectInternalServerError creates DeleteProjectInternalServerError with default headers values
func NewDeleteProjectInternalServerError() *DeleteProjectInternalServerError {

	return &DeleteProjectInternalServerError{}
}

// WithPayload adds the payload to the delete project internal server error response
func (o *DeleteProjectInternalServerError) WithPayload(payload *models.Error) *DeleteProjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project internal server error response
func (o *DeleteProjectInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
