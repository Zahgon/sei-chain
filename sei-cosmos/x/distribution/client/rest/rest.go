package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	govrest "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client/rest"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }

// TODO add proto compatible Handler after x/gov migration
// ProposalRESTHandler returns a ProposalRESTHandler that exposes the community pool spend REST handler with a given sub-route.
func ProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

func postProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
