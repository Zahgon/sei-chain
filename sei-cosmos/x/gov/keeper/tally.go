package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// TODO: Break into several smaller functions for clarity

// Tally iterates over the votes and updates the tally of a proposal based on the voting power of the
// voters
func (keeper Keeper) Tally(ctx sdk.Context, proposal types.Proposal) (passes bool, burnDeposits bool, tallyResults types.TallyResult) {
	_ = "STUB: not implemented"
	return false, false, *new(types.TallyResult)
}

// fetch all the bonded validators, insert them into currValidators

// if validator, just record it in the map

// iterate over all delegations from voter, deduct from any delegated-to validators

// There is no need to handle the special case that validator address equal to voter address.
// Because voter's voting power will tally again even if there will deduct voter's voting power from validator.

// delegation shares * bonded / total shares

// iterate over the validators again to tally their voting power

// TODO: Upgrade the spec to cover all of these cases & remove pseudocode.
// If there is no staked coins, the proposal fails

// If there is not enough quorum of votes, the proposal fails

// Get the quorum threshold based on if the proposal is expedited or not

// If no one votes (everyone abstains), proposal fails

// If more than 1/3 of voters veto, proposal fails

// If more than threshold of non-abstaining voters vote Yes, proposal passes
// default value for regular proposals is 1/2. For expedited 2/3

// Otherwise proposal fails
