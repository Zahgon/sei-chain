package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// RegisterInterchainAccount is the entry point to registering an interchain account.
// It generates a new port identifier using the owner address. It will bind to the
// port identifier and call 04-channel 'ChanOpenInit'. An error is returned if the port
// identifier is already in use. Gaining access to interchain accounts whose channels
// have closed cannot be done with this function. A regular MsgChanOpenInit must be used.
func (k Keeper) RegisterInterchainAccount(ctx sdk.Context, connectionID, owner string) error {
	_ = "STUB: not implemented"
	return nil
}

// if there is an active channel for this portID / connectionID return an error

// NOTE: An empty string is provided for accAddress, to be fulfilled upon OnChanOpenTry handshake step

// NOTE: The sdk msg handler creates a new EventManager, so events must be correctly propagated back to the current context
