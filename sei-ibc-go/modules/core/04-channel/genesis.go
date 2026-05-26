package channel

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/keeper"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// InitGenesis initializes the ibc channel submodule's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the ibc channel submodule's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	_ = "STUB: not implemented"
	return *new(types.GenesisState)
}
