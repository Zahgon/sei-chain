package ibc

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/types"
)

// InitGenesis initializes the ibc state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, createLocalhost bool, gs *types.GenesisState) {
	_ = "STUB: not implemented"
	// Initialize core params with defaults if not set
	return
}

// ExportGenesis returns the ibc exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
