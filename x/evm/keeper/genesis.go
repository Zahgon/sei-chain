package keeper

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/evm/types"
)

var ethReplayInitialied = false

func (k *Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// feeCollector == coinbase

func (k *Keeper) OpenEthDatabase() *ethtypes.Header { _ = "STUB: not implemented"; return nil }
