package rest

import (
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	govrest "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client/rest"
)

// ProposalRESTHandler returns a ProposalRESTHandler that exposes the param
// change REST handler with a given sub-route.
func ProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

func postProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
