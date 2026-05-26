package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) { _ = "STUB: not implemented"; return }

type (
	// DelegateRequest defines the properties of a delegation request's body.
	DelegateRequest struct {
		BaseReq          rest.BaseReq   `json:"base_req" yaml:"base_req"`
		DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"` // in bech32
		ValidatorAddress sdk.ValAddress `json:"validator_address" yaml:"validator_address"` // in bech32
		Amount           sdk.Coin       `json:"amount" yaml:"amount"`
	}

	// RedelegateRequest defines the properties of a redelegate request's body.
	RedelegateRequest struct {
		BaseReq             rest.BaseReq   `json:"base_req" yaml:"base_req"`
		DelegatorAddress    sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`         // in bech32
		ValidatorSrcAddress sdk.ValAddress `json:"validator_src_address" yaml:"validator_src_address"` // in bech32
		ValidatorDstAddress sdk.ValAddress `json:"validator_dst_address" yaml:"validator_dst_address"` // in bech32
		Amount              sdk.Coin       `json:"amount" yaml:"amount"`
	}

	// UndelegateRequest defines the properties of a undelegate request's body.
	UndelegateRequest struct {
		BaseReq          rest.BaseReq   `json:"base_req" yaml:"base_req"`
		DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"` // in bech32
		ValidatorAddress sdk.ValAddress `json:"validator_address" yaml:"validator_address"` // in bech32
		Amount           sdk.Coin       `json:"amount" yaml:"amount"`
	}
)

func newPostDelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func newPostRedelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func newPostUnbondingDelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
