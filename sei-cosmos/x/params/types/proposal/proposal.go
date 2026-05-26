package proposal

import (
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	// ProposalTypeChange defines the type for a ParameterChangeProposal
	ProposalTypeChange = "ParameterChange"
)

// Assert ParameterChangeProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &ParameterChangeProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeChange)
	govtypes.RegisterProposalTypeCodec(&ParameterChangeProposal{}, "cosmos-sdk/ParameterChangeProposal")
}

func NewParameterChangeProposal(title, description string, changes []ParamChange, isExpedited bool) *ParameterChangeProposal {
	_ = "STUB: not implemented"
	return nil
}

// GetTitle returns the title of a parameter change proposal.
func (pcp *ParameterChangeProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the description of a parameter change proposal.
	return ""
}

func (pcp *ParameterChangeProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a parameter change proposal.
func (pcp *ParameterChangeProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// ProposalType returns the type of a parameter change proposal.
	return ""
}

func (pcp *ParameterChangeProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the parameter change proposal
func (pcp *ParameterChangeProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (pcp ParameterChangeProposal) String() string { _ = "STUB: not implemented"; return "" }

func NewParamChange(subspace, key, value string) ParamChange {
	_ = "STUB: not implemented"
	return *new(ParamChange)
}

// String implements the Stringer interface.
func (pc ParamChange) String() string { _ = "STUB: not implemented"; return "" }

// ValidateChanges performs basic validation checks over a set of ParamChange. It
// returns an error if any ParamChange is invalid.
func ValidateChanges(changes []ParamChange) error { _ = "STUB: not implemented"; return nil }

// We need to verify ConsensusParams since they are only validated once the proposal passes.
// If any of them are invalid at time of passing, this will cause a chain halt since validation is done during
// ApplyBlock: https://github.com/sei-protocol/sei-tendermint/blob/d426f1fe475eb0c406296770ff5e9f8869b3887e/internal/state/execution.go#L320
// Therefore, we validate when we get a param-change msg for ConsensusParams

func verifyConsensusParamsUsingDefault(changes []ParamChange) error {
	_ = "STUB: not implemented"
	// Start with a default (valid) set of parameters, and update based on proposal then check
	return nil
}

// Note: BlockParams seems to be the only support ConsensusParams available for modifying with proposal
