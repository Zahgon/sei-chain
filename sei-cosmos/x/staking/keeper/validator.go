package keeper

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// get a single validator
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator types.Validator, found bool) {
	_ = "STUB: not implemented"
	return *new(types.Validator), false
}

func (k Keeper) mustGetValidator(ctx sdk.Context, addr sdk.ValAddress) types.Validator {
	_ = "STUB: not implemented"
	return *new(types.Validator)
}

// get a single validator by consensus address
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	_ = "STUB: not implemented"
	return *new(types.Validator), false
}

func (k Keeper) mustGetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) types.Validator {
	_ = "STUB: not implemented"
	return *new(types.Validator)
}

// set the main record holding validator details
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	_ = "STUB: not implemented"
	return
}

// validator index
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) error {
	_ = "STUB: not implemented"
	return nil
}

// validator index
func (k Keeper) SetValidatorByPowerIndex(ctx sdk.Context, validator types.Validator) {
	_ = "STUB: not implemented"
	// jailed validators are not kept in the power index
	return
}

// validator index
func (k Keeper) DeleteValidatorByPowerIndex(ctx sdk.Context, validator types.Validator) {
	_ = "STUB: not implemented"
	return
}

// validator index
func (k Keeper) SetNewValidatorByPowerIndex(ctx sdk.Context, validator types.Validator) {
	_ = "STUB: not implemented"
	return
}

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) AddValidatorTokensAndShares(ctx sdk.Context, validator types.Validator,
	tokensToAdd sdk.Int) (valOut types.Validator, addedShares sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(types.Validator), *new(sdk.Dec)
}

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) RemoveValidatorTokensAndShares(ctx sdk.Context, validator types.Validator,
	sharesToRemove sdk.Dec) (valOut types.Validator, removedTokens sdk.Int) {
	_ = "STUB: not implemented"
	return *new(types.Validator), *new(sdk.Int)
}

// Update the tokens of an existing validator, update the validators power index key
func (k Keeper) RemoveValidatorTokens(ctx sdk.Context,
	validator types.Validator, tokensToRemove sdk.Int) types.Validator {
	_ = "STUB: not implemented"
	return *new(types.Validator)
}

// UpdateValidatorCommission attempts to update a validator's commission rate.
// An error is returned if the new commission rate is invalid.
func (k Keeper) UpdateValidatorCommission(ctx sdk.Context,
	validator types.Validator, newRate sdk.Dec) (types.Commission, error) {
	_ = "STUB: not implemented"
	return *new(types.Commission), nil
}

// remove the validator record and associated indexes
// except for the bonded validator index which is only handled in ApplyAndReturnTendermintUpdates
// TODO, this function panics, and it's not good.
func (k Keeper) RemoveValidator(ctx sdk.Context, address sdk.ValAddress) {
	_ = "STUB: not implemented"
	// first retrieve the old validator record
	return
}

// delete the old validator record

// call hooks

// get groups of validators

// get the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	_ = "STUB: not implemented"
	return nil
}

// return a given amount of all the validators
func (k Keeper) GetValidators(ctx sdk.Context, maxRetrieve uint32) (validators []types.Validator) {
	_ = "STUB: not implemented"
	return nil
}

// trim if the array length < maxRetrieve

// get the current group of bonded validators sorted by power-rank
func (k Keeper) GetBondedValidatorsByPower(ctx sdk.Context) []types.Validator {
	_ = "STUB: not implemented"
	return nil
}

// trim

// GetBondedValidators returns the bonded validator set.
func (k Keeper) GetBondedValidators(ctx sdk.Context) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// returns an iterator for the current validator power store
func (k Keeper) ValidatorsPowerStoreIterator(ctx sdk.Context) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

// Last Validator Index

// Load the last validator power.
// Returns zero if the operator was not a validator last block.
func (k Keeper) GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64) {
	_ = "STUB: not implemented"
	return 0
}

// Set the last validator power.
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress, power int64) {
	_ = "STUB: not implemented"
	return
}

// Delete the last validator power.
func (k Keeper) DeleteLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// returns an iterator for the consensus validators in the last block
func (k Keeper) LastValidatorsIterator(ctx sdk.Context) (iterator sdk.Iterator) {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

// Iterate over last validator powers.
func (k Keeper) IterateLastValidatorPowers(ctx sdk.Context, handler func(operator sdk.ValAddress, power int64) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// get the group of the bonded validators
func (k Keeper) GetLastValidators(ctx sdk.Context) (validators []types.Validator) {
	_ = "STUB: not implemented"
	return nil
}

// add the actual validator power sorted store

// sanity check

// trim

// GetUnbondingValidators returns a slice of mature validator addresses that
// complete their unbonding at a given time and height.
func (k Keeper) GetUnbondingValidators(ctx sdk.Context, endTime time.Time, endHeight int64) []string {
	_ = "STUB: not implemented"
	return nil
}

// SetUnbondingValidatorsQueue sets a given slice of validator addresses into
// the unbonding validator queue by a given height and time.
func (k Keeper) SetUnbondingValidatorsQueue(ctx sdk.Context, endTime time.Time, endHeight int64, addrs []string) {
	_ = "STUB: not implemented"
	return
}

// InsertUnbondingValidatorQueue inserts a given unbonding validator address into
// the unbonding validator queue for a given height and time.
func (k Keeper) InsertUnbondingValidatorQueue(ctx sdk.Context, val types.Validator) {
	_ = "STUB: not implemented"
	return
}

// DeleteValidatorQueueTimeSlice deletes all entries in the queue indexed by a
// given height and time.
func (k Keeper) DeleteValidatorQueueTimeSlice(ctx sdk.Context, endTime time.Time, endHeight int64) {
	_ = "STUB: not implemented"
	return
}

// DeleteValidatorQueue removes a validator by address from the unbonding queue
// indexed by a given height and time.
func (k Keeper) DeleteValidatorQueue(ctx sdk.Context, val types.Validator) {
	_ = "STUB: not implemented"
	return
}

// ValidatorQueueIterator returns an interator ranging over validators that are
// unbonding whose unbonding completion occurs at the given height and time.
func (k Keeper) ValidatorQueueIterator(ctx sdk.Context, endTime time.Time, endHeight int64) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

// UnbondAllMatureValidators unbonds all the mature unbonding validators that
// have finished their unbonding period.
func (k Keeper) UnbondAllMatureValidators(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// unbondingValIterator will contain all validator addresses indexed under
// the ValidatorQueueKey prefix. Note, the entire index key is composed as
// ValidatorQueueKey | timeBzLen (8-byte big endian) | timeBz | heightBz (8-byte big endian),
// so it may be possible that certain validator addresses that are iterated
// over are not ready to unbond, so an explicit check is required.

// All addresses for the given key have the same unbonding height and time.
// We only unbond if the height and time are less than the current height
// and time.
