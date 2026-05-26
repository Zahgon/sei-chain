package oracle

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

// NewHandler returns a handler for "oracle" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler { _ = "STUB: not implemented"; return *new(sdk.Handler) }
