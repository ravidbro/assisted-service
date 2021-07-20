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

// GetDiscoveryIgnitionReader is a Reader for the GetDiscoveryIgnition structure.
type GetDiscoveryIgnitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoveryIgnitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoveryIgnitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDiscoveryIgnitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDiscoveryIgnitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDiscoveryIgnitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDiscoveryIgnitionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDiscoveryIgnitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoveryIgnitionOK creates a GetDiscoveryIgnitionOK with default headers values
func NewGetDiscoveryIgnitionOK() *GetDiscoveryIgnitionOK {
	return &GetDiscoveryIgnitionOK{}
}

/*GetDiscoveryIgnitionOK handles this case with default header values.

Success.
*/
type GetDiscoveryIgnitionOK struct {
	Payload *models.DiscoveryIgnitionParams
}

func (o *GetDiscoveryIgnitionOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/discovery-ignition][%d] getDiscoveryIgnitionOK  %+v", 200, o.Payload)
}

func (o *GetDiscoveryIgnitionOK) GetPayload() *models.DiscoveryIgnitionParams {
	return o.Payload
}

func (o *GetDiscoveryIgnitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiscoveryIgnitionParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryIgnitionUnauthorized creates a GetDiscoveryIgnitionUnauthorized with default headers values
func NewGetDiscoveryIgnitionUnauthorized() *GetDiscoveryIgnitionUnauthorized {
	return &GetDiscoveryIgnitionUnauthorized{}
}

/*GetDiscoveryIgnitionUnauthorized handles this case with default header values.

Unauthorized.
*/
type GetDiscoveryIgnitionUnauthorized struct {
	Payload *models.InfraError
}

func (o *GetDiscoveryIgnitionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/discovery-ignition][%d] getDiscoveryIgnitionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDiscoveryIgnitionUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetDiscoveryIgnitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryIgnitionForbidden creates a GetDiscoveryIgnitionForbidden with default headers values
func NewGetDiscoveryIgnitionForbidden() *GetDiscoveryIgnitionForbidden {
	return &GetDiscoveryIgnitionForbidden{}
}

/*GetDiscoveryIgnitionForbidden handles this case with default header values.

Forbidden.
*/
type GetDiscoveryIgnitionForbidden struct {
	Payload *models.InfraError
}

func (o *GetDiscoveryIgnitionForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/discovery-ignition][%d] getDiscoveryIgnitionForbidden  %+v", 403, o.Payload)
}

func (o *GetDiscoveryIgnitionForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetDiscoveryIgnitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryIgnitionNotFound creates a GetDiscoveryIgnitionNotFound with default headers values
func NewGetDiscoveryIgnitionNotFound() *GetDiscoveryIgnitionNotFound {
	return &GetDiscoveryIgnitionNotFound{}
}

/*GetDiscoveryIgnitionNotFound handles this case with default header values.

Error.
*/
type GetDiscoveryIgnitionNotFound struct {
	Payload *models.Error
}

func (o *GetDiscoveryIgnitionNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/discovery-ignition][%d] getDiscoveryIgnitionNotFound  %+v", 404, o.Payload)
}

func (o *GetDiscoveryIgnitionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryIgnitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryIgnitionMethodNotAllowed creates a GetDiscoveryIgnitionMethodNotAllowed with default headers values
func NewGetDiscoveryIgnitionMethodNotAllowed() *GetDiscoveryIgnitionMethodNotAllowed {
	return &GetDiscoveryIgnitionMethodNotAllowed{}
}

/*GetDiscoveryIgnitionMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type GetDiscoveryIgnitionMethodNotAllowed struct {
	Payload *models.Error
}

func (o *GetDiscoveryIgnitionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/discovery-ignition][%d] getDiscoveryIgnitionMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetDiscoveryIgnitionMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryIgnitionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryIgnitionInternalServerError creates a GetDiscoveryIgnitionInternalServerError with default headers values
func NewGetDiscoveryIgnitionInternalServerError() *GetDiscoveryIgnitionInternalServerError {
	return &GetDiscoveryIgnitionInternalServerError{}
}

/*GetDiscoveryIgnitionInternalServerError handles this case with default header values.

Error.
*/
type GetDiscoveryIgnitionInternalServerError struct {
	Payload *models.Error
}

func (o *GetDiscoveryIgnitionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/discovery-ignition][%d] getDiscoveryIgnitionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDiscoveryIgnitionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryIgnitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
