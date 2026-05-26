package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

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
