package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// LaneVote .
type LaneVote struct {
	utils.ReadOnly
	header *BlockHeader
}

// NewLaneVote creates a new LaneVote.
func NewLaneVote(header *BlockHeader) *LaneVote { _ = "STUB: not implemented"; return nil }

// Header .
func (m *LaneVote) Header() *BlockHeader {
	_ = "STUB: not implemented"

	// Verify verifies that the LaneVote is consistent with the Committee.
	return nil
}

func (m *LaneVote) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// LaneVoteConv is the protobuf converter for LaneVote.
var LaneVoteConv = protoutils.Conv[*LaneVote, *pb.BlockHeader]{
	Encode: func(m *LaneVote) *pb.BlockHeader {
		return BlockHeaderConv.Encode(m.header)
	},
	Decode: func(m *pb.BlockHeader) (*LaneVote, error) {
		header, err := BlockHeaderConv.DecodeReq(m)
		if err != nil {
			return nil, fmt.Errorf("header: %w", err)
		}
		return &LaneVote{header: header}, nil
	},
}
