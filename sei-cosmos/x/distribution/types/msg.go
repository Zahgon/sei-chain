package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// distribution message types
const (
	TypeMsgSetWithdrawAddress          = "set_withdraw_address"
	TypeMsgWithdrawDelegatorReward     = "withdraw_delegator_reward"
	TypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
	TypeMsgFundCommunityPool           = "fund_community_pool"
)

// Verify interface at compile time
var _, _, _ sdk.Msg = &MsgSetWithdrawAddress{}, &MsgWithdrawDelegatorReward{}, &MsgWithdrawValidatorCommission{}

func NewMsgSetWithdrawAddress(delAddr, withdrawAddr sdk.AccAddress) *MsgSetWithdrawAddress {
	_ = "STUB: not implemented"
	return nil
}

func (msg MsgSetWithdrawAddress) Route() string { _ = "STUB: not implemented"; return "" }
func (msg MsgSetWithdrawAddress) Type() string  { _ = "STUB: not implemented"; return "" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgSetWithdrawAddress) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// get the bytes for the message signer to sign on
func (msg MsgSetWithdrawAddress) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// quick validity check
func (msg MsgSetWithdrawAddress) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func NewMsgWithdrawDelegatorReward(delAddr sdk.AccAddress, valAddr sdk.ValAddress) *MsgWithdrawDelegatorReward {
	_ = "STUB: not implemented"
	return nil
}

func (msg MsgWithdrawDelegatorReward) Route() string { _ = "STUB: not implemented"; return "" }
func (msg MsgWithdrawDelegatorReward) Type() string  { _ = "STUB: not implemented"; return "" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawDelegatorReward) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawDelegatorReward) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// quick validity check
func (msg MsgWithdrawDelegatorReward) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func NewMsgWithdrawValidatorCommission(valAddr sdk.ValAddress) *MsgWithdrawValidatorCommission {
	_ = "STUB: not implemented"
	return nil
}

func (msg MsgWithdrawValidatorCommission) Route() string { _ = "STUB: not implemented"; return "" }
func (msg MsgWithdrawValidatorCommission) Type() string  { _ = "STUB: not implemented"; return "" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawValidatorCommission) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawValidatorCommission) GetSignBytes() []byte {
	_ = "STUB: not implemented"
	return nil
}

// quick validity check
func (msg MsgWithdrawValidatorCommission) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgFundCommunityPool returns a new MsgFundCommunityPool with a sender and
// a funding amount.
func NewMsgFundCommunityPool(amount sdk.Coins, depositor sdk.AccAddress) *MsgFundCommunityPool {
	_ = "STUB: not implemented"
	return nil
}

// Route returns the MsgFundCommunityPool message route.
func (msg MsgFundCommunityPool) Route() string {
	_ = "STUB: not implemented"

	// Type returns the MsgFundCommunityPool message type.
	return ""
}

func (msg MsgFundCommunityPool) Type() string { _ = "STUB: not implemented"; return "" }

// GetSigners returns the signer addresses that are expected to sign the result
// of GetSignBytes.
func (msg MsgFundCommunityPool) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// GetSignBytes returns the raw bytes for a MsgFundCommunityPool message that
// the expected signer needs to sign.
func (msg MsgFundCommunityPool) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// ValidateBasic performs basic MsgFundCommunityPool message validation.
func (msg MsgFundCommunityPool) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
