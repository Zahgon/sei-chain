package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// PrepareQC .
type PrepareQC struct {
	utils.ReadOnly
	vote *Hashed[*PrepareVote]
	sigs []*Signature
}

// NewPrepareQC creates a new PrepareQC.
// PANICS if votes is empty.
func NewPrepareQC(votes []*Signed[*PrepareVote]) *PrepareQC { _ = "STUB: not implemented"; return nil }

// Proposal .
func (m *PrepareQC) Proposal() *Proposal { _ = "STUB: not implemented"; return nil }

// View .
func (m *PrepareQC) View() View { _ = "STUB: not implemented"; return *new(View) }

// Verify verifies the PrepareQC against the committee.
// Currently it doesn't require the previous CommitQC.
func (m *PrepareQC) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// PrepareQCConv is a protobuf converter for PrepareQC.
var PrepareQCConv = protoutils.Conv[*PrepareQC, *pb.PrepareQC]{
	Encode: func(m *PrepareQC) *pb.PrepareQC {
		return &pb.PrepareQC{
			Vote: PrepareVoteConv.Encode(m.vote.Msg()),
			Sigs: SignatureConv.EncodeSlice(m.sigs),
		}
	},
	Decode: func(m *pb.PrepareQC) (*PrepareQC, error) {
		vote, err := PrepareVoteConv.DecodeReq(m.Vote)
		if err != nil {
			return nil, fmt.Errorf("vote: %w", err)
		}
		sigs, err := SignatureConv.DecodeSlice(m.Sigs)
		if err != nil {
			return nil, fmt.Errorf("sigs: %w", err)
		}
		return &PrepareQC{vote: NewHashed(vote), sigs: sigs}, nil
	},
}
