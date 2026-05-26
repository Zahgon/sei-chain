package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
)

// get the delegator withdraw address, defaulting to the delegator address
func (k Keeper) GetDelegatorWithdrawAddr(ctx sdk.Context, delAddr sdk.AccAddress) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// set the delegator withdraw address
func (k Keeper) SetDelegatorWithdrawAddr(ctx sdk.Context, delAddr, withdrawAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// delete a delegator withdraw addr
func (k Keeper) DeleteDelegatorWithdrawAddr(ctx sdk.Context, delAddr, withdrawAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// iterate over delegator withdraw addrs
func (k Keeper) IterateDelegatorWithdrawAddrs(ctx sdk.Context, handler func(del sdk.AccAddress, addr sdk.AccAddress) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// get the global fee pool distribution info
func (k Keeper) GetFeePool(ctx sdk.Context) (feePool types.FeePool) {
	_ = "STUB: not implemented"
	return *new(types.FeePool)
}

// set the global fee pool distribution info
func (k Keeper) SetFeePool(ctx sdk.Context, feePool types.FeePool) {
	_ = "STUB: not implemented"
	return
}

// GetPreviousProposerConsAddr returns the proposer consensus address for the
// current block.
func (k Keeper) GetPreviousProposerConsAddr(ctx sdk.Context) sdk.ConsAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ConsAddress)
}

// set the proposer public key for this block
func (k Keeper) SetPreviousProposerConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) {
	_ = "STUB: not implemented"
	return
}

// get the starting info associated with a delegator
func (k Keeper) GetDelegatorStartingInfo(ctx sdk.Context, val sdk.ValAddress, del sdk.AccAddress) (period types.DelegatorStartingInfo) {
	_ = "STUB: not implemented"
	return *new(types.DelegatorStartingInfo)
}

// set the starting info associated with a delegator
func (k Keeper) SetDelegatorStartingInfo(ctx sdk.Context, val sdk.ValAddress, del sdk.AccAddress, period types.DelegatorStartingInfo) {
	_ = "STUB: not implemented"
	return
}

// check existence of the starting info associated with a delegator
func (k Keeper) HasDelegatorStartingInfo(ctx sdk.Context, val sdk.ValAddress, del sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

// delete the starting info associated with a delegator
func (k Keeper) DeleteDelegatorStartingInfo(ctx sdk.Context, val sdk.ValAddress, del sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// iterate over delegator starting infos
func (k Keeper) IterateDelegatorStartingInfos(ctx sdk.Context, handler func(val sdk.ValAddress, del sdk.AccAddress, info types.DelegatorStartingInfo) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// get historical rewards for a particular period
func (k Keeper) GetValidatorHistoricalRewards(ctx sdk.Context, val sdk.ValAddress, period uint64) (rewards types.ValidatorHistoricalRewards) {
	_ = "STUB: not implemented"
	return *new(types.ValidatorHistoricalRewards)
}

// set historical rewards for a particular period
func (k Keeper) SetValidatorHistoricalRewards(ctx sdk.Context, val sdk.ValAddress, period uint64, rewards types.ValidatorHistoricalRewards) {
	_ = "STUB: not implemented"
	return
}

// iterate over historical rewards
func (k Keeper) IterateValidatorHistoricalRewards(ctx sdk.Context, handler func(val sdk.ValAddress, period uint64, rewards types.ValidatorHistoricalRewards) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// delete a historical reward
func (k Keeper) DeleteValidatorHistoricalReward(ctx sdk.Context, val sdk.ValAddress, period uint64) {
	_ = "STUB: not implemented"
	return
}

// delete historical rewards for a validator
func (k Keeper) DeleteValidatorHistoricalRewards(ctx sdk.Context, val sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// delete all historical rewards
func (k Keeper) DeleteAllValidatorHistoricalRewards(ctx sdk.Context) {
	_ = "STUB: not implemented"
	return
}

// historical reference count (used for testcases)
func (k Keeper) GetValidatorHistoricalReferenceCount(ctx sdk.Context) (count uint64) {
	_ = "STUB: not implemented"
	return 0
}

// get current rewards for a validator
func (k Keeper) GetValidatorCurrentRewards(ctx sdk.Context, val sdk.ValAddress) (rewards types.ValidatorCurrentRewards) {
	_ = "STUB: not implemented"
	return *new(types.ValidatorCurrentRewards)
}

// set current rewards for a validator
func (k Keeper) SetValidatorCurrentRewards(ctx sdk.Context, val sdk.ValAddress, rewards types.ValidatorCurrentRewards) {
	_ = "STUB: not implemented"
	return
}

// delete current rewards for a validator
func (k Keeper) DeleteValidatorCurrentRewards(ctx sdk.Context, val sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// iterate over current rewards
func (k Keeper) IterateValidatorCurrentRewards(ctx sdk.Context, handler func(val sdk.ValAddress, rewards types.ValidatorCurrentRewards) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// get accumulated commission for a validator
func (k Keeper) GetValidatorAccumulatedCommission(ctx sdk.Context, val sdk.ValAddress) (commission types.ValidatorAccumulatedCommission) {
	_ = "STUB: not implemented"
	return *new(types.ValidatorAccumulatedCommission)
}

// set accumulated commission for a validator
func (k Keeper) SetValidatorAccumulatedCommission(ctx sdk.Context, val sdk.ValAddress, commission types.ValidatorAccumulatedCommission) {
	_ = "STUB: not implemented"
	return
}

// delete accumulated commission for a validator
func (k Keeper) DeleteValidatorAccumulatedCommission(ctx sdk.Context, val sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// iterate over accumulated commissions
func (k Keeper) IterateValidatorAccumulatedCommissions(ctx sdk.Context, handler func(val sdk.ValAddress, commission types.ValidatorAccumulatedCommission) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// get validator outstanding rewards
func (k Keeper) GetValidatorOutstandingRewards(ctx sdk.Context, val sdk.ValAddress) (rewards types.ValidatorOutstandingRewards) {
	_ = "STUB: not implemented"
	return *new(types.ValidatorOutstandingRewards)
}

// set validator outstanding rewards
func (k Keeper) SetValidatorOutstandingRewards(ctx sdk.Context, val sdk.ValAddress, rewards types.ValidatorOutstandingRewards) {
	_ = "STUB: not implemented"
	return
}

// delete validator outstanding rewards
func (k Keeper) DeleteValidatorOutstandingRewards(ctx sdk.Context, val sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// iterate validator outstanding rewards
func (k Keeper) IterateValidatorOutstandingRewards(ctx sdk.Context, handler func(val sdk.ValAddress, rewards types.ValidatorOutstandingRewards) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// get slash event for height
func (k Keeper) GetValidatorSlashEvent(ctx sdk.Context, val sdk.ValAddress, height, period uint64) (event types.ValidatorSlashEvent, found bool) {
	_ = "STUB: not implemented"
	return *new(types.ValidatorSlashEvent), false
}

// set slash event for height
func (k Keeper) SetValidatorSlashEvent(ctx sdk.Context, val sdk.ValAddress, height, period uint64, event types.ValidatorSlashEvent) {
	_ = "STUB: not implemented"
	return
}

// iterate over slash events between heights, inclusive
func (k Keeper) IterateValidatorSlashEventsBetween(ctx sdk.Context, val sdk.ValAddress, startingHeight uint64, endingHeight uint64,
	handler func(height uint64, event types.ValidatorSlashEvent) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// iterate over all slash events
func (k Keeper) IterateValidatorSlashEvents(ctx sdk.Context, handler func(val sdk.ValAddress, height uint64, event types.ValidatorSlashEvent) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// delete slash events for a particular validator
func (k Keeper) DeleteValidatorSlashEvents(ctx sdk.Context, val sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// delete all slash events
func (k Keeper) DeleteAllValidatorSlashEvents(ctx sdk.Context) { _ = "STUB: not implemented"; return }
