// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/seanrmurphy/lambda-swagger-test/models"
)

// GetPathParamErrorResponseIMATeapotCode is the HTTP code returned for type GetPathParamErrorResponseIMATeapot
const GetPathParamErrorResponseIMATeapotCode int = 418

/*GetPathParamErrorResponseIMATeapot Error generated when executing API call

swagger:response getPathParamErrorResponseIMATeapot
*/
type GetPathParamErrorResponseIMATeapot struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetPathParamErrorResponseIMATeapot creates GetPathParamErrorResponseIMATeapot with default headers values
func NewGetPathParamErrorResponseIMATeapot() *GetPathParamErrorResponseIMATeapot {

	return &GetPathParamErrorResponseIMATeapot{}
}

// WithPayload adds the payload to the get path param error response i m a teapot response
func (o *GetPathParamErrorResponseIMATeapot) WithPayload(payload *models.ErrorResponse) *GetPathParamErrorResponseIMATeapot {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get path param error response i m a teapot response
func (o *GetPathParamErrorResponseIMATeapot) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPathParamErrorResponseIMATeapot) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(418)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPathParamErrorResponseInternalServerErrorCode is the HTTP code returned for type GetPathParamErrorResponseInternalServerError
const GetPathParamErrorResponseInternalServerErrorCode int = 500

/*GetPathParamErrorResponseInternalServerError General Failure

swagger:response getPathParamErrorResponseInternalServerError
*/
type GetPathParamErrorResponseInternalServerError struct {
}

// NewGetPathParamErrorResponseInternalServerError creates GetPathParamErrorResponseInternalServerError with default headers values
func NewGetPathParamErrorResponseInternalServerError() *GetPathParamErrorResponseInternalServerError {

	return &GetPathParamErrorResponseInternalServerError{}
}

// WriteResponse to the client
func (o *GetPathParamErrorResponseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}