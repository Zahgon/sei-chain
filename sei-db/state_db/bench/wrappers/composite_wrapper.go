package wrappers

import (
	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/composite"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ DBWrapper = (*compositeWrapper)(nil)

// compositeWrapper wraps a composite commit store to implement the DBWrapper interface.
type compositeWrapper struct {
	base *composite.CompositeCommitStore
}

// NewCompositeWrapper creates a new compositeWrapper with a given composite commit store.
func NewCompositeWrapper(store *composite.CompositeCommitStore) DBWrapper {
	_ = "STUB: not implemented"
	return *new(DBWrapper)
}

func (c *compositeWrapper) ApplyChangeSets(entry *proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *compositeWrapper) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *compositeWrapper) LoadVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (c *compositeWrapper) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (c *compositeWrapper) Importer(version int64) (types.Importer, error) {
	_ = "STUB: not implemented"
	return *new(types.Importer), nil
}

func (c *compositeWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (c *compositeWrapper) Read(key []byte) (data []byte, found bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *compositeWrapper) GetPhaseTimer() *metrics.PhaseTimer {
	_ = "STUB: not implemented"
	return nil
}
