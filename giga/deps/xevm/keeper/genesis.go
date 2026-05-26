package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/giga/deps/xevm/types"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// feeCollector == coinbase
