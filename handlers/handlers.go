package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/seanrmurphy/lambda-swagger-test/models"
	"github.com/seanrmurphy/lambda-swagger-test/restapi/operations/open"
)

func TestEndpoint(params open.GetTestEndpointParams) middleware.Responder {

	str := "Success from Test Endpoint"
	r := models.TestResponse{
		Message: &str,
	}

	resp := open.NewGetTestEndpointOK().WithPayload(&r)

	return resp

}

func AlternativeTestEndpoint(params open.GetAlternativeTestEndpointParams) middleware.Responder {

	str := "Success from Alternative Test Endpoint"

	r := models.TestResponse{
		Message: &str,
	}

	resp := open.NewGetAlternativeTestEndpointOK().WithPayload(&r)

	return resp
}

func TestWithParametersEndpoint(params open.GetTestWithParameterEndpointParams) middleware.Responder {

	param := params.Param
	str := "Success from Test With Params Endpoint - param = " + param

	r := models.TestResponse{
		Message: &str,
	}

	resp := open.NewGetTestWithParameterEndpointOK().WithPayload(&r)

	return resp
}
