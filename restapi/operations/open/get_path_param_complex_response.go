// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPathParamComplexResponseHandlerFunc turns a function with the right signature into a get path param complex response handler
type GetPathParamComplexResponseHandlerFunc func(GetPathParamComplexResponseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPathParamComplexResponseHandlerFunc) Handle(params GetPathParamComplexResponseParams) middleware.Responder {
	return fn(params)
}

// GetPathParamComplexResponseHandler interface for that can handle valid get path param complex response params
type GetPathParamComplexResponseHandler interface {
	Handle(GetPathParamComplexResponseParams) middleware.Responder
}

// NewGetPathParamComplexResponse creates a new http.Handler for the get path param complex response operation
func NewGetPathParamComplexResponse(ctx *middleware.Context, handler GetPathParamComplexResponseHandler) *GetPathParamComplexResponse {
	return &GetPathParamComplexResponse{Context: ctx, Handler: handler}
}

/*GetPathParamComplexResponse swagger:route GET /path-param/complex-response/{param} open getPathParamComplexResponse

An endpoint which takes an input parameters and returns a HTTP response code and a complex JSON object in a body

An endpoint which takes an input parameter and returns a HTTP response code and a complex JSON object in a body

*/
type GetPathParamComplexResponse struct {
	Context *middleware.Context
	Handler GetPathParamComplexResponseHandler
}

func (o *GetPathParamComplexResponse) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPathParamComplexResponseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
