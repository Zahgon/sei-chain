package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
)

const (
	MinGasEVMTx = 21000
)

type GasDecorator struct {
	evmKeeper *evmkeeper.Keeper
}

func NewGasDecorator(evmKeeper *evmkeeper.Keeper) *GasDecorator {
	_ = "STUB: not implemented"
	return nil
}

// Called at the end of the ante chain to set gas limit and gas used estimate properly
func (gl GasDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

//nolint:gosec
