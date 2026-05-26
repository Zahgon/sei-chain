package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func registerQueryRoutes(cliCtx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }

func queryExchangeRateHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryExchangeRatesHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryPriceSnapshotHistoryHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryTwapsHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryActivesHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryParamsHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryFeederDelegationHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryVotePenaltyCounterHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryVoteTargetsHandlerFunction(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func checkDenomVar(w http.ResponseWriter, r *http.Request) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func checkLookbackSecondsVar(r *http.Request) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func checkVoterAddressVar(w http.ResponseWriter, r *http.Request) (sdk.ValAddress, bool) {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress), false
}
