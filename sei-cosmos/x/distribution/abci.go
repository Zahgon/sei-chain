package distribution

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// BeginBlocker sets the proposer for determining distribution during endblock
// and distribute rewards for the previous block
func BeginBlocker(ctx sdk.Context, votes []abci.VoteInfo, k keeper.Keeper) {
	_ = "STUB: not implemented"
	return
}

// determine the total power signing the block

// TODO this is Tendermint-dependent
// ref https://github.com/cosmos/cosmos-sdk/issues/3095

// record the proposer for when we payout on the next block
