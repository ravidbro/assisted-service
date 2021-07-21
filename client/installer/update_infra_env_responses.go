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

// UpdateInfraEnvReader is a Reader for the UpdateInfraEnv structure.
type UpdateInfraEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInfraEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateInfraEnvCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInfraEnvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateInfraEnvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInfraEnvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInfraEnvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateInfraEnvMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateInfraEnvConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInfraEnvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewUpdateInfraEnvNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInfraEnvCreated creates a UpdateInfraEnvCreated with default headers values
func NewUpdateInfraEnvCreated() *UpdateInfraEnvCreated {
	return &UpdateInfraEnvCreated{}
}

/*UpdateInfraEnvCreated handles this case with default header values.

Success.
*/
type UpdateInfraEnvCreated struct {
	Payload *models.InfraEnv
}

func (o *UpdateInfraEnvCreated) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvCreated  %+v", 201, o.Payload)
}

func (o *UpdateInfraEnvCreated) GetPayload() *models.InfraEnv {
	return o.Payload
}

func (o *UpdateInfraEnvCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraEnv)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvBadRequest creates a UpdateInfraEnvBadRequest with default headers values
func NewUpdateInfraEnvBadRequest() *UpdateInfraEnvBadRequest {
	return &UpdateInfraEnvBadRequest{}
}

/*UpdateInfraEnvBadRequest handles this case with default header values.

Error.
*/
type UpdateInfraEnvBadRequest struct {
	Payload *models.Error
}

func (o *UpdateInfraEnvBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInfraEnvBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInfraEnvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvUnauthorized creates a UpdateInfraEnvUnauthorized with default headers values
func NewUpdateInfraEnvUnauthorized() *UpdateInfraEnvUnauthorized {
	return &UpdateInfraEnvUnauthorized{}
}

/*UpdateInfraEnvUnauthorized handles this case with default header values.

Unauthorized.
*/
type UpdateInfraEnvUnauthorized struct {
	Payload *models.InfraError
}

func (o *UpdateInfraEnvUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateInfraEnvUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateInfraEnvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvForbidden creates a UpdateInfraEnvForbidden with default headers values
func NewUpdateInfraEnvForbidden() *UpdateInfraEnvForbidden {
	return &UpdateInfraEnvForbidden{}
}

/*UpdateInfraEnvForbidden handles this case with default header values.

Forbidden.
*/
type UpdateInfraEnvForbidden struct {
	Payload *models.InfraError
}

func (o *UpdateInfraEnvForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInfraEnvForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateInfraEnvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvNotFound creates a UpdateInfraEnvNotFound with default headers values
func NewUpdateInfraEnvNotFound() *UpdateInfraEnvNotFound {
	return &UpdateInfraEnvNotFound{}
}

/*UpdateInfraEnvNotFound handles this case with default header values.

Error.
*/
type UpdateInfraEnvNotFound struct {
	Payload *models.Error
}

func (o *UpdateInfraEnvNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInfraEnvNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInfraEnvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvMethodNotAllowed creates a UpdateInfraEnvMethodNotAllowed with default headers values
func NewUpdateInfraEnvMethodNotAllowed() *UpdateInfraEnvMethodNotAllowed {
	return &UpdateInfraEnvMethodNotAllowed{}
}

/*UpdateInfraEnvMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type UpdateInfraEnvMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateInfraEnvMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateInfraEnvMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInfraEnvMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvConflict creates a UpdateInfraEnvConflict with default headers values
func NewUpdateInfraEnvConflict() *UpdateInfraEnvConflict {
	return &UpdateInfraEnvConflict{}
}

/*UpdateInfraEnvConflict handles this case with default header values.

Error.
*/
type UpdateInfraEnvConflict struct {
	Payload *models.Error
}

func (o *UpdateInfraEnvConflict) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvConflict  %+v", 409, o.Payload)
}

func (o *UpdateInfraEnvConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInfraEnvConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvInternalServerError creates a UpdateInfraEnvInternalServerError with default headers values
func NewUpdateInfraEnvInternalServerError() *UpdateInfraEnvInternalServerError {
	return &UpdateInfraEnvInternalServerError{}
}

/*UpdateInfraEnvInternalServerError handles this case with default header values.

Error.
*/
type UpdateInfraEnvInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateInfraEnvInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInfraEnvInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInfraEnvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraEnvNotImplemented creates a UpdateInfraEnvNotImplemented with default headers values
func NewUpdateInfraEnvNotImplemented() *UpdateInfraEnvNotImplemented {
	return &UpdateInfraEnvNotImplemented{}
}

/*UpdateInfraEnvNotImplemented handles this case with default header values.

Not implemented.
*/
type UpdateInfraEnvNotImplemented struct {
	Payload *models.Error
}

func (o *UpdateInfraEnvNotImplemented) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}][%d] updateInfraEnvNotImplemented  %+v", 501, o.Payload)
}

func (o *UpdateInfraEnvNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInfraEnvNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
