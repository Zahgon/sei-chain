package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	_ = "STUB: not implemented"
	// Get all delegations from a delegator
	return
}

// Get all unbonding delegations from a delegator

// Get all staking txs (i.e msgs) from a delegator

// Query all validators that a delegator is bonded to

// Query a validator that a delegator is bonded to

// Query a delegation between a delegator and a validator

// Query all unbonding delegations between a delegator and a validator

// Query redelegations (filters in query params)

// Get all validators

// Get a single validator info

// Get all delegations to a validator

// Get all unbonding delegations from a validator

// Get HistoricalInfo at a given height

// Get the current state of the staking pool

// Get the current staking parameter values

// HTTP request handler to query a delegator delegations
func delegatorDelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query a delegator unbonding delegations
func delegatorUnbondingDelegationsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query all staking txs (msgs) from a delegator
func delegatorTxsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// For each case, we search txs for both:
// - legacy messages: their Type() is a custom string, e.g. "delegate"
// - service Msgs: their Type() is their FQ method name, e.g. "/cosmos.staking.v1beta1.MsgDelegate"
// and we combine the results.

// HTTP request handler to query an unbonding-delegation
func unbondingDelegationHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query redelegations
func redelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query a delegation
func delegationHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query all delegator bonded validators
func delegatorValidatorsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to get information from a currently bonded validator
func delegatorValidatorHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query list of validators
func validatorsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// These are query params that were available in =<0.39. We show a nice
// error message for this breaking change.

// HTTP request handler to query the validator information from a given validator address
func validatorHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query all unbonding delegations from a validator
func validatorDelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query all unbonding delegations from a validator
func validatorUnbondingDelegationsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query historical info at a given height
func historicalInfoHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query the pool information
func poolHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query the staking params values
func paramsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
