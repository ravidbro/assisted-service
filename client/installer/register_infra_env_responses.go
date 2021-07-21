// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// RegisterInfraEnvReader is a Reader for the RegisterInfraEnv structure.
type RegisterInfraEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterInfraEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRegisterInfraEnvCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterInfraEnvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRegisterInfraEnvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegisterInfraEnvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegisterInfraEnvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewRegisterInfraEnvMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRegisterInfraEnvConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegisterInfraEnvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewRegisterInfraEnvNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterInfraEnvCreated creates a RegisterInfraEnvCreated with default headers values
func NewRegisterInfraEnvCreated() *RegisterInfraEnvCreated {
	return &RegisterInfraEnvCreated{}
}

/*RegisterInfraEnvCreated handles this case with default header values.

Success.
*/
type RegisterInfraEnvCreated struct {
	Payload *models.InfraEnv
}

func (o *RegisterInfraEnvCreated) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvCreated  %+v", 201, o.Payload)
}

func (o *RegisterInfraEnvCreated) GetPayload() *models.InfraEnv {
	return o.Payload
}

func (o *RegisterInfraEnvCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraEnv)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvBadRequest creates a RegisterInfraEnvBadRequest with default headers values
func NewRegisterInfraEnvBadRequest() *RegisterInfraEnvBadRequest {
	return &RegisterInfraEnvBadRequest{}
}

/*RegisterInfraEnvBadRequest handles this case with default header values.

Error.
*/
type RegisterInfraEnvBadRequest struct {
	Payload *models.Error
}

func (o *RegisterInfraEnvBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvBadRequest  %+v", 400, o.Payload)
}

func (o *RegisterInfraEnvBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterInfraEnvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvUnauthorized creates a RegisterInfraEnvUnauthorized with default headers values
func NewRegisterInfraEnvUnauthorized() *RegisterInfraEnvUnauthorized {
	return &RegisterInfraEnvUnauthorized{}
}

/*RegisterInfraEnvUnauthorized handles this case with default header values.

Unauthorized.
*/
type RegisterInfraEnvUnauthorized struct {
	Payload *models.InfraError
}

func (o *RegisterInfraEnvUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvUnauthorized  %+v", 401, o.Payload)
}

func (o *RegisterInfraEnvUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *RegisterInfraEnvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvForbidden creates a RegisterInfraEnvForbidden with default headers values
func NewRegisterInfraEnvForbidden() *RegisterInfraEnvForbidden {
	return &RegisterInfraEnvForbidden{}
}

/*RegisterInfraEnvForbidden handles this case with default header values.

Forbidden.
*/
type RegisterInfraEnvForbidden struct {
	Payload *models.InfraError
}

func (o *RegisterInfraEnvForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvForbidden  %+v", 403, o.Payload)
}

func (o *RegisterInfraEnvForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *RegisterInfraEnvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvNotFound creates a RegisterInfraEnvNotFound with default headers values
func NewRegisterInfraEnvNotFound() *RegisterInfraEnvNotFound {
	return &RegisterInfraEnvNotFound{}
}

/*RegisterInfraEnvNotFound handles this case with default header values.

Error.
*/
type RegisterInfraEnvNotFound struct {
	Payload *models.Error
}

func (o *RegisterInfraEnvNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvNotFound  %+v", 404, o.Payload)
}

func (o *RegisterInfraEnvNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterInfraEnvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvMethodNotAllowed creates a RegisterInfraEnvMethodNotAllowed with default headers values
func NewRegisterInfraEnvMethodNotAllowed() *RegisterInfraEnvMethodNotAllowed {
	return &RegisterInfraEnvMethodNotAllowed{}
}

/*RegisterInfraEnvMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type RegisterInfraEnvMethodNotAllowed struct {
	Payload *models.Error
}

func (o *RegisterInfraEnvMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *RegisterInfraEnvMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterInfraEnvMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvConflict creates a RegisterInfraEnvConflict with default headers values
func NewRegisterInfraEnvConflict() *RegisterInfraEnvConflict {
	return &RegisterInfraEnvConflict{}
}

/*RegisterInfraEnvConflict handles this case with default header values.

Error.
*/
type RegisterInfraEnvConflict struct {
	Payload *models.Error
}

func (o *RegisterInfraEnvConflict) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvConflict  %+v", 409, o.Payload)
}

func (o *RegisterInfraEnvConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterInfraEnvConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvInternalServerError creates a RegisterInfraEnvInternalServerError with default headers values
func NewRegisterInfraEnvInternalServerError() *RegisterInfraEnvInternalServerError {
	return &RegisterInfraEnvInternalServerError{}
}

/*RegisterInfraEnvInternalServerError handles this case with default header values.

Error.
*/
type RegisterInfraEnvInternalServerError struct {
	Payload *models.Error
}

func (o *RegisterInfraEnvInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvInternalServerError  %+v", 500, o.Payload)
}

func (o *RegisterInfraEnvInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterInfraEnvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterInfraEnvNotImplemented creates a RegisterInfraEnvNotImplemented with default headers values
func NewRegisterInfraEnvNotImplemented() *RegisterInfraEnvNotImplemented {
	return &RegisterInfraEnvNotImplemented{}
}

/*RegisterInfraEnvNotImplemented handles this case with default header values.

Not implemented.
*/
type RegisterInfraEnvNotImplemented struct {
	Payload *models.Error
}

func (o *RegisterInfraEnvNotImplemented) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs][%d] registerInfraEnvNotImplemented  %+v", 501, o.Payload)
}

func (o *RegisterInfraEnvNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterInfraEnvNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
