package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
)

// InitGenesis sets distribution information for genesis
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// check if the module account exists

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
