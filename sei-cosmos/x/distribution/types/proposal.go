package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	// ProposalTypeCommunityPoolSpend defines the type for a CommunityPoolSpendProposal
	ProposalTypeCommunityPoolSpend = "CommunityPoolSpend"
)

// Assert CommunityPoolSpendProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &CommunityPoolSpendProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCommunityPoolSpend)
	govtypes.RegisterProposalTypeCodec(&CommunityPoolSpendProposal{}, "cosmos-sdk/CommunityPoolSpendProposal")
}

// NewCommunityPoolSpendProposal creates a new community pool spned proposal.
func NewCommunityPoolSpendProposal(title, description string, recipient sdk.AccAddress, amount sdk.Coins) *CommunityPoolSpendProposal {
	_ = "STUB: not implemented"
	return nil
}

// GetTitle returns the title of a community pool spend proposal.
func (csp *CommunityPoolSpendProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the description of a community pool spend proposal.
	return ""
}

func (csp *CommunityPoolSpendProposal) GetDescription() string {
	_ = "STUB: not implemented"
	return ""

	// GetDescription returns the routing key of a community pool spend proposal.
}

func (csp *CommunityPoolSpendProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// ProposalType returns the type of a community pool spend proposal.
	return ""
}

func (csp *CommunityPoolSpendProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic runs basic stateless validity checks
func (csp *CommunityPoolSpendProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (csp CommunityPoolSpendProposal) String() string { _ = "STUB: not implemented"; return "" }
