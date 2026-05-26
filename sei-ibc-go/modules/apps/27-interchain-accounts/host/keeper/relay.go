package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// OnRecvPacket handles a given interchain accounts packet on a destination host chain.
// If the transaction is successfully executed, the transaction response bytes will be returned.
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON errors are indeterminate and therefore are not wrapped and included in failed acks

// executeTx attempts to execute the provided transaction. It begins by authenticating the transaction signer.
// If authentication succeeds, it does basic validation of the messages before attempting to deliver each message
// into state. The state changes will only be committed if all messages in the transaction succeed. Thus the
// execution of the transaction is atomic, all state changes are reverted if a single message fails.
func (k Keeper) executeTx(ctx sdk.Context, sourcePort, destPort, destChannel string, msgs []sdk.Msg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CacheContext returns a new context with the multi-store branched into a cached storage object
// writeCache is called only if all msgs succeed, performing state transitions atomically

// NOTE: The context returned by CacheContext() creates a new EventManager, so events must be correctly propagated back to the current context

// authenticateTx ensures the provided msgs contain the correct interchain account signer address retrieved
// from state using the provided controller port identifier
func (k Keeper) authenticateTx(ctx sdk.Context, msgs []sdk.Msg, connectionID, portID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Attempts to get the message handler from the router and if found will then execute the message.
// If the message execution is successful, the proto marshaled message response will be returned.
func (k Keeper) executeMsg(ctx sdk.Context, msg sdk.Msg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: The sdk msg handler creates a new EventManager, so events must be correctly propagated back to the current context
