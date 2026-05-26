package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	//nolint:staticcheck
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	_ = "STUB: not implemented"
	return
}

// Deprecated: http request handler to query signing info
func signingInfoHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// http request handler to query signing info
func signingInfoHandlerListFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryParamsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
