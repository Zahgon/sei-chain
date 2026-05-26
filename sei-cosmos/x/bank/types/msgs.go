package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// bank message types
const (
	TypeMsgSend      = "send"
	TypeMsgMultiSend = "multisend"
)

var _ sdk.Msg = &MsgSend{}

// NewMsgSend - construct a msg to send coins from one account to another.
func NewMsgSend(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) *MsgSend {
	_ = "STUB: not implemented"
	return nil
}

// Route Implements Msg.
func (msg MsgSend) Route() string {
	_ = "STUB: not implemented"

	// Type Implements Msg.
	return ""
}

func (msg MsgSend) Type() string {
	_ = "STUB: not implemented"

	// ValidateBasic Implements Msg.
	return ""
}

func (msg MsgSend) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSignBytes Implements Msg.
func (msg MsgSend) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners Implements Msg.
func (msg MsgSend) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgMultiSend{}

// NewMsgMultiSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgMultiSend(in []Input, out []Output) *MsgMultiSend { _ = "STUB: not implemented"; return nil }

// Route Implements Msg
func (msg MsgMultiSend) Route() string {
	_ = "STUB: not implemented"

	// Type Implements Msg
	return ""
}

func (msg MsgMultiSend) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic Implements Msg.
func (msg MsgMultiSend) ValidateBasic() error {
	_ = "STUB: not implemented"
	// this just makes sure all the inputs and outputs are properly formatted,
	// not that they actually have the money inside
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMultiSend) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners Implements Msg.
func (msg MsgMultiSend) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// ValidateBasic - validate transaction input
func (in Input) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewInput - create a transaction input, used with MsgMultiSend
func NewInput(addr sdk.AccAddress, coins sdk.Coins) Input {
	_ = "STUB: not implemented"
	return *new(Input)
}

// ValidateBasic - validate transaction output
func (out Output) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewOutput - create a transaction output, used with MsgMultiSend
func NewOutput(addr sdk.AccAddress, coins sdk.Coins) Output {
	_ = "STUB: not implemented"
	return *new(Output)
}

// ValidateInputsOutputs validates that each respective input and output is
// valid and that the sum of inputs is equal to the sum of outputs.
func ValidateInputsOutputs(inputs []Input, outputs []Output) error {
	_ = "STUB: not implemented"
	return nil
}

// make sure inputs and outputs match
