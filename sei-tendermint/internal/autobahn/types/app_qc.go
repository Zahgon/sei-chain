package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// AppQC .
type AppQC struct {
	utils.ReadOnly
	vote *Hashed[*AppVote]
	sigs []*Signature
}

// NewAppQC create a new stateQC.
func NewAppQC(votes []*Signed[*AppVote]) *AppQC { _ = "STUB: not implemented"; return nil }

// Proposal .
func (m *AppQC) Proposal() *AppProposal { _ = "STUB: not implemented"; return nil }

// Next is the number of the next global block to finalize AppHash for.
func (m *AppQC) Next() RoadIndex { _ = "STUB: not implemented"; return *new(RoadIndex) }

// Verify verifies the AppQC against the committee.
func (m *AppQC) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// AppQCConv is a protobuf converter for AppQC.
var AppQCConv = protoutils.Conv[*AppQC, *pb.AppQC]{
	Encode: func(m *AppQC) *pb.AppQC {
		return &pb.AppQC{
			Vote: AppVoteConv.Encode(m.vote.Msg()),
			Sigs: SignatureConv.EncodeSlice(m.sigs),
		}
	},
	Decode: func(m *pb.AppQC) (*AppQC, error) {
		vote, err := AppVoteConv.DecodeReq(m.Vote)
		if err != nil {
			return nil, fmt.Errorf("proposal: %w", err)
		}
		sigs, err := SignatureConv.DecodeSlice(m.Sigs)
		if err != nil {
			return nil, fmt.Errorf("sigs: %w", err)
		}
		return &AppQC{vote: NewHashed(vote), sigs: sigs}, nil
	},
}
