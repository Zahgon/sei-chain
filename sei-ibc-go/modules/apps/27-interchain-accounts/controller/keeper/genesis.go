package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	icatypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/27-interchain-accounts/types"
)

// InitGenesis initializes the interchain accounts controller application state from a provided genesis state
func InitGenesis(ctx sdk.Context, keeper Keeper, state icatypes.ControllerGenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the interchain accounts controller exported genesis
func ExportGenesis(ctx sdk.Context, keeper Keeper) icatypes.ControllerGenesisState {
	_ = "STUB: not implemented"
	return *new(icatypes.ControllerGenesisState)
}
