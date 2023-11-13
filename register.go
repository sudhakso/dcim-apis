package main

import (
	"log"
	"net/http"

	"github.com/dcim-apis/pkg/dcim"
)

// Logging handler
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handler.ServeHTTP(writer, request)
	})
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si dcim.ServerInterface) http.Handler {
	mw := dcim.ChiServerOptions{}
	mw.Middlewares = append(mw.Middlewares, Logger)
	return dcim.HandlerWithOptions(si, mw)
}
