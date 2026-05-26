package factory

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

func MakeCommit(ctx context.Context, blockID types.BlockID, height int64, round int32, voteSet *types.VoteSet, validators []types.PrivValidator, now time.Time) (*types.Commit, error) {
	_ = "STUB: not implemented"
	// all sign
	return nil, nil
}

//nolint:gosec // i is bounded by len(validators) which is bounded by MaxValidators
