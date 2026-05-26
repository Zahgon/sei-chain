package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// Implements GovHooks interface
var _ types.GovHooks = Keeper{}

// AfterProposalSubmission - call hook if registered
func (keeper Keeper) AfterProposalSubmission(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}

// AfterProposalDeposit - call hook if registered
func (keeper Keeper) AfterProposalDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// AfterProposalVote - call hook if registered
func (keeper Keeper) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// AfterProposalFailedMinDeposit - call hook if registered
func (keeper Keeper) AfterProposalFailedMinDeposit(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}

// AfterProposalVotingPeriodEnded - call hook if registered
func (keeper Keeper) AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}
