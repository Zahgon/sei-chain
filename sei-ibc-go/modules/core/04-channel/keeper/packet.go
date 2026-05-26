package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	sdkerrors "github.com/sei-protocol/sei-chain/sei-cosmos/types/errors"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// ErrOutboundDisabled is the error for when outbound is disabled
var ErrOutboundDisabled = sdkerrors.Register("ibc-channel", 102, "ibc outbound disabled")

// ErrInboundDisabled is the error for when inbound is disabled
var ErrInboundDisabled = sdkerrors.Register("ibc-channel", 103, "ibc inbound disabled")

// SendPacket is called by a module in order to send an IBC packet on a channel
// end owned by the calling module to the corresponding module on the counterparty
// chain.
func (k Keeper) SendPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	packet exported.PacketI,
) error {
	_ = "STUB: not implemented"
	// outbound gating: disallow sending packets when outbound disabled
	return nil
}

// prevent accidental sends with clients that cannot be updated

// check if packet is timed out on the receiving chain

// NOTE: this is a temporary fix. Solo machine does not support usage of 'GetTimestampAtHeight'
// A future change should move this function to be a ClientState callback.

func GetPacketTimeoutErrorMessage(message string, latestTimestamp uint64, timeoutTimestamp uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// RecvPacket is called by a module in order to receive & process an IBC packet
// sent on the corresponding channel end on the counterparty chain.
func (k Keeper) RecvPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	proof []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	// inbound gating: disallow processing inbound packets when inbound disabled
	return nil
}

// Authenticate capability to ensure caller has authority to receive packet on this channel

// packet must come from the channel's counterparty

// Connection must be OPEN to receive a packet. It is possible for connection to not yet be open if packet was
// sent optimistically before connection and channel handshake completed. However, to receive a packet,
// connection and channel must both be open

// check if packet timeouted by comparing it with the latest height of the chain

// check if packet timeouted by comparing it with the latest timestamp of the chain

// #nosec G115 -- block time is checked above to be non-negative

// verify that the counterparty did commit to sending this packet

// check if the packet receipt has been received already for unordered channels

// This error indicates that the packet has already been relayed. Core IBC will
// treat this error as a no-op in order to prevent an entire relay transaction
// from failing and consuming unnecessary fees.

// All verification complete, update state
// For unordered channels we must set the receipt so it can be verified on the other side.
// This receipt does not contain any data, since the packet has not yet been processed,
// it's just a single store key set to an empty string to indicate that the packet has been received

// check if the packet is being received in order

// This error indicates that the packet has already been relayed. Core IBC will
// treat this error as a no-op in order to prevent an entire relay transaction
// from failing and consuming unnecessary fees.

// All verification complete, update state
// In ordered case, we must increment nextSequenceRecv

// incrementing nextSequenceRecv and storing under this chain's channelEnd identifiers
// Since this is the receiving chain, our channelEnd is packet's destination port and channel

// log that a packet has been received & executed

// emit an event that the relayer can query for

// WriteAcknowledgement writes the packet execution acknowledgement to the state,
// which will be verified by the counterparty chain using AcknowledgePacket.
//
// CONTRACT:
//
// 1) For synchronous execution, this function is be called in the IBC handler .
// For async handling, it needs to be called directly by the module which originally
// processed the packet.
//
// 2) Assumes that packet receipt has been written (unordered), or nextSeqRecv was incremented (ordered)
// previously by RecvPacket.
func (k Keeper) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	acknowledgement exported.Acknowledgement,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Authenticate capability to ensure caller has authority to receive packet on this channel

// NOTE: IBC app modules might have written the acknowledgement synchronously on
// the OnRecvPacket callback so we need to check if the acknowledgement is already
// set on the store and return an error if so.

// set the acknowledgement so that it can be verified on the other side

// log that a packet acknowledgement has been written

// AcknowledgePacket is called by a module to process the acknowledgement of a
// packet previously sent by the calling module on a channel to a counterparty
// module on the counterparty chain. Its intended usage is within the ante
// handler. AcknowledgePacket will clean up the packet commitment,
// which is no longer necessary since the packet has been received and acted upon.
// It will also increment NextSequenceAck in case of ORDERED channels.
func (k Keeper) AcknowledgePacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	acknowledgement []byte,
	proof []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Authenticate capability to ensure caller has authority to receive packet on this channel

// packet must have been sent to the channel's counterparty

// This error indicates that the acknowledgement has already been relayed
// or there is a misconfigured relayer attempting to prove an acknowledgement
// for a packet never sent. Core IBC will treat this error as a no-op in order to
// prevent an entire relay transaction from failing and consuming unnecessary fees.

// verify we sent the packet and haven't cleared it out yet

// assert packets acknowledged in order

// All verification complete, in the case of ORDERED channels we must increment nextSequenceAck

// incrementing NextSequenceAck and storing under this chain's channelEnd identifiers
// Since this is the original sending chain, our channelEnd is packet's source port and channel

// Delete packet commitment, since the packet has been acknowledged, the commitement is no longer necessary

// log that a packet has been acknowledged

// emit an event marking that we have processed the acknowledgement
