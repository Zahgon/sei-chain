package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgVerifyInvariant{}

// NewMsgVerifyInvariant creates a new MsgVerifyInvariant object
func NewMsgVerifyInvariant(sender sdk.AccAddress, invModeName, invRoute string) *MsgVerifyInvariant {
	_ = "STUB: not implemented"
	return nil
}

func (msg MsgVerifyInvariant) Route() string { _ = "STUB: not implemented"; return "" }
func (msg MsgVerifyInvariant) Type() string  { _ = "STUB: not implemented"; return "" }

// get the bytes for the message signer to sign on
func (msg MsgVerifyInvariant) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// GetSignBytes gets the sign bytes for the msg MsgVerifyInvariant
func (msg MsgVerifyInvariant) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// quick validity check
func (msg MsgVerifyInvariant) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// FullInvariantRoute - get the messages full invariant route
func (msg MsgVerifyInvariant) FullInvariantRoute() string { _ = "STUB: not implemented"; return "" }
