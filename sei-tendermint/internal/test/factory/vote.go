package factory

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

func MakeVote(
	ctx context.Context,
	val types.PrivValidator,
	chainID string,
	valIndex int32,
	height int64,
	round int32,
	step int,
	blockID types.BlockID,
	time time.Time,
) (*types.Vote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // step is a small enum value (prevote/precommit/commit); no overflow risk
