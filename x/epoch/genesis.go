package epoch

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/epoch/keeper"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	_ = "STUB: not implemented"
	// this line is used by starport scaffolding # genesis/module/init
	return
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
