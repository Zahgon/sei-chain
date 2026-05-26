package evm

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genState types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// TODO: move to better location
var GENESIS_EXPORT_STREAM_SERIALIZED_LEN_MAX = 1000

func ExportGenesisStream(ctx sdk.Context, k *keeper.Keeper) <-chan *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// GetGenesisStateFromAppState returns x/evm GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
