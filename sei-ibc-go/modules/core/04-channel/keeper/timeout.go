package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// TimeoutPacket is called by a module which originally attempted to send a
// packet to a counterparty module, where the timeout height has passed on the
// counterparty chain without the packet being committed, to prove that the
// packet can no longer be executed and to allow the calling module to safely
// perform appropriate state transitions. Its intended usage is within the
// ante handler.
func (k Keeper) TimeoutPacket(
	ctx sdk.Context,
	packet exported.PacketI,
	proof []byte,
	proofHeight exported.Height,
	nextSequenceRecv uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: TimeoutPacket is called by the AnteHandler which acts upon the packet.Route(),
// so the capability authentication can be omitted here

// check that timeout height or timeout timestamp has passed on the other end

// This error indicates that the timeout has already been relayed
// or there is a misconfigured relayer attempting to prove a timeout
// for a packet never sent. Core IBC will treat this error as a no-op in order to
// prevent an entire relay transaction from failing and consuming unnecessary fees.

// verify we sent the packet and haven't cleared it out yet

// check that packet has not been received

// check that the recv sequence is as claimed

// NOTE: the remaining code is located in the TimeoutExecuted function

// TimeoutExecuted deletes the commitment send from this chain after it verifies timeout.
// If the timed-out packet came from an ORDERED channel then this channel will be closed.
//
// CONTRACT: this function must be called in the IBC handler
func (k Keeper) TimeoutExecuted(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// emit an event marking that we have processed the timeout

// TimeoutOnClose is called by a module in order to prove that the channel to
// which an unreceived packet was addressed has been closed, so the packet will
// never be received (even if the timeoutHeight has not yet been reached).
func (k Keeper) TimeoutOnClose(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	proof,
	proofClosed []byte,
	proofHeight exported.Height,
	nextSequenceRecv uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// This error indicates that the timeout has already been relayed
// or there is a misconfigured relayer attempting to prove a timeout
// for a packet never sent. Core IBC will treat this error as a no-op in order to
// prevent an entire relay transaction from failing and consuming unnecessary fees.

// verify we sent the packet and haven't cleared it out yet

// check that the opposing channel end has closed

// check that packet has not been received

// check that the recv sequence is as claimed

// NOTE: the remaining code is located in the TimeoutExecuted function
