package rest

import (
	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

// RegisterHandlers registers all x/bank transaction and query HTTP REST handlers
// on the provided mux router.
func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }
