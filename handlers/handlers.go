package handlers

import (
	"fmt"
	"math/rand"

	"github.com/go-openapi/runtime/middleware"
	"github.com/seanrmurphy/lambda-swagger-test/models"
	"github.com/seanrmurphy/lambda-swagger-test/restapi/operations/open"
)

func GetApiIdentifier(params open.GetAPIIdentifierParams) middleware.Responder {
	resp := open.NewGetAPIIdentifierOK()

	return resp
}

func GetNoParamsEmptyResponse(params open.GetNoParamsEmptyResponseParams) middleware.Responder {
	resp := open.NewGetNoParamsEmptyResponseOK()

	return resp
}

func GetNoParamsSimpleResponse(params open.GetNoParamsSimpleResponseParams) middleware.Responder {
	str := "Response from GetNoParamsSimpleResponse function"

	r := models.SimpleMessageResponse{
		Message: &str,
	}
	resp := open.NewGetNoParamsSimpleResponseOK().WithPayload(&r)

	return resp
}

func GetNoParamsComplexResponse(params open.GetNoParamsComplexResponseParams) middleware.Responder {
	stringArray := []string{"param1", "param2"}
	descriptor := "A descriptive string (object1)"
	intVal := rand.Int63()
	o1 := models.SimpleObjectOne{
		StringArray: stringArray,
		Descriptor:  &descriptor,
		IntVal:      &intVal,
	}

	descriptor = "A descriptive string (object2)"
	intArray := []int64{rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63()}
	o2 := models.SimpleObjectTwo{
		Descriptor: &descriptor,
		IntArray:   intArray,
	}
	r := models.ComplexObjectResponse{
		Object1:       &o1,
		Object2:       &o2,
		OptionalParam: rand.Int63(),
	}
	resp := open.NewGetNoParamsComplexResponseOK().WithPayload(&r)

	return resp
}

func GetNoParamsErrorResponse(params open.GetNoParamsErrorResponseParams) middleware.Responder {
	str := "Calling this endpoint (GetNoParamsErrorResponse) deliberately generates an error case."
	randomInt := rand.Int63()

	r := models.ErrorResponse{
		ErrorNumber: &randomInt,
		ErrorString: &str,
	}
	resp := open.NewGetNoParamsErrorResponseIMATeapot().WithPayload(&r)

	return resp
}

func GetPathParamEmptyResponse(params open.GetPathParamEmptyResponseParams) middleware.Responder {
	resp := open.NewGetPathParamEmptyResponseOK()

	return resp
}

func GetPathParamSimpleResponse(params open.GetPathParamSimpleResponseParams) middleware.Responder {
	pathParam := params.Param
	str := fmt.Sprintf("Response from GetPathParamSimpleResponse function - input param = %v", pathParam)

	r := models.SimpleMessageResponse{
		Message: &str,
	}
	resp := open.NewGetPathParamSimpleResponseOK().WithPayload(&r)

	return resp
}

func GetPathParamComplexResponse(params open.GetPathParamComplexResponseParams) middleware.Responder {
	pathParam := params.Param
	str := fmt.Sprintf("Input param = %v", pathParam)

	stringArray := []string{"param1", "param2", str}
	descriptor := "A descriptive string (object1)"
	intVal := rand.Int63()
	o1 := models.SimpleObjectOne{
		StringArray: stringArray,
		Descriptor:  &descriptor,
		IntVal:      &intVal,
	}

	descriptor = "A descriptive string (object2)"
	intArray := []int64{rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63()}
	o2 := models.SimpleObjectTwo{
		Descriptor: &descriptor,
		IntArray:   intArray,
	}
	r := models.ComplexObjectResponse{
		Object1:       &o1,
		Object2:       &o2,
		OptionalParam: rand.Int63(),
	}
	resp := open.NewGetPathParamComplexResponseOK().WithPayload(&r)

	return resp
}

func GetPathParamErrorResponse(params open.GetPathParamErrorResponseParams) middleware.Responder {
	pathParam := params.Param
	str := fmt.Sprintf("Calling this endpoint (GetPathParamErrorResponse) deliberately generates an error case - Input param = %v", pathParam)
	randomInt := rand.Int63()

	r := models.ErrorResponse{
		ErrorNumber: &randomInt,
		ErrorString: &str,
	}
	resp := open.NewGetPathParamErrorResponseIMATeapot().WithPayload(&r)

	return resp
}

func GetBodyParamEmptyResponse(params open.GetBodyParamEmptyResponseParams) middleware.Responder {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int63()

	r := models.ErrorResponse{
		ErrorNumber: &randomInt,
		ErrorString: &str,
	}
	resp := open.NewGetBodyParamEmptyResponseNotImplemented().WithPayload(&r)

	return resp
}

func PostBodyParamEmptyResponse(params open.PostBodyParamEmptyResponseParams) middleware.Responder {
	// there are input parameters provided here, but since we cannot return them
	// in any way, we just ignore them
	resp := open.NewPostBodyParamEmptyResponseOK()

	return resp
}

func GetBodyParamSimpleResponse(params open.GetBodyParamSimpleResponseParams) middleware.Responder {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int63()

	r := models.ErrorResponse{
		ErrorNumber: &randomInt,
		ErrorString: &str,
	}
	resp := open.NewGetBodyParamSimpleResponseNotImplemented().WithPayload(&r)

	return resp
}

func PostBodyParamSimpleResponse(params open.PostBodyParamSimpleResponseParams) middleware.Responder {
	inputParams := params.BodyParam
	str := fmt.Sprintf("Response from PostBodyParamSimpleResponse function - called with JSON object {'descriptor': '%v', 'int_val': %v, 'string': '%v'}",
		inputParams.Descriptor, inputParams.IntVal, inputParams.String)

	r := models.SimpleMessageResponse{
		Message: &str,
	}
	resp := open.NewPostBodyParamSimpleResponseOK().WithPayload(&r)

	return resp
}

func GetBodyParamComplexResponse(params open.GetBodyParamComplexResponseParams) middleware.Responder {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int63()

	r := models.ErrorResponse{
		ErrorNumber: &randomInt,
		ErrorString: &str,
	}
	resp := open.NewGetBodyParamComplexResponseNotImplemented().WithPayload(&r)

	return resp
}

func PostBodyParamComplexResponse(params open.PostBodyParamComplexResponseParams) middleware.Responder {
	inputParams := params.BodyParam
	str := fmt.Sprintf("Response from PostBodyParamSimpleResponse function - called with JSON object {'descriptor': '%v', 'int_val': %v, 'string': '%v'}",
		inputParams.Descriptor, inputParams.IntVal, inputParams.String)

	stringArray := []string{"param1", "param2", str}
	descriptor := "A descriptive string (object1)"
	intVal := rand.Int63()
	o1 := models.SimpleObjectOne{
		StringArray: stringArray,
		Descriptor:  &descriptor,
		IntVal:      &intVal,
	}

	descriptor = "A descriptive string (object2)"
	intArray := []int64{rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63()}
	o2 := models.SimpleObjectTwo{
		Descriptor: &descriptor,
		IntArray:   intArray,
	}
	r := models.ComplexObjectResponse{
		Object1:       &o1,
		Object2:       &o2,
		OptionalParam: rand.Int63(),
	}
	resp := open.NewPostBodyParamComplexResponseOK().WithPayload(&r)

	return resp
}

func GetBodyParamErrorResponse(params open.GetBodyParamErrorResponseParams) middleware.Responder {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int63()

	r := models.ErrorResponse{
		ErrorNumber: &randomInt,
		ErrorString: &str,
	}
	resp := open.NewGetBodyParamErrorResponseNotImplemented().WithPayload(&r)

	return resp
}

func PostBodyParamErrorResponse(params open.PostBodyParamErrorResponseParams) middleware.Responder {
	inputParams := params.BodyParam
	str := fmt.Sprintf("Calling this endpoint (PostBodyParamErrorResponse) deliberately generates an error case - called with JSON object {'descriptor': '%v', 'int_val': %v, 'string': '%v'}",
		inputParams.Descriptor, inputParams.IntVal, inputParams.String)
	randomInt := rand.Int63()

	r := models.ErrorResponse{
		ErrorNumber: &randomInt,
		ErrorString: &str,
	}
	resp := open.NewPostBodyParamErrorResponseIMATeapot().WithPayload(&r)

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
