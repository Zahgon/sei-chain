package simapp

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/ante"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper *keeper.Keeper
}

// NewAnteHandler creates a new ante handler
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AnteHandler), nil
}

// SetPubKeyDecorator must be called before all signature verification decorators
