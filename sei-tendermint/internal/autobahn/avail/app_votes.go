package avail

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "internal", "autobahn", "avail")

type appVotes struct {
	byKey  map[types.PublicKey]*types.Signed[*types.AppVote]
	byHash map[types.Hash[*types.AppVote]][]*types.Signed[*types.AppVote]
}

func newAppVotes() appVotes { _ = "STUB: not implemented"; return *new(appVotes) }

// Returns qc if a new qc has been reached.
func (av appVotes) pushVote(c *types.Committee, vote *types.Signed[*types.AppVote]) (*types.AppQC, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
