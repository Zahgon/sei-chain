package mock

import (
	"io"

	protoio "github.com/gogo/protobuf/io"
	snapshottypes "github.com/sei-protocol/sei-chain/sei-cosmos/snapshots/types"
	store "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
)

var _ sdk.MultiStore = multiStore{}

type multiStore struct {
	kv map[sdk.StoreKey]kvStore
}

func (ms multiStore) RollbackToVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (ms multiStore) CacheMultiStore() sdk.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(sdk.CacheMultiStore)
}

func (ms multiStore) CacheMultiStoreWithVersion(_ int64) (sdk.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(sdk.CacheMultiStore), nil
}

func (ms multiStore) CacheMultiStoreForExport(version int64) (store.CacheMultiStore, error) {
	_ = "STUB: not implemented"
	return *new(store.CacheMultiStore), nil
}

func (ms multiStore) CacheWrap(_ store.StoreKey) sdk.CacheWrap {
	_ = "STUB: not implemented"
	return *new(sdk.CacheWrap)
}

func (ms multiStore) CacheWrapWithTrace(_ store.StoreKey, _ io.Writer, _ sdk.TraceContext) sdk.CacheWrap {
	_ = "STUB: not implemented"
	return *new(sdk.CacheWrap)
}

func (ms multiStore) TracingEnabled() bool { _ = "STUB: not implemented"; return false }

func (ms multiStore) SetTracingContext(tc sdk.TraceContext) sdk.MultiStore {
	_ = "STUB: not implemented"
	return *new(sdk.MultiStore)
}

func (ms multiStore) SetTracer(w io.Writer) sdk.MultiStore {
	_ = "STUB: not implemented"
	return *new(sdk.MultiStore)
}

func (ms multiStore) Commit(_ bool) sdk.CommitID {
	_ = "STUB: not implemented"
	return *new(sdk.CommitID)
}

func (ms multiStore) LastCommitID() sdk.CommitID {
	_ = "STUB: not implemented"
	return *new(sdk.CommitID)
}

func (ms multiStore) SetPruning(opts sdk.PruningOptions) { _ = "STUB: not implemented"; return }

func (ms multiStore) GetPruning() sdk.PruningOptions {
	_ = "STUB: not implemented"
	return *new(sdk.PruningOptions)
}

func (ms multiStore) GetCommitKVStore(key sdk.StoreKey) sdk.CommitKVStore {
	_ = "STUB: not implemented"
	return *new(sdk.CommitKVStore)
}

func (ms multiStore) GetCommitStore(key sdk.StoreKey) sdk.CommitStore {
	_ = "STUB: not implemented"
	return *new(sdk.CommitStore)
}

func (ms multiStore) MountStoreWithDB(key sdk.StoreKey, typ sdk.StoreType, db dbm.DB) {
	_ = "STUB: not implemented"
	return
}

func (ms multiStore) LoadLatestVersion() error { _ = "STUB: not implemented"; return nil }

func (ms multiStore) LoadLatestVersionAndUpgrade(upgrades *store.StoreUpgrades) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms multiStore) LoadVersionAndUpgrade(ver int64, upgrades *store.StoreUpgrades) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms multiStore) LoadVersion(ver int64) error { _ = "STUB: not implemented"; return nil }

func (ms multiStore) GetKVStore(key sdk.StoreKey) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

func (ms multiStore) GetStore(key sdk.StoreKey) sdk.Store {
	_ = "STUB: not implemented"
	return *new(sdk.Store)
}

func (ms multiStore) GetEvents() []abci.Event { _ = "STUB: not implemented"; return nil }

func (ms multiStore) ResetEvents() { _ = "STUB: not implemented"; return }

func (ms multiStore) GetStoreType() sdk.StoreType {
	_ = "STUB: not implemented"
	return *new(sdk.StoreType)
}

func (ms multiStore) SetInterBlockCache(_ sdk.MultiStorePersistentCache) {
	_ = "STUB: not implemented"
	return
}

func (ms multiStore) SetInitialVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (ms multiStore) Snapshot(height uint64, protoWriter protoio.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms multiStore) Restore(
	height uint64, format uint32, protoReader protoio.Reader,
) (snapshottypes.SnapshotItem, error) {
	_ = "STUB: not implemented"
	return *new(snapshottypes.SnapshotItem), nil
}

func (ms multiStore) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var _ sdk.KVStore = kvStore{}

type kvStore struct {
	store map[string][]byte
}

func (kv kvStore) CacheWrap(_ store.StoreKey) sdk.CacheWrap {
	_ = "STUB: not implemented"
	return *new(sdk.CacheWrap)
}

func (kv kvStore) CacheWrapWithTrace(_ store.StoreKey, w io.Writer, tc sdk.TraceContext) sdk.CacheWrap {
	_ = "STUB: not implemented"
	return *new(sdk.CacheWrap)
}

func (kv kvStore) GetStoreType() sdk.StoreType {
	_ = "STUB: not implemented"
	return *new(sdk.StoreType)
}

func (kv kvStore) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (kv kvStore) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (kv kvStore) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

func (kv kvStore) Set(key, value []byte) { _ = "STUB: not implemented"; return }

func (kv kvStore) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (kv kvStore) Prefix(prefix []byte) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

func (kv kvStore) Gas(meter sdk.GasMeter, config sdk.GasConfig) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

func (kv kvStore) Iterator(start, end []byte) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

func (kv kvStore) ReverseIterator(start, end []byte) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

func (kv kvStore) SubspaceIterator(prefix []byte) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

func (kv kvStore) ReverseSubspaceIterator(prefix []byte) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

func (kv kvStore) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

func (kv kvStore) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (kv kvStore) GetAllKeyStrsInRange(start, end []byte) []string {
	_ = "STUB: not implemented"
	return nil
}

func NewCommitMultiStore() sdk.CommitMultiStore {
	_ = "STUB: not implemented"
	return *new(sdk.CommitMultiStore)
}

func (ms multiStore) Close() error { _ = "STUB: not implemented"; return nil }

func (ms multiStore) SetKVStores(handler func(key store.StoreKey, s sdk.KVStore) store.CacheWrap) store.MultiStore {
	_ = "STUB: not implemented"
	return *new(store.MultiStore)
}

func (ms multiStore) StoreKeys() []sdk.StoreKey { _ = "STUB: not implemented"; return nil }

func (ms multiStore) GetEarliestVersion() int64 {
	_ = "STUB: not implemented"
	// TODO implement me
	return 0
}
