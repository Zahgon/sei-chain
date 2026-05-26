package keeper

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// BlockValidatorUpdates calculates the ValidatorUpdates for the current block
// Called in each EndBlock
func (k Keeper) BlockValidatorUpdates(ctx sdk.Context) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	// Calculate validator set changes.
	//
	// NOTE: ApplyAndReturnValidatorSetUpdates has to come before
	// UnbondAllMatureValidatorQueue.
	// This fixes a bug when the unbonding period is instant (is the case in
	// some of the tests). The test expected the validator to be completely
	// unbonded after the Endblocker (go from Bonded -> Unbonding during
	// ApplyAndReturnValidatorSetUpdates and then Unbonding -> Unbonded during
	// UnbondAllMatureValidatorQueue).
	return nil
}

// unbond all mature validators from the unbonding queue

// Remove all mature unbonding delegations from the ubd queue.

// Remove all mature redelegations from the red queue.

// ApplyAndReturnValidatorSetUpdates applies and return accumulated updates to the bonded validator set. Also,
// * Updates the active valset as keyed by LastValidatorPowerKey.
// * Updates the total power as keyed by LastTotalPowerKey.
// * Updates validator status' according to updated powers.
// * Updates the fee pool bonded vs not-bonded tokens.
// * Updates relevant indices.
// It gets called once after genesis, another time maybe after genesis transactions,
// then once at every EndBlock.
//
// CONTRACT: Only validators with non-zero power or zero-power that were bonded
// at the previous block height or were removed from the validator set entirely
// are returned to Tendermint.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve the last validator set.
// The persistent set is updated later in this function.
// (see LastValidatorPowerKey).

// Iterate over validators, highest power to lowest.

// everything that is iterated in this loop is becoming or already a
// part of the bonded validator set

// if we get to a zero-power validator (which we don't bond),
// there are no more possible bonded validators

// apply the appropriate state change if necessary

// no state change

// fetch the old power bytes

// update the validator set if power has changed

// Update the pools based on the recent updates in the validator set:
// - The tokens from the non-bonded candidates that enter the new validator set need to be transferred
// to the Bonded pool.
// - The tokens from the bonded validators that are being kicked out from the validator set
// need to be transferred to the NotBonded pool.

// Compare and subtract the respective amounts to only perform one transfer.
// This is done in order to avoid doing multiple updates inside each iterator/loop.

// equal amounts of tokens; no update required

// set total power on lookup index if there are any updates

// Validator state transitions

func (k Keeper) bondedToUnbonding(ctx sdk.Context, validator types.Validator) (types.Validator, error) {
	_ = "STUB: not implemented"
	return *new(types.Validator), nil
}

func (k Keeper) unbondingToBonded(ctx sdk.Context, validator types.Validator) (types.Validator, error) {
	_ = "STUB: not implemented"
	return *new(types.Validator), nil
}

func (k Keeper) unbondedToBonded(ctx sdk.Context, validator types.Validator) (types.Validator, error) {
	_ = "STUB: not implemented"
	return *new(types.Validator), nil
}

// UnbondingToUnbonded switches a validator from unbonding state to unbonded state
func (k Keeper) UnbondingToUnbonded(ctx sdk.Context, validator types.Validator) types.Validator {
	_ = "STUB: not implemented"
	return *new(types.Validator)
}

// send a validator to jail
func (k Keeper) jailValidator(ctx sdk.Context, validator types.Validator) {
	_ = "STUB: not implemented"
	return
}

// remove a validator from jail
func (k Keeper) unjailValidator(ctx sdk.Context, validator types.Validator) {
	_ = "STUB: not implemented"
	return
}

// perform all the store operations for when a validator status becomes bonded
func (k Keeper) bondValidator(ctx sdk.Context, validator types.Validator) (types.Validator, error) {
	_ = "STUB: not implemented"
	// delete the validator by power index, as the key will change
	return *new(types.Validator), nil
}

// save the now bonded validator record to the two referenced stores

// delete from queue if present

// trigger hook

// perform all the store operations for when a validator begins unbonding
func (k Keeper) beginUnbondingValidator(ctx sdk.Context, validator types.Validator) (types.Validator, error) {
	_ = "STUB: not implemented"
	return *new(types.Validator), nil
}

// delete the validator by power index, as the key will change

// sanity check

// set the unbonding completion time and completion height appropriately

// save the now unbonded validator record and power index

// Adds to unbonding validator queue

// trigger hook

// perform all the store operations for when a validator status becomes unbonded
func (k Keeper) completeUnbondingValidator(ctx sdk.Context, validator types.Validator) types.Validator {
	_ = "STUB: not implemented"
	return *new(types.Validator)
}

// map of operator bech32-addresses to serialized power
// We use bech32 strings here, because we can't have slices as keys: map[[]byte][]byte
type validatorsByAddr map[string][]byte

// get the last validator set
func (k Keeper) getLastValidatorsByAddr(ctx sdk.Context) (validatorsByAddr, error) {
	_ = "STUB: not implemented"
	return *new(validatorsByAddr), nil
}

// extract the validator address from the key (prefix is 1-byte, addrLen is 1-byte)

// given a map of remaining validators to previous bonded power
// returns the list of validators to be unbonded, sorted by operator address
func sortNoLongerBonded(last validatorsByAddr) ([][]byte, error) {
	_ = "STUB: not implemented"
	// sort the map keys for determinism
	return nil, nil
}

// sorted by address - order doesn't matter

// -1 means strictly less than
