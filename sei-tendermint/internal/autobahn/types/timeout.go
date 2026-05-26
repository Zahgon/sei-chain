package types

import (
	"errors"
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// TimeoutVote .
type TimeoutVote struct {
	utils.ReadOnly
	view            View
	latestPrepareQC utils.Option[ViewNumber]
}

// NewTimeoutVote creates a new TimeoutVote.
func NewTimeoutVote(view View, latestPrepareQC utils.Option[ViewNumber]) *TimeoutVote {
	_ = "STUB: not implemented"
	return nil
}

// View .
func (m *TimeoutVote) View() View {
	_ = "STUB: not implemented"

	// latestPrepareQCView is the highest view number for which a PrepareQC was observed by the node.
	return *new(View)
}

func (m *TimeoutVote) latestPrepareQCView() utils.Option[View] {
	_ = "STUB: not implemented"
	return nil
}

// FullTimeoutVote .
type FullTimeoutVote struct {
	utils.ReadOnly
	vote            *Signed[*TimeoutVote]
	latestPrepareQC utils.Option[*PrepareQC]
}

// NewFullTimeoutVote creates a new FullTimeoutVote.
func NewFullTimeoutVote(key SecretKey, view View, latestPrepareQC utils.Option[*PrepareQC]) *FullTimeoutVote {
	_ = "STUB: not implemented"
	return nil
}

// Vote .
func (m *FullTimeoutVote) Vote() *Signed[*TimeoutVote] {
	_ = "STUB: not implemented"

	// View .
	return nil
}

func (m *FullTimeoutVote) View() View { _ = "STUB: not implemented"; return *new(View) }

// Verify verifies the FullTimeoutVote against the committee.
func (m *FullTimeoutVote) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// TODO: verifying PrepareQC in all Timeout votes might be too inefficient.
// If it is, we can skip duplicated verification.

// TimeoutQC .
type TimeoutQC struct {
	utils.ReadOnly
	votes           []*Signed[*TimeoutVote]
	latestPrepareQC utils.Option[*PrepareQC]
}

// NewTimeoutQC creates a new TimeoutQC.
func NewTimeoutQC(fullVotes []*FullTimeoutVote) *TimeoutQC { _ = "STUB: not implemented"; return nil }

// View .
func (m *TimeoutQC) View() View { _ = "STUB: not implemented"; return *new(View) }

// Votes .
func (m *TimeoutQC) Votes() []*Signed[*TimeoutVote] {
	_ = "STUB: not implemented"

	// LatestPrepareQC returns the highest PrepareQC observed by signers.
	return nil
}

func (m *TimeoutQC) LatestPrepareQC() utils.Option[*PrepareQC] {
	_ = "STUB: not implemented"
	return nil

	// Verify verifies the TimeoutQC against the committee and the previous CommitQC.
	// Verifying TimeoutQC should NOT require previous TimeoutQC,
	// since observing prior TimeoutQCs is not required in the pb.
}

func (m *TimeoutQC) Verify(c *Committee, prev utils.Option[*CommitQC]) error {
	_ = "STUB: not implemented"
	// Verify the signatures.
	return nil
}

// Verify that we have enough votes.

// Check that the TimeoutQC is from the correct consensus instance.

// Check that the votes come from the same view.

// Check that the prepareQC is present iff needed.

func (m *TimeoutQC) reproposal() (*Proposal, bool) { _ = "STUB: not implemented"; return nil, false }

// TODO(gprusak): this unnecessarily accesses internal state and does the copy. Fix it.

// TimeoutVoteConv is the protobuf converter for TimeoutVote.
var TimeoutVoteConv = protoutils.Conv[*TimeoutVote, *pb.TimeoutVote]{
	Encode: func(m *TimeoutVote) *pb.TimeoutVote {
		return &pb.TimeoutVote{
			View: ViewConv.Encode(m.view),
			LatestPrepareQcViewNumber: func() *uint64 {
				if v, ok := m.latestPrepareQC.Get(); ok {
					return utils.Alloc(uint64(v))
				}
				return nil
			}(),
		}
	},
	Decode: func(m *pb.TimeoutVote) (*TimeoutVote, error) {
		view, err := ViewConv.DecodeReq(m.View)
		if err != nil {
			return nil, fmt.Errorf("view: %w", err)
		}
		return &TimeoutVote{
			view: view,
			latestPrepareQC: func() utils.Option[ViewNumber] {
				if v := m.LatestPrepareQcViewNumber; v != nil {
					return utils.Some(ViewNumber(*v))
				}
				return utils.None[ViewNumber]()
			}(),
		}, nil
	},
}

// FullTimeoutVoteConv is the protobuf converter for FullTimeoutVote.
var FullTimeoutVoteConv = protoutils.Conv[*FullTimeoutVote, *pb.FullTimeoutVote]{
	Encode: func(m *FullTimeoutVote) *pb.FullTimeoutVote {
		return &pb.FullTimeoutVote{
			Vote:            SignedMsgConv[*TimeoutVote]().Encode(m.vote),
			LatestPrepareQc: PrepareQCConv.EncodeOpt(m.latestPrepareQC),
		}
	},
	Decode: func(m *pb.FullTimeoutVote) (*FullTimeoutVote, error) {
		vote, err := SignedMsgConv[*TimeoutVote]().DecodeReq(m.Vote)
		if err != nil {
			return nil, fmt.Errorf("timeoutVote: %w", err)
		}
		latestPrepareQC, err := PrepareQCConv.DecodeOpt(m.LatestPrepareQc)
		if err != nil {
			return nil, fmt.Errorf("latestPrepareQc: %w", err)
		}
		return &FullTimeoutVote{
			vote:            vote,
			latestPrepareQC: latestPrepareQC,
		}, nil
	},
}

// TimeoutQCConv is the protobuf converter for TimeoutQC.
var TimeoutQCConv = protoutils.Conv[*TimeoutQC, *pb.TimeoutQC]{
	Encode: func(m *TimeoutQC) *pb.TimeoutQC {
		return &pb.TimeoutQC{
			Votes:           SignedMsgConv[*TimeoutVote]().EncodeSlice(m.votes),
			LatestPrepareQc: PrepareQCConv.EncodeOpt(m.latestPrepareQC),
		}
	},
	Decode: func(m *pb.TimeoutQC) (*TimeoutQC, error) {
		votes, err := SignedMsgConv[*TimeoutVote]().DecodeSlice(m.Votes)
		if err != nil {
			return nil, fmt.Errorf("votes: %w", err)
		}
		if len(votes) == 0 {
			return nil, errors.New("votes: missing")
		}
		latestPrepareQC, err := PrepareQCConv.DecodeOpt(m.LatestPrepareQc)
		if err != nil {
			return nil, fmt.Errorf("latestPrepareQc: %w", err)
		}
		return &TimeoutQC{
			votes:           votes,
			latestPrepareQC: latestPrepareQC,
		}, nil
	},
}
