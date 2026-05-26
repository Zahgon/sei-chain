package avail

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
)

type blockVotes struct {
	byKey    map[types.PublicKey]*types.Signed[*types.LaneVote]
	byHeader map[types.BlockHeaderHash][]*types.Signed[*types.LaneVote]
}

func newBlockVotes() blockVotes { _ = "STUB: not implemented"; return *new(blockVotes) }

// Returns true iff a new QC has been constructed.
func (bv blockVotes) pushVote(c *types.Committee, vote *types.Signed[*types.LaneVote]) (*types.LaneQC, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
