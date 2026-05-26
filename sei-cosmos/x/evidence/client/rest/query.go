package rest

import (
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"

	"github.com/gorilla/mux"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	_ = "STUB: not implemented"
	return
}

func queryEvidenceHandler(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryAllEvidenceHandler(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
