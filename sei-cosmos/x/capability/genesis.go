package capability

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// set owners for each index

// initialize in-memory capabilities

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
