package parquet

import (
	"context"
	"os"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/parquet-go/parquet-go"
	dbwal "github.com/sei-protocol/sei-chain/sei-db/wal"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("db", "ledger-db", "parquet")

const (
	maxInt64  = int64(^uint64(0) >> 1)
	maxUint32 = ^uint32(0)

	defaultBlockFlushInterval uint64 = 1
	defaultMaxBlocksPerFile   uint64 = 500
)

var removeFile = os.Remove

// StoreConfig configures the parquet store.
type StoreConfig struct {
	DBDirectory          string
	KeepRecent           int64
	PruneIntervalSeconds int64
	BlockFlushInterval   uint64
	MaxBlocksPerFile     uint64
	TxIndexBackend       string
}

// DefaultStoreConfig returns the default store configuration.
func DefaultStoreConfig() StoreConfig { _ = "STUB: not implemented"; return *new(StoreConfig) }

// ReceiptInput is the input for storing a receipt.
type ReceiptInput struct {
	BlockNumber  uint64
	Receipt      ReceiptRecord
	Logs         []LogRecord
	ReceiptBytes []byte // For WAL
}

// FaultHooks provides optional hook points for fault injection in tests.
// All fields are nil in production. When non-nil, the hook is called at
// the corresponding point in the write path; returning a non-nil error
// aborts the operation and propagates the error to the caller.
type FaultHooks struct {
	AfterWALWrite     func(blockNumber uint64) error // after WAL writes, before parquet apply
	BeforeFlush       func(blockNumber uint64) error // before writing buffers to parquet
	AfterFlush        func(blockNumber uint64) error // after parquet flush, before buffer clear
	AfterCloseWriters func(blockNumber uint64) error // during rotation, after closing old writers
	AfterWALClear     func(blockNumber uint64) error // during rotation, after WAL truncation
}

// Store is the parquet-based receipt store.
type Store struct {
	basePath      string
	receiptWriter *parquet.GenericWriter[ReceiptRecord]
	logWriter     *parquet.GenericWriter[LogRecord]
	receiptFile   *os.File
	logFile       *os.File

	mu               sync.Mutex
	fileStartBlock   uint64
	receiptsBuffer   []ReceiptRecord
	logsBuffer       []LogRecord
	config           StoreConfig
	lastSeenBlock    uint64
	blocksSinceFlush uint64

	Reader          *Reader
	wal             dbwal.GenericWAL[WALEntry]
	latestVersion   atomic.Int64
	earliestVersion atomic.Int64
	closeOnce       sync.Once

	pruneStop chan struct{}

	// WarmupRecords holds receipts recovered from WAL for cache warming.
	WarmupRecords []ReceiptRecord

	// FaultHooks is nil in production. Tests can set this to inject faults
	// at specific points in the write path for crash recovery testing.
	FaultHooks *FaultHooks
}

// NewStore creates a new parquet store.
func NewStore(cfg StoreConfig) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

func resolveStoreConfig(cfg StoreConfig) StoreConfig {
	_ = "STUB: not implemented"
	return *new(StoreConfig)
}

// LatestVersion returns the latest version stored.
func (s *Store) LatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

// SetLatestVersion sets the latest version.
func (s *Store) SetLatestVersion(version int64) { _ = "STUB: not implemented"; return }

// SetEarliestVersion sets the earliest version.
func (s *Store) SetEarliestVersion(version int64) { _ = "STUB: not implemented"; return }

// CacheRotateInterval returns the interval at which the cache should rotate.
func (s *Store) CacheRotateInterval() uint64 { _ = "STUB: not implemented"; return 0 }

// SetBlockFlushInterval overrides the number of blocks buffered before
// flushing to a parquet file. Intended for testing.
func (s *Store) SetBlockFlushInterval(interval uint64) { _ = "STUB: not implemented"; return }

// SetMaxBlocksPerFile overrides the rotation interval after construction.
// Intended for tests that need a small boundary so they can exercise rotation
// behavior without writing hundreds of blocks. Not safe to call while writes
// are in flight (rotation / WAL invariants may disagree with the reader until
// the store is quiesced). Concurrent reads remain race-safe under the race
// detector because the reader field is updated under Reader.mu.
func (s *Store) SetMaxBlocksPerFile(n uint64) { _ = "STUB: not implemented"; return }

// GetReceiptByTxHash retrieves a receipt by transaction hash via a full scan of
// the closed parquet files tracked by the reader.
func (s *Store) GetReceiptByTxHash(ctx context.Context, txHash common.Hash) (*ReceiptResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetReceiptByTxHashInBlock narrows the parquet search to the file containing
// blockNumber, falling back to a full scan on miss.
func (s *Store) GetReceiptByTxHashInBlock(ctx context.Context, txHash common.Hash, blockNumber uint64) (*ReceiptResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLogs retrieves logs matching the filter.
func (s *Store) GetLogs(ctx context.Context, filter LogFilter) ([]LogResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteReceipts writes multiple receipts, batching WAL writes per block.
//
// Rotation fires on aligned block boundaries (blockNumber % MaxBlocksPerFile == 0).
// WAL.Write must run before rotateFileLocked: rotation clears the WAL, so a
// crash between clear and a later WAL write would lose the boundary block.
// ClearWAL preserves the last entry so the just-written boundary entry
// survives the clear and remains replayable.
func (s *Store) WriteReceipts(inputs []ReceiptInput) error { _ = "STUB: not implemented"; return nil }

// Group receipts by block number, preserving encounter order.

// IsRotationBoundary returns true when blockNumber is aligned to the file
// rotation interval (MaxBlocksPerFile). These are the block numbers that start
// a new parquet file.
func (s *Store) IsRotationBoundary(blockNumber uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// UpdateLatestVersion updates the latest version if the new value is higher.
func (s *Store) UpdateLatestVersion(version int64) { _ = "STUB: not implemented"; return }

// SimulateCrash abandons the store without flushing or finalizing, mimicking
// an abrupt process termination (e.g. kill -9 or power loss). Specifically
// it skips:
//   - flushLocked(): in-memory buffered receipts are lost
//   - receiptWriter.Close() / logWriter.Close(): parquet footers are never
//     written, leaving on-disk files corrupt/unreadable
//   - receiptFile.Sync() / logFile.Sync(): OS-buffered writes may be lost
//
// The raw os.File.Close() and wal.Close() calls below exist only to release
// file descriptors and locks so the test process can reopen the same directory.
func (s *Store) SimulateCrash() { _ = "STUB: not implemented"; return }

// Close closes the store.
func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }

// WAL returns the WAL for replay purposes.
func (s *Store) WAL() dbwal.GenericWAL[WALEntry] {
	_ = "STUB: not implemented"

	// ApplyReceiptFromReplay applies a receipt during WAL replay. If the block
	// number is on a rotation boundary, this rotates the file (without touching the
	// WAL) so replay-recovered blocks land in the same aligned files the write path
	// would have produced. Skipping WAL truncation here is mandatory: the caller is
	// iterating WAL offsets, and truncating mid-iteration would break the scan.
	return nil
}

func (s *Store) ApplyReceiptFromReplay(input ReceiptInput) error {
	_ = "STUB: not implemented"
	return nil
}

// ObserveEmptyBlock signals that a block with no receipts was committed at
// height. WriteReceipts is the normal place rotation fires, but it is skipped
// for empty blocks — so without this hook a boundary-aligned empty block would
// leave the open file accepting writes past MaxBlocksPerFile and break the
// reader's file-pruning logic (which assumes each file spans at most that many
// blocks). Callers that bump LatestVersion for empty blocks should invoke this
// so the rotation invariant stays in lockstep with the chain.
func (s *Store) ObserveEmptyBlock(height uint64) error { _ = "STUB: not implemented"; return nil }

// Only advance lastSeenBlock for strictly greater heights. Out-of-order
// observations must not move the cursor backward, or WriteReceipts could
// mis-handle rotation for blocks already seen.

// No file to rotate yet, or the empty block is not on a boundary.
// Still track it so a later observation of the same height is a no-op.

// FileStartBlock returns the current file start block.
func (s *Store) FileStartBlock() uint64 { _ = "STUB: not implemented"; return 0 }

// ClearWAL truncates the WAL after rotation, preserving the last entry.
// WriteReceipts writes the boundary block's WAL entry before calling rotate,
// and that entry's data has not yet been applied to the new file — losing it
// would lose the block. ObserveEmptyBlock's path has no pending entry, so
// the preserved entry is redundant (already in the closed file) but harmless:
// its blockNumber is < the new fileStartBlock, so replay drops it.
func (s *Store) ClearWAL() error { _ = "STUB: not implemented"; return nil }

func (s *Store) startPruning(pruneIntervalSeconds int64) { _ = "STUB: not implemented"; return }

// Add random jitter (up to 50% of base interval) to avoid thundering herd

// Continue to next iteration

// PruneOldFiles removes parquet file pairs whose data is entirely before
// pruneBeforeBlock. Returns the number of file pairs removed.
func (s *Store) PruneOldFiles(pruneBeforeBlock uint64) int {
	_ = "STUB: not implemented"
	// Get list of files to prune from the reader
	return 0
}

// Step 1: Remove from tracking (brief mu.Lock) so new reader
// snapshots won't include these files.

// Step 2: Wait for in-flight readers to finish, then delete.
// pruneMu.Lock blocks until all current pruneMu.RLock holders
// (active queries) release.

// Re-add to tracking if deletion failed (outside pruneMu to
// avoid holding both locks).

// alignedFileStartBlock returns the parquet filename start block for lazy init:
// the greatest multiple of maxBlocksPerFile not above blockNumber, or blockNumber
// when maxBlocksPerFile is zero (rotation disabled).
func alignedFileStartBlock(blockNumber, maxBlocksPerFile uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (s *Store) applyReceiptLocked(input ReceiptInput) error {
	_ = "STUB: not implemented"
	// Lazy writer initialization: defer file creation until the first receipt
	// arrives. The filename start block is normally snapped down to the
	// rotation interval so it matches Reader assumptions
	// ([start, start+MaxBlocksPerFile)) for pruning and file-range logic.
	// On reopen NewStore pre-sets fileStartBlock to maxBlock+1; if the aligned
	// start falls inside that same rotation window we must keep the preset
	// value, otherwise initWriters would os.Create (and truncate) the existing
	// closed parquet file that still holds the last committed blocks.
	return nil
}

// rotateFileLocked closes the current parquet file, clears older WAL entries
// (ClearWAL keeps the last entry — see WriteReceipts / ClearWAL docs), and opens
// a new file at newBlockNumber. Callers must write newBlockNumber's WAL entry
// before invoking this so rotation never drops the boundary block from the WAL.
func (s *Store) rotateFileLocked(newBlockNumber uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// rotateFileLockedNoWAL performs the file-level rotation without touching the
// WAL. Used during WAL replay where the outer scan would break if entries were
// truncated mid-iteration.
func (s *Store) rotateFileLockedNoWAL(newBlockNumber uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// initWriters must come AFTER fileStartBlock is updated so the new file
// name reflects the new aligned boundary.

// Pending buffer data was flushed into the closed file; nothing carries
// over to the new writer, so reset the flush counter too.

func (s *Store) initWriters() error { _ = "STUB: not implemented"; return nil }

// #nosec G304 -- paths are constructed from configured base directory

// #nosec G304 -- paths are constructed from configured base directory

// Flush acquires the write lock and flushes all buffered data to disk.
// Mostly used for testing and benchmarking.
func (s *Store) Flush() error { _ = "STUB: not implemented"; return nil }

func (s *Store) flushLocked() error { _ = "STUB: not implemented"; return nil }

func (s *Store) closeWritersLocked() error { _ = "STUB: not implemented"; return nil }

func int64FromUint64(value uint64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint32FromUint safely converts uint to uint32.
func Uint32FromUint(value uint) uint32 { _ = "STUB: not implemented"; return 0 }

// CopyBytes creates a copy of a byte slice.
func CopyBytes(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// CopyBytesOrEmpty creates a copy of a byte slice, returning empty slice for nil.
func CopyBytesOrEmpty(src []byte) []byte { _ = "STUB: not implemented"; return nil }
