package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("x", "epoch", "keeper")

func (k Keeper) BeginBlock(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// TODO(PLT-336): remove once epoch_begin_blocker_duration_seconds verified

//nolint:gosec
// TODO(PLT-336): remove once epoch_new verified
