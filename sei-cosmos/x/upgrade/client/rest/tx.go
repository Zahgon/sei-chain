package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	govrest "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client/rest"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

func registerTxHandlers(
	clientCtx client.Context,
	r *mux.Router) {
	_ = "STUB: not implemented"
	return
}

// PlanRequest defines a proposal for a new upgrade plan.
type PlanRequest struct {
	BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title         string       `json:"title" yaml:"title"`
	Description   string       `json:"description" yaml:"description"`
	IsExpedited   bool         `json:"is_expedited" yaml:"is_expedited"`
	Deposit       sdk.Coins    `json:"deposit" yaml:"deposit"`
	UpgradeName   string       `json:"upgrade_name" yaml:"upgrade_name"`
	UpgradeHeight int64        `json:"upgrade_height" yaml:"upgrade_height"`
	UpgradeInfo   string       `json:"upgrade_info" yaml:"upgrade_info"`
}

// CancelRequest defines a proposal to cancel a current plan.
type CancelRequest struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	IsExpedited bool         `json:"is_expedited" yaml:"is_expedited"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
}

func ProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

func ProposalCancelRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}

func newPostPlanHandler(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func newCancelPlanHandler(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
