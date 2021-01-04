package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/seanrmurphy/lambda-swagger-test/restapi/operations/open"
)

func GetApiIdentifier(params open.GetAPIIdentifierParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetNoParamsEmptyResponse(params open.GetNoParamsEmptyResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetNoParamsSimpleResponse(params open.GetNoParamsSimpleResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetNoParamsComplexResponse(params open.GetNoParamsComplexResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetNoParamsErrorResponse(params open.GetNoParamsErrorResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetPathParamEmptyResponse(params open.GetPathParamEmptyResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetPathParamSimpleResponse(params open.GetPathParamSimpleResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetPathParamComplexResponse(params open.GetPathParamComplexResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetPathParamErrorResponse(params open.GetPathParamErrorResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetBodyParamEmptyResponse(params open.GetBodyParamEmptyResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func PostBodyParamEmptyResponse(params open.PostBodyParamEmptyResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetBodyParamSimpleResponse(params open.GetBodyParamSimpleResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func PostBodyParamSimpleResponse(params open.PostBodyParamSimpleResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetBodyParamComplexResponse(params open.GetBodyParamComplexResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func PostBodyParamComplexResponse(params open.PostBodyParamComplexResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetBodyParamErrorResponse(params open.GetBodyParamErrorResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func PostBodyParamErrorResponse(params open.PostBodyParamErrorResponseParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

//func TestEndpoint(params open.GetTestEndpointParams) middleware.Responder {

//str := "Success from Test Endpoint"
//r := models.TestResponse{
//Message: &str,
//}

//resp := open.NewGetTestEndpointOK().WithPayload(&r)

//return resp

//}

//func AlternativeTestEndpoint(params open.GetAlternativeTestEndpointParams) middleware.Responder {

//str := "Success from Alternative Test Endpoint"

//r := models.TestResponse{
//Message: &str,
//}

//resp := open.NewGetAlternativeTestEndpointOK().WithPayload(&r)

//return resp
//}

//func TestWithParametersEndpoint(params open.GetTestWithParameterEndpointParams) middleware.Responder {

//param := params.Param
//str := "Success from Test With Params Endpoint - param = " + param

//r := models.TestResponse{
//Message: &str,
//}

//resp := open.NewGetTestWithParameterEndpointOK().WithPayload(&r)

//return resp
//}
