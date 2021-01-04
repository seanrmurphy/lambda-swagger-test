// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPathParamEmptyResponseHandlerFunc turns a function with the right signature into a get path param empty response handler
type GetPathParamEmptyResponseHandlerFunc func(GetPathParamEmptyResponseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPathParamEmptyResponseHandlerFunc) Handle(params GetPathParamEmptyResponseParams) middleware.Responder {
	return fn(params)
}

// GetPathParamEmptyResponseHandler interface for that can handle valid get path param empty response params
type GetPathParamEmptyResponseHandler interface {
	Handle(GetPathParamEmptyResponseParams) middleware.Responder
}

// NewGetPathParamEmptyResponse creates a new http.Handler for the get path param empty response operation
func NewGetPathParamEmptyResponse(ctx *middleware.Context, handler GetPathParamEmptyResponseHandler) *GetPathParamEmptyResponse {
	return &GetPathParamEmptyResponse{Context: ctx, Handler: handler}
}

/*GetPathParamEmptyResponse swagger:route GET /path-param/empty-response/{param} open getPathParamEmptyResponse

Endpoint which takes a parameter and returns a HTTP response code only

Endpoint which takes a parameter and returns a HTTP response code only

*/
type GetPathParamEmptyResponse struct {
	Context *middleware.Context
	Handler GetPathParamEmptyResponseHandler
}

func (o *GetPathParamEmptyResponse) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPathParamEmptyResponseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
