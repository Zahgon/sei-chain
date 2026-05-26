package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k Keeper) IsVoteTarget(ctx sdk.Context, denom string) bool {
	_ = "STUB: not implemented"
	return false
}

func (k Keeper) GetVoteTargets(ctx sdk.Context) (voteTargets []string) {
	_ = "STUB: not implemented"
	return nil
}
