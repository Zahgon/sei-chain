package migration

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/metric"
)

// MigrationMetrics holds OpenTelemetry metrics for a MigrationManager.
// Metrics are exported via whatever exporter is configured on the global
// OTel MeterProvider. All methods are nil-safe so callers (and tests)
// that do not care about metrics can pass a nil *MigrationMetrics to the
// manager.
//
// Unit convention: durations in "s", bytes in "By", counts via curly-brace
// annotations (UCUM, https://ucum.org/ucum).
type MigrationMetrics struct {
	ctx    context.Context
	cancel context.CancelFunc

	// wg tracks the background boundary-snapshot goroutine (if any) so
	// Close can block until it exits.
	wg sync.WaitGroup

	// targetVersion is captured at construction time so the boundary
	// snapshot goroutine can tell, without any DB access, when the
	// migration has completed and it can stop emitting labeled series.
	targetVersion uint64

	keysMigratedTotal       metric.Int64Counter
	keyBytesMigratedTotal   metric.Int64Counter
	valueBytesMigratedTotal metric.Int64Counter
	applyDuration           metric.Float64Histogram
	version                 metric.Int64Gauge
	boundarySnapshot        metric.Int64Gauge

	mu              sync.Mutex
	currentBoundary MigrationBoundary
	currentVersion  uint64
}

// NewMigrationMetrics constructs a MigrationMetrics using the global OTel
// MeterProvider. The caller must have configured the MeterProvider with a
// Prometheus or other exporter before calling this.
//
// targetVersion is the version the associated migration is transitioning
// to; it is used solely by the boundary-snapshot goroutine to decide when
// to stop emitting labeled series.
//
// When boundarySnapshotInterval <= 0 the snapshot goroutine is not started;
// everything else still works. When ctx is cancelled, or Close is called,
// the snapshot goroutine exits.
func NewMigrationMetrics(
	ctx context.Context,
	targetVersion uint64,
	boundarySnapshotInterval time.Duration,
) *MigrationMetrics {
	_ = "STUB: not implemented"
	return nil
}

// SetBoundary updates the in-memory current boundary. No DB access. Safe
// to call concurrently with the snapshot ticker; not safe to call
// concurrently with itself from multiple goroutines, but the
// MigrationManager only updates the boundary from a single ApplyChangeSets
// caller at a time.
func (m *MigrationMetrics) SetBoundary(b MigrationBoundary) { _ = "STUB: not implemented"; return }

// SetVersion updates the in-memory current migration version and records
// the version gauge immediately so Grafana sees the transition without
// waiting for the next snapshot tick.
func (m *MigrationMetrics) SetVersion(v uint64) { _ = "STUB: not implemented"; return }

//nolint:gosec // version is monotonic and bounded

// ReportKeysMigrated records that a batch of (count) keys totaling
// (keyBytes, valueBytes) were migrated in a single ApplyChangeSets call.
// Pass zeros to skip; the method is a no-op on nil receiver.
func (m *MigrationMetrics) ReportKeysMigrated(count int64, keyBytes int64, valueBytes int64) {
	_ = "STUB: not implemented"
	return
}

// RecordApplyDuration records the wall-clock time spent in a single
// ApplyChangeSets call. Invoked from a defer in the manager so both
// success and error paths are captured.
func (m *MigrationMetrics) RecordApplyDuration(d time.Duration) { _ = "STUB: not implemented"; return }

// snapshot returns a safe copy of the in-memory boundary and version
// under the mutex. The returned boundary shares its internal key slice
// with the stored value; callers must not mutate it.
func (m *MigrationMetrics) snapshot() (MigrationBoundary, uint64) {
	_ = "STUB: not implemented"
	return *new(MigrationBoundary), 0
}

// startBoundarySnapshotLoop starts the background goroutine that
// periodically emits the labeled boundary snapshot gauge. The loop exits
// on ctx cancellation, or after emitting a single "complete" sentinel
// once currentVersion reaches targetVersion.
//
// Cardinality rationale: at a 10-minute interval a month-long migration
// tops out at ~4k unique boundary_hex label values — well within
// Prometheus' comfort zone — and the OTel exporter's staleness markers
// keep only the most recent label active in the scrape set.
func (m *MigrationMetrics) startBoundarySnapshotLoop(interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Release resources held by the metrics collector.
func (m *MigrationMetrics) Close() { _ = "STUB: not implemented"; return }

// recordBoundarySnapshot emits the labeled snapshot gauge with value 1.
// The label is the only payload — the value itself is unused.
func (m *MigrationMetrics) recordBoundarySnapshot(label string) { _ = "STUB: not implemented"; return }
