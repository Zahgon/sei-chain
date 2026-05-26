package crisis

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/keeper"
)

// check all registered invariants
func EndBlocker(ctx sdk.Context, k keeper.Keeper) { _ = "STUB: not implemented"; return }

//nolint:gosec // InvCheckPeriod is a small config value, won't overflow int64
// skip running the invariant check
