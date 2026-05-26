package rootmulti

import (
	"io"
	"sync"

	protoio "github.com/gogo/protobuf/io"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/seilog"
	dbm "github.com/tendermint/tm-db"

	snapshottypes "github.com/sei-protocol/sei-chain/sei-cosmos/snapshots/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

var logger = seilog.NewLogger("cosmos", "store", "rootmulti")

const (
	latestVersionKey = "s/latest"
	pruneHeightsKey  = "s/pruneheights"
	commitInfoKeyFmt = "s/%d" // s/<version>

	proofsPath = "proofs"
)

// Store is composed of many CommitStores. Name contrasts with
// cacheMultiStore which is used for branching other MultiStores. It implements
// the CommitMultiStore interface.
type Store struct {
	db                dbm.DB
	archivalDb        dbm.DB
	lastCommitInfo    *types.CommitInfo
	lastCommitInfoMtx sync.RWMutex
	pruningOpts       types.PruningOptions
	storesParams      map[types.StoreKey]storeParams
	stores            map[types.StoreKey]types.CommitKVStore
	keysByName        map[string]types.StoreKey
	lazyLoading       bool
	pruneHeights      []int64
	initialVersion    int64
	archivalVersion   int64
	earliestVersion   int64

	traceWriter       io.Writer
	traceContext      types.TraceContext
	traceContextMutex sync.Mutex

	interBlockCache types.MultiStorePersistentCache
}

var (
	_ types.CommitMultiStore = (*Store)(nil)
	_ types.Queryable        = (*Store)(nil)
)

// NewStore returns a reference to a new Store object with the provided DB. The
// store will be created with a PruneNothing pruning strategy by default. After
// a store is created, KVStores must be mounted and finally LoadLatestVersion or
// LoadVersion must be called.
func NewStore(db dbm.DB) *Store { _ = "STUB: not implemented"; return nil }

func NewStoreWithArchival(db, archivalDb dbm.DB, archivalVersion int64) *Store {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Store) shouldUseArchivalDb(ver int64) bool { _ = "STUB: not implemented"; return false }

// GetPruning fetches the pruning strategy from the root store.
func (rs *Store) GetPruning() types.PruningOptions {
	_ = "STUB: not implemented"
	return *

	// SetPruning sets the pruning strategy on the root store and all the sub-stores.
	// Note, calling SetPruning on the root store prior to LoadVersion or
	// LoadLatestVersion performs a no-op as the stores aren't mounted yet.
	new(types.PruningOptions)
}

func (rs *Store) SetPruning(pruningOpts types.PruningOptions) { _ = "STUB: not implemented"; return }

// SetLazyLoading sets if the iavl store should be loaded lazily or not
func (rs *Store) SetLazyLoading(lazyLoading bool) { _ = "STUB: not implemented"; return }

// GetStoreType implements Store.
func (rs *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// MountStoreWithDB implements CommitMultiStore.
func (rs *Store) MountStoreWithDB(key types.StoreKey, typ types.StoreType, db dbm.DB) {
	_ = "STUB: not implemented"
	return
}

// GetCommitStore returns a mounted CommitStore for a given StoreKey. If the
// store is wrapped in an inter-block cache, it will be unwrapped before returning.
func (rs *Store) GetCommitStore(key types.StoreKey) types.CommitStore {
	_ = "STUB: not implemented"
	return *new(types.CommitStore)
}

// GetCommitKVStore returns a mounted CommitKVStore for a given StoreKey. If the
// store is wrapped in an inter-block cache, it will be unwrapped before returning.
func (rs *Store) GetCommitKVStore(key types.StoreKey) types.CommitKVStore {
	_ = "STUB: not implemented"
	// If the Store has an inter-block cache, first attempt to lookup and unwrap
	// the underlying CommitKVStore by StoreKey. If it does not exist, fallback to
	// the main mapping of CommitKVStores.
	return *new(types.CommitKVStore)
}

// GetStores returns mounted stores
func (rs *Store) GetStores() map[types.StoreKey]types.CommitKVStore {
	_ = "STUB: not implemented"

	// GetStores returns mounted stores
	return nil
}

func (rs *Store) GetEvents() []abci.Event { _ = "STUB: not implemented"; return nil }

func (rs *Store) ResetEvents() { _ = "STUB: not implemented"; return }

// LoadLatestVersionAndUpgrade implements CommitMultiStore
func (rs *Store) LoadLatestVersionAndUpgrade(upgrades *types.StoreUpgrades) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadVersionAndUpgrade allows us to rename substores while loading an older version
func (rs *Store) LoadVersionAndUpgrade(ver int64, upgrades *types.StoreUpgrades) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadLatestVersion implements CommitMultiStore.
func (rs *Store) LoadLatestVersion() error { _ = "STUB: not implemented"; return nil }

// LoadVersion implements CommitMultiStore.
func (rs *Store) LoadVersion(ver int64) error { _ = "STUB: not implemented"; return nil }

func (rs *Store) loadVersion(ver int64, upgrades *types.StoreUpgrades) error {
	_ = "STUB: not implemented"
	return nil
}

// load old data if we are not version 0

// convert StoreInfos slice to map

// load each Store (note this doesn't panic on unmounted keys now)

// TODO: is this safe

// deterministic iteration order for upgrades
// (as the underlying store may change and
// upgrades make store changes where the execution order may matter)

// If it has been added, set the initial version

//nolint:gosec // bounds checked above

// If it was deleted, remove all data

// drop deleted KV store from stores

// handle renames specially
// make an unregistered key to satify loadCommitStore params

// load from the old name

// move all data

// load any pruned heights we missed from disk to be pruned on the next run

func (rs *Store) getCommitID(infos map[string]types.StoreInfo, name string) types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

func deleteKVStore(kv types.KVStore) {
	_ = "STUB: not implemented"
	// Note that we cannot write while iterating, so load all keys here, delete below
	return
}

// we simulate move by a copy and delete
func moveKVStoreData(oldDB types.KVStore, newDB types.KVStore) {
	_ = "STUB: not implemented"
	// we read from one and write to another
	return
}

// then delete the old store

// SetInterBlockCache sets the Store's internal inter-block (persistent) cache.
// When this is defined, all CommitKVStores will be wrapped with their respective
// inter-block cache.
func (rs *Store) SetInterBlockCache(c types.MultiStorePersistentCache) {
	_ = "STUB: not implemented"
	return

	// SetTracer sets the tracer for the MultiStore that the underlying
	// stores will utilize to trace operations. A MultiStore is returned.
}

func (rs *Store) SetTracer(w io.Writer) types.MultiStore {
	_ = "STUB: not implemented"
	return *new(types.MultiStore)
}

// SetTracingContext updates the tracing context for the MultiStore by merging
// the given context with the existing context by key. Any existing keys will
// be overwritten. It is implied that the caller should update the context when
// necessary between tracing operations. It returns a modified MultiStore.
func (rs *Store) SetTracingContext(tc types.TraceContext) types.MultiStore {
	_ = "STUB: not implemented"
	return *new(types.MultiStore)
}

func (rs *Store) getTracingContext() types.TraceContext {
	_ = "STUB: not implemented"
	return *new(types.TraceContext)
}

// TracingEnabled returns if tracing is enabled for the MultiStore.
func (rs *Store) TracingEnabled() bool { _ = "STUB: not implemented"; return false }

// LastCommitID implements Committer/CommitStore.
func (rs *Store) LastCommitID() types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

func (rs *Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Commit implements Committer/CommitStore.
func (rs *Store) Commit(bumpVersion bool) types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

// This case means that no commit has been made in the store, we
// start from initialVersion.

// This case can means two things:
// - either there was already a previous commit in the store, in which
// case we increment the version from there,
// - or there was no previous commit, and initial version was not set,
// in which case we start at version 1.

//nolint:gosec // pruning config values are small, won't overflow int64
//nolint:gosec // pruning config values are small, won't overflow int64
//nolint:gosec // pruning config values are small, won't overflow int64

// batch prune if the current height is a pruning interval height

// PruneStores will batch delete a list of heights from each mounted sub-store.
// If clearStorePruningHeihgts is true, store's pruneHeights is appended to the
// pruningHeights and reset after finishing pruning.
func (rs *Store) PruneStores(clearStorePruningHeights bool, pruningHeights []int64) {
	_ = "STUB: not implemented"
	return
}

// CacheWrap implements CacheWrapper/Store/CommitStore.
func (rs *Store) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements the CacheWrapper interface.
func (rs *Store) CacheWrapWithTrace(storeKey types.StoreKey, _ io.Writer, _ types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheMultiStore creates ephemeral branch of the multi-store and returns a CacheMultiStore.
// It implements the MultiStore interface.
func (rs *Store) CacheMultiStore() types.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore)
}

// CacheMultiStoreWithVersion is analogous to CacheMultiStore except that it
// attempts to load stores at a given version (height). An error is returned if
// any store cannot be loaded. This should only be used for querying and
// iterating at past heights.
func (rs *Store) CacheMultiStoreWithVersion(version int64) (types.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore), nil
}

func (rs *Store) CacheMultiStoreForExport(version int64) (types.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore), nil
}

// GetStore returns a mounted Store for a given StoreKey. If the StoreKey does
// not exist, it will panic. If the Store is wrapped in an inter-block cache, it
// will be unwrapped prior to being returned.
//
// TODO: This isn't used directly upstream. Consider returning the Store as-is
// instead of unwrapping.
func (rs *Store) GetStore(key types.StoreKey) types.Store {
	_ = "STUB: not implemented"
	return *new(types.Store)
}

// GetKVStore returns a mounted KVStore for a given StoreKey. If tracing is
// enabled on the KVStore, a wrapped TraceKVStore will be returned with the root
// store's tracer, otherwise, the original KVStore will be returned.
//
// NOTE: The returned KVStore may be wrapped in an inter-block cache if it is
// set on the root store.
func (rs *Store) GetKVStore(key types.StoreKey) types.KVStore {
	_ = "STUB: not implemented"
	return *new(types.KVStore)
}

// GetStoreByName performs a lookup of a StoreKey given a store name typically
// provided in a path. The StoreKey is then used to perform a lookup and return
// a Store. If the Store is wrapped in an inter-block cache, it will be unwrapped
// prior to being returned. If the StoreKey does not exist, nil is returned.
func (rs *Store) GetStoreByName(name string) types.Store {
	_ = "STUB: not implemented"
	return *new(types.Store)
}

// Query calls substore.Query with the same `req` where `req.Path` is
// modified to remove the substore prefix.
// Ie. `req.Path` here is `/<substore>/<path>`, and trimmed to `/<path>` for the substore.
// Special case: if `req.Path` is `/proofs`, the commit hash is included
// as response value. In addition, proofs of every store are appended to the response for
// the requested height
func (rs *Store) Query(req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// trim the path and make the query

// If the request's height is the latest height we've committed, then utilize
// the store's lastCommitInfo as this commit info may not be flushed to disk.
// Otherwise, we query for the commit info from disk.

// Restore origin path and append proof op.

// SetInitialVersion sets the initial version of the IAVL tree. It is used when
// starting a new chain at an arbitrary height.
// NOTE: this never errors. Can we fix the function signature ?
func (rs *Store) SetInitialVersion(version int64) error { _ = "STUB: not implemented"; return nil }

// parsePath expects a format like /<storeName>[/<subpath>]
// Must start with /, subpath may be empty
// Returns error if it doesn't start with /
func parsePath(path string) (storeName string, subpath string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

//---------------------- Snapshotting ------------------

// Snapshot implements snapshottypes.Snapshotter. The snapshot output for a given format must be
// identical across nodes such that chunks from different sources fit together. If the output for a
// given format changes (at the byte level), the snapshot format must be bumped - see
// TestMultistoreSnapshot_Checksum test.
func (rs *Store) Snapshot(height uint64, protoWriter protoio.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// Restore implements snapshottypes.Snapshotter.
// returns next snapshot item and error.
func (rs *Store) Restore(
	height uint64, format uint32, protoReader protoio.Reader,
) (snapshottypes.SnapshotItem, error) {
	_ = "STUB: not implemented"
	return *new(snapshottypes.SnapshotItem), nil
}

func (rs *Store) loadCommitStoreFromParams(key types.StoreKey, id types.CommitID, params storeParams) (_ types.CommitKVStore, _err error) {
	_ = "STUB: not implemented"
	return *new(types.CommitKVStore), nil
}

//nolint:gosec // bounds checked above

// RollbackToVersion delete the versions after `target` and update the latest version.
func (rs *Store) RollbackToVersion(target int64) error { _ = "STUB: not implemented"; return nil }

func (rs *Store) flushMetadata(db dbm.DB, version int64, cInfo *types.CommitInfo) {
	_ = "STUB: not implemented"
	return
}

func (rs *Store) LastCommitInfo() *types.CommitInfo { _ = "STUB: not implemented"; return nil }

func (rs *Store) SetLastCommitInfo(c *types.CommitInfo) { _ = "STUB: not implemented"; return }

type storeParams struct {
	key            types.StoreKey
	db             dbm.DB
	typ            types.StoreType
	initialVersion uint64
}

func GetLatestVersion(db dbm.DB) int64 { _ = "STUB: not implemented"; return 0 }

// Commits each store and returns a new commitInfo.
func commitStores(version int64, storeMap map[types.StoreKey]types.CommitKVStore, bumpVersion bool) *types.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Store) doProofsQuery(req abci.RequestQuery) abci.ResponseQuery {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// Gets commitInfo from disk.
func getCommitInfo(db dbm.DB, ver int64) (*types.CommitInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPruningHeights(db dbm.DB) ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // deserialized block heights stored by flushPruningHeights, always non-negative

func flushCommitInfo(batch dbm.Batch, version int64, cInfo *types.CommitInfo) {
	_ = "STUB: not implemented"
	return
}

func flushLatestVersion(batch dbm.Batch, version int64) { _ = "STUB: not implemented"; return }

func flushPruningHeights(batch dbm.Batch, pruneHeights []int64) { _ = "STUB: not implemented"; return }

//nolint:gosec // pruning heights are always non-negative block heights

func (rs *Store) Close() error { _ = "STUB: not implemented"; return nil }

func (rs *Store) SetKVStores(handler func(key types.StoreKey, s types.KVStore) types.CacheWrap) types.MultiStore {
	_ = "STUB: not implemented"
	return *new(types.MultiStore)
}

func (rs *Store) StoreKeys() []types.StoreKey { _ = "STUB: not implemented"; return nil }

func (rs *Store) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }
