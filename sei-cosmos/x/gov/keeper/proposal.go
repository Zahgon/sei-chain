package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// SubmitProposal create new proposal given a content
func (keeper Keeper) SubmitProposal(ctx sdk.Context, content types.Content) (types.Proposal, error) {
	_ = "STUB: not implemented"
	return *new(types.Proposal), nil
}

// SubmitProposalWithExpedite create new proposal given a content and whether expedited or not
func (keeper Keeper) SubmitProposalWithExpedite(ctx sdk.Context, content types.Content, isExpedited bool) (types.Proposal, error) {
	_ = "STUB: not implemented"
	return *new(types.Proposal), nil
}

// Ensure that the parameter exists

// Validate each parameter change exists

// called right after a proposal is submitted

// GetProposal get proposal from store by ProposalID
func (keeper Keeper) GetProposal(ctx sdk.Context, proposalID uint64) (types.Proposal, bool) {
	_ = "STUB: not implemented"
	return *new(types.Proposal), false
}

// SetProposal set a proposal to store
func (keeper Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	_ = "STUB: not implemented"
	return
}

// DeleteProposal deletes a proposal from store
func (keeper Keeper) DeleteProposal(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}

// IterateProposals iterates over the all the proposals and performs a callback function
func (keeper Keeper) IterateProposals(ctx sdk.Context, cb func(proposal types.Proposal) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// GetProposals returns all the proposals from store
func (keeper Keeper) GetProposals(ctx sdk.Context) (proposals types.Proposals) {
	_ = "STUB: not implemented"
	return *new(types.Proposals)
}

// GetProposalsFiltered retrieves proposals filtered by a given set of params which
// include pagination parameters along with voter and depositor addresses and a
// proposal status. The voter address will filter proposals by whether or not
// that address has voted on proposals. The depositor address will filter proposals
// by whether or not that address has deposited to them. Finally, status will filter
// proposals by status.
//
// NOTE: If no filters are provided, all proposals will be returned in paginated
// form.
func (keeper Keeper) GetProposalsFiltered(ctx sdk.Context, params types.QueryProposalsParams) types.Proposals {
	_ = "STUB: not implemented"
	return *new(types.Proposals)
}

// match status (if supplied/valid)

// match voter address (if supplied)

// match depositor (if supplied)

// GetProposalID gets the highest proposal ID
func (keeper Keeper) GetProposalID(ctx sdk.Context) (proposalID uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SetProposalID sets the new proposal ID to the store
func (keeper Keeper) SetProposalID(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}

func (keeper Keeper) ActivateVotingPeriod(ctx sdk.Context, proposal types.Proposal) {
	_ = "STUB: not implemented"
	return
}

func (keeper Keeper) MarshalProposal(proposal types.Proposal) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (keeper Keeper) UnmarshalProposal(bz []byte, proposal *types.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (keeper Keeper) MustMarshalProposal(proposal types.Proposal) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (keeper Keeper) MustUnmarshalProposal(bz []byte, proposal *types.Proposal) {
	_ = "STUB: not implemented"
	return
}
