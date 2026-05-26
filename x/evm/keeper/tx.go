package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const DefaultTxHashesToRemove = 100

func (k *Keeper) RemoveFirstNTxHashes(ctx sdk.Context, n int) { _ = "STUB: not implemented"; return }
