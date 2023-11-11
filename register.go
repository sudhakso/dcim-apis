package main

import (
	"net/http"

	"github.com/dcim-apis/pkg/dcim"
)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si dcim.ServerInterface) http.Handler {

	return dcim.HandlerWithOptions(si, dcim.ChiServerOptions{})
}
