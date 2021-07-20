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

// GetCredentialsReader is a Reader for the GetCredentials structure.
type GetCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetCredentialsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetCredentialsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCredentialsOK creates a GetCredentialsOK with default headers values
func NewGetCredentialsOK() *GetCredentialsOK {
	return &GetCredentialsOK{}
}

/*GetCredentialsOK handles this case with default header values.

Success.
*/
type GetCredentialsOK struct {
	Payload *models.Credentials
}

func (o *GetCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/credentials][%d] getCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetCredentialsOK) GetPayload() *models.Credentials {
	return o.Payload
}

func (o *GetCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Credentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsUnauthorized creates a GetCredentialsUnauthorized with default headers values
func NewGetCredentialsUnauthorized() *GetCredentialsUnauthorized {
	return &GetCredentialsUnauthorized{}
}

/*GetCredentialsUnauthorized handles this case with default header values.

Unauthorized.
*/
type GetCredentialsUnauthorized struct {
	Payload *models.InfraError
}

func (o *GetCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/credentials][%d] getCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCredentialsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsForbidden creates a GetCredentialsForbidden with default headers values
func NewGetCredentialsForbidden() *GetCredentialsForbidden {
	return &GetCredentialsForbidden{}
}

/*GetCredentialsForbidden handles this case with default header values.

Forbidden.
*/
type GetCredentialsForbidden struct {
	Payload *models.InfraError
}

func (o *GetCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/credentials][%d] getCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *GetCredentialsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsNotFound creates a GetCredentialsNotFound with default headers values
func NewGetCredentialsNotFound() *GetCredentialsNotFound {
	return &GetCredentialsNotFound{}
}

/*GetCredentialsNotFound handles this case with default header values.

Error.
*/
type GetCredentialsNotFound struct {
	Payload *models.Error
}

func (o *GetCredentialsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/credentials][%d] getCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *GetCredentialsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsMethodNotAllowed creates a GetCredentialsMethodNotAllowed with default headers values
func NewGetCredentialsMethodNotAllowed() *GetCredentialsMethodNotAllowed {
	return &GetCredentialsMethodNotAllowed{}
}

/*GetCredentialsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type GetCredentialsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *GetCredentialsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/credentials][%d] getCredentialsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetCredentialsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsConflict creates a GetCredentialsConflict with default headers values
func NewGetCredentialsConflict() *GetCredentialsConflict {
	return &GetCredentialsConflict{}
}

/*GetCredentialsConflict handles this case with default header values.

Error.
*/
type GetCredentialsConflict struct {
	Payload *models.Error
}

func (o *GetCredentialsConflict) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/credentials][%d] getCredentialsConflict  %+v", 409, o.Payload)
}

func (o *GetCredentialsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsInternalServerError creates a GetCredentialsInternalServerError with default headers values
func NewGetCredentialsInternalServerError() *GetCredentialsInternalServerError {
	return &GetCredentialsInternalServerError{}
}

/*GetCredentialsInternalServerError handles this case with default header values.

Error.
*/
type GetCredentialsInternalServerError struct {
	Payload *models.Error
}

func (o *GetCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/credentials][%d] getCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
