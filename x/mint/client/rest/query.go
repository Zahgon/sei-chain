package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	_ = "STUB: not implemented"
	return
}

func queryParamsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryMinterHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
