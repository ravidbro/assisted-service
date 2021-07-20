// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// InstallHostsHandlerFunc turns a function with the right signature into a install hosts handler
type InstallHostsHandlerFunc func(InstallHostsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn InstallHostsHandlerFunc) Handle(params InstallHostsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// InstallHostsHandler interface for that can handle valid install hosts params
type InstallHostsHandler interface {
	Handle(InstallHostsParams, interface{}) middleware.Responder
}

// NewInstallHosts creates a new http.Handler for the install hosts operation
func NewInstallHosts(ctx *middleware.Context, handler InstallHostsHandler) *InstallHosts {
	return &InstallHosts{Context: ctx, Handler: handler}
}

/*InstallHosts swagger:route POST /v1/clusters/{cluster_id}/actions/install_hosts installer installHosts

Installs the OpenShift cluster.

*/
type InstallHosts struct {
	Context *middleware.Context
	Handler InstallHostsHandler
}

func (o *InstallHosts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewInstallHostsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
