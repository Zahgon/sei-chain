package client

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/keeper"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

// InitGenesis initializes the ibc client submodule's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// Set all client metadata first. This will allow client keeper to overwrite client and consensus state keys
// if clients accidentally write to ClientKeeper reserved keys.

// NOTE: localhost creation is specifically disallowed for the time being.
// Issue: https://github.com/cosmos/cosmos-sdk/issues/7871

// ExportGenesis returns the ibc client submodule's exported genesis.
// NOTE: CreateLocalhost should always be false on export since a
// created localhost will be included in the exported clients.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	_ = "STUB: not implemented"
	return *new(types.GenesisState)
}
