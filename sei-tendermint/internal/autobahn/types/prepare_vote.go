package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// PrepareVote .
type PrepareVote struct {
	utils.ReadOnly
	proposal *Proposal
}

// NewPrepareVote creates a new PrepareVote.
func NewPrepareVote(proposal *Proposal) *PrepareVote { _ = "STUB: not implemented"; return nil }

// Proposal .
func (m *PrepareVote) Proposal() *Proposal {
	_ = "STUB: not implemented"

	// PrepareVoteConv is the protobuf converter for PrepareVote.
	return nil
}

var PrepareVoteConv = protoutils.Conv[*PrepareVote, *pb.Proposal]{
	Encode: func(m *PrepareVote) *pb.Proposal {
		return ProposalConv.Encode(m.proposal)
	},
	Decode: func(m *pb.Proposal) (*PrepareVote, error) {
		proposal, err := ProposalConv.DecodeReq(m)
		if err != nil {
			return nil, fmt.Errorf("proposal: %w", err)
		}
		return &PrepareVote{proposal: proposal}, nil
	},
}
