package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

var _ sdk.Msg = &MsgChannelOpenInit{}

// NewMsgChannelOpenInit creates a new MsgChannelOpenInit. It sets the counterparty channel
// identifier to be empty.
func NewMsgChannelOpenInit(
	portID, version string, channelOrder Order, connectionHops []string,
	counterpartyPortID string, signer string,
) *MsgChannelOpenInit {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenInit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgChannelOpenInit) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgChannelOpenTry{}

// NewMsgChannelOpenTry creates a new MsgChannelOpenTry instance
// The version string is deprecated and will be ignored by core IBC.
// It is left as an argument for go API backwards compatibility.
func NewMsgChannelOpenTry(
	portID, previousChannelID, version string, channelOrder Order, connectionHops []string,
	counterpartyPortID, counterpartyChannelID, counterpartyVersion string,
	proofInit []byte, proofHeight clienttypes.Height, signer string,
) *MsgChannelOpenTry {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenTry) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// counterparty validate basic allows empty counterparty channel identifiers

// GetSigners implements sdk.Msg
func (msg MsgChannelOpenTry) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgChannelOpenAck{}

// NewMsgChannelOpenAck creates a new MsgChannelOpenAck instance
func NewMsgChannelOpenAck(
	portID, channelID, counterpartyChannelID string, cpv string, proofTry []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgChannelOpenAck {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenAck) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgChannelOpenAck) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgChannelOpenConfirm{}

// NewMsgChannelOpenConfirm creates a new MsgChannelOpenConfirm instance
func NewMsgChannelOpenConfirm(
	portID, channelID string, proofAck []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgChannelOpenConfirm {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenConfirm) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgChannelOpenConfirm) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

var _ sdk.Msg = &MsgChannelCloseInit{}

// NewMsgChannelCloseInit creates a new MsgChannelCloseInit instance
func NewMsgChannelCloseInit(
	portID string, channelID string, signer string,
) *MsgChannelCloseInit {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelCloseInit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgChannelCloseInit) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgChannelCloseConfirm{}

// NewMsgChannelCloseConfirm creates a new MsgChannelCloseConfirm instance
func NewMsgChannelCloseConfirm(
	portID, channelID string, proofInit []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgChannelCloseConfirm {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelCloseConfirm) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgChannelCloseConfirm) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

var _ sdk.Msg = &MsgRecvPacket{}

// NewMsgRecvPacket constructs new MsgRecvPacket
func NewMsgRecvPacket(
	packet Packet, proofCommitment []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgRecvPacket {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgRecvPacket) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetDataSignBytes returns the base64-encoded bytes used for the
// data field when signing the packet.
func (msg MsgRecvPacket) GetDataSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgRecvPacket) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgTimeout{}

// NewMsgTimeout constructs new MsgTimeout
func NewMsgTimeout(
	packet Packet, nextSequenceRecv uint64, proofUnreceived []byte,
	proofHeight clienttypes.Height, signer string,
) *MsgTimeout {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgTimeout) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgTimeout) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// NewMsgTimeoutOnClose constructs new MsgTimeoutOnClose
func NewMsgTimeoutOnClose(
	packet Packet, nextSequenceRecv uint64,
	proofUnreceived, proofClose []byte,
	proofHeight clienttypes.Height, signer string,
) *MsgTimeoutOnClose {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgTimeoutOnClose) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgTimeoutOnClose) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgAcknowledgement{}

// NewMsgAcknowledgement constructs a new MsgAcknowledgement
func NewMsgAcknowledgement(
	packet Packet,
	ack, proofAcked []byte,
	proofHeight clienttypes.Height,
	signer string,
) *MsgAcknowledgement {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgAcknowledgement) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgAcknowledgement) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }
