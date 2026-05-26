package types

import (
	"time"

	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

type RequestInfo struct {
	Version      string
	BlockVersion uint64
	P2PVersion   uint64
	AbciVersion  string
}

// RequestInitChain carries the genesis-time initialization inputs passed from
// consensus into the application when bootstrapping a chain.
type RequestInitChain struct {
	Time            time.Time
	ChainId         string
	ConsensusParams *tmproto.ConsensusParams
	AppStateBytes   []byte
	InitialHeight   int64
}

type RequestQuery struct {
	Data   []byte
	Path   string
	Height int64
	Prove  bool
}

type RequestBeginBlock struct {
	Hash                []byte
	Header              tmproto.Header
	LastCommitInfo      LastCommitInfo
	ByzantineValidators []Evidence
	Simulate            bool
}

type RequestEndBlock struct {
	Height       int64
	BlockGasUsed int64
}

type ResponseInitChain struct {
	Validators []ValidatorUpdate
	AppHash    []byte
}

type ResponseGetTxPriorityHint struct {
	Priority int64
}

// RequestListSnapshots is emitted at the start of state sync to ask the application
// which previously committed snapshots are available for peers to download.
type RequestListSnapshots struct{}

// ResponseListSnapshots returns the snapshot metadata advertised by the
// application so Tendermint can decide which snapshot to fetch.
type ResponseListSnapshots struct {
	Snapshots []*Snapshot
}

// RequestOfferSnapshot propagates a snapshot offered by a peer so the
// application can inspect the metadata and the peer's app hash before accepting.
type RequestOfferSnapshot struct {
	Snapshot *Snapshot
	AppHash  []byte
}

// ResponseOfferSnapshot instructs Tendermint how to proceed with the offered
// snapshot (accept, reject, retry, etc.).
type ResponseOfferSnapshot struct {
	Result ResponseOfferSnapshot_Result
}

// ResponseOfferSnapshot_Result enumerates the possible decisions an
// application can take when a snapshot offer is received.
type ResponseOfferSnapshot_Result int32

const (
	ResponseOfferSnapshot_UNKNOWN ResponseOfferSnapshot_Result = iota
	ResponseOfferSnapshot_ACCEPT
	ResponseOfferSnapshot_ABORT
	ResponseOfferSnapshot_REJECT
	ResponseOfferSnapshot_REJECT_FORMAT
	ResponseOfferSnapshot_REJECT_SENDER
)

func (r ResponseOfferSnapshot_Result) String() string { _ = "STUB: not implemented"; return "" }

//nolint:goconst

//nolint:goconst

//nolint:goconst

// RequestLoadSnapshotChunk asks the application to load a specific chunk from
// the accepted snapshot so Tendermint can forward it to peers.
type RequestLoadSnapshotChunk struct {
	Height uint64
	Format uint32
	Chunk  uint32
}

// ResponseLoadSnapshotChunk carries the raw bytes for the requested snapshot chunk.
type ResponseLoadSnapshotChunk struct {
	Chunk []byte
}

// RequestApplySnapshotChunk delivers a snapshot chunk to the application so it
// can reconstruct state during state sync.
type RequestApplySnapshotChunk struct {
	Index  uint32
	Chunk  []byte
	Sender string
}

// ResponseApplySnapshotChunk lets the application signal whether it accepted
// the chunk or needs Tendermint to resend or reject certain senders/chunks.
type ResponseApplySnapshotChunk struct {
	Result        ResponseApplySnapshotChunk_Result
	RefetchChunks []uint32
	RejectSenders []string
}

// ResponseApplySnapshotChunk_Result captures the application-side outcome when
// applying a chunk (accept, abort, retry, etc.).
type ResponseApplySnapshotChunk_Result int32

const (
	ResponseApplySnapshotChunk_UNKNOWN ResponseApplySnapshotChunk_Result = iota
	ResponseApplySnapshotChunk_ACCEPT
	ResponseApplySnapshotChunk_ABORT
	ResponseApplySnapshotChunk_RETRY
	ResponseApplySnapshotChunk_RETRY_SNAPSHOT
	ResponseApplySnapshotChunk_REJECT_SNAPSHOT
)

func (r ResponseApplySnapshotChunk_Result) String() string { _ = "STUB: not implemented"; return "" }

//nolint:goconst

//nolint:goconst

//nolint:goconst

// ResponseProcessProposal communicates the application's decision after
// evaluating a proposed block before votes are cast in the ProcessProposal step.
type ResponseProcessProposal struct {
	Status                ResponseProcessProposal_ProposalStatus
	AppHash               []byte
	TxResults             []*ExecTxResult
	ValidatorUpdates      []ValidatorUpdate
	ConsensusParamUpdates *tmproto.ConsensusParams
}

// ResponseProcessProposal_ProposalStatus lists the possible verdicts when an
// application inspects a proposed block (accept, reject, or unknown).
type ResponseProcessProposal_ProposalStatus int32

const (
	ResponseProcessProposal_UNKNOWN ResponseProcessProposal_ProposalStatus = iota
	ResponseProcessProposal_ACCEPT
	ResponseProcessProposal_REJECT
)

func (s ResponseProcessProposal_ProposalStatus) String() string {
	_ = "STUB: not implemented"
	return ""
}

//nolint:goconst

//nolint:goconst

//nolint:goconst

// RequestProcessProposal bundles all of the proposed block data that the
// application can inspect to decide whether the block should move forward in
// consensus.
type RequestProcessProposal struct {
	Txs                 [][]byte
	ProposedLastCommit  CommitInfo
	ByzantineValidators []Misbehavior
	Hash                []byte

	Header *tmproto.Header
}

// RequestFinalizeBlock is emitted after a proposal is committed so the
// application can run FinalizeBlock logic over the same block data it agreed to.
type RequestFinalizeBlock struct {
	Txs                 [][]byte
	DecidedLastCommit   CommitInfo
	ByzantineValidators []Misbehavior
	Hash                []byte

	Header *tmproto.Header
}

type CommitInfo struct {
	Round int32
	Votes []VoteInfo
}

type LastCommitInfo struct {
	Round int32
	Votes []VoteInfo
}

type Validator struct {
	Address []byte
	Power   int64
}

type VoteInfo struct {
	Validator       Validator
	SignedLastBlock bool
}

type Misbehavior struct {
	Type             MisbehaviorType
	Validator        Validator
	Height           int64
	Time             time.Time
	TotalVotingPower int64
}

type Evidence struct {
	Type             MisbehaviorType
	Validator        Validator
	Height           int64
	Time             time.Time
	TotalVotingPower int64
}

type Snapshot struct {
	Height   uint64
	Format   uint32
	Chunks   uint32
	Hash     []byte
	Metadata []byte
}
