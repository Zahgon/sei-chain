package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	sdkerrors "github.com/sei-protocol/sei-chain/sei-cosmos/types/errors"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var logger = seilog.NewLogger("ibc-go", "modules", "core", "02-client", "keeper")

// ErrInboundDisabled is the error for when inbound is disabled
var ErrInboundDisabled = sdkerrors.Register("ibc-client", 101, "ibc inbound disabled")

// CreateClient creates a new client state and populates it with a given consensus
// state as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-002-client-semantics#create
func (k Keeper) CreateClient(
	ctx sdk.Context, clientState exported.ClientState, consensusState exported.ConsensusState,
) (string, error) {
	_ = "STUB: not implemented"
	// inbound gating: disallow client creation as part of inbound handshakes when inbound disabled
	return "", nil
}

// verifies initial consensus state against client state and initializes client store with any client-specific metadata
// e.g. set ProcessedTime in Tendermint clients

// check if consensus state is nil in case the created client is Localhost

// UpdateClient updates the consensus state and the state root from a provided header.
func (k Keeper) UpdateClient(ctx sdk.Context, clientID string, header exported.Header) error {
	_ = "STUB: not implemented"
	return nil
}

// Any writes made in CheckHeaderAndUpdateState are persisted on both valid updates and misbehaviour updates.
// Light client implementations are responsible for writing the correct metadata (if any) in either case.

// emit the full header in events

// Marshal the Header as an Any and encode the resulting bytes to hex.
// This prevents the event value from containing invalid UTF-8 characters
// which may cause data to be lost when JSON encoding/decoding.

// set default consensus height with header height

// set new client state regardless of if update is valid update or misbehaviour

// If client state is not frozen after clientState CheckHeaderAndUpdateState,
// then update was valid. Write the update state changes, and set new consensus state.
// Else the update was proof of misbehaviour and we must emit appropriate misbehaviour events.

// if update is not misbehaviour then update the consensus state
// we don't set consensus state for localhost client

// emitting events in the keeper emits for both begin block and handler client updates

// UpgradeClient upgrades the client to a new client state if this new client was committed to
// by the old client at the specified upgrade height
func (k Keeper) UpgradeClient(ctx sdk.Context, clientID string, upgradedClient exported.ClientState, upgradedConsState exported.ConsensusState,
	proofUpgradeClient, proofUpgradeConsState []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// emitting events in the keeper emits for client upgrades

// CheckMisbehaviourAndUpdateState checks for client misbehaviour and freezes the
// client if so.
func (k Keeper) CheckMisbehaviourAndUpdateState(ctx sdk.Context, misbehaviour exported.Misbehaviour) error {
	_ = "STUB: not implemented"
	return nil
}
