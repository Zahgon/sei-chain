package rest

import (
	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

// REST query and parameter values
const (
	MethodGet = "GET"
)

// RegisterRoutes registers the auth module REST routes.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router, storeName string) {
	_ = "STUB: not implemented"
	return
}

// RegisterTxRoutes registers all transaction routes on the provided router.
func RegisterTxRoutes(clientCtx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }
