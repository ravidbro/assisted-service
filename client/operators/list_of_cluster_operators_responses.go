// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// ListOfClusterOperatorsReader is a Reader for the ListOfClusterOperators structure.
type ListOfClusterOperatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOfClusterOperatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOfClusterOperatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListOfClusterOperatorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListOfClusterOperatorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOfClusterOperatorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewListOfClusterOperatorsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListOfClusterOperatorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListOfClusterOperatorsOK creates a ListOfClusterOperatorsOK with default headers values
func NewListOfClusterOperatorsOK() *ListOfClusterOperatorsOK {
	return &ListOfClusterOperatorsOK{}
}

/*ListOfClusterOperatorsOK handles this case with default header values.

Success.
*/
type ListOfClusterOperatorsOK struct {
	Payload models.MonitoredOperatorsList
}

func (o *ListOfClusterOperatorsOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/monitored_operators][%d] listOfClusterOperatorsOK  %+v", 200, o.Payload)
}

func (o *ListOfClusterOperatorsOK) GetPayload() models.MonitoredOperatorsList {
	return o.Payload
}

func (o *ListOfClusterOperatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOfClusterOperatorsUnauthorized creates a ListOfClusterOperatorsUnauthorized with default headers values
func NewListOfClusterOperatorsUnauthorized() *ListOfClusterOperatorsUnauthorized {
	return &ListOfClusterOperatorsUnauthorized{}
}

/*ListOfClusterOperatorsUnauthorized handles this case with default header values.

Unauthorized.
*/
type ListOfClusterOperatorsUnauthorized struct {
	Payload *models.InfraError
}

func (o *ListOfClusterOperatorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/monitored_operators][%d] listOfClusterOperatorsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListOfClusterOperatorsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListOfClusterOperatorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOfClusterOperatorsForbidden creates a ListOfClusterOperatorsForbidden with default headers values
func NewListOfClusterOperatorsForbidden() *ListOfClusterOperatorsForbidden {
	return &ListOfClusterOperatorsForbidden{}
}

/*ListOfClusterOperatorsForbidden handles this case with default header values.

Forbidden.
*/
type ListOfClusterOperatorsForbidden struct {
	Payload *models.InfraError
}

func (o *ListOfClusterOperatorsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/monitored_operators][%d] listOfClusterOperatorsForbidden  %+v", 403, o.Payload)
}

func (o *ListOfClusterOperatorsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListOfClusterOperatorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOfClusterOperatorsNotFound creates a ListOfClusterOperatorsNotFound with default headers values
func NewListOfClusterOperatorsNotFound() *ListOfClusterOperatorsNotFound {
	return &ListOfClusterOperatorsNotFound{}
}

/*ListOfClusterOperatorsNotFound handles this case with default header values.

Error.
*/
type ListOfClusterOperatorsNotFound struct {
	Payload *models.Error
}

func (o *ListOfClusterOperatorsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/monitored_operators][%d] listOfClusterOperatorsNotFound  %+v", 404, o.Payload)
}

func (o *ListOfClusterOperatorsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOfClusterOperatorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOfClusterOperatorsMethodNotAllowed creates a ListOfClusterOperatorsMethodNotAllowed with default headers values
func NewListOfClusterOperatorsMethodNotAllowed() *ListOfClusterOperatorsMethodNotAllowed {
	return &ListOfClusterOperatorsMethodNotAllowed{}
}

/*ListOfClusterOperatorsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type ListOfClusterOperatorsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *ListOfClusterOperatorsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/monitored_operators][%d] listOfClusterOperatorsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListOfClusterOperatorsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOfClusterOperatorsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOfClusterOperatorsInternalServerError creates a ListOfClusterOperatorsInternalServerError with default headers values
func NewListOfClusterOperatorsInternalServerError() *ListOfClusterOperatorsInternalServerError {
	return &ListOfClusterOperatorsInternalServerError{}
}

/*ListOfClusterOperatorsInternalServerError handles this case with default header values.

Error.
*/
type ListOfClusterOperatorsInternalServerError struct {
	Payload *models.Error
}

func (o *ListOfClusterOperatorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/monitored_operators][%d] listOfClusterOperatorsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListOfClusterOperatorsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListOfClusterOperatorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
