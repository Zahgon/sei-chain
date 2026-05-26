package wrappers

import (
	"sync/atomic"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	scTypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ DBWrapper = (*noOpWrapper)(nil)

// noOpWrapper lets the benchmark measure its own overhead without DB read/write cost.
type noOpWrapper struct {
	version atomic.Int64
}

func NewNoOpWrapper() DBWrapper { _ = "STUB: not implemented"; return *new(DBWrapper) }

func (n *noOpWrapper) ApplyChangeSets(entry *proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *noOpWrapper) Read(_ []byte) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (n *noOpWrapper) Commit() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *noOpWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (n *noOpWrapper) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (n *noOpWrapper) LoadVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (n *noOpWrapper) Importer(_ int64) (scTypes.Importer, error) {
	_ = "STUB: not implemented"
	return *new(scTypes.Importer), nil
}

func (n *noOpWrapper) GetPhaseTimer() *metrics.PhaseTimer { _ = "STUB: not implemented"; return nil }
