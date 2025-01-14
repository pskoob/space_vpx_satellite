// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// CreateChassisOKCode is the HTTP code returned for type CreateChassisOK
const CreateChassisOKCode int = 200

/*
CreateChassisOK Create Chassis Response

swagger:response createChassisOK
*/
type CreateChassisOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateChassisResponse `json:"body,omitempty"`
}

// NewCreateChassisOK creates CreateChassisOK with default headers values
func NewCreateChassisOK() *CreateChassisOK {

	return &CreateChassisOK{}
}

// WithPayload adds the payload to the create chassis o k response
func (o *CreateChassisOK) WithPayload(payload *models.CreateChassisResponse) *CreateChassisOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create chassis o k response
func (o *CreateChassisOK) SetPayload(payload *models.CreateChassisResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChassisOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateChassisBadRequestCode is the HTTP code returned for type CreateChassisBadRequest
const CreateChassisBadRequestCode int = 400

/*
CreateChassisBadRequest Bad request

swagger:response createChassisBadRequest
*/
type CreateChassisBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateChassisBadRequest creates CreateChassisBadRequest with default headers values
func NewCreateChassisBadRequest() *CreateChassisBadRequest {

	return &CreateChassisBadRequest{}
}

// WithPayload adds the payload to the create chassis bad request response
func (o *CreateChassisBadRequest) WithPayload(payload *models.Error) *CreateChassisBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create chassis bad request response
func (o *CreateChassisBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChassisBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateChassisForbiddenCode is the HTTP code returned for type CreateChassisForbidden
const CreateChassisForbiddenCode int = 403

/*
CreateChassisForbidden Forbidden

swagger:response createChassisForbidden
*/
type CreateChassisForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateChassisForbidden creates CreateChassisForbidden with default headers values
func NewCreateChassisForbidden() *CreateChassisForbidden {

	return &CreateChassisForbidden{}
}

// WithPayload adds the payload to the create chassis forbidden response
func (o *CreateChassisForbidden) WithPayload(payload *models.Error) *CreateChassisForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create chassis forbidden response
func (o *CreateChassisForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChassisForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateChassisUnprocessableEntityCode is the HTTP code returned for type CreateChassisUnprocessableEntity
const CreateChassisUnprocessableEntityCode int = 422

/*
CreateChassisUnprocessableEntity Unprocessable Entity

swagger:response createChassisUnprocessableEntity
*/
type CreateChassisUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateChassisUnprocessableEntity creates CreateChassisUnprocessableEntity with default headers values
func NewCreateChassisUnprocessableEntity() *CreateChassisUnprocessableEntity {

	return &CreateChassisUnprocessableEntity{}
}

// WithPayload adds the payload to the create chassis unprocessable entity response
func (o *CreateChassisUnprocessableEntity) WithPayload(payload *models.Error) *CreateChassisUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create chassis unprocessable entity response
func (o *CreateChassisUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChassisUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateChassisInternalServerErrorCode is the HTTP code returned for type CreateChassisInternalServerError
const CreateChassisInternalServerErrorCode int = 500

/*
CreateChassisInternalServerError Internal server error

swagger:response createChassisInternalServerError
*/
type CreateChassisInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateChassisInternalServerError creates CreateChassisInternalServerError with default headers values
func NewCreateChassisInternalServerError() *CreateChassisInternalServerError {

	return &CreateChassisInternalServerError{}
}

// WithPayload adds the payload to the create chassis internal server error response
func (o *CreateChassisInternalServerError) WithPayload(payload *models.Error) *CreateChassisInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create chassis internal server error response
func (o *CreateChassisInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChassisInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
