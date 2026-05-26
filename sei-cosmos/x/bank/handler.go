package bank

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
)

// NewHandler returns a handler for "bank" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler { _ = "STUB: not implemented"; return *new(sdk.Handler) }
