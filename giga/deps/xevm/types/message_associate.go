package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const TypeMsgAssociate = "evm_associate"

var (
	_ sdk.Msg = &MsgAssociate{}
)

func NewMsgAssociate(sender sdk.AccAddress, customMsg string) *MsgAssociate {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgAssociate) Route() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgAssociate) Type() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgAssociate) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgAssociate) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg *MsgAssociate) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func IsTxMsgAssociate(tx sdk.Tx) bool { _ = "STUB: not implemented"; return false }
