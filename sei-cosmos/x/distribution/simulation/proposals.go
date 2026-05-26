package simulation

import (
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
)

// OpWeightSubmitCommunitySpendProposal app params key for community spend proposal
const OpWeightSubmitCommunitySpendProposal = "op_weight_submit_community_spend_proposal"

// ProposalContents defines the module weighted proposals' contents
func ProposalContents(k keeper.Keeper) []simtypes.WeightedProposalContent {
	_ = "STUB: not implemented"
	return nil
}

// SimulateCommunityPoolSpendProposalContent generates random community-pool-spend proposal content
func SimulateCommunityPoolSpendProposalContent(k keeper.Keeper) simtypes.ContentSimulatorFn {
	_ = "STUB: not implemented"
	return *new(simtypes.ContentSimulatorFn)
}
