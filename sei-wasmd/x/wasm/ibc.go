package wasm

import (
	ibcexported "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	porttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"

	types "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var _ porttypes.IBCModule = IBCHandler{}

type IBCHandler struct {
	keeper        types.IBCContractKeeper
	channelKeeper types.ChannelKeeper
}

func NewIBCHandler(k types.IBCContractKeeper, ck types.ChannelKeeper) IBCHandler {
	_ = "STUB: not implemented"
	return *new(IBCHandler)
}

// OnChanOpenInit implements the IBCModule interface
func (i IBCHandler) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterParty channeltypes.Counterparty,
	version string,
) error {
	_ = "STUB: not implemented"
	// ensure port, version, capability
	return nil
}

// DESIGN V3: this may be "" ??

// At the moment this list must be of length 1. In the future multi-hop channels may be supported.

// Claim channel capability passed back by IBC module

// OnChanOpenTry implements the IBCModule interface
func (i IBCHandler) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID, channelID string,
	chanCap *capabilitytypes.Capability,
	counterParty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	_ = "STUB: not implemented"
	// ensure port, version, capability
	return "", nil
}

// At the moment this list must be of length 1. In the future multi-hop channels may be supported.

// Allow contracts to return a version (or default to counterpartyVersion if unset)

// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
// If module can already authenticate the capability then module already owns it so we don't need to claim
// Otherwise, module does not have channel capability and we must claim it from IBC

// Only claim channel capability passed back by IBC module if we do not already own it

// OnChanOpenAck implements the IBCModule interface
func (i IBCHandler) OnChanOpenAck(
	ctx sdk.Context,
	portID, channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// This is a bit ugly, but it is set AFTER the callback is done, yet we want to provide the contract
// access to the channel in queries. We can revisit how to better integrate with ibc-go in the future,
// but this is the best/safest we can do now. (If you remove this, you error when sending a packet during the
// OnChanOpenAck entry point)
// https://github.com/cosmos/ibc-go/pull/647/files#diff-54b5be375a2333c56f2ae1b5b4dc13ac9c734561e30286505f39837ee75762c7R25

// OnChanOpenConfirm implements the IBCModule interface
func (i IBCHandler) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (i IBCHandler) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// emit events?

// OnChanCloseConfirm implements the IBCModule interface
func (i IBCHandler) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	// counterparty has closed the channel
	return nil
}

// emit events?

func toWasmVMChannel(portID, channelID string, channelInfo channeltypes.Channel) wasmvmtypes.IBCChannel {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.IBCChannel)
}

// At the moment this list must be of length 1. In the future multi-hop channels may be supported.

// OnRecvPacket implements the IBCModule interface
func (i IBCHandler) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

var _ ibcexported.Acknowledgement = ContractConfirmStateAck{}

type ContractConfirmStateAck []byte

func (w ContractConfirmStateAck) Success() bool {
	_ = "STUB: not implemented"
	// always commit state
	return false
}

func (w ContractConfirmStateAck) Acknowledgement() []byte {
	_ = "STUB: not implemented"

	// OnAcknowledgementPacket implements the IBCModule interface
	return nil
}

func (i IBCHandler) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (i IBCHandler) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IBCHandler) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (version string, err error) {
	_ = "STUB: not implemented"
	return "",
		// accept all
		nil
}

func newIBCPacket(packet channeltypes.Packet) wasmvmtypes.IBCPacket {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.IBCPacket)
}

func ValidateChannelParams(channelID string) error {
	_ = "STUB: not implemented"
	// NOTE: for escrow address security only 2^32 channels are allowed to be created
	// Issue: https://github.com/cosmos/cosmos-sdk/issues/7737
	return nil
}
