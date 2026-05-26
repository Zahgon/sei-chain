package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// CommitQC .
type CommitQC struct {
	utils.ReadOnly
	vote *Hashed[*CommitVote]
	sigs []*Signature
}

// NewCommitQC constructs a new CommitQC.
func NewCommitQC(votes []*Signed[*CommitVote]) *CommitQC { _ = "STUB: not implemented"; return nil }

// Proposal .
func (m *CommitQC) Proposal() *Proposal { _ = "STUB: not implemented"; return nil }

// Index .
func (m *CommitQC) Index() RoadIndex { _ = "STUB: not implemented"; return *new(RoadIndex) }

// LaneRange returns the range of lane blocks.
func (m *CommitQC) LaneRange(lane LaneID) *LaneRange { _ = "STUB: not implemented"; return nil }

// GlobalRange returns the finalized global block range.
func (m *CommitQC) GlobalRange(c *Committee) GlobalRange {
	_ = "STUB: not implemented"
	return *new(GlobalRange)
}

// Verify verifies the CommitQC against the committee.
// Currently it doesn't require the previous CommitQC.
func (m *CommitQC) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// FullCommitQC is a CommitQC with the headers of the blocks finalized by it.
type FullCommitQC struct {
	utils.ReadOnly
	qc      *CommitQC
	headers []*BlockHeader
}

// NewFullCommitQC constructs a new FullCommitQC.
func NewFullCommitQC(qc *CommitQC, headers []*BlockHeader) *FullCommitQC {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // total lane range len is a small bounded value representing block count in a QC

// QC CommitQC.
func (m *FullCommitQC) QC() *CommitQC {
	_ = "STUB: not implemented"

	// Headers of the blocks finalized by the QC.
	return nil
}

func (m *FullCommitQC) Headers() []*BlockHeader {
	_ = "STUB: not implemented"

	// Index .
	return nil
}

func (m *FullCommitQC) Index() RoadIndex {
	_ = "STUB: not implemented"
	return *

	// Verify verifies the FullCommitQC against the committee.
	new(RoadIndex)
}

func (m *FullCommitQC) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // global range len is a small bounded value representing block count in a QC

// CommitQCConv is a protobuf converter for CommitQC.
var CommitQCConv = protoutils.Conv[*CommitQC, *pb.CommitQC]{
	Encode: func(m *CommitQC) *pb.CommitQC {
		return &pb.CommitQC{
			Vote: CommitVoteConv.Encode(m.vote.Msg()),
			Sigs: SignatureConv.EncodeSlice(m.sigs),
		}
	},
	Decode: func(m *pb.CommitQC) (*CommitQC, error) {
		vote, err := CommitVoteConv.DecodeReq(m.Vote)
		if err != nil {
			return nil, fmt.Errorf("vote: %w", err)
		}
		sigs, err := SignatureConv.DecodeSlice(m.Sigs)
		if err != nil {
			return nil, fmt.Errorf("sigs: %w", err)
		}
		return &CommitQC{vote: NewHashed(vote), sigs: sigs}, nil
	},
}

// FullCommitQCConv is a protobuf converter for FullCommitQC.
var FullCommitQCConv = protoutils.Conv[*FullCommitQC, *pb.FullCommitQC]{
	Encode: func(m *FullCommitQC) *pb.FullCommitQC {
		return &pb.FullCommitQC{
			Qc:      CommitQCConv.Encode(m.qc),
			Headers: BlockHeaderConv.EncodeSlice(m.headers),
		}
	},
	Decode: func(m *pb.FullCommitQC) (*FullCommitQC, error) {
		qc, err := CommitQCConv.DecodeReq(m.Qc)
		if err != nil {
			return nil, fmt.Errorf("qC: %w", err)
		}
		headers, err := BlockHeaderConv.DecodeSlice(m.Headers)
		if err != nil {
			return nil, fmt.Errorf("headers: %w", err)
		}
		return &FullCommitQC{qc: qc, headers: headers}, nil
	},
}
