package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Keeper of the global paramstore
type Keeper struct {
	cdc         codec.BinaryCodec
	legacyAmino *codec.LegacyAmino
	key         sdk.StoreKey
	tkey        sdk.StoreKey
	spaces      map[string]*types.Subspace
}

// NewKeeper constructs a params keeper
func NewKeeper(cdc codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

func (k Keeper) SetFeesParams(ctx sdk.Context, feesParams types.FeesParams) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) GetFeesParams(ctx sdk.Context) types.FeesParams {
	_ = "STUB: not implemented"
	return *new(types.FeesParams)
}

func (k Keeper) SetCosmosGasParams(ctx sdk.Context, cosmosGasParams types.CosmosGasParams) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) GetCosmosGasParams(ctx sdk.Context) types.CosmosGasParams {
	_ = "STUB: not implemented"
	return *new(types.CosmosGasParams)
}

// Allocate subspace used for keepers
func (k Keeper) Subspace(s string) types.Subspace {
	_ = "STUB: not implemented"
	return *new(types.Subspace)
}

// Get existing substore from keeper
func (k Keeper) GetSubspace(s string) (types.Subspace, bool) {
	_ = "STUB: not implemented"
	return *new(types.Subspace), false
}
