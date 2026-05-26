package wrappers

import (
	"sync/atomic"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	dbTypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	scTypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ DBWrapper = (*combinedWrapper)(nil)

// combinedWrapper drives both a State Commit (SC) and State Store (SS) backend
// from the same changeset stream, mirroring production where SC and SS receive
// identical writes.
type combinedWrapper struct {
	sc        DBWrapper
	ss        dbTypes.StateStore
	ssVersion atomic.Int64
}

func NewCombinedWrapper(sc DBWrapper, ss dbTypes.StateStore) DBWrapper {
	_ = "STUB: not implemented"
	return *new(DBWrapper)
}

func (c *combinedWrapper) ApplyChangeSets(entry *proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *combinedWrapper) Read(key []byte) (data []byte, found bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *combinedWrapper) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *combinedWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (c *combinedWrapper) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (c *combinedWrapper) LoadVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (c *combinedWrapper) Importer(version int64) (scTypes.Importer, error) {
	_ = "STUB: not implemented"
	return *new(scTypes.Importer), nil
}

func (c *combinedWrapper) GetPhaseTimer() *metrics.PhaseTimer {
	_ = "STUB: not implemented"
	return nil
}
