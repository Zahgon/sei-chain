package evidence

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "evidence")

// BeginBlocker iterates through and handles any newly discovered evidence of
// misbehavior submitted by Tendermint. Currently, only equivocation is handled.
func BeginBlocker(ctx sdk.Context, byzantineValidators []abci.Misbehavior, k keeper.Keeper) {
	_ = "STUB: not implemented"
	return
}

// It's still ongoing discussion how should we treat and slash attacks with
// premeditation. So for now we agree to treat them in the same way.
