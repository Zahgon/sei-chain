package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Return all validators that a delegator is bonded to. If maxRetrieve is supplied, the respective amount will be returned.
func (k Keeper) GetDelegatorValidators(
	ctx sdk.Context, delegatorAddr sdk.AccAddress, maxRetrieve uint32,
) types.Validators {
	_ = "STUB: not implemented"
	return *new(types.Validators)
}

// smallest to largest

// trim

// return a validator that a delegator is bonded to
func (k Keeper) GetDelegatorValidator(
	ctx sdk.Context, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
) (validator types.Validator, err error) {
	_ = "STUB: not implemented"
	return *new(types.Validator), nil
}

// return all delegations for a delegator
func (k Keeper) GetAllDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress) []types.Delegation {
	_ = "STUB: not implemented"
	return nil
}

// smallest to largest

// return all unbonding-delegations for a delegator
func (k Keeper) GetAllUnbondingDelegations(ctx sdk.Context, delegator sdk.AccAddress) []types.UnbondingDelegation {
	_ = "STUB: not implemented"
	return nil
}

// smallest to largest

// return all redelegations for a delegator
func (k Keeper) GetAllRedelegations(
	ctx sdk.Context, delegator sdk.AccAddress, srcValAddress, dstValAddress sdk.ValAddress,
) []types.Redelegation {
	_ = "STUB: not implemented"
	return nil
}

// smallest to largest
