package composite

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ types.Exporter = (*SnapshotExporter)(nil)

type exportPhase int

const (
	phaseCosmos exportPhase = iota
	phaseFlatKV
	phaseDone
)

// SnapshotExporter coordinates export from cosmos (memiavl) and flatKV backends.
//
// Next() returns items in stream order. Each item is either:
//   - string: a module name header that starts a new module section
//   - *types.SnapshotNode: a leaf key/value belonging to the current module
//
// FlatKV data is exported as a separate "flatkv" module appended after all
// cosmos modules complete. This keeps the two backends fully independent in the
// snapshot stream.
type SnapshotExporter struct {
	cosmosExporter types.Exporter
	flatkvExporter types.Exporter
	phase          exportPhase
}

// NewExporter creates a composite exporter. cosmosExporter must not be nil.
// flatkvExporter may be nil when FlatKV is not active.
func NewExporter(cosmosExporter types.Exporter, flatkvExporter types.Exporter) (*SnapshotExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Next returns the next item in the composite snapshot stream.
//
// The stream is split into two sequential phases:
//  1. phaseCosmos — drains all items from the cosmos (memiavl) exporter.
//     When the cosmos exporter is exhausted, if a FlatKV exporter is present,
//     the phase transitions to phaseFlatKV and emits the keys.FlatKVStoreKey
//     module header as the first item.
//  2. phaseFlatKV — drains all items from the FlatKV exporter.
//
// Returns ErrorExportDone when both phases are complete.
func (s *SnapshotExporter) Next() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// nextFromCosmos pulls items from the cosmos exporter. On exhaustion, it
// transitions to phaseFlatKV (emitting the module header) or phaseDone.
func (s *SnapshotExporter) nextFromCosmos() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cosmos done. Append flatKV as a separate module.

// nextFromFlatKV pulls items from the FlatKV exporter. On exhaustion, it
// transitions to phaseDone.
func (s *SnapshotExporter) nextFromFlatKV() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SnapshotExporter) Close() error { _ = "STUB: not implemented"; return nil }
