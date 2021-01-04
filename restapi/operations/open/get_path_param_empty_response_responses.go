// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetPathParamEmptyResponseOKCode is the HTTP code returned for type GetPathParamEmptyResponseOK
const GetPathParamEmptyResponseOKCode int = 200

/*GetPathParamEmptyResponseOK Successful execution of API call

swagger:response getPathParamEmptyResponseOK
*/
type GetPathParamEmptyResponseOK struct {
}

// NewGetPathParamEmptyResponseOK creates GetPathParamEmptyResponseOK with default headers values
func NewGetPathParamEmptyResponseOK() *GetPathParamEmptyResponseOK {

	return &GetPathParamEmptyResponseOK{}
}

// WriteResponse to the client
func (o *GetPathParamEmptyResponseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetPathParamEmptyResponseInternalServerErrorCode is the HTTP code returned for type GetPathParamEmptyResponseInternalServerError
const GetPathParamEmptyResponseInternalServerErrorCode int = 500

/*GetPathParamEmptyResponseInternalServerError General Failure

swagger:response getPathParamEmptyResponseInternalServerError
*/
type GetPathParamEmptyResponseInternalServerError struct {
}

// NewGetPathParamEmptyResponseInternalServerError creates GetPathParamEmptyResponseInternalServerError with default headers values
func NewGetPathParamEmptyResponseInternalServerError() *GetPathParamEmptyResponseInternalServerError {

	return &GetPathParamEmptyResponseInternalServerError{}
}

// WriteResponse to the client
func (o *GetPathParamEmptyResponseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
