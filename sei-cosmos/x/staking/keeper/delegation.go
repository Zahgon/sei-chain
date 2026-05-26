package keeper

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "staking", "keeper")

// GetDelegation returns a specific delegation.
func (k Keeper) GetDelegation(ctx sdk.Context,
	delAddr sdk.AccAddress, valAddr sdk.ValAddress,
) (delegation types.Delegation, found bool) {
	_ = "STUB: not implemented"
	return *new(types.Delegation), false
}

// IterateAllDelegations iterates through all of the delegations.
func (k Keeper) IterateAllDelegations(ctx sdk.Context, cb func(delegation types.Delegation) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// GetAllDelegations returns all delegations used during genesis dump.
func (k Keeper) GetAllDelegations(ctx sdk.Context) (delegations []types.Delegation) {
	_ = "STUB: not implemented"
	return nil
}

// GetValidatorDelegations returns all delegations to a specific validator.
// Useful for querier.
func (k Keeper) GetValidatorDelegations(ctx sdk.Context, valAddr sdk.ValAddress) (delegations []types.Delegation) {
	_ = "STUB: not implemented"
	return nil
}

// GetDelegatorDelegations returns a given amount of all the delegations from a
// delegator.
func (k Keeper) GetDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress,
	maxRetrieve uint16,
) (delegations []types.Delegation) {
	_ = "STUB: not implemented"
	return nil
}

// trim if the array length < maxRetrieve

// SetDelegation sets a delegation.
func (k Keeper) SetDelegation(ctx sdk.Context, delegation types.Delegation) {
	_ = "STUB: not implemented"
	return
}

// RemoveDelegation removes a delegation.
func (k Keeper) RemoveDelegation(ctx sdk.Context, delegation types.Delegation) {
	_ = "STUB: not implemented"
	return
}

// GetUnbondingDelegations returns a given amount of all the delegator unbonding-delegations.
func (k Keeper) GetUnbondingDelegations(ctx sdk.Context, delegator sdk.AccAddress,
	maxRetrieve uint16,
) (unbondingDelegations []types.UnbondingDelegation) {
	_ = "STUB: not implemented"
	return nil
}

// trim if the array length < maxRetrieve

// GetUnbondingDelegation returns a unbonding delegation.
func (k Keeper) GetUnbondingDelegation(
	ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress,
) (ubd types.UnbondingDelegation, found bool) {
	_ = "STUB: not implemented"
	return *new(types.UnbondingDelegation), false
}

// GetUnbondingDelegationsFromValidator returns all unbonding delegations from a
// particular validator.
func (k Keeper) GetUnbondingDelegationsFromValidator(ctx sdk.Context, valAddr sdk.ValAddress) (ubds []types.UnbondingDelegation) {
	_ = "STUB: not implemented"
	return nil
}

// IterateUnbondingDelegations iterates through all of the unbonding delegations.
func (k Keeper) IterateUnbondingDelegations(ctx sdk.Context, fn func(index int64, ubd types.UnbondingDelegation) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// GetDelegatorUnbonding returns the total amount a delegator has unbonding.
func (k Keeper) GetDelegatorUnbonding(ctx sdk.Context, delegator sdk.AccAddress) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// IterateDelegatorUnbondingDelegations iterates through a delegator's unbonding delegations.
func (k Keeper) IterateDelegatorUnbondingDelegations(ctx sdk.Context, delegator sdk.AccAddress, cb func(ubd types.UnbondingDelegation) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// GetDelegatorBonded returs the total amount a delegator has bonded.
func (k Keeper) GetDelegatorBonded(ctx sdk.Context, delegator sdk.AccAddress) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// shouldn't happen

// IterateDelegatorDelegations iterates through one delegator's delegations.
func (k Keeper) IterateDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress, cb func(delegation types.Delegation) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// IterateDelegatorRedelegations iterates through one delegator's redelegations.
func (k Keeper) IterateDelegatorRedelegations(ctx sdk.Context, delegator sdk.AccAddress, cb func(red types.Redelegation) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// HasMaxUnbondingDelegationEntries - check if unbonding delegation has maximum number of entries.
func (k Keeper) HasMaxUnbondingDelegationEntries(ctx sdk.Context,
	delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
) bool {
	_ = "STUB: not implemented"
	return false
}

// SetUnbondingDelegation sets the unbonding delegation and associated index.
func (k Keeper) SetUnbondingDelegation(ctx sdk.Context, ubd types.UnbondingDelegation) {
	_ = "STUB: not implemented"
	return
}

// index, store empty bytes

// RemoveUnbondingDelegation removes the unbonding delegation object and associated index.
func (k Keeper) RemoveUnbondingDelegation(ctx sdk.Context, ubd types.UnbondingDelegation) {
	_ = "STUB: not implemented"
	return
}

// SetUnbondingDelegationEntry adds an entry to the unbonding delegation at
// the given addresses. It creates the unbonding delegation if it does not exist.
func (k Keeper) SetUnbondingDelegationEntry(
	ctx sdk.Context, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
	creationHeight int64, minTime time.Time, balance sdk.Int,
) types.UnbondingDelegation {
	_ = "STUB: not implemented"
	return *new(types.UnbondingDelegation)
}

// unbonding delegation queue timeslice operations

// GetUBDQueueTimeSlice gets a specific unbonding queue timeslice. A timeslice
// is a slice of DVPairs corresponding to unbonding delegations that expire at a
// certain time.
func (k Keeper) GetUBDQueueTimeSlice(ctx sdk.Context, timestamp time.Time) (dvPairs []types.DVPair) {
	_ = "STUB: not implemented"
	return nil
}

// SetUBDQueueTimeSlice sets a specific unbonding queue timeslice.
func (k Keeper) SetUBDQueueTimeSlice(ctx sdk.Context, timestamp time.Time, keys []types.DVPair) {
	_ = "STUB: not implemented"
	return
}

// InsertUBDQueue inserts an unbonding delegation to the appropriate timeslice
// in the unbonding queue.
func (k Keeper) InsertUBDQueue(ctx sdk.Context, ubd types.UnbondingDelegation,
	completionTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// UBDQueueIterator returns all the unbonding queue timeslices from time 0 until endTime.
func (k Keeper) UBDQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

// DequeueAllMatureUBDQueue returns a concatenated list of all the timeslices inclusively previous to
// currTime, and deletes the timeslices from the queue.
func (k Keeper) DequeueAllMatureUBDQueue(ctx sdk.Context, currTime time.Time) (matureUnbonds []types.DVPair) {
	_ = "STUB: not implemented"
	return nil
}

// gets an iterator for all timeslices from time 0 until the current Blockheader time

// GetRedelegations returns a given amount of all the delegator redelegations.
func (k Keeper) GetRedelegations(ctx sdk.Context, delegator sdk.AccAddress,
	maxRetrieve uint16,
) (redelegations []types.Redelegation) {
	_ = "STUB: not implemented"
	return nil
}

// trim if the array length < maxRetrieve

// GetRedelegation returns a redelegation.
func (k Keeper) GetRedelegation(ctx sdk.Context,
	delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress,
) (red types.Redelegation, found bool) {
	_ = "STUB: not implemented"
	return *new(types.Redelegation), false
}

// GetRedelegationsFromSrcValidator returns all redelegations from a particular
// validator.
func (k Keeper) GetRedelegationsFromSrcValidator(ctx sdk.Context, valAddr sdk.ValAddress) (reds []types.Redelegation) {
	_ = "STUB: not implemented"
	return nil
}

// HasReceivingRedelegation checks if validator is receiving a redelegation.
func (k Keeper) HasReceivingRedelegation(ctx sdk.Context,
	delAddr sdk.AccAddress, valDstAddr sdk.ValAddress,
) bool {
	_ = "STUB: not implemented"
	return false
}

// HasMaxRedelegationEntries checks if redelegation has maximum number of entries.
func (k Keeper) HasMaxRedelegationEntries(ctx sdk.Context,
	delegatorAddr sdk.AccAddress, validatorSrcAddr,
	validatorDstAddr sdk.ValAddress,
) bool {
	_ = "STUB: not implemented"
	return false
}

// SetRedelegation set a redelegation and associated index.
func (k Keeper) SetRedelegation(ctx sdk.Context, red types.Redelegation) {
	_ = "STUB: not implemented"
	return
}

// SetRedelegationEntry adds an entry to the unbonding delegation at the given
// addresses. It creates the unbonding delegation if it does not exist.
func (k Keeper) SetRedelegationEntry(ctx sdk.Context,
	delegatorAddr sdk.AccAddress, validatorSrcAddr,
	validatorDstAddr sdk.ValAddress, creationHeight int64,
	minTime time.Time, balance sdk.Int,
	sharesSrc, sharesDst sdk.Dec,
) types.Redelegation {
	_ = "STUB: not implemented"
	return *new(types.Redelegation)
}

// IterateRedelegations iterates through all redelegations.
func (k Keeper) IterateRedelegations(ctx sdk.Context, fn func(index int64, red types.Redelegation) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// RemoveRedelegation removes a redelegation object and associated index.
func (k Keeper) RemoveRedelegation(ctx sdk.Context, red types.Redelegation) {
	_ = "STUB: not implemented"
	return
}

// redelegation queue timeslice operations

// GetRedelegationQueueTimeSlice gets a specific redelegation queue timeslice. A
// timeslice is a slice of DVVTriplets corresponding to redelegations that
// expire at a certain time.
func (k Keeper) GetRedelegationQueueTimeSlice(ctx sdk.Context, timestamp time.Time) (dvvTriplets []types.DVVTriplet) {
	_ = "STUB: not implemented"
	return nil
}

// SetRedelegationQueueTimeSlice sets a specific redelegation queue timeslice.
func (k Keeper) SetRedelegationQueueTimeSlice(ctx sdk.Context, timestamp time.Time, keys []types.DVVTriplet) {
	_ = "STUB: not implemented"
	return
}

// InsertRedelegationQueue insert an redelegation delegation to the appropriate
// timeslice in the redelegation queue.
func (k Keeper) InsertRedelegationQueue(ctx sdk.Context, red types.Redelegation,
	completionTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// RedelegationQueueIterator returns all the redelegation queue timeslices from
// time 0 until endTime.
func (k Keeper) RedelegationQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

// DequeueAllMatureRedelegationQueue returns a concatenated list of all the
// timeslices inclusively previous to currTime, and deletes the timeslices from
// the queue.
func (k Keeper) DequeueAllMatureRedelegationQueue(ctx sdk.Context, currTime time.Time) (matureRedelegations []types.DVVTriplet) {
	_ = "STUB: not implemented"
	return nil
}

// gets an iterator for all timeslices from time 0 until the current Blockheader time

// Delegate performs a delegation, set/update everything necessary within the store.
// tokenSrc indicates the bond status of the incoming funds.
func (k Keeper) Delegate(
	ctx sdk.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc types.BondStatus,
	validator types.Validator, subtractAccount bool,
) (newShares sdk.Dec, err error) {
	_ = "STUB: not implemented"
	// In some situations, the exchange rate becomes invalid, e.g. if
	// Validator loses all tokens due to slashing. In this case,
	// make all future delegations invalid.
	return *new(sdk.Dec), nil
}

// check if the validator voting power exceeds the upper bound after the delegation
// validator.Tokens

// 1 power = Bond Amount / Power Reduction

// If it's beyond genesis then enforce power ratio per validator if there's more than maxVotingPowerEnforcementThreshold

// Convert bond amount to power first

// Validator's new total power cannot exceed the max power ratio that's allowed

// Get or create the delegation object

// call the appropriate hook if present

// if subtractAccount is true then we are
// performing a delegation and not a redelegation, thus the source tokens are
// all non bonded

// potentially transfer tokens between pools, if

// do nothing

// do nothing

// transfer pools

// transfer pools

// Update delegation

// Call the after-modification hook

// Unbond unbonds a particular delegation and perform associated store operations.
func (k Keeper) Unbond(
	ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec,
) (amount sdk.Int, err error) {
	_ = "STUB: not implemented"
	// check if a delegation object exists in the store
	return *new(sdk.Int), nil
}

// call the before-delegation-modified hook

// ensure that we have enough shares to remove

// get validator

// subtract shares from delegation

// If the delegation is the operator of the validator and undelegating will decrease the validator's
// self-delegation below their minimum, we jail the validator.

// remove the delegation

// call the after delegation modification hook

// remove the shares and coins from the validator
// NOTE that the amount is later (in keeper.Delegation) moved between staking module pools

// if not unbonded, we must instead remove validator in EndBlocker once it finishes its unbonding period

// getBeginInfo returns the completion time and height of a redelegation, along
// with a boolean signaling if the redelegation is complete based on the source
// validator.
func (k Keeper) getBeginInfo(
	ctx sdk.Context, valSrcAddr sdk.ValAddress,
) (completionTime time.Time, height int64, completeNow bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), 0, false
}

// TODO: When would the validator not be found?

// the longest wait - just unbonding period from now

// Undelegate unbonds an amount of delegator shares from a given validator. It
// will verify that the unbonding entries between the delegator and validator
// are not exceeded and unbond the staked tokens (based on shares) by creating
// an unbonding object and inserting it into the unbonding queue which will be
// processed during the staking EndBlocker.
func (k Keeper) Undelegate(
	ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, sharesAmount sdk.Dec,
) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// transfer the validator tokens to the not bonded pool

// CompleteUnbonding completes the unbonding of all mature entries in the
// retrieved unbonding delegation object and returns the total unbonding balance
// or an error upon failure.
func (k Keeper) CompleteUnbonding(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

// loop through all the entries and complete unbonding mature entries

// track undelegation only when remaining or truncated shares are non-zero

// set the unbonding delegation or remove it if there are no more entries

// BeginRedelegation begins unbonding / redelegation and creates a redelegation
// record.
func (k Keeper) BeginRedelegation(
	ctx sdk.Context, delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress, sharesAmount sdk.Dec,
) (completionTime time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// check if this is a transitive redelegation

// create the unbonding delegation

// no need to create the redelegation object

// CompleteRedelegation completes the redelegations of all mature entries in the
// retrieved redelegation object and returns the total redelegation (initial)
// balance or an error upon failure.
func (k Keeper) CompleteRedelegation(
	ctx sdk.Context, delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress,
) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

// loop through all the entries and complete mature redelegation entries

// set the redelegation or remove it if there are no more entries

// ValidateUnbondAmount validates that a given unbond or redelegation amount is
// valied based on upon the converted shares. If the amount is valid, the total
// amount of respective shares is returned, otherwise an error is returned.
func (k Keeper) ValidateUnbondAmount(
	ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, amt sdk.Int,
) (shares sdk.Dec, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec), nil
}

// Cap the shares at the delegation's shares. Shares being greater could occur
// due to rounding, however we don't want to truncate the shares or take the
// minimum because we want to allow for the full withdraw of shares from a
// delegation.
