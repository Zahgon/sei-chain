package state

import (
	"context"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
	otrace "go.opentelemetry.io/otel/trace"
)

var logger = seilog.NewLogger("tendermint", "internal", "state")

// proposerPriorityHashInterval is how often (in heights) the
// ProposerPriorityHash metric is exported. Used so operators can compare
// hashes across validators and detect ProposerPriority divergence.
const proposerPriorityHashInterval = 1024

//-----------------------------------------------------------------------------
// BlockExecutor handles block execution and state updates.
// It exposes ApplyBlock(), which validates & executes the block, updates state w/ ABCI responses,
// then commits and updates the mempool atomically, then saves state.

// BlockExecutor provides the context and accessories for properly executing a block.
type BlockExecutor struct {
	// save state, validators, consensus params, abci responses here
	store Store

	// use blockstore for the pruning functions.
	blockStore BlockStore

	// execute the app against this
	app *proxy.Proxy

	// events
	eventBus types.BlockEventPublisher

	// manage the mempool lock during commit
	// and update both with block results after commit.
	mempool *mempool.TxMempool
	evpool  EvidencePool

	metrics *Metrics

	// consensusPolicy is a compile-time validation bypass that only takes
	// effect in mock_block_validation builds; production binaries always see
	// the zero-value (no bypass). Distinct from types.SkipLastResultsHashValidation
	// below, which is a runtime atomic.Bool flipped on for the Giga executor.
	consensusPolicy types.ConsensusPolicy

	// cache the verification results over a single height
	cache map[string]struct{}
}

// NewBlockExecutor returns a new BlockExecutor with the passed-in EventBus.
func NewBlockExecutor(
	stateStore Store,
	app *proxy.Proxy,
	pool *mempool.TxMempool,
	evpool EvidencePool,
	blockStore BlockStore,
	eventBus *eventbus.EventBus,
	metrics *Metrics,
	consensusPolicy types.ConsensusPolicy,
) *BlockExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (blockExec *BlockExecutor) Store() Store {
	_ = "STUB: not implemented"
	return *

	// CreateProposalBlock calls state.MakeBlock with evidence from the evpool
	// and txs from the mempool. The max bytes must be big enough to fit the commit.
	// Up to 1/10th of the block space is allcoated for maximum sized evidence.
	// The rest is given to txs, up to the max gas.
	//
	// Contract: application will not return more bytes than are sent over the wire.
	new(Store)
}

func (blockExec *BlockExecutor) CreateProposalBlock(
	ctx context.Context,
	height int64,
	state State,
	lastCommit *types.Commit,
	proposerAddr []byte,
) (block *types.Block, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert panic to error

// Fetch a limited amount of valid txs

func (blockExec *BlockExecutor) GetTxsForHashes(txHashes []types.TxHash) types.Txs {
	_ = "STUB: not implemented"
	return *new(types.Txs)
}

func (blockExec *BlockExecutor) ProcessProposal(
	ctx context.Context,
	block *types.Block,
	state State,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ValidateBlock validates the given block against the given state.
// If the block is invalid, it returns an error.
// Validation does not mutate state, but does require historical information from the stateDB,
// ie. to verify evidence from a validator at an old height.
func (blockExec *BlockExecutor) ValidateBlock(ctx context.Context, state State, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this is a LastResultsHash mismatch and log detailed info

// ApplyBlock validates the block against the state, executes it against the app,
// fires the relevant events, commits the app, and saves the new state and responses.
// It returns the new state.
// It's the only function that needs to be called
// from outside this package to process and commit an entire block.
// It takes a blockID to avoid recomputing the parts hash.
func (blockExec *BlockExecutor) ApplyBlock(ctx context.Context, state State, blockID types.BlockID, block *types.Block, tracer otrace.Tracer) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// validate the block if we haven't already

// Save the results before we commit.

// It is correct to have an empty ResponseFinalizeBlock for ApplyBlock,
// but not for saving it to the state store

// validate the validator updates and convert to tendermint types

// Update the state with the block and responses.

// Log LastResultsHash computation details for debugging consensus issues

// Log per-tx deterministic fields (Code, Data, GasWanted, GasUsed) for debugging

// Export ProposerPriorityHash every proposerPriorityHashInterval heights so
// operators can detect ProposerPriority divergence between validators by
// comparing gauge values across nodes at the same height.
//
// Why emit the hash as a numeric *value* rather than a label?
// A label-based design (gauge with hash as label) would create a new
// Prometheus time series every time the hash changes — since validator
// priorities change every block, each emission would yield a brand-new
// series. Over time this accumulates unbounded cardinality in the
// metrics backend. Exporting as a numeric value keeps cardinality
// constant at one series per node.
//
// Why take only the first 8 bytes?
// Prometheus gauges are float64, which only represents integers up to
// 2^53 exactly. We take the first 8 bytes of the SHA-256 hash and cast
// to float64; the top 11 bits are lost to the mantissa, effectively
// giving us 53 bits of entropy. Collision probability across 40
// validators is ~40^2/2^54 ≈ 9e-14, effectively zero.
//
// Paired with ProposerPriorityHashHeight so operators know which height
// the hash corresponds to. A log line also emits the full 32-byte hash
// for grep-based debugging.
//
// Note on restart staleness: Prometheus Gauges live in memory. After a
// process restart the gauges reset to zero until the next emission at
// the following multiple of proposerPriorityHashInterval — up to ~8.5
// min of stale/zero data at Sei's block times. Acceptable for a
// monitoring signal that is only checked in response to incidents.

// Log both the full 32-byte hash (for unambiguous comparison)
// and the packed value (to correlate with the Prometheus gauge).

// Lock mempool, commit app state, update mempoool.

// Update evpool with the latest state.

// Update the app hash and save the state.

// Prune old heights, if requested by ABCI app.

// reset the verification cache

// Events are fired after everything else.
// NOTE: if we crash between Commit and Save, events wont be fired during replay

// Commit locks the mempool, runs the ABCI Commit message, and updates the
// mempool.
// It returns the result of calling abci.Commit (the AppHash) and the height to retain (if any).
// The Mempool must be locked during commit and update because state is
// typically reset on Commit and old txs must be replayed against committed
// state before new txs are run in the mempool, lest they be invalid.
func (blockExec *BlockExecutor) Commit(
	ctx context.Context,
	state State,
	block *types.Block,
	txResults []*abci.ExecTxResult,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Commit block, get hash back

// ResponseCommit has no error code - just data

// Update mempool.

func (blockExec *BlockExecutor) GetMissingTxs(txHashes []types.TxHash) []types.TxHash {
	_ = "STUB: not implemented"
	return nil
}

func (blockExec *BlockExecutor) SafeGetTxsByHashes(txHashes []types.TxHash) (types.Txs, []types.TxHash) {
	_ = "STUB: not implemented"
	return *new(types.Txs), nil
}

func buildLastCommitInfo(block *types.Block, store Store, initialHeight int64) abci.CommitInfo {
	_ = "STUB: not implemented"
	return *new(abci.CommitInfo)
}

// there is no last commit for the initial height.
// return an empty value.

// ensure that the size of the validator set in the last commit matches
// the size of the validator set in the state store.

func validateValidatorUpdates(abciUpdates []abci.ValidatorUpdate, params types.ValidatorParams) error {
	_ = "STUB: not implemented"
	return nil
}

// continue, since this is deleting the validator, and thus there is no
// pubkey to check

// Check if validator's pubkey matches an ABCI type in the consensus params

// Update returns a copy of state with the fields set using the arguments passed in.
func (state State) Update(
	blockID types.BlockID,
	header *types.Header,
	resultsHash []byte,
	consensusParamUpdates *tmtypes.ConsensusParams,
	validatorUpdates []*types.Validator,
) (State, error) {
	_ = "STUB: not implemented"

	// Copy the valset so we can apply changes from FinalizeBlock
	// and update s.LastValidators and s.Validators.
	return *new(State), nil
}

// Update the validator set with the latest responses to FinalizeBlock.

// Change results from this height but only applies to the next next height.

// Update validator proposer priority and set state variables.

// Update the params with the latest responses to FinalizeBlock.

// NOTE: must not mutate state.ConsensusParams

// Change results from this height but only applies to the next height.

// NOTE: the AppHash has not been populated.
// It will be filled on state.Save.

// Fire NewBlock, NewBlockHeader.
// Fire TxEvent for every tx.
// NOTE: if Tendermint crashes before commit, some or all of these events may be published again.
func FireEvents(
	eventBus types.BlockEventPublisher,
	block *types.Block,
	blockID types.BlockID,
	finalizeBlockResponse *abci.ResponseFinalizeBlock,
	validatorUpdates []*types.Validator,
) {
	_ = "STUB: not implemented"
	return
}

// sanity check

//nolint:gosec // i is bounded by block.Txs length which fits in uint32

//----------------------------------------------------------------------------------------------------
// Execute block without state. TODO: eliminate

// ExecCommitBlock executes and commits a block on the proxyApp without validating or mutating the state.
// It returns the application root hash (result of abci.Commit).
func ExecCommitBlock(
	ctx context.Context,
	be *BlockExecutor,
	appConn *proxy.Proxy,
	block *types.Block,
	store Store,
	initialHeight int64,
	s State,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the BlockExecutor condition is using for the final block replay process.

// Commit block

// ResponseCommit has no error or log

func (blockExec *BlockExecutor) pruneBlocks(retainHeight int64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
