package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	_ sdk.Msg = &MsgConnectionOpenInit{}
	_ sdk.Msg = &MsgConnectionOpenConfirm{}
	_ sdk.Msg = &MsgConnectionOpenAck{}
	_ sdk.Msg = &MsgConnectionOpenTry{}

	_ codectypes.UnpackInterfacesMessage = MsgConnectionOpenTry{}
	_ codectypes.UnpackInterfacesMessage = MsgConnectionOpenAck{}
)

// NewMsgConnectionOpenInit creates a new MsgConnectionOpenInit instance. It sets the
// counterparty connection identifier to be empty.
func NewMsgConnectionOpenInit(
	clientID, counterpartyClientID string,
	counterpartyPrefix commitmenttypes.MerklePrefix,
	version *Version, delayPeriod uint64, signer string,
) *MsgConnectionOpenInit {
	_ = "STUB: not implemented"
	// counterparty must have the same delay period
	return nil
}

// ValidateBasic implements sdk.Msg.
func (msg MsgConnectionOpenInit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Version can be nil on MsgConnectionOpenInit

// GetSigners implements sdk.Msg
func (msg MsgConnectionOpenInit) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgConnectionOpenTry creates a new MsgConnectionOpenTry instance
func NewMsgConnectionOpenTry(
	previousConnectionID, clientID, counterpartyConnectionID,
	counterpartyClientID string, counterpartyClient exported.ClientState,
	counterpartyPrefix commitmenttypes.MerklePrefix,
	counterpartyVersions []*Version, delayPeriod uint64,
	proofInit, proofClient, proofConsensus []byte,
	proofHeight, consensusHeight clienttypes.Height, signer string,
) *MsgConnectionOpenTry {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgConnectionOpenTry) ValidateBasic() error {
	_ = "STUB: not implemented"
	// an empty connection identifier indicates that a connection identifier should be generated
	return nil
}

// counterparty validate basic allows empty counterparty connection identifiers

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgConnectionOpenTry) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgConnectionOpenTry) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgConnectionOpenAck creates a new MsgConnectionOpenAck instance
func NewMsgConnectionOpenAck(
	connectionID, counterpartyConnectionID string, counterpartyClient exported.ClientState,
	proofTry, proofClient, proofConsensus []byte,
	proofHeight, consensusHeight clienttypes.Height,
	version *Version,
	signer string,
) *MsgConnectionOpenAck {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgConnectionOpenAck) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgConnectionOpenAck) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgConnectionOpenAck) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgConnectionOpenConfirm creates a new MsgConnectionOpenConfirm instance
func NewMsgConnectionOpenConfirm(
	connectionID string, proofAck []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgConnectionOpenConfirm {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgConnectionOpenConfirm) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners implements sdk.Msg
func (msg MsgConnectionOpenConfirm) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}
