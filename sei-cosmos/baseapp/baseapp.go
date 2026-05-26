package baseapp

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/snapshots"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/utils/tracing"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/seilog"
	dbm "github.com/tendermint/tm-db"
)

const (
	runTxModeCheck    runTxMode = iota // Check a transaction
	runTxModeReCheck                   // Recheck a (pending) transaction after a commit
	runTxModeSimulate                  // Simulate a transaction
	runTxModeDeliver                   // Deliver a transaction
)

var modeKeyToString = map[runTxMode]string{
	runTxModeCheck:    "check",
	runTxModeReCheck:  "recheck",
	runTxModeSimulate: "simulate",
	runTxModeDeliver:  "deliver",
}

const (
	// archival related flags
	FlagArchivalVersion                = "archival-version"
	FlagArchivalDBType                 = "archival-db-type"
	FlagArchivalArweaveIndexDBFullPath = "archival-arweave-index-db-full-path"
	FlagArchivalArweaveNodeURL         = "archival-arweave-node-url"

	FlagChainID            = "chain-id"
	FlagConcurrencyWorkers = "concurrency-workers"
	FlagOccEnabled         = "occ-enabled"
)

var (
	logger = seilog.NewLogger("cosmos", "baseapp")

	_ abci.Application = (*BaseApp)(nil)
)

type (
	// Enum mode for app.runTx
	runTxMode uint8

	// StoreLoader defines a customizable function to control how we load the CommitMultiStore
	// from disk. This is useful for state migration, when loading a datastore written with
	// an older version of the software. In particular, if a module changed the substore key name
	// (or removed a substore) between two versions of the software.
	StoreLoader func(ms sdk.CommitMultiStore) error

	DeliverTxHook func(sdk.Context, sdk.Tx, [32]byte, sdk.DeliverTxHookInput)
)

func (app *BaseApp) EvmNonce(_ common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

func (app *BaseApp) EvmBalance(_ common.Address, _ []byte) *big.Int {
	_ = "STUB: not implemented"
	return nil

	// BaseApp reflects the ABCI application implementation.
}

type BaseApp struct {
	// initialized on creation
	name              string // application name from abci.Info
	interfaceRegistry types.InterfaceRegistry
	txDecoder         sdk.TxDecoder // unmarshal []byte into sdk.Tx

	processProposalHandler    sdk.ProcessProposalHandler
	finalizeBlocker           sdk.FinalizeBlocker
	anteHandler               sdk.AnteHandler // ante handler for fee and auth
	loadVersionHandler        sdk.LoadVersionHandler
	preCommitHandler          sdk.PreCommitHandler
	closeHandler              sdk.CloseHandler
	inplaceTestnetInitializer sdk.InplaceTestnetInitializer
	txPrioritizer             sdk.TxPrioritizer

	appStore
	baseappVersions
	snapshotData
	abciData
	moduleRouter

	// volatile states:
	//
	// checkState is set on InitChain and reset on Commit
	// deliverState is set on InitChain and BeginBlock and set to nil on Commit
	checkState              *state // for CheckTx
	deliverState            *state // for DeliverTx
	processProposalState    *state
	processProposalCleanCtx sdk.Context // snapshot before optimistic processing
	stateToCommit           *state

	// paramStore is used to query for ABCI consensus parameters from an
	// application parameter store.
	paramStore ParamStore

	// The minimum gas prices a validator is willing to accept for processing a
	// transaction. This is mainly used for DoS and spam prevention.
	minGasPrices sdk.DecCoins

	// initialHeight is the initial height at which we start the baseapp
	initialHeight int64

	// flag for sealing options and parameters to a BaseApp
	sealed bool

	// block height at which to halt the chain and gracefully shutdown
	haltHeight uint64

	// minimum block time (in Unix seconds) at which to halt the chain and gracefully shutdown
	haltTime uint64

	// minRetainBlocks defines the minimum block height offset from the current
	// block being committed, such that all blocks past this offset are pruned
	// from Tendermint. It is used as part of the process of determining the
	// ResponseCommit.RetainHeight value during ABCI Commit. A value of 0 indicates
	// that no blocks should be pruned.
	//
	// Note: Tendermint block pruning is dependant on this parameter in conunction
	// with the unbonding (safety threshold) period, state pruning and state sync
	// snapshot parameters to determine the correct minimum value of
	// ResponseCommit.RetainHeight.
	minRetainBlocks uint64

	// recovery handler for app.runTx method
	runTxRecoveryMiddleware recoveryMiddleware

	// trace set will return full stack traces for errors in ABCI Log field
	trace bool

	// indexEvents defines the set of events in the form {eventType}.{attributeKey},
	// which informs Tendermint what to index. If empty, all events will be indexed.
	IndexEvents map[string]struct{}

	ChainID string

	commitLock       *sync.Mutex
	checkTxStateLock *sync.RWMutex

	compactionInterval uint64

	TmConfig *tmcfg.Config

	TracingInfo *tracing.Info

	concurrencyWorkers int
	occEnabled         bool

	deliverTxHooks []DeliverTxHook

	execProcessProposalMs int64
	execFinalizeBlockMs   int64
	execBlockTxCount      int
}

type appStore struct {
	db              dbm.DB               // common DB backend
	cms             sdk.CommitMultiStore // Main (uncached) state
	qms             sdk.CommitMultiStore // Query multistore used for migration only
	migrationHeight int64
	storeLoader     StoreLoader // function to handle store loading, may be overridden with SetStoreLoader()

	// an inter-block write-through cache provided to the context during deliverState
	interBlockCache sdk.MultiStorePersistentCache

	fauxMerkleMode bool // if true, IAVL MountStores uses MountStoresDB for simulation speed.
}

type moduleRouter struct {
	router           sdk.Router        // handle any kind of message
	queryRouter      sdk.QueryRouter   // router for redirecting query calls
	grpcQueryRouter  *GRPCQueryRouter  // router for redirecting gRPC query calls
	msgServiceRouter *MsgServiceRouter // router for redirecting Msg service messages
}

type abciData struct {
	initChainer sdk.InitChainer // initialize state with validators and state blob
	midBlocker  sdk.MidBlocker  // logic to run after all txs, and to determine valset changes
	endBlocker  sdk.EndBlocker  // logic to run after all txs, and to determine valset changes
}

type baseappVersions struct {
	// application's version string
	version string

	// application's protocol version that increments on every upgrade
	// if BaseApp is passed to the upgrade keeper's NewKeeper method.
	appVersion uint64
}

// should really get handled in some db struct
// which then has a sub-item, persistence fields
type snapshotData struct {
	// manages snapshots, i.e. dumps of app state at certain intervals
	snapshotManager    *snapshots.Manager
	snapshotInterval   uint64 // block interval between state sync snapshots
	snapshotKeepRecent uint32 // recent state sync snapshots to keep
	snapshotDirectory  string //  state sync snapshots directory
}

// NewBaseApp returns a reference to an initialized BaseApp. It accepts a
// variadic number of option functions, which act on the BaseApp to set
// configuration choices.
//
// NOTE: The db is used to store the version number for now.
func NewBaseApp(
	name string, db dbm.DB, txDecoder sdk.TxDecoder, tmConfig *tmcfg.Config, appOpts servertypes.AppOptions, options ...func(*BaseApp),
) *BaseApp {
	_ = "STUB: not implemented"
	return nil
}

// Enable Tracing

// if no option overrode already, initialize to the flags value
// this avoids forcing every implementation to pass an option, but allows it

// safely default this to the default value if 0

// Name returns the name of the BaseApp.
func (app *BaseApp) Name() string {
	_ = "STUB: not implemented"

	// AppVersion returns the application's protocol version.
	return ""
}

func (app *BaseApp) AppVersion() uint64 { _ = "STUB: not implemented"; return 0 }

// ConcurrencyWorkers returns the number of concurrent workers for the BaseApp.
func (app *BaseApp) ConcurrencyWorkers() int { _ = "STUB: not implemented"; return 0 }

// OccEnabled returns whether OCC is enabled for the BaseApp.
func (app *BaseApp) OccEnabled() bool { _ = "STUB: not implemented"; return false }

// Version returns the application's version string.
func (app *BaseApp) Version() string {
	_ = "STUB: not implemented"

	// Trace returns the boolean value for logging error stack traces.
	return ""
}

func (app *BaseApp) Trace() bool {
	_ = "STUB: not implemented"

	// MsgServiceRouter returns the MsgServiceRouter of a BaseApp.
	return false
}

func (app *BaseApp) MsgServiceRouter() *MsgServiceRouter { _ = "STUB: not implemented"; return nil }

// MountStores mounts all IAVL or DB stores to the provided keys in the BaseApp
// multistore.
func (app *BaseApp) MountStores(keys ...sdk.StoreKey) { _ = "STUB: not implemented"; return }

// StoreTypeDB doesn't do anything upon commit, and it doesn't
// retain history, but it's useful for faster simulation.

// MountKVStores mounts all IAVL or DB stores to the provided keys in the
// BaseApp multistore.
func (app *BaseApp) MountKVStores(keys map[string]*sdk.KVStoreKey) {
	_ = "STUB: not implemented"
	return
}

// StoreTypeDB doesn't do anything upon commit, and it doesn't
// retain history, but it's useful for faster simulation.

// MountTransientStores mounts all transient stores to the provided keys in
// the BaseApp multistore.
func (app *BaseApp) MountTransientStores(keys map[string]*sdk.TransientStoreKey) {
	_ = "STUB: not implemented"
	return
}

// MountMemoryStores mounts all in-memory KVStores with the BaseApp's internal
// commit multi-store.
func (app *BaseApp) MountMemoryStores(keys map[string]*sdk.MemoryStoreKey) {
	_ = "STUB: not implemented"
	return
}

// MountStore mounts a store to the provided key in the BaseApp multistore,
// using the default DB.
func (app *BaseApp) MountStore(key sdk.StoreKey, typ sdk.StoreType) {
	_ = "STUB: not implemented"
	return
}

// LoadLatestVersion loads the latest application version. It will panic if
// called more than once on a running BaseApp.
func (app *BaseApp) LoadLatestVersion() error { _ = "STUB: not implemented"; return nil }

// DefaultStoreLoader will be used by default and loads the latest version
func DefaultStoreLoader(ms sdk.CommitMultiStore) error { _ = "STUB: not implemented"; return nil }

// CommitMultiStore returns the root multi-store.
// App constructor can use this to access the `cms`.
// UNSAFE: must not be used during the abci life cycle.
func (app *BaseApp) CommitMultiStore() sdk.CommitMultiStore {
	_ = "STUB: not implemented"

	// SnapshotManager returns the snapshot manager.
	// application use this to register extra extension snapshotters.
	return *new(sdk.CommitMultiStore)
}

func (app *BaseApp) SnapshotManager() *snapshots.Manager { _ = "STUB: not implemented"; return nil }

// LoadVersion loads the BaseApp application version. It will panic if called
// more than once on a running baseapp.
func (app *BaseApp) LoadVersion(version int64) error { _ = "STUB: not implemented"; return nil }

// LoadVersionWithoutInit loads the BaseApp application version, it doesn't call app.init any more,
// specifically used by export genesis command.
func (app *BaseApp) LoadVersionWithoutInit(version int64) error {
	_ = "STUB: not implemented"
	return nil
}

// LastCommitID returns the last CommitID of the multistore.
func (app *BaseApp) LastCommitID() sdk.CommitID {
	_ = "STUB: not implemented"
	return *new(sdk.CommitID)
}

// LastBlockHeight returns the last committed block height.
func (app *BaseApp) LastBlockHeight() int64 { _ = "STUB: not implemented"; return 0 }

func (app *BaseApp) init() error {
	if app.sealed {
		panic("cannot call initFromMainStore: baseapp already sealed")
	}

	// needed for the export command which inits from store but never calls initchain
	app.setCheckState(tmproto.Header{})
	app.Seal()

	return nil
}

func (app *BaseApp) setMinGasPrices(gasPrices sdk.DecCoins) { _ = "STUB: not implemented"; return }

func (app *BaseApp) setHaltHeight(haltHeight uint64) { _ = "STUB: not implemented"; return }

func (app *BaseApp) setHaltTime(haltTime uint64) { _ = "STUB: not implemented"; return }

func (app *BaseApp) setMinRetainBlocks(minRetainBlocks uint64) { _ = "STUB: not implemented"; return }

func (app *BaseApp) setInterBlockCache(cache sdk.MultiStorePersistentCache) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) setCompactionInterval(compactionInterval uint64) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) setTrace(trace bool) { _ = "STUB: not implemented"; return }

func (app *BaseApp) setIndexEvents(ie []string) { _ = "STUB: not implemented"; return }

// Router returns the router of the BaseApp.
func (app *BaseApp) Router() sdk.Router {
	_ = "STUB: not implemented"

	// We cannot return a Router when the app is sealed because we can't have
	// any routes modified which would cause unexpected routing behavior.
	return *new(sdk.Router)
}

// QueryRouter returns the QueryRouter of a BaseApp.
func (app *BaseApp) QueryRouter() sdk.QueryRouter {
	_ = "STUB: not implemented"
	return *

	// Seal seals a BaseApp. It prohibits any further modifications to a BaseApp.
	new(sdk.QueryRouter)
}

func (app *BaseApp) Seal() {
	_ = "STUB: not implemented"

	// IsSealed returns true if the BaseApp is sealed and false otherwise.
	return
}

func (app *BaseApp) IsSealed() bool {
	_ = "STUB: not implemented"

	// setCheckState sets the BaseApp's checkState with a branched multi-store
	// (i.e. a CacheMultiStore) and a new Context with the same multi-store branch,
	// provided header, and minimum gas prices set. It is set on InitChain and reset
	// on Commit.
	return false
}

func (app *BaseApp) setCheckState(header tmproto.Header) { _ = "STUB: not implemented"; return }

// setDeliverState sets the BaseApp's deliverState with a branched multi-store
// (i.e. a CacheMultiStore) and a new Context with the same multi-store branch,
// and provided header. It is set on InitChain and BeginBlock and set to nil on
// Commit.
func (app *BaseApp) setDeliverState(header tmproto.Header) { _ = "STUB: not implemented"; return }

func (app *BaseApp) setProcessProposalState(header tmproto.Header) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) resetStatesExceptCheckState() { _ = "STUB: not implemented"; return }

func (app *BaseApp) setProcessProposalHeader(header tmproto.Header) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) setDeliverStateHeader(header tmproto.Header) { _ = "STUB: not implemented"; return }

// GetProcessProposalCleanContext returns a context snapshotted at the start of
// ProcessProposal, before the handler runs. It has the correct store state,
// consensus params, and header, but is immune to speculative writes from
// optimistic processing.
func (app *BaseApp) GetProcessProposalCleanContext() sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

func (app *BaseApp) prepareProcessProposalState(headerHash []byte) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) prepareDeliverState(headerHash []byte) { _ = "STUB: not implemented"; return }

// GetConsensusParams returns the current consensus parameters from the BaseApp's
// ParamStore. If the BaseApp has no ParamStore defined, nil is returned.
func (app *BaseApp) GetConsensusParams(ctx sdk.Context) *tmproto.ConsensusParams {
	_ = "STUB: not implemented"
	return nil
}

// AddRunTxRecoveryHandler adds custom app.runTx method panic handlers.
func (app *BaseApp) AddRunTxRecoveryHandler(handlers ...RecoveryHandler) {
	_ = "STUB: not implemented"
	return
}

// StoreConsensusParams sets the consensus parameters to the baseapp's param store.
func (app *BaseApp) StoreConsensusParams(ctx sdk.Context, cp *tmproto.ConsensusParams) {
	_ = "STUB: not implemented"
	return
}

func (app *BaseApp) ValidateHeight(height int64) error { _ = "STUB: not implemented"; return nil }

// expectedHeight holds the expected height to validate.

// In this case, we're validating the first block of the chain (no
// previous commit). The height we're expecting is the initial height.

// This case can means two things:
// - either there was already a previous commit in the store, in which
// case we increment the version from there,
// - or there was no previous commit, and initial version was not set,
// in which case we start at version 1.

// validateBasicTxMsgs executes basic validator calls for messages.
func validateBasicTxMsgs(msgs []sdk.Msg) error { _ = "STUB: not implemented"; return nil }

// Returns the applications's deliverState if app is in runTxModeDeliver,
// otherwise it returns the application's checkstate.
func (app *BaseApp) getState(mode runTxMode) *state { _ = "STUB: not implemented"; return nil }

// retrieve the context for the tx w/ txBytes and other memoized values.
func (app *BaseApp) getContextForTx(mode runTxMode, txBytes []byte) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

func (app *BaseApp) GetCheckTxContext(txBytes []byte, recheck bool) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// CacheTxContext returns a new context based off of the provided context with
// a branched multi-store.
func (app *BaseApp) CacheTxContext(ctx sdk.Context, checksum [32]byte) (sdk.Context, sdk.CacheMultiStore) {
	_ = "STUB: not implemented"
	return *

	// TODO: https://github.com/cosmos/cosmos-sdk/issues/2824
	new(sdk.Context), *new(sdk.CacheMultiStore)
}

type runTxResult struct {
	gasInfo    sdk.GasInfo
	result     *sdk.Result
	anteEvents []abci.Event
	ctx        sdk.Context
}

// runTx processes a transaction within a given execution mode, encoded transaction
// bytes, and the decoded transaction itself. All state transitions occur through
// a cached Context depending on the mode provided. State only gets persisted
// if all messages get executed successfully and the execution mode is DeliverTx.
// Note, gas execution info is always returned. A reference to a Result is
// returned if the tx does not run out of gas and if all the messages are valid
// and execute successfully. An error is returned otherwise.
func (app *BaseApp) runTx(ctx sdk.Context, mode runTxMode, tx sdk.Tx, checksum [32]byte) (runTxRes runTxResult, err error) {
	_ = "STUB: not implemented"
	return *new(runTxResult), nil
}

// check for existing parent tracer, and if applicable, use it

// NOTE: GasWanted should be returned by the AnteHandler. GasUsed is
// determined by the GasMeter. We need access to the context to get the gas
// meter so we initialize upfront.

// TODO: do we have to wrap with occ enabled check?

// trace AnteHandler

// Branch context before AnteHandler call in case it aborts.
// This is required for both CheckTx and DeliverTx.
// Ref: https://github.com/cosmos/cosmos-sdk/issues/2772
//
// NOTE: Alternatively, we could require that AnteHandler ensures that
// writes do not happen if aborted/failed.  This may have some
// performance benefits, but it'll be more difficult to get right.

// At this point, newCtx.MultiStore() is a store branch, or something else
// replaced by the AnteHandler. We want the original multistore.
//
// Also, in the case of the tx aborting, we need to track gas consumed via
// the instantiated gas meter in the AnteHandler, so we update the context
// prior to returning.
//
// This also replaces the GasMeter in the context where GasUsed was initialized 0
// and updated with gas consumed in the ante handler runs
// The GasMeter is a pointer and its passed to the RunMsg and tracks the consumed
// gas there too.

// GasMeter expected to be set in AnteHandler

// Create a new Context based off of the existing Context with a MultiStore branch
// in case message processing fails. At this point, the MultiStore
// is a branch of a branch.

// Attempt to execute all messages and only update state if all messages pass
// and we're in DeliverTx. Note, RunMsgs will never return a reference to a
// Result if any single message fails or does not have a registered Handler.

// we do this since we will only be looking at result in DeliverTx

// append the events in the order of occurrence

// only apply hooks if no error

// RunMsgs iterates through a list of messages and executes them with the provided
// Context and execution mode. Messages will only be executed during simulation
// and DeliverTx. An error is returned if any single message fails or if a
// Handler does not exist for a given message route. Otherwise, a reference to a
// Result is returned. The caller must not commit state if an error is returned.
func (app *BaseApp) RunMsgs(ctx sdk.Context, msgs []sdk.Msg) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: GasWanted is determined by the AnteHandler and GasUsed by the GasMeter.

// name to use as value in event `message.action`

// ADR 031 request type routing

// legacy sdk.Msg routing
// Assuming that the app developer has migrated all their Msgs to
// proto messages and has registered all `Msg services`, then this
// path should never be called, because all those Msgs should be
// registered within the `msgServiceRouter` already.

// append message events, data and logs
//
// Note: Each message result's data must be length-prefixed in order to
// separate each result.

//nolint:gosec // loop range index

func (app *BaseApp) startCompactionRoutine(db dbm.DB) { _ = "STUB: not implemented"; return }

//nolint:gosec // compactionInterval is a small config value

func (app *BaseApp) Close() error {
	_ = "STUB: not implemented"
	// we do not want to close when a commit is ongoing since commit writes to stores
	// and metadata in a non-atomic way
	return nil
}

func (app *BaseApp) GetCheckCtx() sdk.Context { _ = "STUB: not implemented"; return *new(sdk.Context) }

func (app *BaseApp) RegisterDeliverTxHook(hook DeliverTxHook) { _ = "STUB: not implemented"; return }

func (app *BaseApp) InplaceTestnetInitialize(pk cryptotypes.PubKey) {
	_ = "STUB: not implemented"
	return
}
