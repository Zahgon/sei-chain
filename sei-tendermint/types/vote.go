package types

import (
	"errors"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

const (
	nilVoteStr string = "nil-Vote"
)

var (
	ErrVoteUnexpectedStep            = errors.New("unexpected step")
	ErrVoteInvalidValidatorIndex     = errors.New("invalid validator index")
	ErrVoteInvalidValidatorAddress   = errors.New("invalid validator address")
	ErrVoteInvalidSignature          = errors.New("invalid signature")
	ErrVoteInvalidBlockHash          = errors.New("invalid block hash")
	ErrVoteNonDeterministicSignature = errors.New("non-deterministic signature")
	ErrVoteNil                       = errors.New("nil vote")
)

type ErrVoteConflictingVotes struct {
	VoteA *Vote
	VoteB *Vote
}

func (err *ErrVoteConflictingVotes) Error() string { _ = "STUB: not implemented"; return "" }

func NewConflictingVoteError(vote1, vote2 *Vote) *ErrVoteConflictingVotes {
	_ = "STUB: not implemented"
	return nil
}

// Address is hex bytes.
type Address = crypto.Address

// Vote represents a prevote, precommit, or commit vote from validators for
// consensus.
type Vote struct {
	Type             tmproto.SignedMsgType    `json:"type"`
	Height           int64                    `json:"height,string"`
	Round            int32                    `json:"round"`    // assume there will not be greater than 2_147_483_647 rounds
	BlockID          BlockID                  `json:"block_id"` // zero if vote is nil.
	Timestamp        time.Time                `json:"timestamp"`
	ValidatorAddress Address                  `json:"validator_address"`
	ValidatorIndex   int32                    `json:"validator_index"`
	Signature        utils.Option[crypto.Sig] `json:"signature"`
}

// VoteFromProto attempts to convert the given serialization (Protobuf) type to
// our Vote domain type. No validation is performed on the resulting vote -
// this is left up to the caller to decide whether to call ValidateBasic or
// ValidateWithExtension.
func VoteFromProto(pv *tmproto.Vote) (*Vote, error) { _ = "STUB: not implemented"; return nil, nil }

// CommitSig converts the Vote to a CommitSig.
func (vote *Vote) CommitSig() CommitSig { _ = "STUB: not implemented"; return *new(CommitSig) }

// VoteSignBytes returns the proto-encoding of the canonicalized Vote, for
// signing. Panics if the marshaling fails.
//
// The encoded Protobuf message is varint length-prefixed (using MarshalDelimited)
// for backwards-compatibility with the Amino encoding, due to e.g. hardware
// devices that rely on this encoding.
//
// See CanonicalizeVote
func VoteSignBytes(chainID string, vote *tmproto.Vote) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (vote *Vote) Copy() *Vote { _ = "STUB: not implemented"; return nil }

// String returns a string representation of Vote.
//
// 1. validator index
// 2. first 6 bytes of validator address
// 3. height
// 4. round,
// 5. type byte
// 6. type string
// 7. first 6 bytes of block hash
// 8. first 6 bytes of signature
// 9. first 6 bytes of vote extension
// 10. timestamp
func (vote *Vote) String() string { _ = "STUB: not implemented"; return "" }

func (vote *Vote) verifyAndReturnProto(chainID string, pubKey crypto.PubKey) (*tmproto.Vote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify checks whether the signature associated with this vote corresponds to
// the given chain ID and public key. This function does not validate vote
// extension signatures - to do so, use VerifyWithExtension instead.
func (vote *Vote) Verify(chainID string, pubKey crypto.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic checks whether the vote is well-formed. It does not, however,
// check vote extensions - for vote validation with vote extension validation,
// use ValidateWithExtension.
func (vote *Vote) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Timestamp validation is subtle and handled elsewhere.

// BlockID.ValidateBasic would not err if we for instance have an empty hash but a
// non-empty PartsSetHeader:

// ToProto converts the handwritten type to proto generated type
// return type, nil if everything converts safely, otherwise nil, error
func (vote *Vote) ToProto() *tmproto.Vote { _ = "STUB: not implemented"; return nil }

func VotesToProto(votes []*Vote) []*tmproto.Vote { _ = "STUB: not implemented"; return nil }

// protobuf crashes when serializing "repeated" fields with nil elements
