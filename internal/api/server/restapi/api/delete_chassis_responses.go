// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteChassisNoContentCode is the HTTP code returned for type DeleteChassisNoContent
const DeleteChassisNoContentCode int = 204

/*
DeleteChassisNoContent Chassis deleted successfully

swagger:response deleteChassisNoContent
*/
type DeleteChassisNoContent struct {
}

// NewDeleteChassisNoContent creates DeleteChassisNoContent with default headers values
func NewDeleteChassisNoContent() *DeleteChassisNoContent {

	return &DeleteChassisNoContent{}
}

// WriteResponse to the client
func (o *DeleteChassisNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteChassisForbiddenCode is the HTTP code returned for type DeleteChassisForbidden
const DeleteChassisForbiddenCode int = 403

/*
DeleteChassisForbidden Forbidden

swagger:response deleteChassisForbidden
*/
type DeleteChassisForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteChassisForbidden creates DeleteChassisForbidden with default headers values
func NewDeleteChassisForbidden() *DeleteChassisForbidden {

	return &DeleteChassisForbidden{}
}

// WithPayload adds the payload to the delete chassis forbidden response
func (o *DeleteChassisForbidden) WithPayload(payload *models.Error) *DeleteChassisForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete chassis forbidden response
func (o *DeleteChassisForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteChassisForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteChassisNotFoundCode is the HTTP code returned for type DeleteChassisNotFound
const DeleteChassisNotFoundCode int = 404

/*
DeleteChassisNotFound Chassis not found

swagger:response deleteChassisNotFound
*/
type DeleteChassisNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteChassisNotFound creates DeleteChassisNotFound with default headers values
func NewDeleteChassisNotFound() *DeleteChassisNotFound {

	return &DeleteChassisNotFound{}
}

// WithPayload adds the payload to the delete chassis not found response
func (o *DeleteChassisNotFound) WithPayload(payload *models.Error) *DeleteChassisNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete chassis not found response
func (o *DeleteChassisNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteChassisNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteChassisInternalServerErrorCode is the HTTP code returned for type DeleteChassisInternalServerError
const DeleteChassisInternalServerErrorCode int = 500

/*
DeleteChassisInternalServerError Internal server error

swagger:response deleteChassisInternalServerError
*/
type DeleteChassisInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteChassisInternalServerError creates DeleteChassisInternalServerError with default headers values
func NewDeleteChassisInternalServerError() *DeleteChassisInternalServerError {

	return &DeleteChassisInternalServerError{}
}

// WithPayload adds the payload to the delete chassis internal server error response
func (o *DeleteChassisInternalServerError) WithPayload(payload *models.Error) *DeleteChassisInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete chassis internal server error response
func (o *DeleteChassisInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteChassisInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
