package consensus

import (
	cstypes "github.com/sei-protocol/sei-chain/sei-tendermint/internal/consensus/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bits"
	tmcons "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/consensus"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// Message defines an interface that the consensus domain types implement. When
// a proto message is received on a consensus p2p Channel, it is wrapped and then
// converted to a Message via MsgFromProto.
type Message interface {
	ValidateBasic() error
}

// NewRoundStepMessage is sent for every step taken in the ConsensusState.
// For every height/round/step transition
type NewRoundStepMessage struct {
	cstypes.HRS
	SecondsSinceStartTime int64
	LastCommitRound       int32
}

// ValidateBasic performs basic validation.
func (m *NewRoundStepMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: SecondsSinceStartTime may be negative

// LastCommitRound will be -1 for the initial height, but we don't know what height this is
// since it can be specified in genesis. The reactor will have to validate this via
// ValidateHeight().

// ValidateHeight validates the height given the chain's initial height.
func (m *NewRoundStepMessage) ValidateHeight(initialHeight int64) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns a string representation.
func (m *NewRoundStepMessage) String() string { _ = "STUB: not implemented"; return "" }

// NewValidBlockMessage is sent when a validator observes a valid block B in some round r,
// i.e., there is a Proposal for block B and 2/3+ prevotes for the block B in the round r.
// In case the block is also committed, then IsCommit flag is set to true.
type NewValidBlockMessage struct {
	Height             int64
	Round              int32
	BlockPartSetHeader types.PartSetHeader
	BlockParts         *bits.BitArray
	IsCommit           bool
}

// ValidateBasic performs basic validation.
func (m *NewValidBlockMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *NewValidBlockMessage) String() string { _ = "STUB: not implemented"; return "" }

// ProposalMessage is sent when a new block is proposed.
type ProposalMessage struct {
	Proposal *types.Proposal
}

// ValidateBasic performs basic validation.
func (m *ProposalMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *ProposalMessage) String() string { _ = "STUB: not implemented"; return "" }

// ProposalPOLMessage is sent when a node needs POL round votes.
type ProposalPOLMessage struct {
	Height           int64
	ProposalPOLRound int32
	ProposalPOL      *bits.BitArray
}

// ValidateBasic performs basic validation.
func (m *ProposalPOLMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *ProposalPOLMessage) String() string { _ = "STUB: not implemented"; return "" }

// BlockPartMessage is sent when gossipping a piece of the proposed block.
type BlockPartMessage struct {
	Height int64
	Round  int32
	Part   *types.Part
}

// ValidateBasic performs basic validation.
func (m *BlockPartMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *BlockPartMessage) String() string { _ = "STUB: not implemented"; return "" }

// VoteMessage is sent when voting for a proposal (or lack thereof).
type VoteMessage struct {
	Vote *types.Vote
}

// ValidateBasic checks whether the vote within the message is well-formed.
func (m *VoteMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *VoteMessage) String() string { _ = "STUB: not implemented"; return "" }

// HasVoteMessage is sent to indicate that a particular vote has been received.
type HasVoteMessage struct {
	Height int64
	Round  int32
	Type   tmproto.SignedMsgType
	Index  int32
}

// ValidateBasic performs basic validation.
func (m *HasVoteMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *HasVoteMessage) String() string { _ = "STUB: not implemented"; return "" }

// VoteSetMaj23Message is sent to indicate that a given BlockID has seen +2/3 votes.
type VoteSetMaj23Message struct {
	Height  int64
	Round   int32
	Type    tmproto.SignedMsgType
	BlockID types.BlockID
}

// ValidateBasic performs basic validation.
func (m *VoteSetMaj23Message) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
func (m *VoteSetMaj23Message) String() string { _ = "STUB: not implemented"; return "" }

// VoteSetBitsMessage is sent to communicate the bit-array of votes seen for the
// BlockID.
type VoteSetBitsMessage struct {
	Height  int64
	Round   int32
	Type    tmproto.SignedMsgType
	BlockID types.BlockID
	Votes   *bits.BitArray
}

// ValidateBasic performs basic validation.
func (m *VoteSetBitsMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Votes.Size() can be zero if the node does not have any

// String returns a string representation.
func (m *VoteSetBitsMessage) String() string { _ = "STUB: not implemented"; return "" }

func newRoundStepMessageFromProto(pb *tmcons.NewRoundStep) (*NewRoundStepMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newValidBlockMessageFromProto(pb *tmcons.NewValidBlock) (*NewValidBlockMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func proposalMessageFromProto(pb *tmcons.Proposal) (*ProposalMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func proposalPOLMessageFromProto(pb *tmcons.ProposalPOL) (*ProposalPOLMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blockPartMessageFromProto(pb *tmcons.BlockPart) (*BlockPartMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func voteMessageFromProto(pb *tmcons.Vote) (*VoteMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasVoteMessageFromProto(pb *tmcons.HasVote) (*HasVoteMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func voteSetMaj23MessageFromProto(pb *tmcons.VoteSetMaj23) (*VoteSetMaj23Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func voteSetBitsMessageFromProto(pb *tmcons.VoteSetBits) (*VoteSetBitsMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (msg *NewRoundStepMessage) ToProto() *tmcons.NewRoundStep {
	_ = "STUB: not implemented"
	return nil
}

func (msg *NewValidBlockMessage) ToProto() *tmcons.NewValidBlock {
	_ = "STUB: not implemented"
	return nil
}

func (msg *ProposalMessage) ToProto() *tmcons.Proposal { _ = "STUB: not implemented"; return nil }

func (msg *ProposalPOLMessage) ToProto() *tmcons.ProposalPOL { _ = "STUB: not implemented"; return nil }

func (msg *BlockPartMessage) ToProto() *tmcons.BlockPart { _ = "STUB: not implemented"; return nil }

func (msg *VoteMessage) ToProto() *tmcons.Vote { _ = "STUB: not implemented"; return nil }

func (msg *HasVoteMessage) ToProto() *tmcons.HasVote { _ = "STUB: not implemented"; return nil }

func (msg *VoteSetMaj23Message) ToProto() *tmcons.VoteSetMaj23 {
	_ = "STUB: not implemented"
	return nil
}

func (msg *VoteSetBitsMessage) ToProto() *tmcons.VoteSetBits { _ = "STUB: not implemented"; return nil }
