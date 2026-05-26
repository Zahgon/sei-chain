package types

import (
	"context"
	"time"
)

func MakeCommit(ctx context.Context, blockID BlockID, height int64, round int32,
	voteSet *VoteSet, validators []PrivValidator, now time.Time) (*Commit, error) {
	_ = "STUB: not implemented"

	// all sign
	return nil, nil
}

//nolint:gosec // i is bounded by len(validators) which is bounded by MaxValidators

func signAddVote(ctx context.Context, privVal PrivValidator, vote *Vote, voteSet *VoteSet) (signed bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
