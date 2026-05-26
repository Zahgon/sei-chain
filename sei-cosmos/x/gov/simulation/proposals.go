package simulation

import (
	"math/rand"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

// OpWeightSubmitTextProposal app params key for text proposal
const OpWeightSubmitTextProposal = "op_weight_submit_text_proposal"

// ProposalContents defines the module weighted proposals' contents
func ProposalContents() []simtypes.WeightedProposalContent { _ = "STUB: not implemented"; return nil }

// SimulateTextProposalContent returns a random text proposal content.
func SimulateTextProposalContent(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) simtypes.Content {
	_ = "STUB: not implemented"
	return *new(simtypes.Content)
}
