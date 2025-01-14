// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetUserMeReader is a Reader for the GetUserMe structure.
type GetUserMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetUserMeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user/get_me] GetUserMe", response, response.Code())
	}
}

// NewGetUserMeOK creates a GetUserMeOK with default headers values
func NewGetUserMeOK() *GetUserMeOK {
	return &GetUserMeOK{}
}

/*
GetUserMeOK describes a response with status code 200, with default header values.

Get User Me Response
*/
type GetUserMeOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get user me o k response has a 2xx status code
func (o *GetUserMeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user me o k response has a 3xx status code
func (o *GetUserMeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user me o k response has a 4xx status code
func (o *GetUserMeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user me o k response has a 5xx status code
func (o *GetUserMeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user me o k response a status code equal to that given
func (o *GetUserMeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user me o k response
func (o *GetUserMeOK) Code() int {
	return 200
}

func (o *GetUserMeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeOK %s", 200, payload)
}

func (o *GetUserMeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeOK %s", 200, payload)
}

func (o *GetUserMeOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserMeBadRequest creates a GetUserMeBadRequest with default headers values
func NewGetUserMeBadRequest() *GetUserMeBadRequest {
	return &GetUserMeBadRequest{}
}

/*
GetUserMeBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUserMeBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user me bad request response has a 2xx status code
func (o *GetUserMeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user me bad request response has a 3xx status code
func (o *GetUserMeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user me bad request response has a 4xx status code
func (o *GetUserMeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user me bad request response has a 5xx status code
func (o *GetUserMeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user me bad request response a status code equal to that given
func (o *GetUserMeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get user me bad request response
func (o *GetUserMeBadRequest) Code() int {
	return 400
}

func (o *GetUserMeBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeBadRequest %s", 400, payload)
}

func (o *GetUserMeBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeBadRequest %s", 400, payload)
}

func (o *GetUserMeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserMeForbidden creates a GetUserMeForbidden with default headers values
func NewGetUserMeForbidden() *GetUserMeForbidden {
	return &GetUserMeForbidden{}
}

/*
GetUserMeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUserMeForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user me forbidden response has a 2xx status code
func (o *GetUserMeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user me forbidden response has a 3xx status code
func (o *GetUserMeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user me forbidden response has a 4xx status code
func (o *GetUserMeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user me forbidden response has a 5xx status code
func (o *GetUserMeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user me forbidden response a status code equal to that given
func (o *GetUserMeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user me forbidden response
func (o *GetUserMeForbidden) Code() int {
	return 403
}

func (o *GetUserMeForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeForbidden %s", 403, payload)
}

func (o *GetUserMeForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeForbidden %s", 403, payload)
}

func (o *GetUserMeForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserMeUnprocessableEntity creates a GetUserMeUnprocessableEntity with default headers values
func NewGetUserMeUnprocessableEntity() *GetUserMeUnprocessableEntity {
	return &GetUserMeUnprocessableEntity{}
}

/*
GetUserMeUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetUserMeUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user me unprocessable entity response has a 2xx status code
func (o *GetUserMeUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user me unprocessable entity response has a 3xx status code
func (o *GetUserMeUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user me unprocessable entity response has a 4xx status code
func (o *GetUserMeUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user me unprocessable entity response has a 5xx status code
func (o *GetUserMeUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get user me unprocessable entity response a status code equal to that given
func (o *GetUserMeUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get user me unprocessable entity response
func (o *GetUserMeUnprocessableEntity) Code() int {
	return 422
}

func (o *GetUserMeUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeUnprocessableEntity %s", 422, payload)
}

func (o *GetUserMeUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeUnprocessableEntity %s", 422, payload)
}

func (o *GetUserMeUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserMeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserMeInternalServerError creates a GetUserMeInternalServerError with default headers values
func NewGetUserMeInternalServerError() *GetUserMeInternalServerError {
	return &GetUserMeInternalServerError{}
}

/*
GetUserMeInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUserMeInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user me internal server error response has a 2xx status code
func (o *GetUserMeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user me internal server error response has a 3xx status code
func (o *GetUserMeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user me internal server error response has a 4xx status code
func (o *GetUserMeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user me internal server error response has a 5xx status code
func (o *GetUserMeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user me internal server error response a status code equal to that given
func (o *GetUserMeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user me internal server error response
func (o *GetUserMeInternalServerError) Code() int {
	return 500
}

func (o *GetUserMeInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeInternalServerError %s", 500, payload)
}

func (o *GetUserMeInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/get_me][%d] getUserMeInternalServerError %s", 500, payload)
}

func (o *GetUserMeInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
