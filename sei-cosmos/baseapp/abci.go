package baseapp

import (
	"context"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// InitChain implements the ABCI interface. It runs the initialization logic
// directly on the CommitMultiStore.
func (app *BaseApp) InitChain(ctx context.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	// On a new chain, we consider the init chain block height as 0, even though
	// req.InitialHeight is 1 by default.
	return nil, nil
}

// If req.InitialHeight is > 1, then we set the initial version in the
// stores.

// initialize the deliver state and check state with a correct header

// Store the consensus params in the BaseApp's paramstore. Note, this must be
// done after the deliver state and context have been set as it's persisted
// to state.

// In the case of a new chain, AppHash will be the hash of an empty string.
// During an upgrade, it'll be the hash of the last committed block.

// $ echo -n '' | sha256sum
// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

// NOTE: We don't commit, but BeginBlock for block `initial_height` starts from this
// deliverState.

// Info implements the ABCI interface.
func (app *BaseApp) Info(ctx context.Context, req *abci.RequestInfo) (*abci.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *BaseApp) MidBlock(ctx sdk.Context, height int64) (events []abci.Event) {
	_ = "STUB: not implemented"
	return nil
}

// EndBlock implements the ABCI interface.
func (app *BaseApp) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) (res abci.ResponseEndBlock) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseEndBlock)
}

// CheckTx implements the ABCI interface and executes a tx in CheckTx mode. In
// CheckTx mode, messages are not executed. This means messages are only validated
// and only the AnteHandler is executed. State is persisted to the BaseApp's
// internal CheckTx state if the AnteHandler passes. Otherwise, the ResponseCheckTx
// will contain releveant error information. Regardless of tx execution outcome,
// the ResponseCheckTx will contain relevant gas execution context.
func (app *BaseApp) CheckTx(ctx context.Context, req *abci.RequestCheckTxV2) *abci.ResponseCheckTxV2 {
	_ = "STUB: not implemented"
	return nil
}

// DeliverTxBatch executes multiple txs
func (app *BaseApp) DeliverTxBatch(ctx sdk.Context, req sdk.DeliverTxBatchRequest) (res sdk.DeliverTxBatchResponse) {
	_ = "STUB: not implemented"
	return *new(sdk.DeliverTxBatchResponse)
}

// avoid overhead for empty batches

// DeliverTx implements the ABCI interface and executes a tx in DeliverTx mode.
// State only gets persisted if all messages are valid and get executed successfully.
// Otherwise, the ResponseDeliverTx will contain relevant error information.
// Regardless of tx execution outcome, the ResponseDeliverTx will contain relevant
// gas execution context.
func (app *BaseApp) DeliverTx(ctx sdk.Context, req abci.RequestDeliverTxV2, tx sdk.Tx, checksum [32]byte) (res abci.ResponseDeliverTx) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseDeliverTx)
}

// if we have a result, use those events instead of just the anteEvents

//nolint:gosec // gas values are practically bounded; TODO: Should type accept unsigned ints?
//nolint:gosec // gas values are practically bounded; TODO: Should type accept unsigned ints?

// TODO: populate error data for EVM err

func (app *BaseApp) WriteState() sdk.CommitMultiStore {
	_ = "STUB: not implemented"
	return *new(sdk.CommitMultiStore)
}

func (app *BaseApp) GetWorkingHash() []byte { _ = "STUB: not implemented"; return nil }

// this should never happen

func (app *BaseApp) SetProcessProposalStateToCommit() { _ = "STUB: not implemented"; return }

func (app *BaseApp) SetDeliverStateToCommit() { _ = "STUB: not implemented"; return }

// Commit implements the ABCI interface. It will commit all state that exists in
// the deliver state's multi-store and includes the resulting commit ID in the
// returned abci.ResponseCommit. Commit will set the check state based on the
// latest header and reset the deliver state. Also, if a non-zero halt height is
// defined in config, Commit will execute a deferred function call to check
// against that height and gracefully halt if it matches the latest committed
// height.
func (app *BaseApp) Commit(ctx context.Context) (res *abci.ResponseCommit, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset the Check state to the latest committed.
//
// NOTE: This is safe because Tendermint holds a lock on the mempool for
// Commit. Use the header from this latest block.

// empty/reset the deliver state

//nolint:gosec // block heights are always non-negative

//nolint:gosec // haltTime is a small config value, won't overflow int64

// Halt the binary and allow Tendermint to receive the ResponseCommit
// response with the commit ID hash. This will allow the node to successfully
// restart and process blocks assuming the halt configuration has been
// reset or moved to a more distant value.

//nolint:gosec // bounds checked above

func (app *BaseApp) SnapshotIfApplicable(height uint64) { _ = "STUB: not implemented"; return }

//nolint:gosec // bounds checked above

// halt attempts to gracefully shutdown the node via SIGINT and SIGTERM falling
// back on os.Exit if both fail.
func (app *BaseApp) halt() { _ = "STUB: not implemented"; return }

// attempt cascading signals in case SIGINT fails (os dependent)

// Resort to exiting immediately if the process could not be found or killed
// via SIGINT/SIGTERM signals.

// Snapshot takes a snapshot of the current state and prunes any old snapshottypes.
func (app *BaseApp) Snapshot(height int64) { _ = "STUB: not implemented"; return }

//nolint:gosec // bounds checked above

// Query implements the ABCI interface. It delegates to CommitMultiStore if it
// implements Queryable.
func (app *BaseApp) Query(ctx context.Context, req *abci.RequestQuery) (res *abci.ResponseQuery, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add panic recovery for all queries.
// ref: https://github.com/cosmos/cosmos-sdk/pull/8039

// when a client did not provide a query height, manually inject the latest

// handle gRPC routes first rather than calling splitPath because '/' characters
// are used as part of gRPC paths

// "/app" prefix for special application queries

func (app *BaseApp) GetValidators() []abci.ValidatorUpdate {
	_ = "STUB: not implemented"

	// ListSnapshots implements the ABCI interface. It delegates to app.snapshotManager if set.
	return nil
}

func (app *BaseApp) ListSnapshots(context context.Context, req *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadSnapshotChunk implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) LoadSnapshotChunk(context context.Context, req *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OfferSnapshot implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) OfferSnapshot(context context.Context, req *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We currently don't support resetting the IAVL stores and retrying a different snapshot,
// so we ask Tendermint to abort all snapshot restoration.

// ApplySnapshotChunk implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) ApplySnapshotChunk(context context.Context, req *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *BaseApp) handleQueryGRPC(handler GRPCQueryHandler, req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

func gRPCErrorToSDKError(err error) error { _ = "STUB: not implemented"; return nil }

func checkNegativeHeight(height int64) error {
	_ = "STUB: not implemented"

	// Reject invalid heights.
	return nil
}

// CreateQueryContext creates a new sdk.Context for a query, taking as args
// the block height and whether the query needs a proof or not.
func (app *BaseApp) CreateQueryContext(height int64, prove bool) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// when a client did not provide a query height, manually inject the latest

// branch the commit-multistore for safety

// GetBlockRetentionHeight returns the height for which all blocks below this height
// are pruned from Tendermint. Given a commitment height and a non-zero local
// minRetainBlocks configuration, the retentionHeight is the smallest height that
// satisfies:
//
// - Unbonding (safety threshold) time: The block interval in which validators
// can be economically punished for misbehavior. Blocks in this interval must be
// auditable e.g. by the light client.
//
// - Logical store snapshot interval: The block interval at which the underlying
// logical store database is persisted to disk, e.g. every 10000 heights. Blocks
// since the last IAVL snapshot must be available for replay on application restart.
//
// - State sync snapshots: Blocks since the oldest available snapshot must be
// available for state sync nodes to catch up (oldest because a node may be
// restoring an old snapshot while a new snapshot was taken).
//
// - Local (minRetainBlocks) config: Archive nodes may want to retain more or
// all blocks, e.g. via a local config option min-retain-blocks. There may also
// be a need to vary retention for other nodes, e.g. sentry nodes which do not
// need historical blocks.
func (app *BaseApp) GetBlockRetentionHeight(commitHeight int64) (int64, error) {
	_ = "STUB: not implemented"
	// pruning is disabled if minRetainBlocks is zero
	return 0, nil
}

// Define retentionHeight as the minimum value that satisfies all non-zero
// constraints. All blocks below (commitHeight-retentionHeight) are pruned
// from Tendermint.

// Define the number of blocks needed to protect against misbehaving validators
// which allows light clients to operate safely. Note, we piggy back of the
// evidence parameters instead of computing an estimated nubmer of blocks based
// on the unbonding period and block commitment time as the two should be
// equivalent.

// Define the state pruning offset, i.e. the block offset at which the
// underlying logical database is persisted to disk.

//nolint:gosec // bounds checked above

// Hitting this case means we have persisting enabled but have yet to reach
// a height in which we persist state, so we return zero regardless of other
// conditions. Otherwise, we could end up pruning blocks without having
// any state committed to disk.

//nolint:gosec // snapshotKeepRecent is a small config value

//nolint:gosec // bounds checked above

//nolint:gosec // bounds checked above

// prune nothing in the case of a non-positive height

func (app *BaseApp) Simulate(txBytes []byte) (sdk.GasInfo, *sdk.Result, error) {
	_ = "STUB: not implemented"
	return *new(sdk.GasInfo), nil, nil
}

func handleQueryApp(app *BaseApp, path []string, req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

func handleQueryStore(app *BaseApp, path []string, req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// Check if online migration is enabled for fallback read

// "/store" prefix for store queries

func handleQueryCustom(app *BaseApp, path []string, req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	// path[0] should be "custom" because "/custom" prefix is required for keeper
	// queries.
	//
	// The QueryRouter routes using path[1]. For example, in the path
	// "custom/gov/proposal", QueryRouter routes using "gov".
	return *new(abci.ResponseQuery)
}

// Passes the rest of the path as an argument to the querier.
//
// For example, in the path "custom/gov/proposal/test", the gov querier gets
// []string{"proposal", "test"} as the path.

// splitPath splits a string path using the delimiter '/'.
//
// e.g. "this/is/funny" becomes []string{"this", "is", "funny"}
func splitPath(requestPath string) (path []string) { _ = "STUB: not implemented"; return nil }

// first element is empty string

// ABCI++
func (app *BaseApp) ProcessProposal(ctx context.Context, req *abci.RequestProcessProposal) (resp *abci.ResponseProcessProposal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In the first block, app.processProposalState.ctx will already be initialized
// by InitChain. Context is now updated with Header information.

// NOTE: header hash is not set in NewContext, so we manually set it here

// Snapshot a clean context for read-only validation (e.g. gas checks).
// Branch from the source store (cms or deliverState) rather than from
// processProposalState, so that speculative writes from the optimistic
// goroutine are not visible.

// Block 1: deliverState has InitChain genesis writes not yet committed

// Blocks 2+: committed root store has everything

func (app *BaseApp) FinalizeBlock(ctx context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In the first block, app.deliverState.ctx will already be initialized
// by InitChain. Context is now updated with Header information.

// NOTE: header hash is not set in NewContext, so we manually set it here

// we also set block gas meter to checkState in case the application needs to
// verify gas consumption during (Re)CheckTx

func (app *BaseApp) GetTxPriorityHint(_ context.Context, req *abci.RequestGetTxPriorityHintV2) (_resp *abci.ResponseGetTxPriorityHint, _err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fall back to no-op priority if we panic for any reason. This is to avoid DoS
// vectors where a malicious actor crafts a transaction that panics the
// prioritizer. Since the prioritizer is used as a hint only, it's safe to fall
// back to zero priority in this case and log the panic for monitoring purposes.

// Do not overwrite an existing error if one was already set to keep panics a
// non-event at this stage but safeguard against them.

// TODO: should we bother validating the messages here?
