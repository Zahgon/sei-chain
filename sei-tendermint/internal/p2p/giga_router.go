package p2p

import (
	"context"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/consensus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/data"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/producer"
	atypes "github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/giga"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/rpc"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils/tcp"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type GigaNodeAddr struct {
	Key      NodePublicKey
	HostPort tcp.HostPort
	EVMRPC   utils.Option[*url.URL]
}

func (a GigaNodeAddr) String() string { _ = "STUB: not implemented"; return "" }

type GigaRouterConfig struct {
	DialInterval   time.Duration
	ValidatorAddrs map[atypes.PublicKey]GigaNodeAddr
	Consensus      *consensus.Config
	Producer       *producer.Config
	TxMempool      *mempool.TxMempool
	GenDoc         *types.GenesisDoc
}

type GigaRouter struct {
	cfg       *GigaRouterConfig
	key       NodeSecretKey
	data      *data.State
	producer  *producer.State
	consensus *consensus.State
	service   *giga.Service
	poolIn    *giga.Pool[NodePublicKey, rpc.Server[giga.API]]
	poolOut   *giga.Pool[NodePublicKey, rpc.Client[giga.API]]

	// lastCommitQCRecv is subscribed once at construction and reused for the
	// lifetime of the GigaRouter. Load() is lock-free (a single
	// atomic.Pointer.Load).
	//
	// Staleness-safety: the receiver points at the same atomicWatch held inside
	// avail.inner.latestCommitQC — a value field on a heap-allocated *inner
	// that is never replaced for the lifetime of the State, only Store()d
	// into. Every Load therefore observes the most recent Store. A
	// reconstructed avail.State (only on process restart) would also
	// reconstruct this GigaRouter, so the receiver can't outlive its watch.
	lastCommitQCRecv utils.AtomicRecv[utils.Option[*atypes.CommitQC]]
}

func NewGigaRouter(cfg *GigaRouterConfig, key NodeSecretKey) (*GigaRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint:gosec // verified to be positive.

// Automated pruning is disabled, because it is controlled by the application.
// The data WAL piggybacks on Consensus.PersistentStateDir: the two layers
// share the same on-disk root and write to distinct subdirectories under
// it (inner / blocks / commitqcs for consensus, globalblocks /
// fullcommitqcs for data).
//
// TODO(autobahn): once sei-db/ledger_db/block.BlockDB has a writer wired
// (see BlockByNumber's TODO), the data layer's WAL is redundant —
// BlockDB is the long-term home for the block read path and survives
// process restarts on its own. At that point this NewDataWAL call can
// drop the directory and become a no-op.

// Subscribe once here (takes avail's internal lock once); subsequent
// Load() calls from RPC handlers are lock-free atomic pointer reads.

// LastCommittedBlockNumber returns the highest global block number finalized
// by consensus (derived from the latest CommitQC). When no CommitQC has been
// recorded yet, atypes.GlobalRangeOpt returns the committee's empty default
// range {First: FirstBlock, Next: FirstBlock}, so this returns FirstBlock-1.
// Safe for high-frequency callers — uses a cached lock-free receiver; no
// locks taken on this path.
func (r *GigaRouter) LastCommittedBlockNumber() int64 {
	_ = "STUB: not implemented"
	// GlobalRange is a half-open [First, Next) interval; the highest
	// committed block number is Next-1.
	return 0
}

// nolint:gosec // gr.Next is uint64 but bounded by actual chain height.

// MaxGasPerBlock returns the producer's configured max gas per block (int64).
// Thin pass-through to producer.Config.MaxGasPerBlockI64 — the clamp logic
// lives there. Exposed at the GigaRouter level so the RPC layer can populate
// ResultBlockResults.ConsensusParamUpdates under Autobahn (where
// FinalizeBlock responses are not stored on disk) without reaching into
// the unexported router.cfg.
func (r *GigaRouter) MaxGasPerBlock() int64 { _ = "STUB: not implemented"; return 0 }

// BlockByNumber returns the finalized global block at height n translated
// into the CometBFT coretypes.ResultBlock shape. This lets consumers
// (notably evmrpc, which wraps receipts/logs with block context) keep
// working under Autobahn without CometBFT's BlockStore being populated.
//
// Fields populated when the underlying GlobalBlock is well-formed:
// BlockID.Hash (Autobahn lane-block header hash — the same bytes passed to
// app.FinalizeBlock's Hash param, which the EVM receipt store records as
// blockHash), Block.Header.ChainID/Height/Time, Block.Data.Txs. Other
// fields (AppHash, ProposerAddress, LastCommit, …) stay at zero values —
// evmrpc does not read them on the receipt path. If gb.Header is nil
// BlockID.Hash also stays empty; if gb.Payload is nil Block.Data.Txs
// stays empty (see the malformed-block handling below).
//
// TODO(autobahn): switch this to read from sei-db/ledger_db/block.BlockDB
// once a writer is wired (e.g. from app.FinalizeBlocker or executeBlock).
// Today no production code calls BlockDB.WriteBlock, so Autobahn's in-memory
// data.State is the only place a full block lives — but it's pruned per
// Sei's RetainHeight and exposes only a height index (no GetBlockByHash).
// BlockDB has the right shape (height + hash indexes, async pruning) and
// is the long-term home for this read path.
func (r *GigaRouter) BlockByNumber(ctx context.Context, n atypes.GlobalBlockNumber) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Map Autobahn's pruning sentinel to CometBFT's, so callers
// (env.Block, evmrpc, ops tooling) get the same error type they
// already handle on the CometBFT path. base is None because the
// active lower bound (data.State.inner.first) is internal to
// data.State; both call sites format through the same helper.

// BlockByHash returns the finalized global block keyed by Autobahn block-
// header hash, translated into the CometBFT coretypes.ResultBlock shape
// (same translation as BlockByNumber). Matches CometBFT semantics for
// unknown hashes: returns &ResultBlock{Block: nil} with no error.
//
// Lookup-and-construct happens under a single data.State lock acquire, so
// the returned block matches the requested hash atomically. Hashes below
// the pruning watermark are not indexed and read as "unknown". Wrong-size
// inputs are rejected at the call site (env.BlockByHash) so this method
// can stay strongly typed on atypes.BlockHeaderHash.
//
// TODO(autobahn): replace this with a direct read from
// sei-db/ledger_db/block.BlockDB.GetBlockByHash once a writer is wired into
// block execution. The data.State-side index can also go away at that point.
func (r *GigaRouter) BlockByHash(ctx context.Context, hash atypes.BlockHeaderHash) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reject the unknown-hash case here so translateGlobalBlock can rely
// on the *GlobalBlock type contract (non-nil, with non-nil Header
// and Payload) — same way executeBlock dereferences b.Header
// without checking. Mirrors CometBFT's BlockStore.LoadBlockByHash
// returning &ResultBlock{Block: nil} for an unknown hash.

// translateGlobalBlock converts an Autobahn GlobalBlock to the CometBFT
// coretypes.ResultBlock shape used by env.Block / env.BlockByHash and
// downstream evmrpc consumers. Caller must pass a non-nil *GlobalBlock with
// non-nil Header and Payload — that's the contract data.State guarantees on
// a successful lookup, and matches how executeBlock dereferences b.Header
// without a nil-check on the same type. The "no such block" case is
// rejected at the BlockByHash call site before delegating here.
//
// LastCommit is non-nil with empty Signatures, mirroring executeBlock's
// FinalizeBlock call which passes an empty abci.CommitInfo. Under Autobahn
// the committee is fixed by genesis (no validator-set updates), so the
// application is not in control of jailing — surfacing N "absent sig"
// entries here would make trace replay's BeginBlock bump missed-block
// counters and diverge from production. ToReqBeginBlock skips the per-
// validator loop when Signatures is empty, so empty Votes flow into
// distribution/slashing on both paths.
func (r *GigaRouter) translateGlobalBlock(gb *atypes.GlobalBlock) *coretypes.ResultBlock {
	_ = "STUB: not implemented"
	return nil
}

// Clamp accepts any constraints.Integer for From, so
// gb.GlobalNumber (a typed uint64) goes in directly — no
// intermediate uint64() conversion needed.

func (r *GigaRouter) executeBlock(ctx context.Context, b *atypes.GlobalBlock) (*abci.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deterministically select a proposer from the app's validator committee.
// We need it so that app does not emit error logs.

// TODO: add metrics to understand execution latency.

// Empty DecidedLastCommit does not indicate missing votes.

// WARNING: this is a hash of the autobahn block header.
// It is used to identify block processed optimistically
// and is fed as block hash to EVM contracts.

// nolint:gosec // different representations of the same value

// WARNING: the reward distribution has corner cases where it forgets the proposer,
// because reward is distributed with a delay. This is not our problem here though.

// nolint:gosec // autobahn block numbers fit in int64.

// TODO: We need the constraints to be fixed per epoch, because we don't know where the lane blocks will be sequenced.
// Therefore we disable constraints for now, until epochs are supported AND
// chain state understands that consensus parameters can change only at the epoch boundary.

// recheck=false; see TxMempool.Update doc for why.

func (r *GigaRouter) runExecute(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Fresh start: the CometBFT handshaker is skipped in giga mode
// (see node.go: shouldHandshake = !stateSync && !gigaEnabled), so
// nobody has called InitChain yet. Call it here ourselves; this sets
// up the app's deliverState (matching real SDK: InitChain leaves
// deliverState populated with no intermediate Commit, so the first
// FinalizeBlock below runs against it).
//
// On restart (last > 0, below), InitChain must NOT be called again;
// the app's committed CMS already holds the latest state, and
// BaseApp.FinalizeBlock rebuilds deliverState from it via its
// nil-check fallback.
//
// Note: if a process crashed after InitChain but before the first
// Commit, LastBlockHeight is still 0 and we enter this branch again
// on restart. Re-calling InitChain is safe in that case because
// nothing was committed — it behaves as a fresh init.

// NOTE that with the current implementation losing prefix of appHashes on crash is fine:
// if everyone votes on apphashes of a suffix of finalized blocks, then AppQC will be reached.

func (r *GigaRouter) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Spawn outbound connections dialing.

func (r *GigaRouter) dialAndRunConn(ctx context.Context, key NodePublicKey, hp tcp.HostPort) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: handshake needs a timeout.

func (r *GigaRouter) RunInboundConn(ctx context.Context, hConn *handshakedConn) error {
	_ = "STUB: not implemented"
	return nil
}

// Filter unwanded connections.

func (r *GigaRouter) EvmProxy(sender common.Address) (*url.URL, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
