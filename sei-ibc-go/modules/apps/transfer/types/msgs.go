package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

// msg types
const (
	TypeMsgTransfer   = "transfer"
	MaximumMemoLength = 32768 // maximum length of the memo in bytes (value chosen arbitrarily)
)

// NewMsgTransfer creates a new MsgTransfer instance
func NewMsgTransfer(
	sourcePort, sourceChannel string,
	token sdk.Coin, sender, receiver string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
) *MsgTransfer {
	_ = "STUB: not implemented"
	return nil
}

// Route implements sdk.Msg
func (MsgTransfer) Route() string {
	_ = "STUB: not implemented"

	// Type implements sdk.Msg
	return ""
}

func (MsgTransfer) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic performs a basic check of the MsgTransfer fields.
// NOTE: timeout height or timestamp values can be 0 to disable the timeout.
// NOTE: The recipient addresses format is not validated as the format defined by
// the chain is not known to IBC.
func (msg MsgTransfer) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: sender format must be validated as it is required by the GetSigners function.

// GetSignBytes implements sdk.Msg.
func (msg MsgTransfer) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgTransfer) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }
