package slashing

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
)

// NewHandler creates an sdk.Handler for all the slashing type messages
func NewHandler(k keeper.Keeper) sdk.Handler { _ = "STUB: not implemented"; return *new(sdk.Handler) }
