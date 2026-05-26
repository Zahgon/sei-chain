package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"

	icatypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// SendTx takes pre-built packet data containing messages to be executed on the host chain from an authentication module and attempts to send the packet.
// The packet sequence for the outgoing packet is returned as a result.
// If the base application has the capability to send on the provided portID. An appropriate
// absolute timeoutTimestamp must be provided. If the packet is timed out, the channel will be closed.
// In the case of channel closure, a new channel may be reopened to reconnect to the host chain.
func (k Keeper) SendTx(ctx sdk.Context, chanCap *capabilitytypes.Capability, connectionID, portID string, icaPacketData icatypes.InterchainAccountPacketData, timeoutTimestamp uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// #nosec G115 -- block time is checked above to be non-negative

func (k Keeper) createOutgoingPacket(
	ctx sdk.Context,
	sourcePort,
	sourceChannel,
	destinationPort,
	destinationChannel string,
	chanCap *capabilitytypes.Capability,
	icaPacketData icatypes.InterchainAccountPacketData,
	timeoutTimestamp uint64,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// get the next sequence

// OnTimeoutPacket removes the active channel associated with the provided packet, the underlying channel end is closed
// due to the semantics of ORDERED channels
func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	return nil
}
