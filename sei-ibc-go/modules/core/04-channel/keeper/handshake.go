package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	sdkerrors "github.com/sei-protocol/sei-chain/sei-cosmos/types/errors"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var logger = seilog.NewLogger("ibc-go", "modules", "core", "04-channel", "keeper")

// ErrInboundDisabledHandshake is the error for when inbound is disabled during channel handshake
var ErrInboundDisabledHandshake = sdkerrors.Register("ibc-channel-handshake", 101, "ibc inbound disabled")

// ErrOutboundDisabledHandshake is the error for when outbound is disabled during channel handshake
var ErrOutboundDisabledHandshake = sdkerrors.Register("ibc-channel-handshake", 102, "ibc outbound disabled")

// ChanOpenInit is called by a module to initiate a channel opening handshake with
// a module on another chain. The counterparty channel identifier is validated to be
// empty in msg validation.
func (k Keeper) ChanOpenInit(
	ctx sdk.Context,
	order types.Order,
	connectionHops []string,
	portID string,
	portCap *capabilitytypes.Capability,
	counterparty types.Counterparty,
	version string,
) (string, *capabilitytypes.Capability, error) {
	_ = "STUB: not implemented"
	// outbound gating: disallow outbound channel inits when outbound disabled
	return "", nil, nil
}

// connection hop length checked on msg.ValidateBasic()

// WriteOpenInitChannel writes a channel which has successfully passed the OpenInit handshake step.
// The channel is set in state and all the associated Send and Recv sequences are set to 1.
// An event is emitted for the handshake step.
func (k Keeper) WriteOpenInitChannel(
	ctx sdk.Context,
	portID,
	channelID string,
	order types.Order,
	connectionHops []string,
	counterparty types.Counterparty,
	version string,
) {
	_ = "STUB: not implemented"
	return
}

// ChanOpenTry is called by a module to accept the first step of a channel opening
// handshake initiated by a module on another chain.
func (k Keeper) ChanOpenTry(
	ctx sdk.Context,
	order types.Order,
	connectionHops []string,
	portID,
	previousChannelID string,
	portCap *capabilitytypes.Capability,
	counterparty types.Counterparty,
	counterpartyVersion string,
	proofInit []byte,
	proofHeight exported.Height,
) (string, *capabilitytypes.Capability, error) {
	_ = "STUB: not implemented"
	// inbound gating: disallow inbound channel tries when inbound disabled
	return "", nil, nil
}

// connection hops only supports a single connection

// empty channel identifier indicates continuing a previous channel handshake

// channel identifier and connection hop length checked on msg.ValidateBasic()
// ensure that the previous channel exists

// previous channel must use the same fields

// ChanOpenInit will only set a single connection hop

// generate a new channel

// expectedCounterpaty is the counterparty of the counterparty's channel end
// (i.e self)

// capability initialized in ChanOpenInit

// WriteOpenTryChannel writes a channel which has successfully passed the OpenTry handshake step.
// The channel is set in state. If a previous channel state did not exist, all the Send and Recv
// sequences are set to 1. An event is emitted for the handshake step.
func (k Keeper) WriteOpenTryChannel(
	ctx sdk.Context,
	portID,
	channelID string,
	order types.Order,
	connectionHops []string,
	counterparty types.Counterparty,
	version string,
) {
	_ = "STUB: not implemented"
	return
}

// ChanOpenAck is called by the handshake-originating module to acknowledge the
// acceptance of the initial request by the counterparty module on the other chain.
func (k Keeper) ChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterpartyVersion,
	counterpartyChannelID string,
	proofTry []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	return nil
}

// counterparty of the counterparty channel end (i.e self)

// WriteOpenAckChannel writes an updated channel state for the successful OpenAck handshake step.
// An event is emitted for the handshake step.
func (k Keeper) WriteOpenAckChannel(
	ctx sdk.Context,
	portID,
	channelID,
	counterpartyVersion,
	counterpartyChannelID string,
) {
	_ = "STUB: not implemented"
	return
}

// ChanOpenConfirm is called by the counterparty module to close their end of the
// channel, since the other end has been closed.
func (k Keeper) ChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	proofAck []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteOpenConfirmChannel writes an updated channel state for the successful OpenConfirm handshake step.
// An event is emitted for the handshake step.
func (k Keeper) WriteOpenConfirmChannel(
	ctx sdk.Context,
	portID,
	channelID string,
) {
	_ = "STUB: not implemented"
	return
}

// Closing Handshake
//
// This section defines the set of functions required to close a channel handshake
// as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-004-channel-and-packet-semantics#closing-handshake
//
// ChanCloseInit is called by either module to close their end of the channel. Once
// closed, channels cannot be reopened.
func (k Keeper) ChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ChanCloseConfirm is called by the counterparty module to close their end of the
// channel, since the other end has been closed.
func (k Keeper) ChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	proofInit []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	return nil
}
