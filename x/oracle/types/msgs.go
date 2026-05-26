package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// ensure Msg interface compliance at compile time
var (
	_ sdk.Msg = &MsgDelegateFeedConsent{}
	_ sdk.Msg = &MsgAggregateExchangeRateVote{}
)

// oracle message types
const (
	TypeMsgDelegateFeedConsent       = "delegate_feeder"
	TypeMsgAggregateExchangeRateVote = "aggregate_exchange_rate_vote"
)

//-------------------------------------------------
//-------------------------------------------------

// NewMsgAggregateExchangeRateVote returns MsgAggregateExchangeRateVote instance
func NewMsgAggregateExchangeRateVote(exchangeRates string, feeder sdk.AccAddress, validator sdk.ValAddress) *MsgAggregateExchangeRateVote {
	_ = "STUB: not implemented"
	return nil
}

// Route implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) Route() string {
	_ = "STUB: not implemented"

	// Type implements sdk.Msg
	return ""
}

func (msg MsgAggregateExchangeRateVote) Type() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) GetSignBytes() []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgAggregateExchangeRateVote) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

// Check overflow bit length

// NewMsgDelegateFeedConsent creates a MsgDelegateFeedConsent instance
func NewMsgDelegateFeedConsent(operatorAddress sdk.ValAddress, feederAddress sdk.AccAddress) *MsgDelegateFeedConsent {
	_ = "STUB: not implemented"
	return nil
}

// Route implements sdk.Msg
func (msg MsgDelegateFeedConsent) Route() string {
	_ = "STUB: not implemented"

	// Type implements sdk.Msg
	return ""
}

func (msg MsgDelegateFeedConsent) Type() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements sdk.Msg
func (msg MsgDelegateFeedConsent) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgDelegateFeedConsent) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgDelegateFeedConsent) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
