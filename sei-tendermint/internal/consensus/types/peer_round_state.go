package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bits"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

//-----------------------------------------------------------------------------

type HRS struct {
	Height int64         // Height peer is at
	Round  int32         // Round peer is at, -1 if unknown.
	Step   RoundStepType // Step peer is at
}

func (a HRS) Cmp(b HRS) int { _ = "STUB: not implemented"; return 0 }

// PeerRoundState contains the known state of a peer.
// NOTE: Read-only when returned by PeerState.GetRoundState().
type PeerRoundState struct {
	HRS

	StartTime time.Time // Estimated start of round 0 at this height

	// WARNING: this only partially validated, so logic accessing it should be conservative.
	Proposal                   bool // True if peer has proposal for this round
	ProposalBlockPartSetHeader types.PartSetHeader
	ProposalBlockParts         *bits.BitArray
	ProposalPOLRound           int32          // Proposal's POL round. -1 if none.
	ProposalPOL                *bits.BitArray // nil until ProposalPOLMessage received.
	Prevotes                   *bits.BitArray // All votes peer has for this round
	Precommits                 *bits.BitArray // All precommits peer has for this round
	LastCommitRound            int32          // Round of commit for last height. -1 if none.
	LastCommit                 *bits.BitArray // All commit precommits of commit for last height.

	CatchupCommitRound int32          // Round that we have commit for. Not necessarily unique. -1 if none.
	CatchupCommit      *bits.BitArray // All commit precommits peer has for this height & CatchupCommitRound
}

// Copy provides a deep copy operation. Because many of the fields in
// the PeerRound struct are pointers, we need an explicit deep copy
// operation to avoid a non-obvious shared data situation.
func (prs PeerRoundState) Copy() PeerRoundState {
	_ = "STUB: not implemented"
	// this works because it's not a pointer receiver so it's
	// already, effectively a copy.
	return *new(PeerRoundState)
}
