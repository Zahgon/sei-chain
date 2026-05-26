package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"

	clientkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/keeper"
	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	connectionkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/keeper"
	channelkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/keeper"
	portkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/keeper"
	porttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// Keeper defines each ICS keeper for IBC
type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	cdc codec.BinaryCodec

	ClientKeeper     clientkeeper.Keeper
	ConnectionKeeper connectionkeeper.Keeper
	ChannelKeeper    channelkeeper.Keeper
	PortKeeper       portkeeper.Keeper
	Router           *porttypes.Router

	paramSpace paramtypes.Subspace
}

// NewKeeper creates a new ibc Keeper
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	stakingKeeper clienttypes.StakingKeeper, upgradeKeeper clienttypes.UpgradeKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
) *Keeper {
	_ = "STUB: not implemented"
	// register paramSpace at top level keeper
	// set KeyTable if it has not already been set
	return nil
}

// register core params

// panic if any of the keepers passed in is empty

// Codec returns the IBC module codec.
func (k Keeper) Codec() codec.BinaryCodec {
	_ = "STUB: not implemented"

	// SetRouter sets the Router in IBC Keeper and seals it. The method panics if
	// there is an existing router that's already sealed.
	return *new(codec.BinaryCodec)
}

func (k *Keeper) SetRouter(rtr *porttypes.Router) { _ = "STUB: not implemented"; return }
