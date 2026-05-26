package gov

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, data *types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// check if the deposits pool account exists

// if account has zero balance it probably means it's not set, so we set it

// check if total deposits equals balance, if it doesn't panic because there were export/import errors

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
