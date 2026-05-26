package types

import (
	"context"
	"time"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/jsontypes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

// Evidence represents any provable malicious activity by a validator.
// Verification logic for each evidence is part of the evidence module.
type Evidence interface {
	ABCI() []abci.Misbehavior // forms individual evidence to be sent to the application
	Bytes() []byte            // bytes which comprise the evidence
	Hash() []byte             // hash of the evidence
	Height() int64            // height of the infraction
	String() string           // string format of the evidence
	Time() time.Time          // time of the infraction
	ValidateBasic() error     // basic consistency check

	// Implementations must support tagged encoding in JSON.
	jsontypes.Tagged
}

//--------------------------------------------------------------------------------------

type encodedVote struct {
	*Vote
	proto *tmproto.Vote
}

func (ev *encodedVote) ToProto() *tmproto.Vote { _ = "STUB: not implemented"; return nil }

func encodedVoteFromProto(pb *tmproto.Vote) (*encodedVote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEncodedVote(v *Vote) *encodedVote { _ = "STUB: not implemented"; return nil }

func (ev *encodedVote) Copy() *encodedVote { _ = "STUB: not implemented"; return nil }

// DuplicateVoteEvidence contains evidence of a single validator signing two conflicting votes.
type DuplicateVoteEvidence struct {
	VoteA *encodedVote `json:"vote_a"`
	VoteB *encodedVote `json:"vote_b"`

	// abci specific information
	TotalVotingPower int64 `json:",string"`
	ValidatorPower   int64 `json:",string"`
	Timestamp        time.Time
}

// TypeTag implements the jsontypes.Tagged interface.
func (*DuplicateVoteEvidence) TypeTag() string { _ = "STUB: not implemented"; return "" }

var _ Evidence = &DuplicateVoteEvidence{}

// NewDuplicateVoteEvidence creates DuplicateVoteEvidence with right ordering given
// two conflicting votes. If either of the votes is nil, the val set is nil or the voter is
// not in the val set, an error is returned
func NewDuplicateVoteEvidence(vote1, vote2 *Vote, blockTime time.Time, valSet *ValidatorSet,
) (*DuplicateVoteEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ABCI returns the application relevant representation of the evidence
func (dve *DuplicateVoteEvidence) ABCI() []abci.Misbehavior { _ = "STUB: not implemented"; return nil }

// Bytes returns the proto-encoded evidence as a byte array.
func (dve *DuplicateVoteEvidence) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Hash returns the hash of the evidence.
func (dve *DuplicateVoteEvidence) Hash() []byte { _ = "STUB: not implemented"; return nil }

// Height returns the height of the infraction
func (dve *DuplicateVoteEvidence) Height() int64 { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of the evidence.
func (dve *DuplicateVoteEvidence) String() string { _ = "STUB: not implemented"; return "" }

// Time returns the time of the infraction
func (dve *DuplicateVoteEvidence) Time() time.Time {
	_ = "STUB: not implemented"
	return *

	// ValidateBasic performs basic validation.
	new(time.Time)
}

func (dve *DuplicateVoteEvidence) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Enforce Votes are lexicographically sorted on blockID

// ValidateABCI validates the ABCI component of the evidence by checking the
// timestamp, validator power and total voting power.
func (dve *DuplicateVoteEvidence) ValidateABCI(
	val *Validator,
	valSet *ValidatorSet,
	evidenceTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateABCI populates the ABCI component of the evidence. This includes the
// validator power, timestamp and total voting power.
func (dve *DuplicateVoteEvidence) GenerateABCI(
	val *Validator,
	valSet *ValidatorSet,
	evidenceTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// ToProto encodes DuplicateVoteEvidence to protobuf
func (dve *DuplicateVoteEvidence) ToProto() *tmproto.DuplicateVoteEvidence {
	_ = "STUB: not implemented"
	return nil
}

// DuplicateVoteEvidenceFromProto decodes protobuf into DuplicateVoteEvidence
func DuplicateVoteEvidenceFromProto(pb *tmproto.DuplicateVoteEvidence) (*DuplicateVoteEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//------------------------------------ LIGHT EVIDENCE --------------------------------------

// LightClientAttackEvidence is a generalized evidence that captures all forms of known attacks on
// a light client such that a full node can verify, propose and commit the evidence on-chain for
// punishment of the malicious validators. There are three forms of attacks: Lunatic, Equivocation
// and Amnesia. These attacks are exhaustive. You can find a more detailed overview of this at
// tendermint/docs/architecture/adr-047-handling-evidence-from-light-client.md
//
// CommonHeight is used to indicate the type of attack. If the height is different to the conflicting block
// height, then nodes will treat this as of the Lunatic form, else it is of the Equivocation form.
type LightClientAttackEvidence struct {
	ConflictingBlock *LightBlock
	CommonHeight     int64 `json:",string"`

	// abci specific information
	ByzantineValidators []*Validator // validators in the validator set that misbehaved in creating the conflicting block
	TotalVotingPower    int64        `json:",string"` // total voting power of the validator set at the common height
	Timestamp           time.Time    // timestamp of the block at the common height
}

// TypeTag implements the jsontypes.Tagged interface.
func (*LightClientAttackEvidence) TypeTag() string { _ = "STUB: not implemented"; return "" }

var _ Evidence = &LightClientAttackEvidence{}

// ABCI forms an array of abci.Misbehavior for each byzantine validator
func (l *LightClientAttackEvidence) ABCI() []abci.Misbehavior {
	_ = "STUB: not implemented"
	return nil
}

// Bytes returns the proto-encoded evidence as a byte array
func (l *LightClientAttackEvidence) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// GetByzantineValidators finds out what style of attack LightClientAttackEvidence was and then works out who
// the malicious validators were and returns them. This is used both for forming the ByzantineValidators
// field and for validating that it is correct. Validators are ordered based on validator power
func (l *LightClientAttackEvidence) GetByzantineValidators(commonVals *ValidatorSet,
	trusted *SignedHeader) []*Validator {
	_ = "STUB: not implemented"
	return nil

	// First check if the header is invalid. This means that it is a lunatic attack and therefore we take the
	// validators who are in the commonVals and voted for the lunatic header
}

// validator wasn't in the common validator set

// This is an equivocation attack as both commits are in the same round. We then find the validators
// from the conflicting light block validator set that voted in both headers.
// Validator hashes are the same therefore the indexing order of validators are the same and thus we
// only need a single loop to find the validators that voted twice.

// if the rounds are different then this is an amnesia attack. Unfortunately, given the nature of the attack,
// we aren't able yet to deduce which are malicious validators and which are not hence we return an
// empty validator set.

// ConflictingHeaderIsInvalid takes a trusted header and matches it againt a conflicting header
// to determine whether the conflicting header was the product of a valid state transition
// or not. If it is then all the deterministic fields of the header should be the same.
// If not, it is an invalid header and constitutes a lunatic attack.
func (l *LightClientAttackEvidence) ConflictingHeaderIsInvalid(trustedHeader *Header) bool {
	_ = "STUB: not implemented"
	return false
}

// Hash returns the hash of the header and the commonHeight. This is designed to cause hash collisions
// with evidence that have the same conflicting header and common height but different permutations
// of validator commit signatures. The reason for this is that we don't want to allow several
// permutations of the same evidence to be committed on chain. Ideally we commit the header with the
// most commit signatures (captures the most byzantine validators) but anything greater than 1/3 is
// sufficient.
// TODO: We should change the hash to include the commit, header, total voting power, byzantine
// validators and timestamp
func (l *LightClientAttackEvidence) Hash() []byte { _ = "STUB: not implemented"; return nil }

// Height returns the last height at which the primary provider and witness provider had the same header.
// We use this as the height of the infraction rather than the actual conflicting header because we know
// that the malicious validators were bonded at this height which is important for evidence expiry
func (l *LightClientAttackEvidence) Height() int64 { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of LightClientAttackEvidence
func (l *LightClientAttackEvidence) String() string { _ = "STUB: not implemented"; return "" }

// Time returns the time of the common block where the infraction leveraged off.
func (l *LightClientAttackEvidence) Time() time.Time {
	_ = "STUB: not implemented"
	return *

	// ValidateBasic performs basic validation such that the evidence is consistent and can now be used for verification.
	new(time.Time)
}

func (l *LightClientAttackEvidence) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// this check needs to be done before we can run validate basic

// check that common height isn't ahead of the height of the conflicting block. It
// is possible that they are the same height if the light node witnesses either an
// amnesia or a equivocation attack.

// ValidateABCI validates the ABCI component of the evidence by checking the
// timestamp, byzantine validators and total voting power all match. ABCI
// components are validated separately because they can be re generated if
// invalid.
func (l *LightClientAttackEvidence) ValidateABCI(
	commonVals *ValidatorSet,
	trustedHeader *SignedHeader,
	evidenceTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Find out what type of attack this was and thus extract the malicious
// validators. Note, in the case of an Amnesia attack we don't have any
// malicious validators.

// Ensure this matches the validators that are listed in the evidence. They
// should be ordered based on power.

// GenerateABCI populates the ABCI component of the evidence: the timestamp,
// total voting power and byantine validators
func (l *LightClientAttackEvidence) GenerateABCI(
	commonVals *ValidatorSet,
	trustedHeader *SignedHeader,
	evidenceTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// ToProto encodes LightClientAttackEvidence to protobuf
func (l *LightClientAttackEvidence) ToProto() (*tmproto.LightClientAttackEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LightClientAttackEvidenceFromProto decodes protobuf
func LightClientAttackEvidenceFromProto(lpb *tmproto.LightClientAttackEvidence) (*LightClientAttackEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//------------------------------------------------------------------------------------------

// EvidenceList is a list of Evidence. Evidences is not a word.
type EvidenceList []Evidence

// StringIndented returns a string representation of the evidence.
func (evl EvidenceList) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ByteSize returns the total byte size of all the evidence
func (evl EvidenceList) ByteSize() int64 { _ = "STUB: not implemented"; return 0 }

// FromProto sets a protobuf EvidenceList to the given pointer.
func (evl *EvidenceList) FromProto(eviList *tmproto.EvidenceList) error {
	_ = "STUB: not implemented"
	return nil
}

// ToProto converts EvidenceList to protobuf
func (evl *EvidenceList) ToProto() (*tmproto.EvidenceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (evl EvidenceList) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (evl *EvidenceList) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Hash returns the simple merkle root hash of the EvidenceList.
func (evl EvidenceList) Hash() []byte {
	_ = "STUB: not implemented"
	// These allocations are required because Evidence is not of type Bytes, and
	// golang slices can't be typed cast. This shouldn't be a performance problem since
	// the Evidence size is capped.
	return nil
}

// TODO: We should change this to the hash. Using bytes contains some unexported data that
// may cause different hashes

func (evl EvidenceList) String() string { _ = "STUB: not implemented"; return "" }

// Has returns true if the evidence is in the EvidenceList.
func (evl EvidenceList) Has(evidence Evidence) bool { _ = "STUB: not implemented"; return false }

// ToABCI converts the evidence list to a slice of the ABCI protobuf messages
// for use when communicating the evidence to an application.
func (evl EvidenceList) ToABCI() []abci.Misbehavior { _ = "STUB: not implemented"; return nil }

//------------------------------------------ PROTO --------------------------------------

// EvidenceToProto is a generalized function for encoding evidence that conforms to the
// evidence interface to protobuf
func EvidenceToProto(evidence Evidence) (*tmproto.Evidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EvidenceFromProto is a generalized function for decoding protobuf into the
// evidence interface
func EvidenceFromProto(evidence *tmproto.Evidence) (Evidence, error) {
	_ = "STUB: not implemented"
	return *new(Evidence), nil
}

func init() {
	jsontypes.MustRegister((*DuplicateVoteEvidence)(nil))
	jsontypes.MustRegister((*LightClientAttackEvidence)(nil))
}

//-------------------------------------------- ERRORS --------------------------------------

// ErrInvalidEvidence wraps a piece of evidence and the error denoting how or why it is invalid.
type ErrInvalidEvidence struct {
	Evidence Evidence
	Reason   error
}

// NewErrInvalidEvidence returns a new EvidenceInvalid with the given err.
func NewErrInvalidEvidence(ev Evidence, err error) *ErrInvalidEvidence {
	_ = "STUB: not implemented"
	return nil
}

// Error returns a string representation of the error.
func (err *ErrInvalidEvidence) Error() string { _ = "STUB: not implemented"; return "" }

// ErrEvidenceOverflow is for when there the amount of evidence exceeds the max bytes.
type ErrEvidenceOverflow struct {
	Max int64
	Got int64
}

// NewErrEvidenceOverflow returns a new ErrEvidenceOverflow where got > max.
func NewErrEvidenceOverflow(max, got int64) ErrEvidenceOverflow {
	_ = "STUB: not implemented"
	return *new(ErrEvidenceOverflow)
}

// Error returns a string representation of the error.
func (err ErrEvidenceOverflow) Error() string { _ = "STUB: not implemented"; return "" }

//-------------------------------------------- MOCKING --------------------------------------

// unstable - use only for testing

// assumes the round to be 0 and the validator index to be 0
func NewMockDuplicateVoteEvidence(ctx context.Context, height int64, time time.Time, chainID string) (*DuplicateVoteEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assumes voting power to be 10 and validator to be the only one in the set
func NewMockDuplicateVoteEvidenceWithValidator(ctx context.Context, height int64, time time.Time, pv PrivValidator, chainID string) (*DuplicateVoteEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeMockVote(height int64, round, index int32, addr Address,
	blockID BlockID, time time.Time) *Vote {
	_ = "STUB: not implemented"
	return nil
}

func randBlockID() BlockID { _ = "STUB: not implemented"; return *new(BlockID) }
