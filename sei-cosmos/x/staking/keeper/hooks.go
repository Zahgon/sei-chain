package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Implements StakingHooks interface
var _ types.StakingHooks = Keeper{}

// AfterValidatorCreated - call hook if registered
func (k Keeper) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// BeforeValidatorModified - call hook if registered
func (k Keeper) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// AfterValidatorRemoved - call hook if registered
func (k Keeper) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// AfterValidatorBonded - call hook if registered
func (k Keeper) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// AfterValidatorBeginUnbonding - call hook if registered
func (k Keeper) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// BeforeDelegationCreated - call hook if registered
func (k Keeper) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// BeforeDelegationSharesModified - call hook if registered
func (k Keeper) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// BeforeDelegationRemoved - call hook if registered
func (k Keeper) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// AfterDelegationModified - call hook if registered
func (k Keeper) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// BeforeValidatorSlashed - call hook if registered
func (k Keeper) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {
	_ = "STUB: not implemented"
	return
}
