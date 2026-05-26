package types

import (
	gov "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	ProposalTypeSoftwareUpgrade       string = "SoftwareUpgrade"
	ProposalTypeCancelSoftwareUpgrade string = "CancelSoftwareUpgrade"
)

func NewSoftwareUpgradeProposal(title, description string, plan Plan) gov.Content {
	_ = "STUB: not implemented"
	return *new(gov.Content)
}

// Implements Proposal Interface
var _ gov.Content = &SoftwareUpgradeProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeSoftwareUpgrade)
	gov.RegisterProposalTypeCodec(&SoftwareUpgradeProposal{}, "cosmos-sdk/SoftwareUpgradeProposal")
	gov.RegisterProposalType(ProposalTypeCancelSoftwareUpgrade)
	gov.RegisterProposalTypeCodec(&CancelSoftwareUpgradeProposal{}, "cosmos-sdk/CancelSoftwareUpgradeProposal")
}

func (sup *SoftwareUpgradeProposal) GetTitle() string       { _ = "STUB: not implemented"; return "" }
func (sup *SoftwareUpgradeProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }
func (sup *SoftwareUpgradeProposal) ProposalRoute() string  { _ = "STUB: not implemented"; return "" }
func (sup *SoftwareUpgradeProposal) ProposalType() string   { _ = "STUB: not implemented"; return "" }
func (sup *SoftwareUpgradeProposal) ValidateBasic() error   { _ = "STUB: not implemented"; return nil }

func (sup SoftwareUpgradeProposal) String() string { _ = "STUB: not implemented"; return "" }

func NewCancelSoftwareUpgradeProposal(title, description string) gov.Content {
	_ = "STUB: not implemented"
	return *new(gov.Content)
}

// Implements Proposal Interface
var _ gov.Content = &CancelSoftwareUpgradeProposal{}

func (csup *CancelSoftwareUpgradeProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }
func (csup *CancelSoftwareUpgradeProposal) GetDescription() string {
	_ = "STUB: not implemented"
	return ""
}
func (csup *CancelSoftwareUpgradeProposal) ProposalRoute() string {
	_ = "STUB: not implemented"
	return ""
}
func (csup *CancelSoftwareUpgradeProposal) ProposalType() string {
	_ = "STUB: not implemented"
	return ""
}

func (csup *CancelSoftwareUpgradeProposal) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

func (csup CancelSoftwareUpgradeProposal) String() string { _ = "STUB: not implemented"; return "" }
