package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// slashing message types
const (
	TypeMsgUnjail = "unjail"
)

// verify interface at compile time
var _ sdk.Msg = &MsgUnjail{}

// NewMsgUnjail creates a new MsgUnjail instance
func NewMsgUnjail(validatorAddr sdk.ValAddress) *MsgUnjail { _ = "STUB: not implemented"; return nil }

func (msg MsgUnjail) Route() string                { _ = "STUB: not implemented"; return "" }
func (msg MsgUnjail) Type() string                 { _ = "STUB: not implemented"; return "" }
func (msg MsgUnjail) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUnjail) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// ValidateBasic validity check for the AnteHandler
func (msg MsgUnjail) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
