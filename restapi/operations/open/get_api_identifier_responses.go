// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/seanrmurphy/lambda-swagger-test/models"
)

// GetAPIIdentifierOKCode is the HTTP code returned for type GetAPIIdentifierOK
const GetAPIIdentifierOKCode int = 200

/*GetAPIIdentifierOK Returns API version and running backend version

swagger:response getApiIdentifierOK
*/
type GetAPIIdentifierOK struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleMessageResponse `json:"body,omitempty"`
}

// NewGetAPIIdentifierOK creates GetAPIIdentifierOK with default headers values
func NewGetAPIIdentifierOK() *GetAPIIdentifierOK {

	return &GetAPIIdentifierOK{}
}

// WithPayload adds the payload to the get Api identifier o k response
func (o *GetAPIIdentifierOK) WithPayload(payload *models.SimpleMessageResponse) *GetAPIIdentifierOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api identifier o k response
func (o *GetAPIIdentifierOK) SetPayload(payload *models.SimpleMessageResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIIdentifierOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIIdentifierInternalServerErrorCode is the HTTP code returned for type GetAPIIdentifierInternalServerError
const GetAPIIdentifierInternalServerErrorCode int = 500

/*GetAPIIdentifierInternalServerError General Failure

swagger:response getApiIdentifierInternalServerError
*/
type GetAPIIdentifierInternalServerError struct {
}

// NewGetAPIIdentifierInternalServerError creates GetAPIIdentifierInternalServerError with default headers values
func NewGetAPIIdentifierInternalServerError() *GetAPIIdentifierInternalServerError {

	return &GetAPIIdentifierInternalServerError{}
}

// WriteResponse to the client
func (o *GetAPIIdentifierInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
