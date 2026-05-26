package types

import (
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	ProposalTypeUpdateMinter = "UpdateMinter"
)

func init() {
	// for routing
	govtypes.RegisterProposalType(ProposalTypeUpdateMinter)
	// for marshal and unmarshal
	govtypes.RegisterProposalTypeCodec(&UpdateMinterProposal{}, "mint/UpdateMinterProposal")
}

func (p *UpdateMinterProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *UpdateMinterProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *UpdateMinterProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *UpdateMinterProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *UpdateMinterProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p UpdateMinterProposal) String() string { _ = "STUB: not implemented"; return "" }

func NewUpdateMinterProposalHandler(title, description string, minter Minter) *UpdateMinterProposal {
	_ = "STUB: not implemented"
	return nil
}
