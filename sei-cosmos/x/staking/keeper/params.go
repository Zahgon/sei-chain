package keeper

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// UnbondingTime
func (k Keeper) UnbondingTime(ctx sdk.Context) (res time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// MaxValidators - Maximum number of validators
func (k Keeper) MaxValidators(ctx sdk.Context) (res uint32) { _ = "STUB: not implemented"; return 0 }

// MaxEntries - Maximum number of simultaneous unbonding
// delegations or redelegations (per pair/trio)
func (k Keeper) MaxEntries(ctx sdk.Context) (res uint32) { _ = "STUB: not implemented"; return 0 }

// HistoricalEntries = number of historical info entries
// to persist in store
func (k Keeper) HistoricalEntries(ctx sdk.Context) (res uint32) {
	_ = "STUB: not implemented"
	return 0
}

// BondDenom - Bondable coin denomination
func (k Keeper) BondDenom(ctx sdk.Context) (res string) { _ = "STUB: not implemented"; return "" }

// MaxVotingPowerRatio - maximal allowed voting power ratio of a validator
func (k Keeper) MaxVotingPowerRatio(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// MaxVotingPowerEnforcementThreshold - minimal bonded voting power of the max voting power ratio enforcement
func (k Keeper) MaxVotingPowerEnforcementThreshold(ctx sdk.Context) (res sdk.Int) {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// PowerReduction - is the amount of staking tokens required for 1 unit of consensus-engine power.
// Currently, this returns a global variable that the app developer can tweak.
// TODO: we might turn this into an on-chain param:
// https://github.com/cosmos/cosmos-sdk/issues/8365
func (k Keeper) PowerReduction(ctx sdk.Context) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// MinCommissionRate - Minimum validator commission rate
func (k Keeper) MinCommissionRate(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
