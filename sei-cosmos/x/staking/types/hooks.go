package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// combine multiple staking hooks, all hook functions are run in array sequence
type MultiStakingHooks []StakingHooks

func NewMultiStakingHooks(hooks ...StakingHooks) MultiStakingHooks {
	_ = "STUB: not implemented"
	return *new(MultiStakingHooks)
}

func (h MultiStakingHooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiStakingHooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {
	_ = "STUB: not implemented"
	return
}
