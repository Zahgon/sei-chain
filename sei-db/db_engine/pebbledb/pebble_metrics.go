package pebbledb

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/metric"

	"github.com/cockroachdb/pebble/v2"
)

const pebbleMeterName = "seidb_pebble"

// PebbleMetrics scrapes metrics from a Pebble DB and records them via OTel instruments.
// Instrument names match sei-db/db_engine/pebbledb/mvcc for dashboard compatibility.
// The databaseName is used as the "db" attribute on all recorded metrics.
//
// Multiple instances are safe: OTel instrument registration is idempotent, so each
// NewPebbleMetrics call receives references to the same underlying instruments.
// The "db" attribute distinguishes series (e.g. pebble_compaction_count{db="state"}).
type PebbleMetrics struct {
	db           *pebble.DB
	databaseName string

	getLatency                 metric.Float64Histogram
	applyChangesetLatency      metric.Float64Histogram
	applyChangesetAsyncLatency metric.Float64Histogram
	pruneLatency               metric.Float64Histogram
	importLatency              metric.Float64Histogram
	batchWriteLatency          metric.Float64Histogram

	compactionCount                 metric.Int64Counter
	compactionDuration              metric.Float64Histogram
	compactionBytesRead             metric.Int64Counter
	compactionBytesWritten          metric.Int64Counter
	compactionEstimatedDebt         metric.Int64Gauge
	compactionInProgressBytes       metric.Int64Gauge
	compactionNumInProgress         metric.Int64Gauge
	compactionCancelledCount        metric.Int64Counter
	compactionCancelledBytes        metric.Int64Counter
	compactionFailedCount           metric.Int64Counter
	compactionDefaultCount          metric.Int64Counter
	compactionDeleteOnlyCount       metric.Int64Counter
	compactionElisionOnlyCount      metric.Int64Counter
	compactionCopyCount             metric.Int64Counter
	compactionMoveCount             metric.Int64Counter
	compactionReadCount             metric.Int64Counter
	compactionTombstoneDensityCount metric.Int64Counter
	compactionRewriteCount          metric.Int64Counter
	compactionMultiLevelCount       metric.Int64Counter
	compactionBlobFileRewriteCount  metric.Int64Counter
	compactionCounterLevelCount     metric.Int64Counter
	compactionNumProblemSpans       metric.Int64Gauge
	compactionMarkedFiles           metric.Int64Gauge

	ingestCount metric.Int64Counter

	flushCount              metric.Int64Counter
	flushDuration           metric.Float64Histogram
	flushBytesWritten       metric.Int64Counter
	flushNumInProgress      metric.Int64Gauge
	flushAsIngestCount      metric.Int64Counter
	flushAsIngestTableCount metric.Int64Counter
	flushAsIngestBytes      metric.Int64Counter
	flushIdleDuration       metric.Float64Gauge

	filterHits   metric.Int64Counter
	filterMisses metric.Int64Counter

	sstableCount                   metric.Int64Gauge
	sstableTotalSize               metric.Int64Gauge
	sstableSublevels               metric.Int64Gauge
	sstableScore                   metric.Float64Gauge
	sstableFillFactor              metric.Float64Gauge
	sstableVirtualCount            metric.Int64Gauge
	sstableVirtualSize             metric.Int64Gauge
	sstableBytesIngested           metric.Int64Counter
	sstableBytesMoved              metric.Int64Counter
	sstableBytesRead               metric.Int64Counter
	sstableBytesFlushed            metric.Int64Counter
	sstableTablesCompacted         metric.Int64Counter
	sstableTablesFlushed           metric.Int64Counter
	sstableTablesIngested          metric.Int64Counter
	sstableTablesMoved             metric.Int64Counter
	sstableCompensatedFillFactor   metric.Float64Gauge
	sstableEstimatedReferencesSize metric.Int64Gauge
	sstableTablesDeleted           metric.Int64Counter
	sstableTablesExcised           metric.Int64Counter
	sstableBlobBytesReadEstimate   metric.Int64Counter
	sstableBlobBytesCompacted      metric.Int64Counter
	sstableBlobBytesFlushed        metric.Int64Counter
	sstableMultiLevelBytesInTop    metric.Int64Counter
	sstableMultiLevelBytesIn       metric.Int64Counter
	sstableMultiLevelBytesRead     metric.Int64Counter
	sstableValueBlocksSize         metric.Int64Gauge
	sstableBytesWrittenDataBlocks  metric.Int64Counter
	sstableBytesWrittenValueBlocks metric.Int64Counter

	memtableCount       metric.Int64Gauge
	memtableTotalSize   metric.Int64Gauge
	memtableZombieSize  metric.Int64Gauge
	memtableZombieCount metric.Int64Gauge

	walSize                 metric.Int64Gauge
	walFiles                metric.Int64Gauge
	walObsoleteFiles        metric.Int64Gauge
	walObsoletePhysicalSize metric.Int64Gauge
	walPhysicalSize         metric.Int64Gauge
	walBytesIn              metric.Int64Counter
	walBytesWritten         metric.Int64Counter

	tableObsoleteSize                  metric.Int64Gauge
	tableObsoleteCount                 metric.Int64Gauge
	tableZombieSize                    metric.Int64Gauge
	tableZombieCount                   metric.Int64Gauge
	tableLiveSize                      metric.Int64Gauge
	tableLiveCount                     metric.Int64Gauge
	tableBackingCount                  metric.Int64Gauge
	tableBackingSize                   metric.Int64Gauge
	tableCompressedUnknown             metric.Int64Gauge
	tableCompressedSnappy              metric.Int64Gauge
	tableCompressedZstd                metric.Int64Gauge
	tableCompressedMinLZ               metric.Int64Gauge
	tableCompressedNone                metric.Int64Gauge
	tableLocalObsoleteSize             metric.Int64Gauge
	tableLocalObsoleteCount            metric.Int64Gauge
	tableLocalZombieSize               metric.Int64Gauge
	tableLocalZombieCount              metric.Int64Gauge
	tableGarbagePointDeletionsEstimate metric.Int64Gauge
	tableGarbageRangeDeletionsEstimate metric.Int64Gauge
	tableInitialStatsComplete          metric.Int64Gauge
	tablePendingStatsCount             metric.Int64Gauge

	blobFilesLiveCount           metric.Int64Gauge
	blobFilesLiveSize            metric.Int64Gauge
	blobFilesValueSize           metric.Int64Gauge
	blobFilesReferencedValueSize metric.Int64Gauge
	blobFilesObsoleteCount       metric.Int64Gauge
	blobFilesObsoleteSize        metric.Int64Gauge
	blobFilesZombieCount         metric.Int64Gauge
	blobFilesZombieSize          metric.Int64Gauge
	blobFilesLocalLiveSize       metric.Int64Gauge
	blobFilesLocalLiveCount      metric.Int64Gauge
	blobFilesLocalObsoleteSize   metric.Int64Gauge
	blobFilesLocalObsoleteCount  metric.Int64Gauge
	blobFilesLocalZombieSize     metric.Int64Gauge
	blobFilesLocalZombieCount    metric.Int64Gauge

	fileCacheSize          metric.Int64Gauge
	fileCacheTableCount    metric.Int64Gauge
	fileCacheBlobFileCount metric.Int64Gauge
	fileCacheHits          metric.Int64Counter
	fileCacheMisses        metric.Int64Counter

	// prev* track last scraped cumulative values so we Add(delta) not Add(total).
	prevCompactionCount                 int64
	prevCompactionCancelledCount        int64
	prevCompactionCancelledBytes        int64
	prevCompactionFailedCount           int64
	prevCompactionDefaultCount          int64
	prevCompactionDeleteOnlyCount       int64
	prevCompactionElisionOnlyCount      int64
	prevCompactionCopyCount             int64
	prevCompactionMoveCount             int64
	prevCompactionReadCount             int64
	prevCompactionTombstoneDensityCount int64
	prevCompactionRewriteCount          int64
	prevCompactionMultiLevelCount       int64
	prevCompactionBlobFileRewriteCount  int64
	prevCompactionCounterLevelCount     int64
	prevIngestCount                     int64
	prevFlushCount                      int64
	prevFlushBytesWritten               int64
	prevFlushAsIngestCount              int64
	prevFlushAsIngestTableCount         int64
	prevFlushAsIngestBytes              int64
	prevFilterHits                      int64
	prevFilterMisses                    int64
	prevWalBytesIn                      int64
	prevWalBytesWritten                 int64
	prevWalFailoverDirSwitchCount       int64
	prevKeysMissizedTombstonesCount     int64
	prevSnapshotPinnedKeys              int64
	prevSnapshotPinnedSize              int64
	prevFileCacheHits                   int64
	prevFileCacheMisses                 int64
	prevCacheHits                       int64
	prevCacheMisses                     int64

	// prev*ByLevel hold previous cumulative values per level (index = level).
	prevCompactionBytesReadByLevel            []int64
	prevCompactionBytesWrittenByLevel         []int64
	prevSstableBytesIngestedByLevel           []int64
	prevSstableBytesMovedByLevel              []int64
	prevSstableBytesReadByLevel               []int64
	prevSstableBytesFlushedByLevel            []int64
	prevSstableTablesCompactedByLevel         []int64
	prevSstableTablesFlushedByLevel           []int64
	prevSstableTablesIngestedByLevel          []int64
	prevSstableTablesMovedByLevel             []int64
	prevSstableTablesDeletedByLevel           []int64
	prevSstableTablesExcisedByLevel           []int64
	prevSstableBlobBytesReadEstimateByLevel   []int64
	prevSstableBlobBytesCompactedByLevel      []int64
	prevSstableBlobBytesFlushedByLevel        []int64
	prevSstableMultiLevelBytesInTopByLevel    []int64
	prevSstableMultiLevelBytesInByLevel       []int64
	prevSstableMultiLevelBytesReadByLevel     []int64
	prevSstableBytesWrittenDataBlocksByLevel  []int64
	prevSstableBytesWrittenValueBlocksByLevel []int64

	walFailoverDirSwitchCount    metric.Int64Counter
	walFailoverPrimaryDuration   metric.Float64Gauge
	walFailoverSecondaryDuration metric.Float64Gauge

	numVirtual        metric.Int64Gauge
	virtualSize       metric.Int64Gauge
	remoteTablesCount metric.Int64Gauge
	remoteTablesSize  metric.Int64Gauge

	keysRangeKeySetsCount       metric.Int64Gauge
	keysTombstoneCount          metric.Int64Gauge
	keysMissizedTombstonesCount metric.Int64Counter

	snapshotCount          metric.Int64Gauge
	snapshotPinnedKeys     metric.Int64Counter
	snapshotPinnedSize     metric.Int64Counter
	snapshotEarliestSeqNum metric.Int64Gauge

	tableIters     metric.Int64Gauge
	uptimeSeconds  metric.Float64Gauge
	readAmp        metric.Int64Gauge
	diskSpaceUsage metric.Int64Gauge

	cacheHits   metric.Int64Counter
	cacheMisses metric.Int64Counter
	cacheSize   metric.Int64Gauge

	batchSize                metric.Int64Histogram
	pendingChangesQueueDepth metric.Int64Gauge
	iteratorIterations       metric.Float64Histogram
}

// NewPebbleMetrics creates a PebbleMetrics that scrapes metrics from the given Pebble DB
// and records them to OTel. A background goroutine runs every scrapeInterval until
// ctx is cancelled. The databaseName is attached as the "db" attribute to all recorded
// metrics, enabling multi-DB setups to distinguish series in Prometheus/Grafana.
//
// Multiple instances (e.g. one per DB) are safe: OTel returns the same instruments
// for duplicate registrations, and the "db" attribute separates series.
func NewPebbleMetrics(
	ctx context.Context,
	db *pebble.DB,
	databaseName string,
	scrapeInterval time.Duration,
) *PebbleMetrics {
	_ = "STUB: not implemented"
	return nil
}

// collectLoop runs a ticker that periodically calls recordFromPebble. It exits when ctx is cancelled.
func (pm *PebbleMetrics) collectLoop(ctx context.Context, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func uint64ToInt64Clamped(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

// addDelta computes the difference between current and prev, updates prev to current,
// and adds the positive delta to the counter. Used to convert cumulative scraped
// values into rate/counter increments.
func addDelta(ctx context.Context, counter metric.Int64Counter, current int64, prev *int64, opts ...metric.AddOption) {
	_ = "STUB: not implemented"
	return
}

// recordFromPebble fetches the current metrics from the Pebble DB via Metrics(), then
// records compaction, flush, level, memtable, WAL, and cache metrics with the configured
// database name as the "db" attribute.
func (pm *PebbleMetrics) recordFromPebble(ctx context.Context) { _ = "STUB: not implemented"; return }

// Grow prev slices if needed.
