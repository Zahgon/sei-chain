package wrappers

import (
	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/memiavl"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ DBWrapper = (*memIAVLWrapper)(nil)

// A light wrapper around a memiavl commit store to implement the DBWrapper interface.
type memIAVLWrapper struct {
	base *memiavl.CommitStore
}

// NewMemIAVLWrapper creates a new memIAVLWrapper with a given memiavl commit store.
func NewMemIAVLWrapper(commitStore *memiavl.CommitStore) DBWrapper {
	_ = "STUB: not implemented"
	return *new(DBWrapper)
}

func (m *memIAVLWrapper) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *memIAVLWrapper) LoadVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (m *memIAVLWrapper) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (m *memIAVLWrapper) ApplyChangeSets(entry *proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memIAVLWrapper) Importer(version int64) (types.Importer, error) {
	_ = "STUB: not implemented"
	// Close DB first to release lock
	return *new(types.Importer), nil
}

func (m *memIAVLWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (m *memIAVLWrapper) Read(key []byte) (data []byte, found bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *memIAVLWrapper) GetPhaseTimer() *metrics.PhaseTimer { _ = "STUB: not implemented"; return nil }
