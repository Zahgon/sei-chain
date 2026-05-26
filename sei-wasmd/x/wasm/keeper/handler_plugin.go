package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// msgEncoder is an extension point to customize encodings
type msgEncoder interface {
	// Encode converts wasmvm message to n cosmos message types
	Encode(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]sdk.Msg, error)
}

// MessageRouter ADR 031 request type routing
type MessageRouter interface {
	Handler(msg sdk.Msg) baseapp.MsgServiceHandler
}

// SDKMessageHandler can handles messages that can be encoded into sdk.Message types and routed.
type SDKMessageHandler struct {
	router   MessageRouter
	encoders msgEncoder
}

func NewDefaultMessageHandler(
	router MessageRouter,
	channelKeeper types.ChannelKeeper,
	capabilityKeeper types.CapabilityKeeper,
	bankKeeper types.Burner,
	unpacker codectypes.AnyUnpacker,
	portSource types.ICS20TransferPortSource,
	customEncoders ...*MessageEncoders,
) Messenger {
	_ = "STUB: not implemented"
	return *new(Messenger)
}

func NewSDKMessageHandler(router MessageRouter, encoders msgEncoder) SDKMessageHandler {
	_ = "STUB: not implemented"
	return *new(SDKMessageHandler)
}

func (h SDKMessageHandler) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) (events []sdk.Event, data [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// append data

// append events

func (h SDKMessageHandler) handleSdkMessage(ctx sdk.Context, contractAddr sdk.Address, msg sdk.Msg) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// make sure this account can send it

// find the handler and execute it

// ADR 031 request type routing

// legacy sdk.Msg routing
// Assuming that the app developer has migrated all their Msgs to
// proto messages and has registered all `Msg services`, then this
// path should never be called, because all those Msgs should be
// registered within the `msgServiceRouter` already.

// MessageHandlerChain defines a chain of handlers that are called one by one until it can be handled.
type MessageHandlerChain struct {
	handlers []Messenger
}

func NewMessageHandlerChain(first Messenger, others ...Messenger) *MessageHandlerChain {
	_ = "STUB: not implemented"
	return nil
}

// DispatchMsg dispatch message and calls chained handlers one after another in
// order to find the right one to process given message. If a handler cannot
// process given message (returns ErrUnknownMsg), its result is ignored and the
// next handler is executed.
func (m MessageHandlerChain) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]sdk.Event, [][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// IBCRawPacketHandler handels IBC.SendPacket messages which are published to an IBC channel.
type IBCRawPacketHandler struct {
	channelKeeper    types.ChannelKeeper
	capabilityKeeper types.CapabilityKeeper
}

func NewIBCRawPacketHandler(chk types.ChannelKeeper, cak types.CapabilityKeeper) IBCRawPacketHandler {
	_ = "STUB: not implemented"
	return *new(IBCRawPacketHandler)
}

// DispatchMsg publishes a raw IBC packet onto the channel.
func (h IBCRawPacketHandler) DispatchMsg(ctx sdk.Context, _ sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, _ wasmvmtypes.MessageInfo, _ types.CodeInfo) (events []sdk.Event, data [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var _ Messenger = MessageHandlerFunc(nil)

// MessageHandlerFunc is a helper to construct a function based message handler.
type MessageHandlerFunc func(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) (events []sdk.Event, data [][]byte, err error)

// DispatchMsg delegates dispatching of provided message into the MessageHandlerFunc.
func (m MessageHandlerFunc) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) (events []sdk.Event, data [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NewBurnCoinMessageHandler handles wasmvm.BurnMsg messages
func NewBurnCoinMessageHandler(burner types.Burner) MessageHandlerFunc {
	_ = "STUB: not implemented"
	return *new(MessageHandlerFunc)
}
