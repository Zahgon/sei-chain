package evidence

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/types"
)

// InitGenesis initializes the evidence module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the evidence module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
