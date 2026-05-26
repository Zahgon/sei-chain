package consensus

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type timeoutVotes struct {
	byKey  map[types.PublicKey]*types.FullTimeoutVote
	byView map[types.View]map[types.PublicKey]*types.FullTimeoutVote
	qc     utils.AtomicSend[utils.Option[*types.TimeoutQC]]
}

func newTimeoutVotes() *timeoutVotes { _ = "STUB: not implemented"; return nil }

func (tv *timeoutVotes) pushVote(c *types.Committee, vote *types.FullTimeoutVote) {
	_ = "STUB: not implemented"
	// TODO: verify the vote.
	return
}

// Check if the old vote is newer than the new one.

// Prune the old vote.

// Insert the new vote.

// Check if we have enough votes for a TimeoutQC.

// Construct a TimeoutQC from the votes.
