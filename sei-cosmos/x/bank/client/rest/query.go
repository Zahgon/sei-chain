package rest

import (
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

// QueryBalancesRequestHandlerFn returns a REST handler that queries for all
// account balances or a specific balance by denomination.
func QueryBalancesRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query the total supply of coins
func totalSupplyHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query the supply of a single denom
func supplyOfHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
