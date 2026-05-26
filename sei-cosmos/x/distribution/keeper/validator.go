package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val stakingtypes.ValidatorI) {
	_ = "STUB: not implemented"
	// set initial historical rewards (period 0) with reference count of 1
	return
}

// set current rewards (starting at period 1)

// set accumulated commission

// set outstanding rewards

// increment validator period, returning the period just ended
func (k Keeper) IncrementValidatorPeriod(ctx sdk.Context, val stakingtypes.ValidatorI) uint64 {
	_ = "STUB: not implemented"
	// fetch current rewards
	return 0
}

// calculate current ratio

// can't calculate ratio for zero-token validators
// ergo we instead add to the community pool

// note: necessary to truncate so we don't allow withdrawing more rewards than owed

// fetch historical rewards for last period

// decrement reference count

// set new historical rewards with reference count of 1

// set current rewards, incrementing period by 1

// increment the reference count for a historical rewards value
func (k Keeper) incrementReferenceCount(ctx sdk.Context, valAddr sdk.ValAddress, period uint64) {
	_ = "STUB: not implemented"
	return
}

// decrement the reference count for a historical rewards value, and delete if zero references remain
func (k Keeper) decrementReferenceCount(ctx sdk.Context, valAddr sdk.ValAddress, period uint64) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) updateValidatorSlashFraction(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

// increment current period

// increment reference count on period we need to track

//nolint:gosec // block heights are always non-negative
