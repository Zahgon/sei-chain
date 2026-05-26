package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// TypeMsgCreateVestingAccount defines the type value for a MsgCreateVestingAccount.
const TypeMsgCreateVestingAccount = "msg_create_vesting_account"

var _ sdk.Msg = &MsgCreateVestingAccount{}

// NewMsgCreateVestingAccount returns a reference to a new MsgCreateVestingAccount.
func NewMsgCreateVestingAccount(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins, endTime int64, delayed bool, admin sdk.AccAddress) *MsgCreateVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// Route returns the message route for a MsgCreateVestingAccount.
func (msg MsgCreateVestingAccount) Route() string {
	_ = "STUB: not implemented"

	// Type returns the message type for a MsgCreateVestingAccount.
	return ""
}

func (msg MsgCreateVestingAccount) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic Implements Msg.
func (msg MsgCreateVestingAccount) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSignBytes returns the bytes all expected signers must sign over for a
// MsgCreateVestingAccount.
func (msg MsgCreateVestingAccount) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners returns the expected signers for a MsgCreateVestingAccount.
func (msg MsgCreateVestingAccount) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}
