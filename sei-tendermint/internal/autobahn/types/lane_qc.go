package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// LaneQC .
type LaneQC struct {
	utils.ReadOnly
	vote *Hashed[*LaneVote]
	sigs []*Signature
}

// NewLaneQC constructs a new LaneQC.
func NewLaneQC(votes []*Signed[*LaneVote]) *LaneQC { _ = "STUB: not implemented"; return nil }

// Header .
func (m *LaneQC) Header() *BlockHeader { _ = "STUB: not implemented"; return nil }

// Next is the number of the first block not known to be available.
func (m *LaneQC) Next() BlockNumber {
	_ = "STUB: not implemented"
	return *

	// Verify verifies LaneQC against the committee.
	new(BlockNumber)
}

func (m *LaneQC) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// LaneQCConv is a protobuf converter for LaneQC.
var LaneQCConv = protoutils.Conv[*LaneQC, *pb.LaneQC]{
	Encode: func(m *LaneQC) *pb.LaneQC {
		return &pb.LaneQC{
			Vote: LaneVoteConv.Encode(m.vote.Msg()),
			Sigs: SignatureConv.EncodeSlice(m.sigs),
		}
	},
	Decode: func(m *pb.LaneQC) (*LaneQC, error) {
		vote, err := LaneVoteConv.DecodeReq(m.Vote)
		if err != nil {
			return nil, fmt.Errorf("vote: %w", err)
		}
		sigs, err := SignatureConv.DecodeSlice(m.Sigs)
		if err != nil {
			return nil, fmt.Errorf("sigs: %w", err)
		}
		return &LaneQC{vote: NewHashed(vote), sigs: sigs}, nil
	},
}
