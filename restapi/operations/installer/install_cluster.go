// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// InstallClusterHandlerFunc turns a function with the right signature into a install cluster handler
type InstallClusterHandlerFunc func(InstallClusterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn InstallClusterHandlerFunc) Handle(params InstallClusterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// InstallClusterHandler interface for that can handle valid install cluster params
type InstallClusterHandler interface {
	Handle(InstallClusterParams, interface{}) middleware.Responder
}

// NewInstallCluster creates a new http.Handler for the install cluster operation
func NewInstallCluster(ctx *middleware.Context, handler InstallClusterHandler) *InstallCluster {
	return &InstallCluster{Context: ctx, Handler: handler}
}

/*InstallCluster swagger:route POST /v1/clusters/{cluster_id}/actions/install installer installCluster

Installs the OpenShift cluster.

*/
type InstallCluster struct {
	Context *middleware.Context
	Handler InstallClusterHandler
}

func (o *InstallCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewInstallClusterParams()

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
