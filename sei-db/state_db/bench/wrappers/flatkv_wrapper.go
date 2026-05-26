package wrappers

import (
	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ DBWrapper = (*flatKVWrapper)(nil)

// flatKVWrapper wraps a flatkv commit store to implement the DBWrapper interface.
// Version() returns the working version (committedVersion+1) when there are pending
// changes, matching memiavl's WorkingCommitInfo().Version semantics for benchmarks.
type flatKVWrapper struct {
	base       flatkv.Store
	hasPending bool
}

// NewFlatKVWrapper creates a new flatKVWrapper with a given flatkv store.
func NewFlatKVWrapper(store flatkv.Store) DBWrapper {
	_ = "STUB: not implemented"
	return *new(DBWrapper)
}

func (f *flatKVWrapper) ApplyChangeSets(entry *proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *flatKVWrapper) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *flatKVWrapper) LoadVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (f *flatKVWrapper) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (f *flatKVWrapper) Importer(version int64) (types.Importer, error) {
	_ = "STUB: not implemented"
	return *new(types.Importer), nil
}

func (f *flatKVWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (f *flatKVWrapper) Read(key []byte) (data []byte, found bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (f *flatKVWrapper) GetPhaseTimer() *metrics.PhaseTimer { _ = "STUB: not implemented"; return nil }
