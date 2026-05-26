package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Unjail calls the staking Unjail function to unjail a validator if the
// jailed period has concluded
func (k Keeper) Unjail(ctx sdk.Context, validatorAddr sdk.ValAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// cannot be unjailed if no self-delegation exists

// cannot be unjailed if not jailed

// If the validator has a ValidatorSigningInfo object that signals that the
// validator was bonded and so we must check that the validator is not tombstoned
// and can be unjailed at the current block.
//
// A validator that is jailed but has no ValidatorSigningInfo object signals
// that the validator was never bonded and must've been jailed due to falling
// below their minimum self-delegation. The validator can unjail at any point
// assuming they've now bonded above their minimum self-delegation.

// cannot be unjailed if tombstoned

// cannot be unjailed until out of jail
