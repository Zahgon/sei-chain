package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	_ sdk.Msg = &MsgInternalEVMDelegateCall{}
)

func NewMessageInternalEVMDelegateCall(from sdk.AccAddress, to string, codeHash []byte, data []byte, fromContract string) *MsgInternalEVMDelegateCall {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgInternalEVMDelegateCall) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgInternalEVMDelegateCall) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
