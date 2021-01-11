package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	loads "github.com/go-openapi/loads"
	"github.com/seanrmurphy/lambda-swagger-test/handlers"
	"github.com/seanrmurphy/lambda-swagger-test/restapi"
	"github.com/seanrmurphy/lambda-swagger-test/restapi/operations"
	"github.com/seanrmurphy/lambda-swagger-test/restapi/operations/open"
)

var httpAdapter *httpadapter.HandlerAdapter
var nullHandler = false

func setupHandlers() *operations.LambdaGoSwaggerTestAPIAPI {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	api := operations.NewLambdaGoSwaggerTestAPIAPI(swaggerSpec)
	api.OpenGetAPIIdentifierHandler = open.GetAPIIdentifierHandlerFunc(handlers.GetApiIdentifier)
	api.OpenGetNoParamsEmptyResponseHandler = open.GetNoParamsEmptyResponseHandlerFunc(handlers.GetNoParamsEmptyResponse)
	api.OpenGetNoParamsSimpleResponseHandler = open.GetNoParamsSimpleResponseHandlerFunc(handlers.GetNoParamsSimpleResponse)
	api.OpenGetNoParamsComplexResponseHandler = open.GetNoParamsComplexResponseHandlerFunc(handlers.GetNoParamsComplexResponse)
	api.OpenGetNoParamsErrorResponseHandler = open.GetNoParamsErrorResponseHandlerFunc(handlers.GetNoParamsErrorResponse)
	api.OpenGetPathParamEmptyResponseHandler = open.GetPathParamEmptyResponseHandlerFunc(handlers.GetPathParamEmptyResponse)
	api.OpenGetPathParamSimpleResponseHandler = open.GetPathParamSimpleResponseHandlerFunc(handlers.GetPathParamSimpleResponse)
	api.OpenGetPathParamComplexResponseHandler = open.GetPathParamComplexResponseHandlerFunc(handlers.GetPathParamComplexResponse)
	api.OpenGetPathParamErrorResponseHandler = open.GetPathParamErrorResponseHandlerFunc(handlers.GetPathParamErrorResponse)
	api.OpenGetQueryParamEmptyResponseHandler = open.GetQueryParamEmptyResponseHandlerFunc(handlers.GetQueryParamEmptyResponse)
	api.OpenGetQueryParamSimpleResponseHandler = open.GetQueryParamSimpleResponseHandlerFunc(handlers.GetQueryParamSimpleResponse)
	api.OpenGetQueryParamComplexResponseHandler = open.GetQueryParamComplexResponseHandlerFunc(handlers.GetQueryParamComplexResponse)
	api.OpenGetQueryParamErrorResponseHandler = open.GetQueryParamErrorResponseHandlerFunc(handlers.GetQueryParamErrorResponse)
	api.OpenGetBodyParamEmptyResponseHandler = open.GetBodyParamEmptyResponseHandlerFunc(handlers.GetBodyParamEmptyResponse)
	api.OpenPostBodyParamEmptyResponseHandler = open.PostBodyParamEmptyResponseHandlerFunc(handlers.PostBodyParamEmptyResponse)
	api.OpenGetBodyParamSimpleResponseHandler = open.GetBodyParamSimpleResponseHandlerFunc(handlers.GetBodyParamSimpleResponse)
	api.OpenPostBodyParamSimpleResponseHandler = open.PostBodyParamSimpleResponseHandlerFunc(handlers.PostBodyParamSimpleResponse)
	api.OpenGetBodyParamComplexResponseHandler = open.GetBodyParamComplexResponseHandlerFunc(handlers.GetBodyParamComplexResponse)
	api.OpenPostBodyParamComplexResponseHandler = open.PostBodyParamComplexResponseHandlerFunc(handlers.PostBodyParamComplexResponse)
	api.OpenGetBodyParamErrorResponseHandler = open.GetBodyParamErrorResponseHandlerFunc(handlers.GetBodyParamErrorResponse)
	api.OpenPostBodyParamErrorResponseHandler = open.PostBodyParamErrorResponseHandlerFunc(handlers.PostBodyParamErrorResponse)

	return api
}

func init() {

	api := setupHandlers()
	server := restapi.NewServer(api)
	server.ConfigureAPI()

	httpAdapter = httpadapter.New(server.GetHandler())
}

// Handler handles API requests
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if nullHandler {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       `{"message": "Null handler operational - always returns this message"}`,
		}, nil
	} else {
		return httpAdapter.Proxy(req)
	}
}

func main() {
	lambda.Start(Handler)
}
