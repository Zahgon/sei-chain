package cosmos

import (
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
)

// Compile-time check: CosmosStateStore implements db_engine.StateStore.
var _ types.StateStore = (*CosmosStateStore)(nil)

// CosmosStateStore wraps a single StateStore (MVCC DB) and satisfies db_engine.StateStore.
// It is the SS-layer adapter for the main Cosmos state (all non-EVM modules).
type CosmosStateStore struct {
	db types.StateStore
}

// NewCosmosStateStore wraps an existing StateStore as a CosmosStateStore.
func NewCosmosStateStore(db types.StateStore) types.StateStore {
	_ = "STUB: not implemented"
	return *new(types.StateStore)
}

func (s *CosmosStateStore) Get(storeKey string, version int64, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CosmosStateStore) Has(storeKey string, version int64, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *CosmosStateStore) Iterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (s *CosmosStateStore) ReverseIterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (s *CosmosStateStore) RawIterate(storeKey string, fn func([]byte, []byte, int64) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *CosmosStateStore) GetLatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *CosmosStateStore) SetLatestVersion(version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CosmosStateStore) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *CosmosStateStore) SetEarliestVersion(version int64, ignoreVersion bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CosmosStateStore) ApplyChangesetSync(version int64, changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CosmosStateStore) ApplyChangesetAsync(version int64, changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CosmosStateStore) Prune(version int64) error { _ = "STUB: not implemented"; return nil }

func (s *CosmosStateStore) Import(version int64, ch <-chan types.SnapshotNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *CosmosStateStore) Close() error { _ = "STUB: not implemented"; return nil }
