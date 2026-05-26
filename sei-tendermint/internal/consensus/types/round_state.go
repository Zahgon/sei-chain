package types

import (
	"encoding/hex"
	"encoding/json"
	"sync"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

//-----------------------------------------------------------------------------
// RoundStepType enum type

// RoundStepType enumerates the state of the consensus state machine
type RoundStepType uint8 // These must be numeric, ordered.

// RoundStepType
const (
	RoundStepNewHeight     = RoundStepType(0x01) // Wait til CommitTime + timeoutCommit
	RoundStepNewRound      = RoundStepType(0x02) // Setup new round and go to RoundStepPropose
	RoundStepPropose       = RoundStepType(0x03) // Did propose, gossip proposal
	RoundStepPrevote       = RoundStepType(0x04) // Did prevote, gossip prevotes
	RoundStepPrevoteWait   = RoundStepType(0x05) // Did receive any +2/3 prevotes, start timeout
	RoundStepPrecommit     = RoundStepType(0x06) // Did precommit, gossip precommits
	RoundStepPrecommitWait = RoundStepType(0x07) // Did receive any +2/3 precommits, start timeout
	RoundStepCommit        = RoundStepType(0x08) // Entered commit state machine
	// NOTE: RoundStepNewHeight acts as RoundStepCommitWait.

	// NOTE: Update IsValid method if you change this!
)

// IsValid returns true if the step is valid, false if unknown/undefined.
func (rs RoundStepType) IsValid() bool { _ = "STUB: not implemented"; return false }

// String returns a string
func (rs RoundStepType) String() string { _ = "STUB: not implemented"; return "" }

// Cannot panic.

type SafeRoundState struct {
	internal RoundState
	mtx      sync.RWMutex
}

func NewSafeRoundState() SafeRoundState { _ = "STUB: not implemented"; return *new(SafeRoundState) }

func (s *SafeRoundState) CopyInternal() *RoundState { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) GetInternalPointer() *RoundState { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) Height() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SafeRoundState) SetHeight(h int64) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) Round() int32 { _ = "STUB: not implemented"; return 0 }

func (s *SafeRoundState) SetRound(r int32) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) Step() RoundStepType {
	_ = "STUB: not implemented"
	return *new(RoundStepType)
}

func (s *SafeRoundState) SetStep(t RoundStepType) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) StartTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *SafeRoundState) SetStartTime(t time.Time) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) CommitTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *SafeRoundState) SetCommitTime(t time.Time) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) LastCommit() *types.VoteSet { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetLastCommit(c *types.VoteSet) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) CommitRound() int32 { _ = "STUB: not implemented"; return 0 }

func (s *SafeRoundState) SetCommitRound(r int32) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) Votes() *HeightVoteSet { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetVotes(v *HeightVoteSet) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) Validators() *types.ValidatorSet { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) Leader() crypto.PubKey {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey)
}

func (s *SafeRoundState) SetValidators(v *types.ValidatorSet) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) Proposal() *types.Proposal { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetProposal(p *types.Proposal) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) ProposalReceiveTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (s *SafeRoundState) SetProposalReceiveTime(p time.Time) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) ProposalBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetProposalBlock(p *types.Block) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) ProposalBlockParts() *types.PartSet { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetProposalBlockParts(p *types.PartSet) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) LockedRound() int32 { _ = "STUB: not implemented"; return 0 }

func (s *SafeRoundState) SetLockedRound(p int32) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) LockedBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetLockedBlock(p *types.Block) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) LockedBlockParts() *types.PartSet { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetLockedBlockParts(p *types.PartSet) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) ValidRound() int32 { _ = "STUB: not implemented"; return 0 }

func (s *SafeRoundState) SetValidRound(p int32) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) ValidBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetValidBlock(p *types.Block) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) ValidBlockParts() *types.PartSet { _ = "STUB: not implemented"; return nil }

func (s *SafeRoundState) SetValidBlockParts(p *types.PartSet) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) LastValidators() *types.ValidatorSet {
	_ = "STUB: not implemented"
	return nil
}

func (s *SafeRoundState) SetLastValidators(p *types.ValidatorSet) {
	_ = "STUB: not implemented"
	return
}

func (s *SafeRoundState) TriggeredTimeoutPrecommit() bool { _ = "STUB: not implemented"; return false }

func (s *SafeRoundState) SetTriggeredTimeoutPrecommit(p bool) { _ = "STUB: not implemented"; return }

func (s *SafeRoundState) RoundStateEvent() types.EventDataRoundState {
	_ = "STUB: not implemented"
	return *new(types.EventDataRoundState)
}

func (s *SafeRoundState) NewRoundEvent() types.EventDataNewRound {
	_ = "STUB: not implemented"
	return *new(types.EventDataNewRound)
}

func (s *SafeRoundState) CompleteProposalEvent() types.EventDataCompleteProposal {
	_ = "STUB: not implemented"
	return *new(types.EventDataCompleteProposal)
}

//-----------------------------------------------------------------------------

// RoundState defines the internal consensus state.
// NOTE: Not thread safe. Should only be manipulated by functions downstream
// of the cs.receiveRoutine
type RoundState struct {
	HRS
	StartTime time.Time

	// Subjective time when +2/3 precommits for Block at Round were found
	CommitTime          time.Time
	Validators          *types.ValidatorSet
	Proposal            *types.Proposal
	ProposalReceiveTime time.Time
	ProposalBlock       *types.Block
	ProposalBlockParts  *types.PartSet
	LockedRound         int32
	LockedBlock         *types.Block
	LockedBlockParts    *types.PartSet

	// The variables below starting with "Valid..." derive their name from
	// the algorithm presented in this paper:
	// [The latest gossip on BFT consensus](https://arxiv.org/abs/1807.04938).
	// Therefore, "Valid...":
	//   * means that the block or round that the variable refers to has
	//     received 2/3+ non-`nil` prevotes (a.k.a. a *polka*)
	//   * has nothing to do with whether the Application returned "Accept" in its
	//     response to `ProcessProposal`, or "Reject"

	ValidRound      int32          // Last known round with POL for non-nil valid block.
	ValidBlock      *types.Block   // Last known block of POL mentioned above.
	ValidBlockParts *types.PartSet // Last known block parts of POL mentioned above.

	Votes                     *HeightVoteSet
	CommitRound               int32
	LastCommit                *types.VoteSet // Last precommits at Height-1
	LastValidators            *types.ValidatorSet
	TriggeredTimeoutPrecommit bool
}

// 32 bytes crypto hash seed, generated via random.org.
// THIS IS A PROTOCOL CONSTANT, DO NOT CHANGE.
var leaderElectionSeed = [32]byte(utils.OrPanic1(hex.DecodeString(
	"3793f16d412703e5805755e5282f681c70e771f151c8864c656c6c259243f85f",
)))

// Leader for each round is drawn at random from the validator set with
// probabilities proportional to the voting powers.
//
// The following pseudorandom function is used to select the leader:
// pos(height,round) := hash(height ++ round)
// Validators are assigned subintervals of [0,TotalVotingPower) of length
// equal to their voting poser.
// Validator i is the leader of (height,round) <=> pos(height,round)%TotalVotingPower \in validator_interval[i]
func (rs *RoundState) Leader() crypto.PubKey {
	_ = "STUB: not implemented"
	// sha256 does not support seed natively, so we add it by hand.
	return *new(crypto.PubKey)
}

//nolint:gosec
//nolint:gosec

// Compressed version of the RoundState for use in RPC.
// Used only for JSON representation.
type RoundStateSimple struct {
	HeightRoundStep   string              `json:"height/round/step"`
	StartTime         time.Time           `json:"start_time"`
	ProposalBlockHash bytes.HexBytes      `json:"proposal_block_hash"`
	LockedBlockHash   bytes.HexBytes      `json:"locked_block_hash"`
	ValidBlockHash    bytes.HexBytes      `json:"valid_block_hash"`
	Votes             json.RawMessage     `json:"height_vote_set"`
	Proposer          types.ValidatorInfo `json:"proposer"`
}

// Compress the RoundState to RoundStateSimple.
func (rs *RoundState) RoundStateSimple() RoundStateSimple {
	_ = "STUB: not implemented"
	return *new(RoundStateSimple)
}

// NewRoundEvent returns the RoundState with proposer information as an event.
func (rs *RoundState) NewRoundEvent() types.EventDataNewRound {
	_ = "STUB: not implemented"
	return *new(types.EventDataNewRound)
}

// CompleteProposalEvent returns information about a proposed block as an event.
func (rs *RoundState) CompleteProposalEvent() types.EventDataCompleteProposal {
	_ = "STUB: not implemented"
	// We must construct BlockID from ProposalBlock and ProposalBlockParts
	// cs.Proposal is not guaranteed to be set when this function is called
	return *new(types.EventDataCompleteProposal)
}

// RoundStateEvent returns the H/R/S of the RoundState as an event.
func (rs *RoundState) RoundStateEvent() types.EventDataRoundState {
	_ = "STUB: not implemented"
	return *new(types.EventDataRoundState)
}

// String returns a string
func (rs *RoundState) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns a string
func (rs *RoundState) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// StringShort returns a string
func (rs *RoundState) StringShort() string { _ = "STUB: not implemented"; return "" }
