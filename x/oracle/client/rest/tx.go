package rest

import (
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"

	"github.com/gorilla/mux"
)

func registerTxHandlers(cliCtx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }

type (
	delegateReq struct {
		BaseReq rest.BaseReq   `json:"base_req" yaml:"base_req"`
		Feeder  sdk.AccAddress `json:"feeder" yaml:"feeder"`
	}

	aggregateVoteReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		ExchangeRates string `json:"exchange_rates" yaml:"exchange_rates"`
	}
)

func newDelegateHandlerFunction(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// create the message

func newAggregateVoteHandlerFunction(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Check validation of tuples

// create the message
