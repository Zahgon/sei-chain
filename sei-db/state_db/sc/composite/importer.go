package composite

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ types.Importer = (*SnapshotImporter)(nil)

type SnapshotImporter struct {
	cosmosImporter types.Importer
	flatkvImporter types.Importer
	currentModule  string
}

func NewImporter(cosmosImporter types.Importer, flatkvImporter types.Importer) *SnapshotImporter {
	_ = "STUB: not implemented"
	return nil
}

func (si *SnapshotImporter) AddModule(name string) error { _ = "STUB: not implemented"; return nil }

func (si *SnapshotImporter) AddNode(node *types.SnapshotNode) { _ = "STUB: not implemented"; return }

func (si *SnapshotImporter) Close() error { _ = "STUB: not implemented"; return nil }
