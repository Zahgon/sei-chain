package cachemulti

import (
	"io"
	"sync"

	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

//----------------------------------------
// Store

// Store holds many branched stores.
// Implements MultiStore.
// NOTE: a Store (and MultiStores in general) should never expose the
// keys for the substores.
type Store struct {
	db      types.CacheKVStore
	stores  map[types.StoreKey]types.CacheWrap
	parents map[types.StoreKey]types.CacheWrapper
	keys    map[string]types.StoreKey

	gigaStores map[types.StoreKey]types.KVStore
	gigaKeys   []types.StoreKey

	traceWriter  io.Writer
	traceContext types.TraceContext

	mu              *sync.RWMutex // protects stores and parents during lazy creation
	materializeOnce *sync.Once

	closers []io.Closer
}

var _ types.CacheMultiStore = Store{}

// NewFromKVStore creates a new Store object from a mapping of store keys to
// CacheWrapper objects and a KVStore as the database. Each CacheWrapper store
// is a branched store.
func NewFromKVStore(
	store types.KVStore, stores map[types.StoreKey]types.CacheWrapper,
	gigaStores map[types.StoreKey]types.KVStore,
	keys map[string]types.StoreKey, gigaKeys []types.StoreKey, traceWriter io.Writer, traceContext types.TraceContext,
) Store {
	_ = "STUB: not implemented"
	return *new(Store)
}

// if key is in gigaStores, use it as the parent store

// if not, use regular store as the parent store

func newStoreWithoutGiga(store types.KVStore, stores map[types.StoreKey]types.CacheWrapper, keys map[string]types.StoreKey, gigaKeys []types.StoreKey, traceWriter io.Writer, traceContext types.TraceContext) Store {
	_ = "STUB: not implemented"
	return *new(Store)
}

// NewStore creates a new Store object from a mapping of store keys to
// CacheWrapper objects. Each CacheWrapper store is a branched store.
func NewStore(
	db dbm.DB, stores map[types.StoreKey]types.CacheWrapper, keys map[string]types.StoreKey,
	gigaKeys []types.StoreKey, traceWriter io.Writer, traceContext types.TraceContext,
) Store {
	_ = "STUB: not implemented"
	return *new(Store)
}

func newCacheMultiStoreFromCMS(cms Store) Store {
	_ = "STUB: not implemented"
	// Thread-safe materialization: the OCC scheduler calls CacheMultiStore()
	// concurrently from multiple goroutines on the same block CMS.
	// sync.Once ensures exactly one goroutine materializes, others wait.
	return *new(Store)
}

// Lock held for bulk materialization to avoid per-key lock overhead.

// Inline the creation here — we already hold the write lock.

// After Do returns, cms.parents is empty and cms.stores has all entries.

// cms.parents is now empty — all moved to cms.stores by getOrCreateStore

// getOrCreateStore lazily creates a cachekv store from its parent on first access.
// Thread-safe: concurrent callers (e.g. slashing BeginBlocker goroutines) may
// call GetKVStore on the same CMS simultaneously.
func (cms Store) getOrCreateStore(key types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	// Fast path: store already materialized, read-only check.
	return *new(types.CacheWrap)
}

// Slow path: acquire write lock and create.

// Double-check after acquiring write lock.

// SetTracer sets the tracer for the MultiStore that the underlying
// stores will utilize to trace operations. A MultiStore is returned.
func (cms Store) SetTracer(w io.Writer) types.MultiStore {
	_ = "STUB: not implemented"
	return *new(types.MultiStore)
}

// SetTracingContext updates the tracing context for the MultiStore by merging
// the given context with the existing context by key. Any existing keys will
// be overwritten. It is implied that the caller should update the context when
// necessary between tracing operations. It returns a modified MultiStore.
func (cms Store) SetTracingContext(tc types.TraceContext) types.MultiStore {
	_ = "STUB: not implemented"
	return *new(types.MultiStore)
}

// TracingEnabled returns if tracing is enabled for the MultiStore.
func (cms Store) TracingEnabled() bool { _ = "STUB: not implemented"; return false }

// GetStoreType returns the type of the store.
func (cms Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// Write calls Write on each underlying store.
func (cms Store) Write() { _ = "STUB: not implemented"; return }

func (cms Store) WriteGiga() { _ = "STUB: not implemented"; return }

// Implements CacheWrapper.
func (cms Store) CacheWrap(_ types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements the CacheWrapper interface.
func (cms Store) CacheWrapWithTrace(storeKey types.StoreKey, _ io.Writer, _ types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// Implements MultiStore.
func (cms Store) CacheMultiStore() types.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore)
}

// CacheMultiStoreWithVersion implements the MultiStore interface. It will panic
// as an already cached multi-store cannot load previous versions.
//
// TODO: The store implementation can possibly be modified to support this as it
// seems safe to load previous versions (heights).
func (cms Store) CacheMultiStoreWithVersion(_ int64) (types.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore), nil
}

// GetStore returns an underlying Store by key.
func (cms Store) GetStore(key types.StoreKey) types.Store {
	_ = "STUB: not implemented"
	return *new(types.Store)
}

// GetKVStore returns an underlying KVStore by key.
func (cms Store) GetKVStore(key types.StoreKey) types.KVStore {
	_ = "STUB: not implemented"
	return *new(types.KVStore)
}

func (cms Store) GetGigaKVStore(key types.StoreKey) types.KVStore {
	_ = "STUB: not implemented"
	return *new(types.KVStore)
}

func (cms Store) IsStoreGiga(key types.StoreKey) bool { _ = "STUB: not implemented"; return false }

func (cms Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// StoreKeys returns a list of all store keys
func (cms Store) StoreKeys() []types.StoreKey { _ = "STUB: not implemented"; return nil }

// SetKVStores sets the underlying KVStores via a handler for each key
func (cms Store) SetKVStores(handler func(sk types.StoreKey, s types.KVStore) types.CacheWrap) types.MultiStore {
	_ = "STUB: not implemented"
	// Force-create any lazy stores
	return *new(types.MultiStore)
}

func (cms Store) SetGigaKVStores(handler func(sk types.StoreKey, s types.KVStore) types.KVStore) types.MultiStore {
	_ = "STUB: not implemented"
	return *new(types.MultiStore)
}

func (cms Store) CacheMultiStoreForExport(_ int64) (types.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore), nil
}

func (cms *Store) AddCloser(closer io.Closer) { _ = "STUB: not implemented"; return }

func (cms Store) Close() { _ = "STUB: not implemented"; return }

func (cms Store) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }
