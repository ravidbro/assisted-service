// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewUnbindHostParams creates a new UnbindHostParams object
// no default values defined in spec.
func NewUnbindHostParams() UnbindHostParams {

	return UnbindHostParams{}
}

// UnbindHostParams contains all the bound params for the unbind host operation
// typically these are obtained from a http.Request
//
// swagger:parameters UnbindHost
type UnbindHostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The host that is being bound.
	  Required: true
	  In: path
	*/
	HostID strfmt.UUID
	/*The infra env of the host that is being bound.
	  Required: true
	  In: path
	*/
	InfraEnvID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUnbindHostParams() beforehand.
func (o *UnbindHostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rHostID, rhkHostID, _ := route.Params.GetOK("host_id")
	if err := o.bindHostID(rHostID, rhkHostID, route.Formats); err != nil {
		res = append(res, err)
	}

	rInfraEnvID, rhkInfraEnvID, _ := route.Params.GetOK("infra_env_id")
	if err := o.bindInfraEnvID(rInfraEnvID, rhkInfraEnvID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHostID binds and validates parameter HostID from path.
func (o *UnbindHostParams) bindHostID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("host_id", "path", "strfmt.UUID", raw)
	}
	o.HostID = *(value.(*strfmt.UUID))

	if err := o.validateHostID(formats); err != nil {
		return err
	}

	return nil
}

// validateHostID carries on validations for parameter HostID
func (o *UnbindHostParams) validateHostID(formats strfmt.Registry) error {

	if err := validate.FormatOf("host_id", "path", "uuid", o.HostID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindInfraEnvID binds and validates parameter InfraEnvID from path.
func (o *UnbindHostParams) bindInfraEnvID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("infra_env_id", "path", "strfmt.UUID", raw)
	}
	o.InfraEnvID = *(value.(*strfmt.UUID))

	if err := o.validateInfraEnvID(formats); err != nil {
		return err
	}

	return nil
}

// validateInfraEnvID carries on validations for parameter InfraEnvID
func (o *UnbindHostParams) validateInfraEnvID(formats strfmt.Registry) error {

	if err := validate.FormatOf("infra_env_id", "path", "uuid", o.InfraEnvID.String(), formats); err != nil {
		return err
	}
	return nil
}
