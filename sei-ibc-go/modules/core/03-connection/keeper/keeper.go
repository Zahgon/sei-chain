package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	sdkerrors "github.com/sei-protocol/sei-chain/sei-cosmos/types/errors"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// KeyInboundEnabled is the param key for inbound enabled
var KeyInboundEnabled = []byte("InboundEnabled")

// KeyOutboundEnabled is the param key for outbound enabled
var KeyOutboundEnabled = []byte("OutboundEnabled")

// ErrInboundDisabled is the error for when inbound is disabled
var ErrInboundDisabled = sdkerrors.Register("ibc-connection", 101, "ibc inbound disabled")

// ErrOutboundDisabled is the error for when outbound is disabled
var ErrOutboundDisabled = sdkerrors.Register("ibc-connection", 102, "ibc outbound disabled")

// Keeper defines the IBC connection keeper
type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeKey     sdk.StoreKey
	paramSpace   paramtypes.Subspace
	cdc          codec.BinaryCodec
	clientKeeper types.ClientKeeper
}

// NewKeeper creates a new IBC connection Keeper instance
func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace, ck types.ClientKeeper) Keeper {
	_ = "STUB: not implemented"
	// set KeyTable if it has not already been set
	return *new(Keeper)
}

// IsInboundEnabled returns true if inbound IBC is enabled.
func (k Keeper) IsInboundEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// IsOutboundEnabled returns true if outbound IBC is enabled.
func (k Keeper) IsOutboundEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// GetCommitmentPrefix returns the IBC connection store prefix as a commitment
// Prefix
func (k Keeper) GetCommitmentPrefix() exported.Prefix {
	_ = "STUB: not implemented"
	return *new(exported.Prefix)
}

// GenerateConnectionIdentifier returns the next connection identifier.
func (k Keeper) GenerateConnectionIdentifier(ctx sdk.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// GetConnection returns a connection with a particular identifier
func (k Keeper) GetConnection(ctx sdk.Context, connectionID string) (types.ConnectionEnd, bool) {
	_ = "STUB: not implemented"
	return *new(types.ConnectionEnd), false
}

// SetConnection sets a connection to the store
func (k Keeper) SetConnection(ctx sdk.Context, connectionID string, connection types.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}

// GetTimestampAtHeight returns the timestamp in nanoseconds of the consensus state at the
// given height.
func (k Keeper) GetTimestampAtHeight(ctx sdk.Context, connection types.ConnectionEnd, height exported.Height) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetClientConnectionPaths returns all the connection paths stored under a
// particular client
func (k Keeper) GetClientConnectionPaths(ctx sdk.Context, clientID string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetClientConnectionPaths sets the connections paths for client
func (k Keeper) SetClientConnectionPaths(ctx sdk.Context, clientID string, paths []string) {
	_ = "STUB: not implemented"
	return
}

// GetNextConnectionSequence gets the next connection sequence from the store.
func (k Keeper) GetNextConnectionSequence(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// SetNextConnectionSequence sets the next connection sequence to the store.
func (k Keeper) SetNextConnectionSequence(ctx sdk.Context, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetAllClientConnectionPaths returns all stored clients connection id paths. It
// will ignore the clients that haven't initialized a connection handshake since
// no paths are stored.
func (k Keeper) GetAllClientConnectionPaths(ctx sdk.Context) []types.ConnectionPaths {
	_ = "STUB: not implemented"
	return nil
}

// continue when connection handshake is not initialized

// IterateConnections provides an iterator over all ConnectionEnd objects.
// For each ConnectionEnd, cb will be called. If the cb returns true, the
// iterator will close and stop.
func (k Keeper) IterateConnections(ctx sdk.Context, cb func(types.IdentifiedConnection) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllConnections returns all stored ConnectionEnd objects.
func (k Keeper) GetAllConnections(ctx sdk.Context) (connections []types.IdentifiedConnection) {
	_ = "STUB: not implemented"
	return nil
}

// addConnectionToClient is used to add a connection identifier to the set of
// connections associated with a client.
func (k Keeper) addConnectionToClient(ctx sdk.Context, clientID, connectionID string) error {
	_ = "STUB: not implemented"
	return nil
}
