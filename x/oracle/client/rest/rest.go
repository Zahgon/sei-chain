package rest

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"

	"github.com/gorilla/mux"
)

const (
	RestDenom           = "denom"
	RestVoter           = "voter"
	RestLookbackSeconds = "lookback_seconds"
)

// RegisterRoutes registers oracle-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }
