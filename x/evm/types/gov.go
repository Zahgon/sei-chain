package types

import (
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	ProposalTypeAddERCNativePointer   = "AddERCNativePointer"
	ProposalTypeAddERCCW20Pointer     = "AddERCCW20Pointer"
	ProposalTypeAddERCCW721Pointer    = "AddERCCW721Pointer"
	ProposalTypeAddERCCW1155Pointer   = "AddERCCW1155Pointer"
	ProposalTypeAddCWERC20Pointer     = "AddCWERC20Pointer"
	ProposalTypeAddCWERC721Pointer    = "AddCWERC721Pointer"
	ProposalTypeAddCWERC1155Pointer   = "AddCWERC1155Pointer"
	ProposalTypeAddERCNativePointerV2 = "AddERCNativePointerV2"
)

func init() {
	// for routing
	govtypes.RegisterProposalType(ProposalTypeAddERCNativePointer)
	govtypes.RegisterProposalType(ProposalTypeAddERCCW20Pointer)
	govtypes.RegisterProposalType(ProposalTypeAddERCCW721Pointer)
	govtypes.RegisterProposalType(ProposalTypeAddERCCW1155Pointer)
	govtypes.RegisterProposalType(ProposalTypeAddCWERC20Pointer)
	govtypes.RegisterProposalType(ProposalTypeAddCWERC721Pointer)
	govtypes.RegisterProposalType(ProposalTypeAddCWERC1155Pointer)
	govtypes.RegisterProposalType(ProposalTypeAddERCNativePointerV2)

	// for marshal and unmarshal
	govtypes.RegisterProposalTypeCodec(&AddERCNativePointerProposal{}, "evm/AddERCNativePointerProposal")
	govtypes.RegisterProposalTypeCodec(&AddERCCW20PointerProposal{}, "evm/AddERCCW20PointerProposal")
	govtypes.RegisterProposalTypeCodec(&AddERCCW721PointerProposal{}, "evm/AddERCCW721PointerProposal")
	govtypes.RegisterProposalTypeCodec(&AddERCCW1155PointerProposal{}, "evm/AddERCCW1155PointerProposal")
	govtypes.RegisterProposalTypeCodec(&AddCWERC20PointerProposal{}, "evm/AddCWERC20PointerProposal")
	govtypes.RegisterProposalTypeCodec(&AddCWERC721PointerProposal{}, "evm/AddCWERC721PointerProposal")
	govtypes.RegisterProposalTypeCodec(&AddCWERC1155PointerProposal{}, "evm/AddCWERC1155PointerProposal")
	govtypes.RegisterProposalTypeCodec(&AddERCNativePointerProposalV2{}, "evm/AddERCNativePointerProposalV2")
}

func (p *AddERCNativePointerProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCNativePointerProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCNativePointerProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCNativePointerProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCNativePointerProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p AddERCNativePointerProposal) String() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW20PointerProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW20PointerProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW20PointerProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW20PointerProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW20PointerProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p AddERCCW20PointerProposal) String() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW721PointerProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW721PointerProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW721PointerProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW721PointerProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW721PointerProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p AddERCCW721PointerProposal) String() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW1155PointerProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW1155PointerProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW1155PointerProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW1155PointerProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCCW1155PointerProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p AddERCCW1155PointerProposal) String() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC20PointerProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC20PointerProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC20PointerProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC20PointerProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC20PointerProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p AddCWERC20PointerProposal) String() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC721PointerProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC721PointerProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC721PointerProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC721PointerProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC721PointerProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p AddCWERC721PointerProposal) String() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC1155PointerProposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC1155PointerProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC1155PointerProposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC1155PointerProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddCWERC1155PointerProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (p AddCWERC1155PointerProposal) String() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCNativePointerProposalV2) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCNativePointerProposalV2) GetDescription() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *AddERCNativePointerProposalV2) ProposalRoute() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *AddERCNativePointerProposalV2) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p *AddERCNativePointerProposalV2) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

func (p AddERCNativePointerProposalV2) String() string { _ = "STUB: not implemented"; return "" }
