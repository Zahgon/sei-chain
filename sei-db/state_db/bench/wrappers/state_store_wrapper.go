package wrappers

import (
	"sync/atomic"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	dbTypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	scTypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ DBWrapper = (*stateStoreWrapper)(nil)

// stateStoreWrapper adapts a versioned StateStore (SS layer) to the DBWrapper
// interface used by the cryptosim benchmark. Each ApplyChangeSets call maps to
// a single ApplyChangesetAsync at the benchmark-provided version. The SS layer
// persists on every apply, so Commit is a no-op.
type stateStoreWrapper struct {
	base    dbTypes.StateStore
	version atomic.Int64
}

func NewStateStoreWrapper(store dbTypes.StateStore) DBWrapper {
	_ = "STUB: not implemented"
	return *new(DBWrapper)
}

func (s *stateStoreWrapper) ApplyChangeSets(entry *proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stateStoreWrapper) Read(key []byte) (data []byte, found bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (s *stateStoreWrapper) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stateStoreWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (s *stateStoreWrapper) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (s *stateStoreWrapper) LoadVersion(_ int64) error { _ = "STUB: not implemented"; return nil }

func (s *stateStoreWrapper) Importer(_ int64) (scTypes.Importer, error) {
	_ = "STUB: not implemented"
	return *new(scTypes.Importer), nil
}

func (s *stateStoreWrapper) GetPhaseTimer() *metrics.PhaseTimer {
	_ = "STUB: not implemented"
	return nil
}
