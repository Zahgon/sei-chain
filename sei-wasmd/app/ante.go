package app

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/ante"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"

	wasmTypes "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper         *keeper.Keeper
	WasmConfig        *wasmTypes.WasmConfig
	TXCounterStoreKey sdk.StoreKey
}

func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AnteHandler), nil
}

// outermost AnteDecorator. SetUpContext must be called first
// after setup context to enforce limits early

// SetPubKeyDecorator must be called before all signature verification decorators
