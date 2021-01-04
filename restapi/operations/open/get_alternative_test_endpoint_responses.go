// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/seanrmurphy/lambda-swagger-test/models"
)

// GetAlternativeTestEndpointOKCode is the HTTP code returned for type GetAlternativeTestEndpointOK
const GetAlternativeTestEndpointOKCode int = 200

/*GetAlternativeTestEndpointOK Test Response

swagger:response getAlternativeTestEndpointOK
*/
type GetAlternativeTestEndpointOK struct {

	/*
	  In: Body
	*/
	Payload *models.TestResponse `json:"body,omitempty"`
}

// NewGetAlternativeTestEndpointOK creates GetAlternativeTestEndpointOK with default headers values
func NewGetAlternativeTestEndpointOK() *GetAlternativeTestEndpointOK {

	return &GetAlternativeTestEndpointOK{}
}

// WithPayload adds the payload to the get alternative test endpoint o k response
func (o *GetAlternativeTestEndpointOK) WithPayload(payload *models.TestResponse) *GetAlternativeTestEndpointOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alternative test endpoint o k response
func (o *GetAlternativeTestEndpointOK) SetPayload(payload *models.TestResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlternativeTestEndpointOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAlternativeTestEndpointInternalServerErrorCode is the HTTP code returned for type GetAlternativeTestEndpointInternalServerError
const GetAlternativeTestEndpointInternalServerErrorCode int = 500

/*GetAlternativeTestEndpointInternalServerError General Failure

swagger:response getAlternativeTestEndpointInternalServerError
*/
type GetAlternativeTestEndpointInternalServerError struct {
}

// NewGetAlternativeTestEndpointInternalServerError creates GetAlternativeTestEndpointInternalServerError with default headers values
func NewGetAlternativeTestEndpointInternalServerError() *GetAlternativeTestEndpointInternalServerError {

	return &GetAlternativeTestEndpointInternalServerError{}
}

// WriteResponse to the client
func (o *GetAlternativeTestEndpointInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
