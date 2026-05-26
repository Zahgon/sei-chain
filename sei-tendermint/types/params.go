package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/ed25519"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

const (
	// MaxBlockSizeBytes is the maximum permitted size of the blocks.
	MaxBlockSizeBytes = 104857600 // 100MB

	// BlockPartSizeBytes is the size of one block part.
	BlockPartSizeBytes uint32 = 1048576 // 1MB

	// MaxBlockPartsCount is the maximum number of block parts,
	// this also the maximum number of bits in the bit array
	MaxBlockPartsCount = (MaxBlockSizeBytes / BlockPartSizeBytes) + 1

	ABCIPubKeyTypeEd25519 = ed25519.KeyType
)

var ABCIPubKeyTypesToNames = map[string]string{
	ABCIPubKeyTypeEd25519: ed25519.PublicKeyName,
}

// ConsensusParams contains consensus critical parameters that determine the
// validity of blocks.
type ConsensusParams struct {
	Block     BlockParams     `json:"block"`
	Evidence  EvidenceParams  `json:"evidence"`
	Validator ValidatorParams `json:"validator"`
	Version   VersionParams   `json:"version"`
	Synchrony SynchronyParams `json:"synchrony"`
	Timeout   TimeoutParams   `json:"timeout"`
	ABCI      ABCIParams      `json:"abci"`
}

// HashedParams is a subset of ConsensusParams.
// It is amino encoded and hashed into
// the Header.ConsensusHash.
type HashedParams struct {
	BlockMaxBytes int64
	BlockMaxGas   int64
}

// BlockParams define limits on the block size and gas plus minimum time
// between blocks.
type BlockParams struct {
	MaxBytes      int64 `json:"max_bytes,string"`
	MaxGas        int64 `json:"max_gas,string"`
	MinTxsInBlock int64 `json:"min_txs_in_block,string"` // deprecated
	MaxGasWanted  int64 `json:"max_gas_wanted,string"`
}

// EvidenceParams determine how we handle evidence of malfeasance.
type EvidenceParams struct {
	MaxAgeNumBlocks int64         `json:"max_age_num_blocks,string"` // only accept new evidence more recent than this
	MaxAgeDuration  time.Duration `json:"max_age_duration,string"`
	MaxBytes        int64         `json:"max_bytes,string"`
}

// ValidatorParams restrict the public key types validators can use.
// NOTE: uses ABCI pubkey naming, not Amino names.
type ValidatorParams struct {
	PubKeyTypes []string `json:"pub_key_types"`
}

type VersionParams struct {
	AppVersion uint64 `json:"app_version,string"`
}

// SynchronyParams influence the validity of block timestamps.
// For more information on the relationship of the synchrony parameters to
// block validity, see the Proposer-Based Timestamps specification:
// https://github.com/tendermint/tendermint/blob/master/spec/consensus/proposer-based-timestamp/README.md
type SynchronyParams struct {
	Precision    time.Duration `json:"precision,string"`
	MessageDelay time.Duration `json:"message_delay,string"`
}

// TimeoutParams configure the timings of the steps of the Tendermint consensus algorithm.
type TimeoutParams struct {
	Propose             time.Duration `json:"propose,string"`
	ProposeDelta        time.Duration `json:"propose_delta,string"`
	Vote                time.Duration `json:"vote,string"`
	VoteDelta           time.Duration `json:"vote_delta,string"`
	Commit              time.Duration `json:"commit,string"`
	BypassCommitTimeout bool          `json:"bypass_commit_timeout"`
}

// ABCIParams configure ABCI functionality specific to the Application Blockchain
// Interface.
type ABCIParams struct {
	VoteExtensionsEnableHeight int64 `json:"vote_extensions_enable_height"`
	RecheckTx                  bool  `json:"recheck_tx"`
}

// DefaultConsensusParams returns a default ConsensusParams.
func DefaultConsensusParams() *ConsensusParams { _ = "STUB: not implemented"; return nil }

// DefaultBlockParams returns a default BlockParams.
func DefaultBlockParams() BlockParams { _ = "STUB: not implemented"; return *new(BlockParams) }

// 21MB
// Default, can be increased and tuned as needed

// DefaultEvidenceParams returns a default EvidenceParams.
func DefaultEvidenceParams() EvidenceParams { _ = "STUB: not implemented"; return *new(EvidenceParams) }

// 27.8 hrs at 1block/s

// 1MB

// DefaultValidatorParams returns a default ValidatorParams, which allows
// only ed25519 pubkeys.
func DefaultValidatorParams() ValidatorParams {
	_ = "STUB: not implemented"
	return *new(ValidatorParams)
}

func DefaultVersionParams() VersionParams { _ = "STUB: not implemented"; return *new(VersionParams) }

func DefaultSynchronyParams() SynchronyParams {
	_ = "STUB: not implemented"
	return *

	// 505ms was selected as the default to enable chains that have validators in
	// mixed leap-second handling environments.
	// For more information, see: https://github.com/tendermint/tendermint/issues/7724
	new(SynchronyParams)
}

// SynchronyParamsOrDefaults returns the SynchronyParams, filling in any zero values
// with the Tendermint defined default values.
func (s SynchronyParams) SynchronyParamsOrDefaults() SynchronyParams {
	_ = "STUB: not implemented"
	// TODO: Remove this method and all uses once development on v0.37 begins.
	// See: https://github.com/tendermint/tendermint/issues/8187
	return *new(SynchronyParams)
}

func DefaultTimeoutParams() TimeoutParams { _ = "STUB: not implemented"; return *new(TimeoutParams) }

func DefaultABCIParams() ABCIParams {
	_ = "STUB: not implemented"

	// When set to 0, vote extensions are not required.
	return *new(ABCIParams)
}

// When true, run CheckTx on each transaction in the mempool after each height.

// TimeoutParamsOrDefaults returns the SynchronyParams, filling in any zero values
// with the Tendermint defined default values.
func (t TimeoutParams) TimeoutParamsOrDefaults() TimeoutParams {
	_ = "STUB: not implemented"
	// TODO: Remove this method and all uses once development on v0.37 begins.
	// See: https://github.com/tendermint/tendermint/issues/8187
	return *new(TimeoutParams)
}

// ProposeTimeout returns the amount of time to wait for a proposal.
func (t TimeoutParams) ProposeTimeout(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// VoteTimeout returns the amount of time to wait for remaining votes after receiving any +2/3 votes.
func (t TimeoutParams) VoteTimeout(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// CommitTime accepts ti, the time at which the consensus engine received +2/3
// precommits for a block and returns the point in time at which the consensus
// engine should begin consensus on the next block.
func (t TimeoutParams) CommitTime(ti time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (val *ValidatorParams) IsValidPubkeyType(pubkeyType string) bool {
	_ = "STUB: not implemented"
	return false
}

func (params *ConsensusParams) Complete() { _ = "STUB: not implemented"; return }

// Validate validates the ConsensusParams to ensure all values are within their
// allowed limits, and returns an error if they are not.
func (params ConsensusParams) ValidateConsensusParams() error {
	_ = "STUB: not implemented"
	return nil
}

// Check if keyType is a known ABCIPubKeyType

func (params ConsensusParams) ValidateUpdate(updated *tmproto.ConsensusParams, h int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash returns a hash of a subset of the parameters to store in the block header.
// Only the Block.MaxBytes and Block.MaxGas are included in the hash.
// This allows the ConsensusParams to evolve more without breaking the block
// protocol. No need for a Merkle tree here, just a small struct to hash.
// TODO: We should hash the other parameters as well
func (params ConsensusParams) HashConsensusParams() []byte { _ = "STUB: not implemented"; return nil }

func (params *ConsensusParams) Equals(params2 *ConsensusParams) bool {
	_ = "STUB: not implemented"
	return false
}

// Update returns a copy of the params with updates from the non-zero fields of p2.
// NOTE: note: must not modify the original
func (params ConsensusParams) UpdateConsensusParams(params2 *tmproto.ConsensusParams) ConsensusParams {
	_ = "STUB: not implemented"
	// explicit copy
	return *new(ConsensusParams)
}

// we must defensively consider any structs may be nil

// Copy params2.Validator.PubkeyTypes, and set result's value to the copy.
// This avoids having to initialize the slice to 0 values, and then write to it again.

func (params *ConsensusParams) ToProto() tmproto.ConsensusParams {
	_ = "STUB: not implemented"
	return *new(tmproto.ConsensusParams)
}

func ConsensusParamsFromProto(pbParams tmproto.ConsensusParams) ConsensusParams {
	_ = "STUB: not implemented"
	return *new(ConsensusParams)
}
