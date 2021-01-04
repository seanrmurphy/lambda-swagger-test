package main

import (
	"log"

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

func init() {
	// Initiate business logic implementers.
	// This is the main function, so here the implementers' dependencies can be
	// injected, such as database, parameters from environment variables, or different
	// clients for different APIs.

	//tableName := os.Getenv("TABLE_NAME")

	//d := db.StatusDb{TableName: tableName}

	//p := apiimplementation.New(d)

	// Initiate the http handler, with the objects that are implementing the business logic.
	//h, err := restapi.Handler(restapi.Config{
	//OpenAPI: p,
	//Logger:  log.Printf,
	//})
	//if err != nil {
	//log.Fatal(err)
	//}

	// Run the standard http server
	// log.Fatal(http.ListenAndServe(":8080", h))

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
	api.OpenGetBodyParamEmptyResponseHandler = open.GetBodyParamEmptyResponseHandlerFunc(handlers.GetBodyParamEmptyResponse)
	api.OpenPostBodyParamEmptyResponseHandler = open.PostBodyParamEmptyResponseHandlerFunc(handlers.PostBodyParamEmptyResponse)
	api.OpenGetBodyParamSimpleResponseHandler = open.GetBodyParamSimpleResponseHandlerFunc(handlers.GetBodyParamSimpleResponse)
	api.OpenPostBodyParamSimpleResponseHandler = open.PostBodyParamSimpleResponseHandlerFunc(handlers.PostBodyParamSimpleResponse)
	api.OpenGetBodyParamComplexResponseHandler = open.GetBodyParamComplexResponseHandlerFunc(handlers.GetBodyParamComplexResponse)
	api.OpenPostBodyParamComplexResponseHandler = open.PostBodyParamComplexResponseHandlerFunc(handlers.PostBodyParamComplexResponse)
	api.OpenGetBodyParamErrorResponseHandler = open.GetBodyParamErrorResponseHandlerFunc(handlers.GetBodyParamErrorResponse)
	api.OpenPostBodyParamErrorResponseHandler = open.PostBodyParamErrorResponseHandlerFunc(handlers.PostBodyParamErrorResponse)

	server := restapi.NewServer(api)
	server.ConfigureAPI()

	//log.Printf("handler = %v", reflect.TypeOf(h).String())
	httpAdapter = httpadapter.New(server.GetHandler())
}

// Handler handles API requests
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("In new2 handler - endpoint %v...ready to go...", req.Path)
	return httpAdapter.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
