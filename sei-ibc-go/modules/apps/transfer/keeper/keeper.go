package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
)

// Keeper defines the IBC fungible transfer keeper
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramSpace paramtypes.Subspace

	ics4Wrapper   types.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	authKeeper    types.AccountKeeper
	bankKeeper    types.BankKeeper
	scopedKeeper  capabilitykeeper.ScopedKeeper

	addressHandler types.AddressHandler
}

// NewKeeper creates a new IBC transfer Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	ics4Wrapper types.ICS4Wrapper, channelKeeper types.ChannelKeeper, portKeeper types.PortKeeper,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper, scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {
	_ = "STUB: not implemented"
	// ensure ibc transfer module account is set
	return *new(Keeper)
}

// set KeyTable if it has not already been set

// NewKeeperWithAddressHandler creates a new IBC transfer Keeper instance with an address handler
func NewKeeperWithAddressHandler(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	ics4Wrapper types.ICS4Wrapper, channelKeeper types.ChannelKeeper, portKeeper types.PortKeeper,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper, scopedKeeper capabilitykeeper.ScopedKeeper,
	addressHandler types.AddressHandler,
) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// GetTransferAccount returns the ICS20 - transfers ModuleAccount
func (k Keeper) GetTransferAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	_ = "STUB: not implemented"
	return *new(authtypes.ModuleAccountI)
}

// IsBound checks if the transfer module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_ = "STUB: not implemented"
	return false
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPort returns the portID for the transfer module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string { _ = "STUB: not implemented"; return "" }

// SetPort sets the portID for the transfer module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) { _ = "STUB: not implemented"; return }

// GetDenomTrace retreives the full identifiers trace and base denomination from the store.
func (k Keeper) GetDenomTrace(ctx sdk.Context, denomTraceHash tmbytes.HexBytes) (types.DenomTrace, bool) {
	_ = "STUB: not implemented"
	return *new(types.DenomTrace), false
}

// HasDenomTrace checks if a the key with the given denomination trace hash exists on the store.
func (k Keeper) HasDenomTrace(ctx sdk.Context, denomTraceHash tmbytes.HexBytes) bool {
	_ = "STUB: not implemented"
	return false
}

// SetDenomTrace sets a new {trace hash -> denom trace} pair to the store.
func (k Keeper) SetDenomTrace(ctx sdk.Context, denomTrace types.DenomTrace) {
	_ = "STUB: not implemented"
	return
}

// GetAllDenomTraces returns the trace information for all the denominations.
func (k Keeper) GetAllDenomTraces(ctx sdk.Context) types.Traces {
	_ = "STUB: not implemented"
	return *new(types.Traces)
}

// IterateDenomTraces iterates over the denomination traces in the store
// and performs a callback function.
func (k Keeper) IterateDenomTraces(ctx sdk.Context, cb func(denomTrace types.DenomTrace) bool) {
	_ = "STUB: not implemented"
	return
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// ClaimCapability allows the transfer module that can claim a capability that IBC module
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	_ = "STUB: not implemented"
	return nil
}
