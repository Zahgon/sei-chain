package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// EmitChannelOpenInitEvent emits a channel open init event
func EmitChannelOpenInitEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// EmitChannelOpenTryEvent emits a channel open try event
func EmitChannelOpenTryEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// EmitChannelOpenAckEvent emits a channel open acknowledge event
func EmitChannelOpenAckEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// EmitChannelOpenConfirmEvent emits a channel open confirm event
func EmitChannelOpenConfirmEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// EmitChannelCloseInitEvent emits a channel close init event
func EmitChannelCloseInitEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// EmitChannelCloseConfirmEvent emits a channel close confirm event
func EmitChannelCloseConfirmEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// EmitSendPacketEvent emits an event with packet data along with other packet information for relayer
// to pick up and relay to other chain
func EmitSendPacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel, timeoutHeight exported.Height) {
	_ = "STUB: not implemented"
	return
}

// DEPRECATED

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)

// EmitRecvPacketEvent emits a receive packet event. It will be emitted both the first time a packet
// is received for a certain sequence and for all duplicate receives.
func EmitRecvPacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// DEPRECATED

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)

// EmitWriteAcknowledgementEvent emits an event that the relayer can query for
func EmitWriteAcknowledgementEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel, acknowledgement []byte) {
	_ = "STUB: not implemented"
	return
}

// DEPRECATED

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)

// EmitAcknowledgePacketEvent emits an acknowledge packet event. It will be emitted both the first time
// a packet is acknowledged for a certain sequence and for all duplicate acknowledgements.
func EmitAcknowledgePacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)

// EmitTimeoutPacketEvent emits a timeout packet event. It will be emitted both the first time a packet
// is timed out for a certain sequence and for all duplicate timeouts.
func EmitTimeoutPacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// EmitChannelClosedEvent emits a channel closed event.
func EmitChannelClosedEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}
