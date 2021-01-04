// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostBodyParamEmptyResponseOKCode is the HTTP code returned for type PostBodyParamEmptyResponseOK
const PostBodyParamEmptyResponseOKCode int = 200

/*PostBodyParamEmptyResponseOK Successful execution of API call

swagger:response postBodyParamEmptyResponseOK
*/
type PostBodyParamEmptyResponseOK struct {
}

// NewPostBodyParamEmptyResponseOK creates PostBodyParamEmptyResponseOK with default headers values
func NewPostBodyParamEmptyResponseOK() *PostBodyParamEmptyResponseOK {

	return &PostBodyParamEmptyResponseOK{}
}

// WriteResponse to the client
func (o *PostBodyParamEmptyResponseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostBodyParamEmptyResponseInternalServerErrorCode is the HTTP code returned for type PostBodyParamEmptyResponseInternalServerError
const PostBodyParamEmptyResponseInternalServerErrorCode int = 500

/*PostBodyParamEmptyResponseInternalServerError General Failure

swagger:response postBodyParamEmptyResponseInternalServerError
*/
type PostBodyParamEmptyResponseInternalServerError struct {
}

// NewPostBodyParamEmptyResponseInternalServerError creates PostBodyParamEmptyResponseInternalServerError with default headers values
func NewPostBodyParamEmptyResponseInternalServerError() *PostBodyParamEmptyResponseInternalServerError {

	return &PostBodyParamEmptyResponseInternalServerError{}
}

// WriteResponse to the client
func (o *PostBodyParamEmptyResponseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}