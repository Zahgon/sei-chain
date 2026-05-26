package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// AddVote adds a vote on a specific proposal
func (keeper Keeper) AddVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress, options types.WeightedVoteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// called after a vote on a proposal is cast

// GetAllVotes returns all the votes from the store
func (keeper Keeper) GetAllVotes(ctx sdk.Context) (votes types.Votes) {
	_ = "STUB: not implemented"
	return *new(types.Votes)
}

// GetVotes returns all the votes from a proposal
func (keeper Keeper) GetVotes(ctx sdk.Context, proposalID uint64) (votes types.Votes) {
	_ = "STUB: not implemented"
	return *new(types.Votes)
}

// GetVote gets the vote from an address on a specific proposal
func (keeper Keeper) GetVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) (vote types.Vote, found bool) {
	_ = "STUB: not implemented"
	return *new(types.Vote), false
}

// SetVote sets a Vote to the gov store
func (keeper Keeper) SetVote(ctx sdk.Context, vote types.Vote) {
	_ = "STUB: not implemented"
	// vote.Option is a deprecated field, we don't set it in state
	return
}

//nolint
//nolint

// IterateAllVotes iterates over the all the stored votes and performs a callback function
func (keeper Keeper) IterateAllVotes(ctx sdk.Context, cb func(vote types.Vote) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// IterateVotes iterates over the all the proposals votes and performs a callback function
func (keeper Keeper) IterateVotes(ctx sdk.Context, proposalID uint64, cb func(vote types.Vote) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// deleteVote deletes a vote from a given proposalID and voter from the store
func (keeper Keeper) deleteVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// populateLegacyOption adds graceful fallback of deprecated `Option` field, in case
// there's only 1 VoteOption.
func populateLegacyOption(vote *types.Vote) { _ = "STUB: not implemented"; return }

//nolint
