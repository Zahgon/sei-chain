package consensus

import (
	"context"
	"errors"
	"sync"

	cstypes "github.com/sei-protocol/sei-chain/sei-tendermint/internal/consensus/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bits"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var (
	ErrPeerStateHeightRegression = errors.New("peer state height regression")
	ErrPeerStateInvalidStartTime = errors.New("peer state invalid startTime")
	ErrPeerStateSetNilVote       = errors.New("peer state set a nil vote")
	ErrPeerStateInvalidVoteIndex = errors.New("peer sent a vote with an invalid vote index")
)

// peerStateStats holds internal statistics for a peer.
type peerStateStats struct {
	Votes      int
	BlockParts int
}

func (pss peerStateStats) String() string { _ = "STUB: not implemented"; return "" }

// PeerState contains the known state of a peer, including its connection and
// threadsafe access to its PeerRoundState.
// NOTE: THIS GETS DUMPED WITH rpc/core/consensus.go.
// Be mindful of what you Expose.
type PeerState struct {
	peerID types.NodeID

	// NOTE: Modify below using setters, never directly.
	mtx    sync.RWMutex
	cancel context.CancelFunc
	PRS    cstypes.PeerRoundState
	Stats  *peerStateStats
}

// NewPeerState returns a new PeerState for the given node ID.
func NewPeerState(peerID types.NodeID) *PeerState { _ = "STUB: not implemented"; return nil }

// GetRoundState returns a shallow copy of the PeerRoundState. There's no point
// in mutating it since it won't change PeerState.
func (ps *PeerState) GetRoundState() *cstypes.PeerRoundState { _ = "STUB: not implemented"; return nil }

// ToJSON returns a json of PeerState. UNSTABLE.
func (ps *PeerState) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetHeight returns an atomic snapshot of the PeerRoundState's height used by
// the mempool to ensure peers are caught up before broadcasting new txs.
func (ps *PeerState) GetHeight() int64 { _ = "STUB: not implemented"; return 0 }

// SetHasProposal sets the given proposal as known for the peer.
func (ps *PeerState) SetHasProposal(proposal *types.Proposal) {
	_ = "STUB: not implemented"
	// ignore nil proposals
	return
}

// ps.PRS.ProposalBlockParts is set due to NewValidBlockMessage

// Nil until ProposalPOLMessage received.

// InitProposalBlockParts initializes the peer's proposal block parts header
// and bit array.
func (ps *PeerState) InitProposalBlockParts(partSetHeader types.PartSetHeader) {
	_ = "STUB: not implemented"
	return
}

// Apply the same memory exhaustion protection as in SetHasProposal

// SetHasProposalBlockPart sets the given block part index as known for the peer.
func (ps *PeerState) SetHasProposalBlockPart(height int64, round int32, index int) {
	_ = "STUB: not implemented"
	return
}

// PickVoteToSend picks a vote to send to the peer. It will return true if a
// vote was picked.
//
// NOTE: `votes` must be the correct Size() for the Height().
func (ps *PeerState) PickVoteToSend(votes types.VoteSetReader) (*types.Vote, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec // Type() returns a small enum value; no overflow risk

// lazily set data using 'votes'

// not something worth sending

//nolint:gosec // index is bounded by validator set size which fits in int32

func (ps *PeerState) getVoteBitArray(height int64, round int32, votesType tmproto.SignedMsgType) *bits.BitArray {
	_ = "STUB: not implemented"
	return nil
}

// 'round': A round for which we have a +2/3 commit.
func (ps *PeerState) ensureCatchupCommitRound(height int64, round int32, numValidators int) {
	_ = "STUB: not implemented"
	return
}

/*
	NOTE: This is wrong, 'round' could change.
	e.g. if orig round is not the same as block LastCommit round.
	if ps.CatchupCommitRound != -1 && ps.CatchupCommitRound != round {
		panic(fmt.Sprintf(
			"Conflicting CatchupCommitRound. Height: %v,
			Orig: %v,
			New: %v",
			height,
			ps.CatchupCommitRound,
			round))
	}
*/

// Nothing to do!

// EnsureVoteBitArrays ensures the bit-arrays have been allocated for tracking
// what votes this peer has received.
// NOTE: It's important to make sure that numValidators actually matches
// what the node sees as the number of validators for height.
func (ps *PeerState) EnsureVoteBitArrays(height int64, numValidators int) {
	_ = "STUB: not implemented"
	return
}

func (ps *PeerState) ensureVoteBitArrays(height int64, numValidators int) {
	_ = "STUB: not implemented"
	return
}

// RecordVote increments internal votes related statistics for this peer.
// It returns the total number of added votes.
func (ps *PeerState) RecordVote() int { _ = "STUB: not implemented"; return 0 }

// VotesSent returns the number of blocks for which peer has been sending us
// votes.
func (ps *PeerState) VotesSent() int { _ = "STUB: not implemented"; return 0 }

// RecordBlockPart increments internal block part related statistics for this peer.
// It returns the total number of added block parts.
func (ps *PeerState) RecordBlockPart() int { _ = "STUB: not implemented"; return 0 }

// BlockPartsSent returns the number of useful block parts the peer has sent us.
func (ps *PeerState) BlockPartsSent() int { _ = "STUB: not implemented"; return 0 }

// SetHasVote sets the given vote as known by the peer
func (ps *PeerState) SetHasVote(vote *types.Vote) error {
	_ = "STUB: not implemented"
	// sanity check
	return nil
}

// setHasVote will return an error when the index exceeds the bitArray length
func (ps *PeerState) setHasVote(height int64, round int32, voteType tmproto.SignedMsgType, index int32) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: some may be nil BitArrays -> no side effects

// https://github.com/tendermint/tendermint/issues/2871

// ApplyNewRoundStepMessage updates the peer state for the new round.
func (ps *PeerState) ApplyNewRoundStepMessage(msg *NewRoundStepMessage) {
	_ = "STUB: not implemented"
	return
}

// ignore duplicates or decreases

// we'll update the BitArray capacity later

// Peer caught up to CatchupCommitRound.
// Preserve psCatchupCommit!
// NOTE: We prefer to use prs.Precommits if
// pr.Round matches pr.CatchupCommitRound.

// shift Precommits to LastCommit

// we'll update the BitArray capacity later

// ApplyNewValidBlockMessage updates the peer state for the new valid block.
func (ps *PeerState) ApplyNewValidBlockMessage(msg *NewValidBlockMessage) {
	_ = "STUB: not implemented"
	return
}

// ApplyProposalPOLMessage updates the peer state for the new proposal POL.
func (ps *PeerState) ApplyProposalPOLMessage(msg *ProposalPOLMessage) {
	_ = "STUB: not implemented"
	return
}

// TODO: Merge onto existing ps.PRS.ProposalPOL?
// We might have sent some prevotes in the meantime.

// ApplyHasVoteMessage updates the peer state for the new vote.
func (ps *PeerState) ApplyHasVoteMessage(msg *HasVoteMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyVoteSetBitsMessage updates the peer state for the bit-array of votes
// it claims to have for the corresponding BlockID.
// `ourVotes` is a BitArray of votes we have for msg.BlockID
// NOTE: if ourVotes is nil (e.g. msg.Height < rs.Height),
// we conservatively overwrite ps's votes w/ msg.Votes.
func (ps *PeerState) ApplyVoteSetBitsMessage(msg *VoteSetBitsMessage, ourVotes *bits.BitArray) {
	_ = "STUB: not implemented"
	return
}
