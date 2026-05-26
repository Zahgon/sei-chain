package keeper

import (
	"github.com/sei-protocol/sei-chain/x/oracle/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// VotePeriod returns the number of blocks during which voting takes place.
func (k Keeper) VotePeriod(ctx sdk.Context) (res uint64) { _ = "STUB: not implemented"; return 0 }

// VoteThreshold returns the minimum percentage of votes that must be received for a ballot to pass.
func (k Keeper) VoteThreshold(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// RewardBand returns the ratio of allowable exchange rate error that a validator can be rewared
func (k Keeper) RewardBand(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Whitelist returns the denom list that can be activated
func (k Keeper) Whitelist(ctx sdk.Context) (res types.DenomList) {
	_ = "STUB: not implemented"
	return *new(types.DenomList)
}

// SetWhitelist store new whitelist to param store
// this function is only for test purpose
func (k Keeper) SetWhitelist(ctx sdk.Context, whitelist types.DenomList) {
	_ = "STUB: not implemented"
	return
}

// SlashFraction returns oracle voting penalty rate
func (k Keeper) SlashFraction(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// SlashWindow returns # of vote period for oracle slashing
func (k Keeper) SlashWindow(ctx sdk.Context) (res uint64) { _ = "STUB: not implemented"; return 0 }

// MinValidPerWindow returns oracle slashing threshold
func (k Keeper) MinValidPerWindow(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

func (k Keeper) LookbackDuration(ctx sdk.Context) (res uint64) { _ = "STUB: not implemented"; return 0 }

// GetParams returns the total set of oracle parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the total set of oracle parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
