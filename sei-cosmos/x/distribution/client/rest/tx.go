package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

type (
	withdrawRewardsReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	}

	setWithdrawalAddrReq struct {
		BaseReq         rest.BaseReq   `json:"base_req" yaml:"base_req"`
		WithdrawAddress sdk.AccAddress `json:"withdraw_address" yaml:"withdraw_address"`
	}

	fundCommunityPoolReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
		Amount  sdk.Coins    `json:"amount" yaml:"amount"`
	}
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	_ = "STUB: not implemented"
	// Withdraw all delegator rewards
	return
}

// Withdraw delegation rewards

// Replace the rewards withdrawal address

// Withdraw validator rewards and commission

// Fund the community pool

func newWithdrawDelegatorRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// read and validate URL's variables

func newWithdrawDelegationRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// read and validate URL's variables

func newSetDelegatorWithdrawalAddrHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// read and validate URL's variables

func newWithdrawValidatorRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// read and validate URL's variable

// prepare multi-message transaction

func newFundCommunityPoolHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Auxiliary

func checkDelegatorAddressVar(w http.ResponseWriter, r *http.Request) (sdk.AccAddress, bool) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), false
}

func checkValidatorAddressVar(w http.ResponseWriter, r *http.Request) (sdk.ValAddress, bool) {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress), false
}
