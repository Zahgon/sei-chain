package keeper

import (
	baseapp "github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	icatypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/types"
)

// Keeper defines the IBC interchain accounts controller keeper
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramSpace paramtypes.Subspace

	ics4Wrapper   icatypes.ICS4Wrapper
	channelKeeper icatypes.ChannelKeeper
	portKeeper    icatypes.PortKeeper

	scopedKeeper capabilitykeeper.ScopedKeeper

	msgRouter *baseapp.MsgServiceRouter
}

// NewKeeper creates a new interchain accounts controller Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	ics4Wrapper icatypes.ICS4Wrapper, channelKeeper icatypes.ChannelKeeper, portKeeper icatypes.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper, msgRouter *baseapp.MsgServiceRouter,
) Keeper {
	_ = "STUB: not implemented"
	// set KeyTable if it has not already been set
	return *new(Keeper)
}

// GetAllPorts returns all ports to which the interchain accounts controller module is bound. Used in ExportGenesis
func (k Keeper) GetAllPorts(ctx sdk.Context) []string { _ = "STUB: not implemented"; return nil }

// BindPort stores the provided portID and binds to it, returning the associated capability
func (k Keeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	_ = "STUB: not implemented"
	return nil
}

// IsBound checks if the interchain account controller module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_ = "STUB: not implemented"
	return false
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// ClaimCapability wraps the scopedKeeper's ClaimCapability function
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetActiveChannelID retrieves the active channelID from the store, keyed by the provided connectionID and portID
func (k Keeper) GetActiveChannelID(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetOpenActiveChannel retrieves the active channelID from the store, keyed by the provided connectionID and portID & checks if the channel in question is in state OPEN
func (k Keeper) GetOpenActiveChannel(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetAllActiveChannels returns a list of all active interchain accounts controller channels and their associated connection and port identifiers
func (k Keeper) GetAllActiveChannels(ctx sdk.Context) []icatypes.ActiveChannel {
	_ = "STUB: not implemented"
	return nil
}

// SetActiveChannelID stores the active channelID, keyed by the provided connectionID and portID
func (k Keeper) SetActiveChannelID(ctx sdk.Context, connectionID, portID, channelID string) {
	_ = "STUB: not implemented"
	return
}

// IsActiveChannel returns true if there exists an active channel for the provided connectionID and portID, otherwise false
func (k Keeper) IsActiveChannel(ctx sdk.Context, connectionID, portID string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetInterchainAccountAddress retrieves the InterchainAccount address from the store associated with the provided connectionID and portID
func (k Keeper) GetInterchainAccountAddress(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetAllInterchainAccounts returns a list of all registered interchain account addresses and their associated connection and controller port identifiers
func (k Keeper) GetAllInterchainAccounts(ctx sdk.Context) []icatypes.RegisteredInterchainAccount {
	_ = "STUB: not implemented"
	return nil
}

// SetInterchainAccountAddress stores the InterchainAccount address, keyed by the associated connectionID and portID
func (k Keeper) SetInterchainAccountAddress(ctx sdk.Context, connectionID, portID, address string) {
	_ = "STUB: not implemented"
	return
}
