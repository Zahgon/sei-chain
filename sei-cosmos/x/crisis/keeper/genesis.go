package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/types"
)

// new crisis genesis
func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
