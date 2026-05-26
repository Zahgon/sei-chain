package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type EVMRouterDecorator struct {
	defaultAnteHandler sdk.AnteHandler
	evmAnteHandler     sdk.AnteHandler
}

func NewEVMRouterDecorator(
	defaultAnteHandler sdk.AnteHandler,
	evmAnteHandler sdk.AnteHandler,
) *EVMRouterDecorator {
	_ = "STUB: not implemented"
	return nil
}

func (r EVMRouterDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

func IsEVMMessage(tx sdk.Tx) (bool, error) { _ = "STUB: not implemented"; return false, nil }
