package simulation

import (
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

// OpWeightSubmitParamChangeProposal app params key for param change proposal
const OpWeightSubmitParamChangeProposal = "op_weight_submit_param_change_proposal"

// ProposalContents defines the module weighted proposals' contents
func ProposalContents(paramChanges []simtypes.ParamChange) []simtypes.WeightedProposalContent {
	_ = "STUB: not implemented"
	return nil
}
