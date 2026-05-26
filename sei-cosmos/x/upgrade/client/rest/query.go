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

func getCurrentPlanHandler(clientCtx client.Context) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

// ignore height for now

func getDonePlanHandler(clientCtx client.Context) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // stored by SetDone from block heights which are always non-negative
