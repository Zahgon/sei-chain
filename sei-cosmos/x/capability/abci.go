package capability

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
)

// BeginBlocker will call InitMemStore to initialize the memory stores in the case
// that this is the first time the node is executing a block since restarting (wiping memory).
// In this case, the BeginBlocker method will reinitialize the memory stores locally, so that subsequent
// capability transactions will pass.
// Otherwise BeginBlocker performs a no-op.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) { _ = "STUB: not implemented"; return }
