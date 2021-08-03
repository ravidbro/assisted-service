// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetClusterSupportedPlatformsParams creates a new GetClusterSupportedPlatformsParams object
// with the default values initialized.
func NewGetClusterSupportedPlatformsParams() *GetClusterSupportedPlatformsParams {
	var ()
	return &GetClusterSupportedPlatformsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterSupportedPlatformsParamsWithTimeout creates a new GetClusterSupportedPlatformsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterSupportedPlatformsParamsWithTimeout(timeout time.Duration) *GetClusterSupportedPlatformsParams {
	var ()
	return &GetClusterSupportedPlatformsParams{

		timeout: timeout,
	}
}

// NewGetClusterSupportedPlatformsParamsWithContext creates a new GetClusterSupportedPlatformsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterSupportedPlatformsParamsWithContext(ctx context.Context) *GetClusterSupportedPlatformsParams {
	var ()
	return &GetClusterSupportedPlatformsParams{

		Context: ctx,
	}
}

// NewGetClusterSupportedPlatformsParamsWithHTTPClient creates a new GetClusterSupportedPlatformsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterSupportedPlatformsParamsWithHTTPClient(client *http.Client) *GetClusterSupportedPlatformsParams {
	var ()
	return &GetClusterSupportedPlatformsParams{
		HTTPClient: client,
	}
}

/*GetClusterSupportedPlatformsParams contains all the parameters to send to the API endpoint
for the get cluster supported platforms operation typically these are written to a http.Request
*/
type GetClusterSupportedPlatformsParams struct {

	/*ClusterID
	  The cluster whose platform types should be retrieved.

	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) WithTimeout(timeout time.Duration) *GetClusterSupportedPlatformsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) WithContext(ctx context.Context) *GetClusterSupportedPlatformsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) WithHTTPClient(client *http.Client) *GetClusterSupportedPlatformsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) WithClusterID(clusterID strfmt.UUID) *GetClusterSupportedPlatformsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster supported platforms params
func (o *GetClusterSupportedPlatformsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterSupportedPlatformsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}