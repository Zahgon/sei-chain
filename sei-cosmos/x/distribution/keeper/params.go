package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
)

// GetParams returns the total set of distribution parameters.
func (k Keeper) GetParams(clientCtx sdk.Context) (params types.Params) {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the distribution parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }

// GetCommunityTax returns the current distribution community tax.
func (k Keeper) GetCommunityTax(ctx sdk.Context) (percent sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GetBaseProposerReward returns the current distribution base proposer rate.
func (k Keeper) GetBaseProposerReward(ctx sdk.Context) (percent sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GetBonusProposerReward returns the current distribution bonus proposer reward
// rate.
func (k Keeper) GetBonusProposerReward(ctx sdk.Context) (percent sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GetWithdrawAddrEnabled returns the current distribution withdraw address
// enabled parameter.
func (k Keeper) GetWithdrawAddrEnabled(ctx sdk.Context) (enabled bool) {
	_ = "STUB: not implemented"
	return false
}
