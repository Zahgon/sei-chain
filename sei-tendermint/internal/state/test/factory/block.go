package factory

import (
	"context"
	"testing"

	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

func MakeBlocks(ctx context.Context, t *testing.T, n int, state *sm.State, privVal types.PrivValidator) []*types.Block {
	_ = "STUB: not implemented"
	return nil
}

// update state

func MakeBlock(state sm.State, height int64, c *types.Commit) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

func makeBlockAndPartSet(
	ctx context.Context,
	t *testing.T,
	state sm.State,
	lastBlock *types.Block,
	lastBlockMeta *types.BlockMeta,
	privVal types.PrivValidator,
	height int64,
) (*types.Block, *types.PartSet) {
	_ = "STUB: not implemented"
	return nil, nil
}
