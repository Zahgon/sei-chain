package migrations

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("x", "evm", "migrations")

// This migration is to fix total supply mismatch caused by mishandled
// ante surplus
func FixTotalSupply(ctx sdk.Context, k *keeper.Keeper) error { _ = "STUB: not implemented"; return nil }
