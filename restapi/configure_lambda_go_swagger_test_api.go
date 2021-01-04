// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/seanrmurphy/lambda-swagger-test/restapi/operations"
	"github.com/seanrmurphy/lambda-swagger-test/restapi/operations/open"
)

//go:generate swagger generate server --target ../../lambda-swagger-test --name LambdaGoSwaggerTestAPI --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.LambdaGoSwaggerTestAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LambdaGoSwaggerTestAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.OpenGetAPIIdentifierHandler == nil {
		api.OpenGetAPIIdentifierHandler = open.GetAPIIdentifierHandlerFunc(func(params open.GetAPIIdentifierParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetAPIIdentifier has not yet been implemented")
		})
	}
	if api.OpenGetBodyParamComplexResponseHandler == nil {
		api.OpenGetBodyParamComplexResponseHandler = open.GetBodyParamComplexResponseHandlerFunc(func(params open.GetBodyParamComplexResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetBodyParamComplexResponse has not yet been implemented")
		})
	}
	if api.OpenGetBodyParamEmptyResponseHandler == nil {
		api.OpenGetBodyParamEmptyResponseHandler = open.GetBodyParamEmptyResponseHandlerFunc(func(params open.GetBodyParamEmptyResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetBodyParamEmptyResponse has not yet been implemented")
		})
	}
	if api.OpenGetBodyParamErrorResponseHandler == nil {
		api.OpenGetBodyParamErrorResponseHandler = open.GetBodyParamErrorResponseHandlerFunc(func(params open.GetBodyParamErrorResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetBodyParamErrorResponse has not yet been implemented")
		})
	}
	if api.OpenGetBodyParamSimpleResponseHandler == nil {
		api.OpenGetBodyParamSimpleResponseHandler = open.GetBodyParamSimpleResponseHandlerFunc(func(params open.GetBodyParamSimpleResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetBodyParamSimpleResponse has not yet been implemented")
		})
	}
	if api.OpenGetNoParamsComplexResponseHandler == nil {
		api.OpenGetNoParamsComplexResponseHandler = open.GetNoParamsComplexResponseHandlerFunc(func(params open.GetNoParamsComplexResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetNoParamsComplexResponse has not yet been implemented")
		})
	}
	if api.OpenGetNoParamsEmptyResponseHandler == nil {
		api.OpenGetNoParamsEmptyResponseHandler = open.GetNoParamsEmptyResponseHandlerFunc(func(params open.GetNoParamsEmptyResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetNoParamsEmptyResponse has not yet been implemented")
		})
	}
	if api.OpenGetNoParamsErrorResponseHandler == nil {
		api.OpenGetNoParamsErrorResponseHandler = open.GetNoParamsErrorResponseHandlerFunc(func(params open.GetNoParamsErrorResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetNoParamsErrorResponse has not yet been implemented")
		})
	}
	if api.OpenGetNoParamsSimpleResponseHandler == nil {
		api.OpenGetNoParamsSimpleResponseHandler = open.GetNoParamsSimpleResponseHandlerFunc(func(params open.GetNoParamsSimpleResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetNoParamsSimpleResponse has not yet been implemented")
		})
	}
	if api.OpenGetPathParamComplexResponseHandler == nil {
		api.OpenGetPathParamComplexResponseHandler = open.GetPathParamComplexResponseHandlerFunc(func(params open.GetPathParamComplexResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetPathParamComplexResponse has not yet been implemented")
		})
	}
	if api.OpenGetPathParamEmptyResponseHandler == nil {
		api.OpenGetPathParamEmptyResponseHandler = open.GetPathParamEmptyResponseHandlerFunc(func(params open.GetPathParamEmptyResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetPathParamEmptyResponse has not yet been implemented")
		})
	}
	if api.OpenGetPathParamErrorResponseHandler == nil {
		api.OpenGetPathParamErrorResponseHandler = open.GetPathParamErrorResponseHandlerFunc(func(params open.GetPathParamErrorResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetPathParamErrorResponse has not yet been implemented")
		})
	}
	if api.OpenGetPathParamSimpleResponseHandler == nil {
		api.OpenGetPathParamSimpleResponseHandler = open.GetPathParamSimpleResponseHandlerFunc(func(params open.GetPathParamSimpleResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.GetPathParamSimpleResponse has not yet been implemented")
		})
	}
	if api.OpenPostBodyParamComplexResponseHandler == nil {
		api.OpenPostBodyParamComplexResponseHandler = open.PostBodyParamComplexResponseHandlerFunc(func(params open.PostBodyParamComplexResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.PostBodyParamComplexResponse has not yet been implemented")
		})
	}
	if api.OpenPostBodyParamEmptyResponseHandler == nil {
		api.OpenPostBodyParamEmptyResponseHandler = open.PostBodyParamEmptyResponseHandlerFunc(func(params open.PostBodyParamEmptyResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.PostBodyParamEmptyResponse has not yet been implemented")
		})
	}
	if api.OpenPostBodyParamErrorResponseHandler == nil {
		api.OpenPostBodyParamErrorResponseHandler = open.PostBodyParamErrorResponseHandlerFunc(func(params open.PostBodyParamErrorResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.PostBodyParamErrorResponse has not yet been implemented")
		})
	}
	if api.OpenPostBodyParamSimpleResponseHandler == nil {
		api.OpenPostBodyParamSimpleResponseHandler = open.PostBodyParamSimpleResponseHandlerFunc(func(params open.PostBodyParamSimpleResponseParams) middleware.Responder {
			return middleware.NotImplemented("operation open.PostBodyParamSimpleResponse has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
