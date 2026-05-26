package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	_ sdk.Msg = &MsgInternalEVMCall{}
)

func NewMessageInternalEVMCall(from sdk.AccAddress, to string, value *sdk.Int, data []byte) *MsgInternalEVMCall {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgInternalEVMCall) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgInternalEVMCall) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
