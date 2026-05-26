package keeper

import (
	"context"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

var (
	_ clienttypes.MsgServer     = Keeper{}
	_ connectiontypes.MsgServer = Keeper{}
	_ channeltypes.MsgServer    = Keeper{}
)

// CreateClient defines a rpc handler method for MsgCreateClient.
func (k Keeper) CreateClient(goCtx context.Context, msg *clienttypes.MsgCreateClient) (*clienttypes.MsgCreateClientResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateClient defines a rpc handler method for MsgUpdateClient.
func (k Keeper) UpdateClient(goCtx context.Context, msg *clienttypes.MsgUpdateClient) (*clienttypes.MsgUpdateClientResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpgradeClient defines a rpc handler method for MsgUpgradeClient.
func (k Keeper) UpgradeClient(goCtx context.Context, msg *clienttypes.MsgUpgradeClient) (*clienttypes.MsgUpgradeClientResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SubmitMisbehaviour defines a rpc handler method for MsgSubmitMisbehaviour.
func (k Keeper) SubmitMisbehaviour(goCtx context.Context, msg *clienttypes.MsgSubmitMisbehaviour) (*clienttypes.MsgSubmitMisbehaviourResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionOpenInit defines a rpc handler method for MsgConnectionOpenInit.
func (k Keeper) ConnectionOpenInit(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenInit) (*connectiontypes.MsgConnectionOpenInitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// outbound gating: disallow outbound connection inits when outbound disabled

// ConnectionOpenTry defines a rpc handler method for MsgConnectionOpenTry.
func (k Keeper) ConnectionOpenTry(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenTry) (*connectiontypes.MsgConnectionOpenTryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionOpenAck defines a rpc handler method for MsgConnectionOpenAck.
func (k Keeper) ConnectionOpenAck(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenAck) (*connectiontypes.MsgConnectionOpenAckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionOpenConfirm defines a rpc handler method for MsgConnectionOpenConfirm.
func (k Keeper) ConnectionOpenConfirm(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenConfirm) (*connectiontypes.MsgConnectionOpenConfirmResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChannelOpenInit defines a rpc handler method for MsgChannelOpenInit.
// ChannelOpenInit will perform 04-channel checks, route to the application
// callback, and write an OpenInit channel into state upon successful execution.
func (k Keeper) ChannelOpenInit(goCtx context.Context, msg *channeltypes.MsgChannelOpenInit) (*channeltypes.MsgChannelOpenInitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// outbound gating: disallow outbound channel inits when outbound disabled

// Lookup module by port capability

// Retrieve application callbacks from router

// Perform 04-channel verification

// Perform application logic callback

// Write channel into state

// ChannelOpenTry defines a rpc handler method for MsgChannelOpenTry.
// ChannelOpenTry will perform 04-channel checks, route to the application
// callback, and write an OpenTry channel into state upon successful execution.
func (k Keeper) ChannelOpenTry(goCtx context.Context, msg *channeltypes.MsgChannelOpenTry) (*channeltypes.MsgChannelOpenTryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by port capability

// Retrieve application callbacks from router

// Perform 04-channel verification

// Perform application logic callback

// Write channel into state

// ChannelOpenAck defines a rpc handler method for MsgChannelOpenAck.
// ChannelOpenAck will perform 04-channel checks, route to the application
// callback, and write an OpenAck channel into state upon successful execution.
func (k Keeper) ChannelOpenAck(goCtx context.Context, msg *channeltypes.MsgChannelOpenAck) (*channeltypes.MsgChannelOpenAckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve application callbacks from router

// Perform 04-channel verification

// Perform application logic callback

// Write channel into state

// ChannelOpenConfirm defines a rpc handler method for MsgChannelOpenConfirm.
// ChannelOpenConfirm will perform 04-channel checks, route to the application
// callback, and write an OpenConfirm channel into state upon successful execution.
func (k Keeper) ChannelOpenConfirm(goCtx context.Context, msg *channeltypes.MsgChannelOpenConfirm) (*channeltypes.MsgChannelOpenConfirmResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve application callbacks from router

// Perform 04-channel verification

// Perform application logic callback

// Write channel into state

// ChannelCloseInit defines a rpc handler method for MsgChannelCloseInit.
func (k Keeper) ChannelCloseInit(goCtx context.Context, msg *channeltypes.MsgChannelCloseInit) (*channeltypes.MsgChannelCloseInitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve callbacks from router

// ChannelCloseConfirm defines a rpc handler method for MsgChannelCloseConfirm.
func (k Keeper) ChannelCloseConfirm(goCtx context.Context, msg *channeltypes.MsgChannelCloseConfirm) (*channeltypes.MsgChannelCloseConfirmResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve callbacks from router

// RecvPacket defines a rpc handler method for MsgRecvPacket.
func (k Keeper) RecvPacket(goCtx context.Context, msg *channeltypes.MsgRecvPacket) (*channeltypes.MsgRecvPacketResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve callbacks from router

// Perform TAO verification
//
// If the packet was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// NOTE: The context returned by CacheContext() refers to a new EventManager, so it needs to explicitly set events to the original context.

// Perform application logic callback
//
// Cache context so that we may discard state changes from callback if the acknowledgement is unsuccessful.

// write application state changes for asynchronous and successful acknowledgements

// NOTE: The context returned by CacheContext() refers to a new EventManager, so it needs to explicitly set events to the original context.
// Events from callback are emitted regardless of acknowledgement success

// Set packet acknowledgement only if the acknowledgement is not nil.
// NOTE: IBC applications modules may call the WriteAcknowledgement asynchronously if the
// acknowledgement is nil.

// Timeout defines a rpc handler method for MsgTimeout.
func (k Keeper) Timeout(goCtx context.Context, msg *channeltypes.MsgTimeout) (*channeltypes.MsgTimeoutResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve callbacks from router

// Perform TAO verification
//
// If the timeout was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// NOTE: The context returned by CacheContext() refers to a new EventManager, so it needs to explicitly set events to the original context.

// Perform application logic callback

// Delete packet commitment

// TimeoutOnClose defines a rpc handler method for MsgTimeoutOnClose.
func (k Keeper) TimeoutOnClose(goCtx context.Context, msg *channeltypes.MsgTimeoutOnClose) (*channeltypes.MsgTimeoutOnCloseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve callbacks from router

// Perform TAO verification
//
// If the timeout was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// NOTE: The context returned by CacheContext() refers to a new EventManager, so it needs to explicitly set events to the original context.

// Perform application logic callback
//
// NOTE: MsgTimeout and MsgTimeoutOnClose use the same "OnTimeoutPacket"
// application logic callback.

// Delete packet commitment

// Acknowledgement defines a rpc handler method for MsgAcknowledgement.
func (k Keeper) Acknowledgement(goCtx context.Context, msg *channeltypes.MsgAcknowledgement) (*channeltypes.MsgAcknowledgementResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup module by channel capability

// Retrieve callbacks from router

// Perform TAO verification
//
// If the acknowledgement was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// NOTE: The context returned by CacheContext() refers to a new EventManager, so it needs to explicitly set events to the original context.

// Perform application logic callback
