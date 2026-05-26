package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// CommitVote .
type CommitVote struct {
	utils.ReadOnly
	proposal *Proposal
}

// NewCommitVote creates a new CommitVote.
func NewCommitVote(proposal *Proposal) *CommitVote { _ = "STUB: not implemented"; return nil }

// Proposal .
func (m *CommitVote) Proposal() *Proposal {
	_ = "STUB: not implemented"

	// CommitVoteConv is the protobuf converter for CommitVote.
	return nil
}

var CommitVoteConv = protoutils.Conv[*CommitVote, *pb.Proposal]{
	Encode: func(m *CommitVote) *pb.Proposal {
		return ProposalConv.Encode(m.proposal)
	},
	Decode: func(m *pb.Proposal) (*CommitVote, error) {
		proposal, err := ProposalConv.DecodeReq(m)
		if err != nil {
			return nil, fmt.Errorf("proposal: %w", err)
		}
		return &CommitVote{proposal: proposal}, nil
	},
}
