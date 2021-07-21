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

// DownloadInfraEnvDiscoveryImageReader is a Reader for the DownloadInfraEnvDiscoveryImage structure.
type DownloadInfraEnvDiscoveryImageReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadInfraEnvDiscoveryImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadInfraEnvDiscoveryImageOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadInfraEnvDiscoveryImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDownloadInfraEnvDiscoveryImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadInfraEnvDiscoveryImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadInfraEnvDiscoveryImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDownloadInfraEnvDiscoveryImageMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDownloadInfraEnvDiscoveryImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadInfraEnvDiscoveryImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDownloadInfraEnvDiscoveryImageNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadInfraEnvDiscoveryImageOK creates a DownloadInfraEnvDiscoveryImageOK with default headers values
func NewDownloadInfraEnvDiscoveryImageOK(writer io.Writer) *DownloadInfraEnvDiscoveryImageOK {
	return &DownloadInfraEnvDiscoveryImageOK{
		Payload: writer,
	}
}

/*DownloadInfraEnvDiscoveryImageOK handles this case with default header values.

Success.
*/
type DownloadInfraEnvDiscoveryImageOK struct {
	Payload io.Writer
}

func (o *DownloadInfraEnvDiscoveryImageOK) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageOK  %+v", 200, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadInfraEnvDiscoveryImageBadRequest creates a DownloadInfraEnvDiscoveryImageBadRequest with default headers values
func NewDownloadInfraEnvDiscoveryImageBadRequest() *DownloadInfraEnvDiscoveryImageBadRequest {
	return &DownloadInfraEnvDiscoveryImageBadRequest{}
}

/*DownloadInfraEnvDiscoveryImageBadRequest handles this case with default header values.

Error.
*/
type DownloadInfraEnvDiscoveryImageBadRequest struct {
	Payload *models.Error
}

func (o *DownloadInfraEnvDiscoveryImageBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadInfraEnvDiscoveryImageUnauthorized creates a DownloadInfraEnvDiscoveryImageUnauthorized with default headers values
func NewDownloadInfraEnvDiscoveryImageUnauthorized() *DownloadInfraEnvDiscoveryImageUnauthorized {
	return &DownloadInfraEnvDiscoveryImageUnauthorized{}
}

/*DownloadInfraEnvDiscoveryImageUnauthorized handles this case with default header values.

Unauthorized.
*/
type DownloadInfraEnvDiscoveryImageUnauthorized struct {
	Payload *models.InfraError
}

func (o *DownloadInfraEnvDiscoveryImageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadInfraEnvDiscoveryImageForbidden creates a DownloadInfraEnvDiscoveryImageForbidden with default headers values
func NewDownloadInfraEnvDiscoveryImageForbidden() *DownloadInfraEnvDiscoveryImageForbidden {
	return &DownloadInfraEnvDiscoveryImageForbidden{}
}

/*DownloadInfraEnvDiscoveryImageForbidden handles this case with default header values.

Forbidden.
*/
type DownloadInfraEnvDiscoveryImageForbidden struct {
	Payload *models.InfraError
}

func (o *DownloadInfraEnvDiscoveryImageForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageForbidden  %+v", 403, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadInfraEnvDiscoveryImageNotFound creates a DownloadInfraEnvDiscoveryImageNotFound with default headers values
func NewDownloadInfraEnvDiscoveryImageNotFound() *DownloadInfraEnvDiscoveryImageNotFound {
	return &DownloadInfraEnvDiscoveryImageNotFound{}
}

/*DownloadInfraEnvDiscoveryImageNotFound handles this case with default header values.

Error.
*/
type DownloadInfraEnvDiscoveryImageNotFound struct {
	Payload *models.Error
}

func (o *DownloadInfraEnvDiscoveryImageNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageNotFound  %+v", 404, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadInfraEnvDiscoveryImageMethodNotAllowed creates a DownloadInfraEnvDiscoveryImageMethodNotAllowed with default headers values
func NewDownloadInfraEnvDiscoveryImageMethodNotAllowed() *DownloadInfraEnvDiscoveryImageMethodNotAllowed {
	return &DownloadInfraEnvDiscoveryImageMethodNotAllowed{}
}

/*DownloadInfraEnvDiscoveryImageMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type DownloadInfraEnvDiscoveryImageMethodNotAllowed struct {
}

func (o *DownloadInfraEnvDiscoveryImageMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageMethodNotAllowed ", 405)
}

func (o *DownloadInfraEnvDiscoveryImageMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadInfraEnvDiscoveryImageConflict creates a DownloadInfraEnvDiscoveryImageConflict with default headers values
func NewDownloadInfraEnvDiscoveryImageConflict() *DownloadInfraEnvDiscoveryImageConflict {
	return &DownloadInfraEnvDiscoveryImageConflict{}
}

/*DownloadInfraEnvDiscoveryImageConflict handles this case with default header values.

Error.
*/
type DownloadInfraEnvDiscoveryImageConflict struct {
	Payload *models.Error
}

func (o *DownloadInfraEnvDiscoveryImageConflict) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageConflict  %+v", 409, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadInfraEnvDiscoveryImageInternalServerError creates a DownloadInfraEnvDiscoveryImageInternalServerError with default headers values
func NewDownloadInfraEnvDiscoveryImageInternalServerError() *DownloadInfraEnvDiscoveryImageInternalServerError {
	return &DownloadInfraEnvDiscoveryImageInternalServerError{}
}

/*DownloadInfraEnvDiscoveryImageInternalServerError handles this case with default header values.

Error.
*/
type DownloadInfraEnvDiscoveryImageInternalServerError struct {
	Payload *models.Error
}

func (o *DownloadInfraEnvDiscoveryImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadInfraEnvDiscoveryImageNotImplemented creates a DownloadInfraEnvDiscoveryImageNotImplemented with default headers values
func NewDownloadInfraEnvDiscoveryImageNotImplemented() *DownloadInfraEnvDiscoveryImageNotImplemented {
	return &DownloadInfraEnvDiscoveryImageNotImplemented{}
}

/*DownloadInfraEnvDiscoveryImageNotImplemented handles this case with default header values.

Not implemented.
*/
type DownloadInfraEnvDiscoveryImageNotImplemented struct {
	Payload *models.Error
}

func (o *DownloadInfraEnvDiscoveryImageNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/image][%d] downloadInfraEnvDiscoveryImageNotImplemented  %+v", 501, o.Payload)
}

func (o *DownloadInfraEnvDiscoveryImageNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadInfraEnvDiscoveryImageNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
