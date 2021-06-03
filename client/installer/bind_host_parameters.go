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

	"github.com/openshift/assisted-service/models"
)

// NewBindHostParams creates a new BindHostParams object
// with the default values initialized.
func NewBindHostParams() *BindHostParams {
	var ()
	return &BindHostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBindHostParamsWithTimeout creates a new BindHostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBindHostParamsWithTimeout(timeout time.Duration) *BindHostParams {
	var ()
	return &BindHostParams{

		timeout: timeout,
	}
}

// NewBindHostParamsWithContext creates a new BindHostParams object
// with the default values initialized, and the ability to set a context for a request
func NewBindHostParamsWithContext(ctx context.Context) *BindHostParams {
	var ()
	return &BindHostParams{

		Context: ctx,
	}
}

// NewBindHostParamsWithHTTPClient creates a new BindHostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBindHostParamsWithHTTPClient(client *http.Client) *BindHostParams {
	var ()
	return &BindHostParams{
		HTTPClient: client,
	}
}

/*BindHostParams contains all the parameters to send to the API endpoint
for the bind host operation typically these are written to a http.Request
*/
type BindHostParams struct {

	/*ClusterID
	  The cluster of the host that it's currently belings to.

	*/
	ClusterID strfmt.UUID
	/*HostID
	  The host that that is going to be binded.

	*/
	HostID strfmt.UUID
	/*NewClusterID
	  The destination cluster ID.

	*/
	NewClusterID *models.ClusterID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bind host params
func (o *BindHostParams) WithTimeout(timeout time.Duration) *BindHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bind host params
func (o *BindHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bind host params
func (o *BindHostParams) WithContext(ctx context.Context) *BindHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bind host params
func (o *BindHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bind host params
func (o *BindHostParams) WithHTTPClient(client *http.Client) *BindHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bind host params
func (o *BindHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the bind host params
func (o *BindHostParams) WithClusterID(clusterID strfmt.UUID) *BindHostParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the bind host params
func (o *BindHostParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the bind host params
func (o *BindHostParams) WithHostID(hostID strfmt.UUID) *BindHostParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the bind host params
func (o *BindHostParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithNewClusterID adds the newClusterID to the bind host params
func (o *BindHostParams) WithNewClusterID(newClusterID *models.ClusterID) *BindHostParams {
	o.SetNewClusterID(newClusterID)
	return o
}

// SetNewClusterID adds the newClusterId to the bind host params
func (o *BindHostParams) SetNewClusterID(newClusterID *models.ClusterID) {
	o.NewClusterID = newClusterID
}

// WriteToRequest writes these params to a swagger request
func (o *BindHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if o.NewClusterID != nil {
		if err := r.SetBodyParam(o.NewClusterID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
