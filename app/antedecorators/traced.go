package antedecorators

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/utils/tracing"
)

type TracedAnteDecorator struct {
	wrapped sdk.AnteDecorator

	traceName   string
	tracingInfo *tracing.Info
}

func NewTracedAnteDecorator(wrapped sdk.AnteDecorator, tracingInfo *tracing.Info) TracedAnteDecorator {
	_ = "STUB: not implemented"
	return *new(TracedAnteDecorator)
}

func (d TracedAnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}
