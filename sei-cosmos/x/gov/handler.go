package gov

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
)

// NewHandler creates an sdk.Handler for all the gov type messages
func NewHandler(k keeper.Keeper) sdk.Handler { _ = "STUB: not implemented"; return *new(sdk.Handler) }
