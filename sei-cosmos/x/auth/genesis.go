package auth

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

// InitGenesis - Init store state from genesis data
//
// CONTRACT: old coins from the FeeCollectionKeeper need to be transferred through
// a genesis port script to the new fee collector account
func InitGenesis(ctx sdk.Context, ak keeper.AccountKeeper, data types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns a GenesisState for a given context and keeper
func ExportGenesis(ctx sdk.Context, ak keeper.AccountKeeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
