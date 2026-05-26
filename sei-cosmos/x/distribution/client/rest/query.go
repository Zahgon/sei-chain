package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	_ = "STUB: not implemented"
	// Get the total rewards balance from all delegations
	return
}

// Query a delegation reward

// Get the rewards withdrawal address

// Validator distribution information

// Commission and self-delegation rewards of a single a validator

// Outstanding rewards of a single validator

// Get the current distribution parameter values

// Get the amount held in the community pool

// HTTP request handler to query the total rewards balance from all delegations
func delegatorRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query a delegation rewards
func delegationRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// query for rewards from a particular delegation

// HTTP request handler to query a delegation rewards
func delegatorWithdrawalAddrHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// ValidatorDistInfo defines the properties of
// validator distribution information response.
type ValidatorDistInfo struct {
	OperatorAddress     sdk.AccAddress                       `json:"operator_address" yaml:"operator_address"`
	SelfBondRewards     sdk.DecCoins                         `json:"self_bond_rewards" yaml:"self_bond_rewards"`
	ValidatorCommission types.ValidatorAccumulatedCommission `json:"val_commission" yaml:"val_commission"`
}

// NewValidatorDistInfo creates a new instance of ValidatorDistInfo.
func NewValidatorDistInfo(operatorAddr sdk.AccAddress, rewards sdk.DecCoins,
	commission types.ValidatorAccumulatedCommission) ValidatorDistInfo {
	_ = "STUB: not implemented"
	return *new(ValidatorDistInfo)
}

// HTTP request handler to query validator's distribution information
func validatorInfoHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// query commission

// self bond rewards

// HTTP request handler to query validator's commission and self-delegation rewards
func validatorRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query the distribution params values
func paramsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func communityPoolHandler(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// HTTP request handler to query the outstanding rewards
func outstandingRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func checkResponseQueryDelegationRewards(
	w http.ResponseWriter, clientCtx client.Context, delAddr, valAddr string,
) (res []byte, height int64, ok bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}
