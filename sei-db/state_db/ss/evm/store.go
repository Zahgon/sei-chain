package evm

import (
	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
)

var _ types.StateStore = (*EVMStateStore)(nil)

// EVMStateStore manages either a single MVCC DB for all EVM data or one DB per
// EVM sub-type, depending on config. In both modes, the logical store key and
// key encoding remain unchanged.
type EVMStateStore struct {
	subDBs      map[EVMStoreType]types.StateStore
	managedDBs  []types.StateStore
	dir         string
	separateDBs bool
}

// NewEVMStateStore opens either a single unified MVCC DB for all EVM state
// or one MVCC DB per EVM sub-type.
func NewEVMStateStore(dir string, ssConfig config.StateStoreConfig) (*EVMStateStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func subDBConfig(parent config.StateStoreConfig, dbDir string) config.StateStoreConfig {
	_ = "STUB: not implemented"
	return *new(config.StateStoreConfig)
}

func (s *EVMStateStore) primaryDB() types.StateStore {
	_ = "STUB: not implemented"
	return *new(types.StateStore)
}

func (s *EVMStateStore) routeKey(key []byte) types.StateStore {
	_ = "STUB: not implemented"
	return *new(types.StateStore)
}

func (s *EVMStateStore) Get(_ string, version int64, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *EVMStateStore) Has(_ string, version int64, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *EVMStateStore) Iterator(_ string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (s *EVMStateStore) ReverseIterator(_ string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (s *EVMStateStore) RawIterate(_ string, _ func([]byte, []byte, int64) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *EVMStateStore) GetLatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *EVMStateStore) SetLatestVersion(version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *EVMStateStore) SetEarliestVersion(version int64, ignoreVersion bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) ApplyChangesetSync(version int64, changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) ApplyChangesetAsync(version int64, changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) groupBySubType(changesets []*proto.NamedChangeSet) map[EVMStoreType][]*proto.KVPair {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) applyGrouped(version int64, grouped map[EVMStoreType][]*proto.KVPair, async bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) applyToSubDB(storeType EVMStoreType, version int64, pairs []*proto.KVPair, async bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) Import(version int64, ch <-chan types.SnapshotNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *EVMStateStore) Prune(version int64) error { _ = "STUB: not implemented"; return nil }

func (s *EVMStateStore) Close() error { _ = "STUB: not implemented"; return nil }

func filterEVMChangesets(changesets []*proto.NamedChangeSet) []*proto.NamedChangeSet {
	_ = "STUB: not implemented"
	return nil
}
