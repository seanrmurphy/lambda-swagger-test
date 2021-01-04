// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/seanrmurphy/lambda-swagger-test/models"
)

// GetBodyParamSimpleResponseOKCode is the HTTP code returned for type GetBodyParamSimpleResponseOK
const GetBodyParamSimpleResponseOKCode int = 200

/*GetBodyParamSimpleResponseOK Successful execution of API call returning message in simple JSON object

swagger:response getBodyParamSimpleResponseOK
*/
type GetBodyParamSimpleResponseOK struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleMessageResponse `json:"body,omitempty"`
}

// NewGetBodyParamSimpleResponseOK creates GetBodyParamSimpleResponseOK with default headers values
func NewGetBodyParamSimpleResponseOK() *GetBodyParamSimpleResponseOK {

	return &GetBodyParamSimpleResponseOK{}
}

// WithPayload adds the payload to the get body param simple response o k response
func (o *GetBodyParamSimpleResponseOK) WithPayload(payload *models.SimpleMessageResponse) *GetBodyParamSimpleResponseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get body param simple response o k response
func (o *GetBodyParamSimpleResponseOK) SetPayload(payload *models.SimpleMessageResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBodyParamSimpleResponseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBodyParamSimpleResponseInternalServerErrorCode is the HTTP code returned for type GetBodyParamSimpleResponseInternalServerError
const GetBodyParamSimpleResponseInternalServerErrorCode int = 500

/*GetBodyParamSimpleResponseInternalServerError General Failure

swagger:response getBodyParamSimpleResponseInternalServerError
*/
type GetBodyParamSimpleResponseInternalServerError struct {
}

// NewGetBodyParamSimpleResponseInternalServerError creates GetBodyParamSimpleResponseInternalServerError with default headers values
func NewGetBodyParamSimpleResponseInternalServerError() *GetBodyParamSimpleResponseInternalServerError {

	return &GetBodyParamSimpleResponseInternalServerError{}
}

// WriteResponse to the client
func (o *GetBodyParamSimpleResponseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
