package rest

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"

	"github.com/gorilla/mux"
)

func registerTxRoutes(clientCtx client.Context, r *mux.Router, handlers []EvidenceRESTHandler) {
	_ = "STUB: not implemented"
	// TODO: Register tx handlers.
	return
}
