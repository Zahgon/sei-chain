package multiversion

import (
	"io"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	scheduler "github.com/sei-protocol/sei-chain/sei-cosmos/types/occ"
	dbm "github.com/tendermint/tm-db"
)

// exposes a handler for adding items to readset, useful for iterators
type ReadsetHandler interface {
	UpdateReadSet(key []byte, value []byte)
}

type NoOpHandler struct{}

func (NoOpHandler) UpdateReadSet(key []byte, value []byte) {
	_ = "STUB: not implemented"

	// exposes a handler for adding items to iterateset, to be called upon iterator close
	return
}

type IterateSetHandler interface {
	UpdateIterateSet(*iterationTracker)
}

type iterationTracker struct {
	startKey     []byte              // start of the iteration range
	endKey       []byte              // end of the iteration range
	earlyStopKey []byte              // key that caused early stop
	iteratedKeys map[string]struct{} // TODO: is a map okay because the ordering will be enforced when we replay the iterator?
	ascending    bool

	writeset WriteSet

	// TODO: is it possible that terimation is affected by keys later in iteration that weren't reached? eg. number of keys affecting iteration?
	// TODO: i believe to get number of keys the iteration would need to be done fully so its not a concern?

	// TODO: maybe we need to store keys served from writeset for the transaction? that way if theres OTHER keys within the writeset and the iteration range, and were written to the writeset later, we can discriminate between the groups?
	// keysServedFromWriteset map[string]struct{}

	// actually its simpler to just store a copy of the writeset at the time of iterator creation
}

func NewIterationTracker(startKey, endKey []byte, ascending bool, writeset WriteSet) iterationTracker {
	_ = "STUB: not implemented"
	return *new(iterationTracker)
}

// AddKey adds a key to the iterated keys map and sets the early stop key as the key since it's the latest key iterated
func (item *iterationTracker) AddKey(key []byte) { _ = "STUB: not implemented"; return }

func (item *iterationTracker) SetEarlyStopKey(key []byte) { _ = "STUB: not implemented"; return }

// Version Indexed Store wraps the multiversion store in a way that implements the KVStore interface, but also stores the index of the transaction, and so store actions are applied to the multiversion store using that index
type VersionIndexedStore struct {
	// TODO: this shouldnt NEED a mutex because its used within single transaction execution, therefore no concurrency
	// mtx sync.Mutex
	// used for tracking reads and writes for eventual validation + persistence into multi-version store
	// TODO: does this need sync.Map?
	readset    map[string][][]byte // contains the key -> []value mapping for all keys read from the store (not mvkv, underlying store)
	writeset   map[string][]byte   // contains the key -> value mapping for all keys written to the store
	iterateset Iterateset
	// TODO: need to add iterateset here as well

	// parent stores (both multiversion and underlying parent store)
	multiVersionStore MultiVersionStore
	parent            types.KVStore
	// transaction metadata for versioned operations
	transactionIndex int
	incarnation      int
	// have abort channel here for aborting transactions
	abortChannel chan scheduler.Abort
}

var _ types.KVStore = (*VersionIndexedStore)(nil)
var _ ReadsetHandler = (*VersionIndexedStore)(nil)
var _ IterateSetHandler = (*VersionIndexedStore)(nil)

func NewVersionIndexedStore(parent types.KVStore, multiVersionStore MultiVersionStore, transactionIndex, incarnation int, abortChannel chan scheduler.Abort) *VersionIndexedStore {
	_ = "STUB: not implemented"
	return nil
}

// GetReadset returns the readset
func (store *VersionIndexedStore) GetReadset() map[string][][]byte {
	_ = "STUB: not implemented"
	return nil

	// GetWriteset returns the writeset
}

func (store *VersionIndexedStore) GetWriteset() map[string][]byte {
	_ = "STUB: not implemented"
	return nil

	// WriteAbort writes an abort to the store but only allows one abort to be written PER instance of mvkv. This is because we pair abort channel writes with panics, and if we hit this more than once, it means that the panic was swallowed, so we won't write any aborts after a first abort is written to prevent any potential for deadlocking due to full channels
}

func (store *VersionIndexedStore) WriteAbort(abort scheduler.Abort) {
	_ = "STUB: not implemented"
	return
}

// Get implements types.KVStore.
func (store *VersionIndexedStore) Get(key []byte) []byte {
	_ = "STUB: not implemented"
	// first try to get from writeset cache, if cache miss, then try to get from multiversion store, if that misses, then get from parent store
	// if the key is in the cache, return it
	return nil
}

// don't have RW mutex because we have to update readset
// TODO: remove?
// store.mtx.Lock()
// defer store.mtx.Unlock()
// defer telemetry.MeasureSince(time.Now(), "store", "mvkv", "get")

// first check the MVKV writeset, and return that value if present

// return the value from the cache, no need to update any readset stuff

// read the readset to see if the value exists - and return if applicable

// just return the first one, if there is more than one, we will fail the validation anyways

// if we didn't find it, then we want to check the multivalue store + add to readset if applicable

// This handles both detecting readset conflicts and updating readset if applicable

// if we didn't find it in the multiversion store, then we want to check the parent store + add to readset

// This functions handles reads with deleted items and values and verifies that the data is consistent to what we currently have in the readset (IF we have a readset value for that key)
func (store *VersionIndexedStore) parseValueAndUpdateReadset(strKey string, mvsValue MultiVersionValueItem) []byte {
	_ = "STUB: not implemented"
	return nil
}

// This function iterates over the readset, validating that the values in the readset are consistent with the values in the multiversion store and underlying parent store, and returns a boolean indicating validity
func (store *VersionIndexedStore) ValidateReadset() bool {
	_ = "STUB: not implemented"
	// TODO: remove?
	// store.mtx.Lock()
	// defer store.mtx.Unlock()
	// defer telemetry.MeasureSince(time.Now(), "store", "mvkv", "validate_readset")
	return false
}

// sort the readset keys - this is so we have consistent behavior when theres varying conflicts within the readset (eg. read conflict vs estimate)

// iterate over readset keys and values

// if we have more than one value, we will fail the validation since we dedup when adding to readset

// if we see an estimate, that means that we need to abort and rerun

// check for `nil`

// check for equality

// value is valid, continue to next key

// this shouldnt happen because if we have a conflict it should always happen within multiversion store

// value was correct, we can continue to the next value

// Delete implements types.KVStore.
func (store *VersionIndexedStore) Delete(key []byte) {
	_ = "STUB: not implemented"
	// TODO: remove?
	// store.mtx.Lock()
	// defer store.mtx.Unlock()
	// defer telemetry.MeasureSince(time.Now(), "store", "mvkv", "delete")
	return
}

// Has implements types.KVStore.
func (store *VersionIndexedStore) Has(key []byte) bool {
	_ = "STUB: not implemented"
	// necessary locking happens within store.Get
	return false
}

// Set implements types.KVStore.
func (store *VersionIndexedStore) Set(key []byte, value []byte) {
	_ = "STUB: not implemented"
	// TODO: remove?
	// store.mtx.Lock()
	// defer store.mtx.Unlock()
	// defer telemetry.MeasureSince(time.Now(), "store", "mvkv", "set")
	return
}

// Iterator implements types.KVStore.
func (v *VersionIndexedStore) Iterator(start []byte, end []byte) dbm.Iterator {
	_ = "STUB: not implemented"
	return *new(dbm.Iterator)
}

// ReverseIterator implements types.KVStore.
func (v *VersionIndexedStore) ReverseIterator(start []byte, end []byte) dbm.Iterator {
	_ = "STUB: not implemented"
	return *new(dbm.Iterator)
}

// Iterator implements types.KVStore.
func (store *VersionIndexedStore) iterator(start []byte, end []byte, ascending bool) dbm.Iterator {
	_ = "STUB: not implemented"
	// TODO: remove?
	// store.mtx.Lock()
	// defer store.mtx.Unlock()
	return *new(dbm.Iterator)
}

// get the sorted keys from MVS
// TODO: ideally we take advantage of mvs keys already being sorted
// TODO: ideally merge btree and mvs keys into a single sorted btree

// TODO: ideally we persist writeset keys into a sorted btree for later use
// make a set of total keys across mvkv and mvs to iterate

// also add readset elements such that they fetch from readset instead of parent

// make a memIterator

// mergeIterator

func (v *VersionIndexedStore) VersionExists(version int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *VersionIndexedStore) DeleteAll(start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *VersionIndexedStore) GetAllKeyStrsInRange(start, end []byte) (res []string) {
	_ = "STUB: not implemented"
	return nil
}

// GetStoreType implements types.KVStore.
func (v *VersionIndexedStore) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// CacheWrap implements types.KVStore.
func (*VersionIndexedStore) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements types.KVStore.
func (*VersionIndexedStore) CacheWrapWithTrace(storeKey types.StoreKey, w io.Writer, tc types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// GetWorkingHash implements types.KVStore.
func (v *VersionIndexedStore) GetWorkingHash() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only entrypoint to mutate writeset
func (store *VersionIndexedStore) setValue(key, value []byte) { _ = "STUB: not implemented"; return }

func (store *VersionIndexedStore) WriteToMultiVersionStore() {
	_ = "STUB: not implemented"
	// TODO: remove?
	// store.mtx.Lock()
	// defer store.mtx.Unlock()
	// defer telemetry.MeasureSince(time.Now(), "store", "mvkv", "write_mvs")
	return
}

func (store *VersionIndexedStore) WriteEstimatesToMultiVersionStore() {
	_ = "STUB: not implemented"
	// TODO: remove?
	// store.mtx.Lock()
	// defer store.mtx.Unlock()
	// defer telemetry.MeasureSince(time.Now(), "store", "mvkv", "write_mvs")
	return
}

// TODO: do we need to write readset and iterateset in this case? I don't think so since if this is called it means we aren't doing validation

func (store *VersionIndexedStore) UpdateReadSet(key []byte, value []byte) {
	_ = "STUB: not implemented"
	return
}

// fast path: new key, store value directly (avoids empty slice + append)

// Write implements types.CacheWrap so this store can exist on the cache multi store
func (store *VersionIndexedStore) Write() { _ = "STUB: not implemented"; return }

// GetEvents implements types.CacheWrap so this store can exist on the cache multi store
func (store *VersionIndexedStore) GetEvents() []abci.Event { _ = "STUB: not implemented"; return nil }

// ResetEvents implements types.CacheWrap so this store can exist on the cache multi store
func (store *VersionIndexedStore) ResetEvents() { _ = "STUB: not implemented"; return }

func (store *VersionIndexedStore) UpdateIterateSet(iterationTracker *iterationTracker) {
	_ = "STUB: not implemented"
	// TODO: refactor such that the iterateset is added to the store at the time of iterator creation and updated continuously instead of at Close
	// append to iterateset
	return
}
