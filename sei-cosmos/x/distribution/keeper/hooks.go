package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var _ stakingtypes.StakingHooks = Hooks{}

// Create new distribution hooks
func (k Keeper) Hooks() Hooks {
	_ = "STUB: not implemented"

	// initialize validator distribution record
	return *new(Hooks)
}

func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// AfterValidatorRemoved performs clean up after a validator is removed
func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, _ sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	// fetch outstanding
	return
}

// force-withdraw commission

// subtract from outstanding

// split into integral & remainder

// remainder to community pool

// add to validator account

// Add outstanding to community pool
// The validator is removed only after it has no more delegations.
// This operation sends only the remaining dust to the community pool.

// delete outstanding

// remove commission record

// clear slashes

// clear historical rewards

// clear current rewards

// increment period
func (h Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// withdraw delegation rewards (which also increments period)
func (h Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// create new delegation period record
func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// record the slash event
func (h Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

func (h Hooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
