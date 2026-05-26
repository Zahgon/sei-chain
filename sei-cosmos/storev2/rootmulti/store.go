package rootmulti

import (
	"io"
	"sync"

	"github.com/sei-protocol/seilog"
	"golang.org/x/time/rate"

	protoio "github.com/gogo/protobuf/io"
	snapshottypes "github.com/sei-protocol/sei-chain/sei-cosmos/snapshots/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	"github.com/sei-protocol/sei-chain/sei-db/config"
	seidbtypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	sctypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
)

var (
	logger = seilog.NewLogger("cosmos", "storev2", "rootmulti")

	_ types.CommitMultiStore = (*Store)(nil)
	_ types.Queryable        = (*Store)(nil)
)

type Store struct {
	mtx            sync.RWMutex
	scStore        sctypes.Committer
	ssStore        seidbtypes.StateStore
	lastCommitInfo *types.CommitInfo
	storesParams   map[types.StoreKey]storeParams
	storeKeys      map[string]types.StoreKey
	ckvStores      map[types.StoreKey]types.CommitKVStore
	gigaKeys       []string

	histProofSem     chan struct{}
	histProofLimiter *rate.Limiter

	snapshotSCStoreWarnOnce sync.Once
}

type VersionedChangesets struct {
	Version    int64
	Changesets []*proto.NamedChangeSet
}

func NewStore(
	homeDir string,
	scConfig config.StateCommitConfig,
	ssConfig config.StateStoreConfig,
	gigaKeys []string,
) *Store {
	_ = "STUB: not implemented"
	// Use custom directory if specified, otherwise use homeDir
	return nil
}

// Check whether SC was enabled before but SS was not

// Commit implements interface Committer, called by ABCI Commit
func (rs *Store) Commit(bumpVersion bool) types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

// Commit to SC Store

// The underlying sc store might be reloaded, reload the store as well.

// Flush all the pending changesets to commit store.
func (rs *Store) flush() error { _ = "STUB: not implemented"; return nil }

// it'll unwrap the inter-block cache

// ensure the state store watermark advances even for empty blocks

func (rs *Store) Close() error { _ = "STUB: not implemented"; return nil }

// LastCommitID Implements interface Committer
func (rs *Store) LastCommitID() types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

// Implements interface Committer
func (rs *Store) SetPruning(types.PruningOptions) {
	_ = "STUB: not implemented"

	// Implements interface Committer
	return
}

func (rs *Store) GetPruning() types.PruningOptions {
	_ = "STUB: not implemented"
	return *new(types.PruningOptions)
}

// Implements interface Store
func (rs *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// GetStateStore returns the ssStore instance
func (rs *Store) GetStateStore() seidbtypes.StateStore {
	_ = "STUB: not implemented"

	// Implements interface CacheWrapper
	return *new(seidbtypes.StateStore)
}

func (rs *Store) CacheWrap(_ types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// Implements interface CacheWrapper
func (rs *Store) CacheWrapWithTrace(storeKey types.StoreKey, _ io.Writer, _ types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// Implements interface MultiStore
func (rs *Store) CacheMultiStore() types.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore)
}

// cacheMultiStoreLocked must be called with rs.mtx held (at least RLock).
func (rs *Store) cacheMultiStoreLocked() types.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore)
}

// CacheMultiStoreWithVersion Implements interface MultiStore
// used to createQueryContext, abci_query or grpc query service.
func (rs *Store) CacheMultiStoreWithVersion(version int64) (types.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore), nil
}

// Serve from SS stores for ALL historical queries

// add the transient/mem stores registered in current app.

// Only serve from SC when query latest version and SS not enabled

func (rs *Store) CacheMultiStoreForExport(version int64) (types.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore), nil
}

// Open SC stores for wasm snapshot, this op is blocking and could take a long time

// add the transient/mem stores registered in current app.

// We need this because we need to make sure sc is closed after being used to release the resources

// SnapshotSCStore returns an O(1) SC snapshot, or nil when flatkv is engaged.
func (rs *Store) SnapshotSCStore() sctypes.Committer {
	_ = "STUB: not implemented"
	return *new(sctypes.Committer)
}

// CacheMultiStoreFromCommitter builds a CacheMultiStore backed by snap for
// IAVL stores; non-IAVL stores use their live counterparts.
func (rs *Store) CacheMultiStoreFromCommitter(snap sctypes.Committer) (types.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore), nil
}

// GetStore Implements interface MultiStore
func (rs *Store) GetStore(key types.StoreKey) types.Store {
	_ = "STUB: not implemented"
	return *

	// GetKVStore Implements interface MultiStore
	new(types.Store)
}

func (rs *Store) GetKVStore(key types.StoreKey) types.KVStore {
	_ = "STUB: not implemented"
	return *

	// Implements interface MultiStore
	new(types.KVStore)
}

func (rs *Store) TracingEnabled() bool {
	_ = "STUB: not implemented"

	// Implements interface MultiStore
	return false
}

func (rs *Store) SetTracer(_ io.Writer) types.MultiStore {
	_ = "STUB: not implemented"

	// Implements interface MultiStore
	return *new(types.MultiStore)
}

func (rs *Store) SetTracingContext(types.TraceContext) types.MultiStore {
	_ = "STUB: not implemented"

	// Implements interface Snapshotter
	// not needed, memiavl manage its own snapshot/pruning strategy
	return *new(types.MultiStore)
}

func (rs *Store) PruneSnapshotHeight(_ int64) {
	_ = "STUB: not implemented"

	// Implements interface Snapshotter
	// not needed, memiavl manage its own snapshot/pruning strategy
	return
}

func (rs *Store) SetSnapshotInterval(_ uint64) {
	_ = "STUB: not implemented"

	// Implements interface CommitMultiStore
	return
}

func (rs *Store) MountStoreWithDB(key types.StoreKey, typ types.StoreType, _ dbm.DB) {
	_ = "STUB: not implemented"
	return
}

// Implements interface CommitMultiStore
func (rs *Store) GetCommitStore(key types.StoreKey) types.CommitStore {
	_ = "STUB: not implemented"
	return *new(types.CommitStore)
}

// GetCommitKVStore Implements interface CommitMultiStore
func (rs *Store) GetCommitKVStore(key types.StoreKey) types.CommitKVStore {
	_ = "STUB: not implemented"
	return *

	// Implements interface CommitMultiStore
	// used by normal node startup.
	new(types.CommitKVStore)
}

func (rs *Store) LoadLatestVersion() error { _ = "STUB: not implemented"; return nil }

// Implements interface CommitMultiStore
func (rs *Store) LoadLatestVersionAndUpgrade(upgrades *types.StoreUpgrades) error {
	_ = "STUB: not implemented"
	return nil
}

// Implements interface CommitMultiStore
// used by node startup with UpgradeStoreLoader
func (rs *Store) LoadVersionAndUpgrade(version int64, upgrades *types.StoreUpgrades) error {
	_ = "STUB: not implemented"
	return nil
}

// load storeKeys for deletion

// deterministic iteration order for upgrades

// to keep the root hash compatible with cosmos-sdk 0.46

func (rs *Store) loadCommitStoreFromParams(key types.StoreKey, params storeParams) (types.CommitKVStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CommitKVStore), nil
}

// Implements interface CommitMultiStore
// used by export cmd
func (rs *Store) LoadVersion(ver int64) error { _ = "STUB: not implemented"; return nil }

// SetInterBlockCache is a noop since we do caching on its own, which works well with zero-copy.
func (rs *Store) SetInterBlockCache(_ types.MultiStorePersistentCache) {
	_ = "STUB: not implemented"

	// SetInitialVersion Implements interface CommitMultiStore
	// used by InitChain when the initial height is bigger than 1
	return
}

func (rs *Store) SetInitialVersion(version int64) error { _ = "STUB: not implemented"; return nil }

// Implements interface CommitMultiStore
func (rs *Store) SetLazyLoading(_ bool) {
	_ = "STUB: not implemented"

	// RollbackToVersion delete the versions after `target` and update the latest version.
	// it should only be called in standalone cli commands.
	return
}

func (rs *Store) RollbackToVersion(target int64) error { _ = "STUB: not implemented"; return nil }

// We need to update the lastCommitInfo after rollback

// getStoreByName performs a lookup of a StoreKey given a store name typically
// provided in a path. The StoreKey is then used to perform a lookup and return
// a Store. If the Store is wrapped in an inter-block cache, it will be unwrapped
// prior to being returned. If the StoreKey does not exist, nil is returned.
func (rs *Store) GetStoreByName(name string) types.Store {
	_ = "STUB: not implemented"
	return *new(types.Store)
}

// Implements interface Queryable
func (rs *Store) Query(req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// keep downstream store.Query height consistent

// Fast path: no proof + SS enabled

// latest never needs historical LoadVersion clone

// historical path (this is where RPC pressure happens)

// If underlying query failed (e.g. invalid height/path) or doesn' need proof, return as-is.

// Must have proof ops from underlying store query before appending commit proof.

func (rs *Store) tryAcquireHistProofPermit() error { _ = "STUB: not implemented"; return nil }

func (rs *Store) releaseHistProofPermit() { _ = "STUB: not implemented"; return }

// parsePath expects a format like /<storeName>[/<subpath>]
// Must start with /, subpath may be empty
// Returns error if it doesn't start with /
func parsePath(path string) (storeName string, subpath string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type storeParams struct {
	key types.StoreKey
	typ types.StoreType
}

func newStoreParams(key types.StoreKey, typ types.StoreType) storeParams {
	_ = "STUB: not implemented"
	return *new(storeParams)
}

func mergeStoreInfos(commitInfo *types.CommitInfo, storeInfos []types.StoreInfo) *types.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

// amendCommitInfo add mem stores commit infos to keep it compatible with cosmos-sdk 0.46
func amendCommitInfo(commitInfo *types.CommitInfo, storeParams map[types.StoreKey]storeParams) *types.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

func convertCommitInfo(commitInfo *proto.CommitInfo) *types.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

// GetWorkingHash returns the working app hash
func (rs *Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// for sdk 0.46 and backward compatibility

func (rs *Store) GetEvents() []abci.Event { _ = "STUB: not implemented"; return nil }

func (rs *Store) ResetEvents() { _ = "STUB: not implemented"; return }

// Restore Implements interface Snapshotter
func (rs *Store) Restore(
	height uint64, format uint32, protoReader protoio.Reader,
) (snapshottypes.SnapshotItem, error) {
	_ = "STUB: not implemented"
	return *new(snapshottypes.SnapshotItem), nil
}

//nolint:gosec // bounds checked above

func (rs *Store) restore(height int64, protoReader protoio.Reader) (snapshottypes.SnapshotItem, error) {
	_ = "STUB: not implemented"
	return *new(snapshottypes.SnapshotItem), nil
}

//nolint:gosec // bounds checked above against math.MaxInt8

// Protobuf does not differentiate between []byte{} as nil, but fortunately IAVL does
// not allow nil keys nor nil values for leaf nodes, so we can always set them to empty.

// Check if we should also import to SS store

// unknown element, could be an extension

// Initialize SS version metadata. Without SetLatestVersion, GetLatestVersion()
// stays 0 until the first post-sync block commits, which is misleading to any
// caller that reads it in that window.

// Snapshot Implements the interface from Snapshotter
func (rs *Store) Snapshot(height uint64, protoWriter protoio.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// SetKVStores implements types.CommitMultiStore.
func (*Store) SetKVStores(handler func(key types.StoreKey, s types.KVStore) types.CacheWrap) types.MultiStore {
	_ = "STUB: not implemented"
	return *

	// StoreKeys implements types.CommitMultiStore.
	new(types.MultiStore)
}

func (rs *Store) StoreKeys() []types.StoreKey { _ = "STUB: not implemented"; return nil }

// GetEarliestVersion return earliest version for SS or latestVersion if only SC is enabled
func (rs *Store) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }
