package types

import (
	"errors"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

var (
	ErrInvalidBlockPartSignature = errors.New("error invalid block part signature")
	ErrInvalidBlockPartHash      = errors.New("error invalid block part hash")
)

// Proposal defines a block proposal for the consensus.
// It refers to the block by BlockID field.
// It must be signed by the correct proposer for the given Height/Round
// to be considered valid. It may depend on votes from a previous round,
// a so-called Proof-of-Lock (POL) round, as noted in the POLRound.
// If POLRound >= 0, then BlockID corresponds to the block that is locked in POLRound.
type Proposal struct {
	Type            tmproto.SignedMsgType
	Height          int64      `json:"height,string"`
	Round           int32      `json:"round"`     // there can not be greater than 2_147_483_647 rounds
	POLRound        int32      `json:"pol_round"` // -1 if null.
	BlockID         BlockID    `json:"block_id"`
	Timestamp       time.Time  `json:"timestamp"`
	Signature       crypto.Sig `json:"signature"`
	TxHashes        []TxHash   `json:"tx_keys"`
	Header          `json:"header"`
	LastCommit      *Commit      `json:"last_commit"`
	Evidence        EvidenceList `json:"evidence"`
	ProposerAddress Address      `json:"proposer_address"` // proposer of this proposal for the given height/round
}

// NewProposal returns a new Proposal.
// If there is no POLRound, polRound should be -1.
func NewProposal(height int64, round int32, polRound int32, blockID BlockID, ts time.Time, txHashes []TxHash, header Header, lastCommit *Commit, evidenceList EvidenceList, proposerAddress Address) *Proposal {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic validation.
func (p *Proposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ValidateBasic above would pass even if the BlockID was empty:

// NOTE: Timestamp validation is subtle and handled elsewhere.

// IsTimely validates that the block timestamp is 'timely' according to the proposer-based timestamp algorithm.
// To evaluate if a block is timely, its timestamp is compared to the local time of the validator along with the
// configured Precision and MsgDelay parameters.
// Specifically, a proposed block timestamp is considered timely if it is satisfies the following inequalities:
//
// localtime >= proposedBlockTime - Precision
// localtime <= proposedBlockTime + MsgDelay + Precision
//
// For more information on the meaning of 'timely', see the proposer-based timestamp specification:
// https://github.com/tendermint/tendermint/tree/master/spec/consensus/proposer-based-timestamp
func (p *Proposal) IsTimely(recvTime time.Time, sp SynchronyParams, round int32) bool {
	_ = "STUB: not implemented"
	// The message delay values are scaled as rounds progress.
	// Every 10 rounds, the message delay is doubled to allow consensus to
	// proceed in the case that the chosen value was too small for the given network conditions.
	// For more information and discussion on this mechanism, see the relevant github issue:
	// https://github.com/tendermint/spec/issues/371
	return false
}

//nolint:gosec // message delay is non zero
//nolint:gosec // round is validated non-negative above

// if the number of 'doublings' would overflow the size of the int, use the
// maximum instead.

// lhs is `proposedBlockTime - Precision` in the first inequality

// rhs is `proposedBlockTime + MsgDelay + Precision` in the second inequality

// String returns a string representation of the Proposal.
//
// 1. height
// 2. round
// 3. block ID
// 4. POL round
// 5. first 6 bytes of signature
// 6. timestamp
//
// See BlockID#String.
func (p *Proposal) String() string { _ = "STUB: not implemented"; return "" }

// ProposalSignBytes returns the proto-encoding of the canonicalized Proposal,
// for signing. Panics if the marshaling fails.
//
// The encoded Protobuf message is varint length-prefixed (using MarshalDelimited)
// for backwards-compatibility with the Amino encoding, due to e.g. hardware
// devices that rely on this encoding.
//
// See CanonicalizeProposal
func ProposalSignBytes(chainID string, p *tmproto.Proposal) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ToProto converts Proposal to protobuf
func (p *Proposal) ToProto() *tmproto.Proposal { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf Proposal to the given pointer.
// It returns an error if the proposal is invalid.
func ProposalFromProto(pp *tmproto.Proposal) (*Proposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
