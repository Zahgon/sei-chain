package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	types "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Slash a validator for an infraction committed at a known height
// Find the contributing stake at that height and burn the specified slashFactor
// of it, updating unbonding delegations & redelegations appropriately
//
// CONTRACT:
//
//	slashFactor is non-negative
//
// CONTRACT:
//
//	Infraction was committed equal to or less than an unbonding period in the past,
//	so all unbonding delegations and redelegations from that height are stored
//
// CONTRACT:
//
//	Slash will not slash unbonded validators (for the above reason)
//
// CONTRACT:
//
//	Infraction was committed at the current height or at a past height,
//	not at a height in the future
func (k Keeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

// Amount of slashing = slash slashFactor * power at time of infraction

// ref https://github.com/cosmos/cosmos-sdk/issues/1348

// If not found, the validator must have been overslashed and removed - so we don't need to do anything
// NOTE:  Correctness dependent on invariant that unbonding delegations / redelegations must also have been completely
//        slashed in this case - which we don't explicitly check, but should be true.
// Log the slash attempt for future reference (maybe we should tag it too)

// should not be slashing an unbonded validator

// call the before-modification hook

// Track remaining slash amount for the validator
// This will decrease when we slash unbondings and
// redelegations, as that stake has since unbonded

// Can't slash infractions in the future

// Special-case slash at current height for efficiency - we don't need to
// look through unbonding delegations or redelegations.

// Iterate through unbonding delegations from slashed validator

// Iterate through redelegations from slashed source validator

// cannot decrease balance below zero

// defensive.

// we need to calculate the *effective* slash fraction for distribution

// possible if power has changed

// call the before-slashed hook

// Deduct from validator's bonded tokens and update the validator.
// Burn the slashed tokens from the pool account and decrease the total supply.

// jail a validator
func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) { _ = "STUB: not implemented"; return }

// unjail a validator
func (k Keeper) Unjail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	_ = "STUB: not implemented"
	return
}

// slash an unbonding delegation and update the pool
// return the amount that would have been slashed assuming
// the unbonding delegation had enough stake to slash
// (the amount actually slashed may be less if there's
// insufficient stake remaining)
func (k Keeper) SlashUnbondingDelegation(ctx sdk.Context, unbondingDelegation types.UnbondingDelegation,
	infractionHeight int64, slashFactor sdk.Dec) (totalSlashAmount sdk.Int) {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// perform slashing on all entries within the unbonding delegation

// If unbonding started before this height, stake didn't contribute to infraction

// Unbonding delegation no longer eligible for slashing, skip it

// Calculate slash amount proportional to stake contributing to infraction

// Don't slash more tokens than held
// Possible since the unbonding delegation may already
// have been slashed, and slash amounts are calculated
// according to stake held at time of infraction

// Update unbonding delegation if necessary

// slash a redelegation and update the pool
// return the amount that would have been slashed assuming
// the unbonding delegation had enough stake to slash
// (the amount actually slashed may be less if there's
// insufficient stake remaining)
// NOTE this is only slashing for prior infractions from the source validator
func (k Keeper) SlashRedelegation(ctx sdk.Context, srcValidator types.Validator, redelegation types.Redelegation,
	infractionHeight int64, slashFactor sdk.Dec) (totalSlashAmount sdk.Int) {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// perform slashing on all entries within the redelegation

// If redelegation started before this height, stake didn't contribute to infraction

// Redelegation no longer eligible for slashing, skip it

// Calculate slash amount proportional to stake contributing to infraction

// Unbond from target validator

// If deleted, delegation has zero shares, and we can't unbond any more

// tokens of a redelegation currently live in the destination validator
// therefor we must burn tokens from the destination-validator's bonding status
