package consensus

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type spv = *types.Signed[*types.PrepareVote]
type hpv = types.Hash[*types.PrepareVote]

// prepareVotes holds the votes for the prepare phase of consensus.
type prepareVotes struct {
	byKey  map[types.PublicKey]spv
	byHash map[hpv]map[types.PublicKey]spv
	qc     utils.AtomicSend[utils.Option[*types.PrepareQC]]
}

// newPrepareVotes initializes a new prepareVotes instance.
func newPrepareVotes() *prepareVotes { _ = "STUB: not implemented"; return nil }

// pushVote processes a new prepare vote and updates the prepare votes state.
func (pv *prepareVotes) pushVote(c *types.Committee, vote *types.Signed[*types.PrepareVote]) {
	_ = "STUB: not implemented"
	return
}

// Check if the key has already voted.

// Ignore older or equal votes.

// Remove the old vote from the view map.

// Insert the new vote.

// Check if we have enough votes for a PrepareQC.

// Construct a PrepareQC from the votes.
