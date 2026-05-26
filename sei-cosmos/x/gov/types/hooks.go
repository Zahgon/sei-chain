package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var _ GovHooks = MultiGovHooks{}

// combine multiple governance hooks, all hook functions are run in array sequence
type MultiGovHooks []GovHooks

func NewMultiGovHooks(hooks ...GovHooks) MultiGovHooks {
	_ = "STUB: not implemented"
	return *new(MultiGovHooks)
}

func (h MultiGovHooks) AfterProposalSubmission(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}

func (h MultiGovHooks) AfterProposalDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiGovHooks) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

func (h MultiGovHooks) AfterProposalFailedMinDeposit(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}

func (h MultiGovHooks) AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}
