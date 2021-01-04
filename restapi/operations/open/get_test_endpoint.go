// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTestEndpointHandlerFunc turns a function with the right signature into a get test endpoint handler
type GetTestEndpointHandlerFunc func(GetTestEndpointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTestEndpointHandlerFunc) Handle(params GetTestEndpointParams) middleware.Responder {
	return fn(params)
}

// GetTestEndpointHandler interface for that can handle valid get test endpoint params
type GetTestEndpointHandler interface {
	Handle(GetTestEndpointParams) middleware.Responder
}

// NewGetTestEndpoint creates a new http.Handler for the get test endpoint operation
func NewGetTestEndpoint(ctx *middleware.Context, handler GetTestEndpointHandler) *GetTestEndpoint {
	return &GetTestEndpoint{Context: ctx, Handler: handler}
}

/*GetTestEndpoint swagger:route GET /test open getTestEndpoint

Gets test endpoint

An endpoint for testing

*/
type GetTestEndpoint struct {
	Context *middleware.Context
	Handler GetTestEndpointHandler
}

func (o *GetTestEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTestEndpointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}