package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	govrest "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client/rest"
	"github.com/sei-protocol/sei-chain/x/mint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	typesrest "github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) { _ = "STUB: not implemented"; return }

// PlanRequest defines a proposal for a new upgrade plan.
type UpdateMinterRequest struct {
	BaseReq     typesrest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string            `json:"title" yaml:"title"`
	Description string            `json:"description" yaml:"description"`
	Deposit     sdk.Coins         `json:"deposit" yaml:"deposit"`
	Minter      types.Minter      `json:"minter" yaml:"minter"`
}

func UpdateResourceDependencyProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

func newUpdateMinterPostHandler(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
