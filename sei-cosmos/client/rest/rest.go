package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DeprecationURL is the URL for migrating deprecated REST endpoints to newer ones.
// TODO Switch to `/` (not `/master`) once v0.40 docs are deployed.
// https://github.com/cosmos/cosmos-sdk/issues/8019
const DeprecationURL = "https://docs.cosmos.network/master/migrations/rest.html"

// addHTTPDeprecationHeaders is a mux middleware function for adding HTTP
// Deprecation headers to a http handler
func addHTTPDeprecationHeaders(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// WithHTTPDeprecationHeaders returns a new *mux.Router, identical to its input
// but with the addition of HTTP Deprecation headers. This is used to mark legacy
// amino REST endpoints as deprecated in the REST API.
// nolint: gocritic
func WithHTTPDeprecationHeaders(r *mux.Router) *mux.Router { _ = "STUB: not implemented"; return nil }
