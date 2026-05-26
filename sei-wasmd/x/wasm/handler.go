package wasm

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewHandler returns a handler for "wasm" type messages.
func NewHandler(k types.ContractOpsKeeper) sdk.Handler {
	_ = "STUB: not implemented"
	return *new(sdk.Handler)
}

//nolint:typecheck

// filterMessageEvents returns the same events with all of type == EventTypeMessage removed except
// for wasm message types.
// this is so only our top-level message event comes through
func filterMessageEvents(ctx sdk.Context) *sdk.EventManager { _ = "STUB: not implemented"; return nil }

func hasWasmModuleAttribute(attrs []abci.EventAttribute) bool {
	_ = "STUB: not implemented"
	return false
}
