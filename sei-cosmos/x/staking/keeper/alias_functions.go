package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Validator Set

// iterate through the validator set and perform the provided function
func (k Keeper) IterateValidators(ctx sdk.Context, fn func(index int64, validator types.ValidatorI) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// XXX is this safe will the validator unexposed fields be able to get written to?

// iterate through the bonded validator set and perform the provided function
func (k Keeper) IterateBondedValidatorsByPower(ctx sdk.Context, fn func(index int64, validator types.ValidatorI) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// XXX is this safe will the validator unexposed fields be able to get written to?

// iterate through the active validator set and perform the provided function
func (k Keeper) IterateLastValidators(ctx sdk.Context, fn func(index int64, validator types.ValidatorI) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// XXX is this safe will the validator unexposed fields be able to get written to?

// Validator gets the Validator interface for a particular address
func (k Keeper) Validator(ctx sdk.Context, address sdk.ValAddress) types.ValidatorI {
	_ = "STUB: not implemented"
	return *new(types.ValidatorI)
}

// ValidatorByConsAddr gets the validator interface for a particular pubkey
func (k Keeper) ValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) types.ValidatorI {
	_ = "STUB: not implemented"
	return *new(types.ValidatorI)
}

// Delegation Set

// Returns self as it is both a validatorset and delegationset
func (k Keeper) GetValidatorSet() types.ValidatorSet {
	_ = "STUB: not implemented"

	// Delegation get the delegation interface for a particular set of delegator and validator addresses
	return *new(types.ValidatorSet)
}

func (k Keeper) Delegation(ctx sdk.Context, addrDel sdk.AccAddress, addrVal sdk.ValAddress) types.DelegationI {
	_ = "STUB: not implemented"
	return *new(types.DelegationI)
}

// iterate through all of the delegations from a delegator
func (k Keeper) IterateDelegations(ctx sdk.Context, delAddr sdk.AccAddress,
	fn func(index int64, del types.DelegationI) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// smallest to largest

// return all delegations used during genesis dump
// TODO: remove this func, change all usage for iterate functionality
func (k Keeper) GetAllSDKDelegations(ctx sdk.Context) (delegations []types.Delegation) {
	_ = "STUB: not implemented"
	return nil
}
