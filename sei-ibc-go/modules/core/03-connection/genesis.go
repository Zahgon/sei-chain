package connection

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/keeper"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
)

// InitGenesis initializes the ibc connection submodule's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the ibc connection submodule's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	_ = "STUB: not implemented"
	return *new(types.GenesisState)
}
