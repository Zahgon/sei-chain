package keeper

import (
	"github.com/sei-protocol/sei-chain/giga/deps/xbank/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// InitGenesis initializes the bank module's state from a given genesis state.
func (k BaseKeeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the bank module's genesis state.
func (k BaseKeeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
