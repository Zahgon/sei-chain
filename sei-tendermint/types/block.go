package types

import (
	"sync"
	"sync/atomic"
	"time"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bits"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/version"
)

// SkipLastResultsHashValidation controls whether LastResultsHash validation
// is skipped during block validation. This is set to true when the Giga
// executor is enabled, since it may produce different gas used values.
// Uses atomic.Bool for concurrency safety since NewApp may be called
// multiple times simultaneously.
var SkipLastResultsHashValidation atomic.Bool

const (
	// MaxHeaderBytes is a maximum header size.
	// NOTE: Because app hash can be of arbitrary size, the header is therefore not
	// capped in size and thus this number should be seen as a soft max
	MaxHeaderBytes int64 = 626

	// MaxOverheadForBlock - maximum overhead to encode a block (up to
	// MaxBlockSizeBytes in size) not including it's parts except Data.
	// This means it also excludes the overhead for individual transactions.
	//
	// Uvarint length of MaxBlockSizeBytes: 4 bytes
	// 2 fields (2 embedded):               2 bytes
	// Uvarint length of Data.Txs:          4 bytes
	// Data.Txs field:                      1 byte
	MaxOverheadForBlock int64 = 11
)

// Block defines the atomic unit of a Tendermint blockchain.
type Block struct {
	mtx sync.Mutex

	Header     `json:"header"`
	Data       `json:"data"`
	Evidence   EvidenceList `json:"evidence"`
	LastCommit *Commit      `json:"last_commit"`
}

func (b *Block) GetTxHashes() []TxHash { _ = "STUB: not implemented"; return nil }

// ValidateBasic performs basic validation that doesn't involve state data.
// It checks the internal consistency of the block.
// Further validation is done using state#ValidateBlock.
func (b *Block) ValidateBasic(policy ConsensusPolicy) error { _ = "STUB: not implemented"; return nil }

// Validate the last commit and its hash.

// Fall back to legacy hash calculation pre-6.4.

// NOTE: b.Data.Txs may be nil, but b.Data.Hash() still works fine.

// NOTE: b.Evidence may be nil, but we're just looping.

// fillHeader fills in any remaining header fields that are a function of the block data
func (b *Block) fillHeader() { _ = "STUB: not implemented"; return }

// Hash computes and returns the block hash.
// If the block is incomplete, block hash is nil for safety.
func (b *Block) Hash() tmbytes.HexBytes { _ = "STUB: not implemented"; return *new(tmbytes.HexBytes) }

// MakePartSet returns a PartSet containing parts of a serialized block.
// This is the form in which the block is gossipped to peers.
// CONTRACT: partSize is greater than zero.
func (b *Block) MakePartSet(partSize uint32) (*PartSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HashesTo is a convenience function that checks if a block hashes to the given argument.
// Returns false if the block is nil or the hash is empty.
func (b *Block) HashesTo(hash []byte) bool { _ = "STUB: not implemented"; return false }

// Size returns size of the block in bytes.
func (b *Block) Size() int { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of the block
//
// See StringIndented.
func (b *Block) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented String.
//
// Header
// Data
// Evidence
// LastCommit
// Hash
func (b *Block) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// StringShort returns a shortened string representation of the block.
func (b *Block) StringShort() string { _ = "STUB: not implemented"; return "" }

// ToProto converts Block to protobuf
func (b *Block) ToProto() (*tmproto.Block, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Block) ToReqBeginBlock(vals []*Validator) abci.RequestBeginBlock {
	_ = "STUB: not implemented"
	return *new(abci.RequestBeginBlock)
}

// b.LastCommit.Signatures is only empty on the trace path.

// FromProto sets a protobuf Block to the given pointer.
// It returns an error if the block is invalid.
func BlockFromProto(bp *tmproto.Block) (*Block, error) { _ = "STUB: not implemented"; return nil, nil }

//-----------------------------------------------------------------------------

// MaxDataBytes returns the maximum size of block's data.
//
// XXX: Panics on negative result.
func MaxDataBytes(maxBytes, evidenceBytes int64, valsCount int) int64 {
	_ = "STUB: not implemented"
	return 0
}

// MaxDataBytesNoEvidence returns the maximum size of block's data when
// evidence count is unknown. MaxEvidencePerBlock will be used for the size
// of evidence.
//
// XXX: Panics on negative result.
func MaxDataBytesNoEvidence(maxBytes int64, valsCount int) int64 {
	_ = "STUB: not implemented"
	return 0
}

// MakeBlock returns a new block with an empty header, except what can be
// computed from itself.
// It populates the same set of fields validated by ValidateBasic.
func MakeBlock(height int64, txs []Tx, lastCommit *Commit, evidence []Evidence) *Block {
	_ = "STUB: not implemented"
	return nil
}

//-----------------------------------------------------------------------------

// Header defines the structure of a Tendermint block header.
// NOTE: changes to the Header should be duplicated in:
// - header.Hash()
// - abci.Header
// - https://github.com/tendermint/tendermint/blob/master/spec/core/data_structures.md
type Header struct {
	// basic block info
	Version version.Consensus `json:"version"`
	ChainID string            `json:"chain_id"`
	Height  int64             `json:"height,string"`
	Time    time.Time         `json:"time"`

	// prev block info
	LastBlockID BlockID `json:"last_block_id"`

	// hashes of block data
	LastCommitHash tmbytes.HexBytes `json:"last_commit_hash"` // commit from validators from the last block
	DataHash       tmbytes.HexBytes `json:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash     tmbytes.HexBytes `json:"validators_hash"`      // validators for the current block
	NextValidatorsHash tmbytes.HexBytes `json:"next_validators_hash"` // validators for the next block
	ConsensusHash      tmbytes.HexBytes `json:"consensus_hash"`       // consensus params for current block
	AppHash            tmbytes.HexBytes `json:"app_hash"`             // state after txs from the previous block
	// root hash of all results from the txs from the previous block
	// see `deterministicResponseDeliverTx` to understand which parts of a tx is hashed into here
	LastResultsHash tmbytes.HexBytes `json:"last_results_hash"`

	// consensus info
	EvidenceHash    tmbytes.HexBytes `json:"evidence_hash"`    // evidence included in the block
	ProposerAddress Address          `json:"proposer_address"` // proposer recorded in the block header; preserved across re-proposals
}

// Populate the Header with state-derived data.
// Call this after MakeBlock to complete the Header.
func (h *Header) Populate(
	version version.Consensus, chainID string,
	timestamp time.Time, lastBlockID BlockID,
	valHash, nextValHash []byte,
	consensusHash, appHash, lastResultsHash []byte,
	proposerAddress Address,
) {
	_ = "STUB: not implemented"
	return
}

// ValidateBasic performs stateless validation on a Header returning an error
// if any validation fails.
//
// NOTE: Timestamp validation is subtle and handled elsewhere.
func (h Header) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Basic validation of hashes related to application data.
// Will validate fully against state in state#ValidateBlock.

// NOTE: AppHash is arbitrary length

// Hash returns the hash of the header.
// It computes a Merkle tree from the header fields
// ordered as they appear in the Header.
// Returns nil if ValidatorHash is missing,
// since a Header is not valid unless there is
// a ValidatorsHash (corresponding to the validator set).
func (h *Header) Hash() tmbytes.HexBytes { _ = "STUB: not implemented"; return *new(tmbytes.HexBytes) }

// StringIndented returns an indented string representation of the header.
func (h *Header) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts Header to protobuf
func (h *Header) ToProto() *tmproto.Header { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf Header to the given pointer.
// It returns an error if the header is invalid.
func HeaderFromProto(ph *tmproto.Header) (Header, error) {
	_ = "STUB: not implemented"
	return *new(Header), nil
}

//-------------------------------------

// BlockIDFlag indicates which BlockID the signature is for.
type BlockIDFlag byte

const (
	// BlockIDFlagAbsent - no vote was received from a validator.
	BlockIDFlagAbsent BlockIDFlag = iota + 1
	// BlockIDFlagCommit - voted for the Commit.BlockID.
	BlockIDFlagCommit
	// BlockIDFlagNil - voted for nil.
	BlockIDFlagNil
)

const (
	// Max size of commit without any commitSigs -> 82 for BlockID, 8 for Height, 4 for Round.
	MaxCommitOverheadBytes int64 = 94
	// Commit sig size is made up of 64 bytes for the signature, 20 bytes for the address,
	// 1 byte for the flag and 14 bytes for the timestamp
	MaxCommitSigBytes int64 = 109
)

// CommitSig is a part of the Vote included in a Commit.
type CommitSig struct {
	BlockIDFlag BlockIDFlag `json:"block_id_flag"`
	// WARNING: all fields below should be zeroed if BlockIDFlag == BlockIDFlagAbsent
	ValidatorAddress Address                  `json:"validator_address"`
	Timestamp        time.Time                `json:"timestamp"`
	Signature        utils.Option[crypto.Sig] `json:"signature"`
}

func MaxCommitBytes(valCount int) int64 {
	_ = "STUB: not implemented"
	// From the repeated commit sig field
	return 0
}

// NewCommitSigAbsent returns new CommitSig with BlockIDFlagAbsent. Other
// fields are all empty.
func NewCommitSigAbsent() CommitSig { _ = "STUB: not implemented"; return *new(CommitSig) }

// CommitSig returns a string representation of CommitSig.
//
// 1. first 6 bytes of signature
// 2. first 6 bytes of validator address
// 3. block ID flag
// 4. timestamp
func (cs CommitSig) String() string { _ = "STUB: not implemented"; return "" }

// BlockID returns the Commit's BlockID if CommitSig indicates signing,
// otherwise - empty BlockID.
func (cs CommitSig) BlockID(commitBlockID BlockID) BlockID {
	_ = "STUB: not implemented"
	return *new(BlockID)
}

// ValidateBasic performs basic validation.
func (cs CommitSig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Timestamp validation is subtle and handled elsewhere.

// ToProto converts CommitSig to protobuf
func (cs *CommitSig) ToProto() *tmproto.CommitSig { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf CommitSig to the given pointer.
// It returns an error if the CommitSig is invalid.
func (cs *CommitSig) FromProto(csp tmproto.CommitSig) error { _ = "STUB: not implemented"; return nil }

//-------------------------------------

// Commit contains the evidence that a block was committed by a set of validators.
// NOTE: Commit is empty for height 1, but never nil.
type Commit struct {
	// NOTE: The signatures are in order of address to preserve the bonded
	// ValidatorSet order.
	// Any peer with a block can gossip signatures by index with a peer without
	// recalculating the active ValidatorSet.
	Height     int64       `json:"height,string"`
	Round      int32       `json:"round"`
	BlockID    BlockID     `json:"block_id"`
	Signatures []CommitSig `json:"signatures"`

	// Memoized in first call to corresponding method.
	// NOTE: can't memoize in constructor because constructor isn't used for
	// unmarshaling.
	hash     tmbytes.HexBytes
	bitArray *bits.BitArray
}

// GetVote converts the CommitSig for the given valIdx to a Vote. Commits do
// not contain vote extensions, so the vote extension and vote extension
// signature will not be present in the returned vote.
// Returns nil if the precommit at valIdx is nil.
// Panics if valIdx >= commit.Size().
func (commit *Commit) GetVote(valIdx int32) (*Vote, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// VoteSignBytes returns the bytes of the Vote corresponding to valIdx for
// signing.
//
// The only unique part is the Timestamp - all other fields signed over are
// otherwise the same for all validators.
//
// Panics if valIdx >= commit.Size().
//
// See VoteSignBytes
func (commit *Commit) VoteSignBytes(chainID string, valIdx int32) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Size returns the number of signatures in the commit.
func (commit *Commit) Size() int { _ = "STUB: not implemented"; return 0 }

// ValidateBasic performs basic validation that doesn't involve state data.
// Does not actually check the cryptographic signatures.
func (commit *Commit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Hash returns the hash of the commit.
// It computes a Merkle tree from all commit fields: Height, Round, BlockID, and Signatures.
func (commit *Commit) Hash() tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

// Encode BlockID

// Build slice with metadata fields first, then signatures
// Fields: Height, Round, BlockID, followed by each CommitSig

// Cast to int64 for cdcEncode

// StringIndented returns a string representation of the commit.
func (commit *Commit) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts Commit to protobuf
func (commit *Commit) ToProto() *tmproto.Commit { _ = "STUB: not implemented"; return nil }

// FromProto sets a protobuf Commit to the given pointer.
// It returns an error if the commit is invalid.
func CommitFromProto(cp *tmproto.Commit) (*Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//-------------------------------------

// ToVoteSet constructs a VoteSet from the Commit and validator set.
// Panics if signatures from the commit can't be added to the voteset.
// Inverse of VoteSet.MakeCommit().
func (commit *Commit) ToVoteSet(chainID string, vals *ValidatorSet) *VoteSet {
	_ = "STUB: not implemented"
	return nil
}

// OK, some precommits can be missing.

//nolint:gosec // idx is bounded by commit.Signatures length which fits in int32

func (ec *Commit) Type() byte { _ = "STUB: not implemented"; return 0 }

// GetHeight returns height of the extended commit.
// Implements VoteSetReader.
func (ec *Commit) GetHeight() int64 {
	_ = "STUB: not implemented"

	// GetRound returns height of the extended commit.
	// Implements VoteSetReader.
	return 0
}

func (ec *Commit) GetRound() int32 {
	_ = "STUB: not implemented"

	// BitArray returns a BitArray of which validators voted for BlockID or nil in
	// this extended commit.
	// Implements VoteSetReader.
	return 0
}

func (ec *Commit) BitArray() *bits.BitArray { _ = "STUB: not implemented"; return nil }

// TODO: need to check the BlockID otherwise we could be counting conflicts,
//       not just the one with +2/3 !

// GetByIndex returns the vote corresponding to a given validator index.
// Panics if `index >= extCommit.Size()`.
// Implements VoteSetReader.
func (ec *Commit) GetByIndex(valIdx int32) (*Vote, bool) {
	_ = "STUB: not implemented"
	return nil,

		// IsCommit returns true if there is at least one signature.
		// Implements VoteSetReader.
		false
}

func (ec *Commit) IsCommit() bool { _ = "STUB: not implemented"; return false }

// legacyHash computes the commit hash using the pre-v6.4 algorithm, which
// only includes signatures (not Height, Round, or BlockID). This is needed
// to validate blocks that were created before the CommitHash change.
func (commit *Commit) legacyHash() tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

//-------------------------------------

// Data contains the set of transactions included in the block
type Data struct {

	// Txs that will be applied by state @ block.Height+1.
	// NOTE: not all txs here are valid.  We're just agreeing on the order first.
	// This means that block.AppHash does not include these txs.
	Txs Txs `json:"txs"`

	// Volatile
	hash tmbytes.HexBytes
}

// Hash returns the hash of the data
func (data *Data) Hash(overwrite bool) tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

// NOTE: leaves of merkle tree are TxIDs

// StringIndented returns an indented string representation of the transactions.
func (data *Data) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts Data to protobuf
func (data *Data) ToProto() tmproto.Data { _ = "STUB: not implemented"; return *new(tmproto.Data) }

// DataFromProto takes a protobuf representation of Data &
// returns the native type.
func DataFromProto(dp *tmproto.Data) (Data, error) {
	_ = "STUB: not implemented"
	return *new(Data), nil
}

//--------------------------------------------------------------------------------

// BlockID
type BlockID struct {
	Hash          tmbytes.HexBytes `json:"hash"`
	PartSetHeader PartSetHeader    `json:"parts"`
}

// Equals returns true if the BlockID matches the given BlockID
func (blockID BlockID) Equals(other BlockID) bool { _ = "STUB: not implemented"; return false }

// Key returns a machine-readable string representation of the BlockID
func (blockID BlockID) Key() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic performs basic validation.
func (blockID BlockID) ValidateBasic() error {
	_ = "STUB: not implemented"
	// Hash can be empty in case of POLBlockID in Proposal.
	return nil
}

// IsNil returns true if this is the BlockID of a nil block.
func (blockID BlockID) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsComplete returns true if this is a valid BlockID of a non-nil block.
func (blockID BlockID) IsComplete() bool { _ = "STUB: not implemented"; return false }

// String returns a human readable string representation of the BlockID.
//
// 1. hash
// 2. part set header
//
// See PartSetHeader#String
func (blockID BlockID) String() string { _ = "STUB: not implemented"; return "" }

// ToProto converts BlockID to protobuf
func (blockID *BlockID) ToProto() tmproto.BlockID {
	_ = "STUB: not implemented"
	return *new(tmproto.BlockID)
}

// FromProto sets a protobuf BlockID to the given pointer.
// It returns an error if the block id is invalid.
func BlockIDFromProto(bID *tmproto.BlockID) (*BlockID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProtoBlockIDIsNil is similar to the IsNil function on BlockID, but for the
// Protobuf representation.
func ProtoBlockIDIsNil(bID *tmproto.BlockID) bool { _ = "STUB: not implemented"; return false }
