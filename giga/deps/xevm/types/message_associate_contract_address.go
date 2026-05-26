package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const TypeMsgAssociateContractAddress = "evm_associate_contract_address"

var (
	_ sdk.Msg = &MsgAssociateContractAddress{}
)

func NewMsgAssociateContractAddress(sender sdk.AccAddress, addr sdk.AccAddress) *MsgAssociateContractAddress {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgAssociateContractAddress) Route() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgAssociateContractAddress) Type() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgAssociateContractAddress) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgAssociateContractAddress) GetSignBytes() []byte {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgAssociateContractAddress) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}
