// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UnbindHostHandlerFunc turns a function with the right signature into a unbind host handler
type UnbindHostHandlerFunc func(UnbindHostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UnbindHostHandlerFunc) Handle(params UnbindHostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UnbindHostHandler interface for that can handle valid unbind host params
type UnbindHostHandler interface {
	Handle(UnbindHostParams, interface{}) middleware.Responder
}

// NewUnbindHost creates a new http.Handler for the unbind host operation
func NewUnbindHost(ctx *middleware.Context, handler UnbindHostHandler) *UnbindHost {
	return &UnbindHost{Context: ctx, Handler: handler}
}

/*UnbindHost swagger:route POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind installer unbindHost

Unbind host to a cluster

*/
type UnbindHost struct {
	Context *middleware.Context
	Handler UnbindHostHandler
}

func (o *UnbindHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUnbindHostParams()

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
