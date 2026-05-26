package consensus

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type scv = *types.Signed[*types.CommitVote]
type hcv = types.Hash[*types.CommitVote]

type commitVotes struct {
	byKey  map[types.PublicKey]scv
	byHash map[hcv]map[types.PublicKey]scv
	qc     utils.AtomicSend[utils.Option[*types.CommitQC]]
}

func newCommitVotes() *commitVotes { _ = "STUB: not implemented"; return nil }

func (cv *commitVotes) pushVote(c *types.Committee, vote *types.Signed[*types.CommitVote]) {
	_ = "STUB: not implemented"
	return
}

// Check if the key has already voted.

// Ignore older or equal votes.

// Remove the old vote from the view map.

// Insert the new vote.

// Check if we have enough votes for a CommitQC.

// Construct a CommitQC from the votes.
