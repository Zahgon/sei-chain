package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	upgradetypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

const (
	// ProposalTypeClientUpdate defines the type for a ClientUpdateProposal
	ProposalTypeClientUpdate = "ClientUpdate"
	ProposalTypeUpgrade      = "IBCUpgrade"
)

var (
	_ govtypes.Content                   = &ClientUpdateProposal{}
	_ govtypes.Content                   = &UpgradeProposal{}
	_ codectypes.UnpackInterfacesMessage = &UpgradeProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeClientUpdate)
	govtypes.RegisterProposalType(ProposalTypeUpgrade)
}

// NewClientUpdateProposal creates a new client update proposal.
func NewClientUpdateProposal(title, description, subjectClientID, substituteClientID string) govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

// GetTitle returns the title of a client update proposal.
func (cup *ClientUpdateProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the description of a client update proposal.
	return ""
}

func (cup *ClientUpdateProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a client update proposal.
func (cup *ClientUpdateProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// ProposalType returns the type of a client update proposal.
	return ""
}

func (cup *ClientUpdateProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic runs basic stateless validity checks
func (cup *ClientUpdateProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewUpgradeProposal creates a new IBC breaking upgrade proposal.
func NewUpgradeProposal(title, description string, plan upgradetypes.Plan, upgradedClientState exported.ClientState) (govtypes.Content, error) {
	_ = "STUB: not implemented"
	return *new(govtypes.Content), nil
}

// GetTitle returns the title of a upgrade proposal.
func (up *UpgradeProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the description of a upgrade proposal.
	return ""
}

func (up *UpgradeProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a upgrade proposal.
func (up *UpgradeProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// ProposalType returns the upgrade proposal type.
	return ""
}

func (up *UpgradeProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic runs basic stateless validity checks
func (up *UpgradeProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns the string representation of the UpgradeProposal.
func (up UpgradeProposal) String() string { _ = "STUB: not implemented"; return "" }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (up UpgradeProposal) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
