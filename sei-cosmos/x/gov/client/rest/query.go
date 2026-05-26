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

func queryParamsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryDepositsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// For inactive proposals we must query the txs directly to get the deposits
// as they're no longer in state.

func queryProposerHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryDepositHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// For an empty deposit, either the proposal does not exist or is inactive in
// which case the deposit would be removed from state and should be queried
// for directly via a txs query.

func queryVoteHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// For an empty vote, either the proposal does not exist or is inactive in
// which case the vote would be removed from state and should be queried for
// directly via a txs query.

// todo: Split this functionality into helper functions to remove the above
func queryVotesOnProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// For inactive proposals we must query the txs directly to get the votes
// as they're no longer in state.

// HTTP request handler to query list of governance proposals
func queryProposalsWithParameterFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// todo: Split this functionality into helper functions to remove the above
func queryTallyOnProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
