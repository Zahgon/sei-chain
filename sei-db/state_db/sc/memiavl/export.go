package memiavl

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

// exportBufferSize is the number of nodes to buffer in the exporter. It improves throughput by
// processing multiple nodes per context switch, but take care to avoid excessive memory usage,
// especially since callers may export several IAVL stores in parallel (e.g. the Cosmos SDK).
const exportBufferSize = 32

type MultiTreeExporter struct {
	// only one of them is non-nil
	db    *DB
	mtree *MultiTree

	iTree    int
	exporter *Exporter
}

func NewMultiTreeExporter(dir string, version uint32, onlyAllowExportOnSnapshotVersion bool) (exporter *MultiTreeExporter, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mte *MultiTreeExporter) trees() []NamedTree { _ = "STUB: not implemented"; return nil }

func (mte *MultiTreeExporter) Next() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mte *MultiTreeExporter) Close() error { _ = "STUB: not implemented"; return nil }

type exportWorker func(callback func(*types.SnapshotNode) bool)

type Exporter struct {
	ch     <-chan *types.SnapshotNode
	cancel context.CancelFunc
}

func newExporter(worker exportWorker) *Exporter { _ = "STUB: not implemented"; return nil }

func (e *Exporter) Next() (*types.SnapshotNode, error) { _ = "STUB: not implemented"; return nil, nil }

// Close closes the exporter. It is safe to call multiple times.
func (e *Exporter) Close() { _ = "STUB: not implemented"; return }

// drain channel
