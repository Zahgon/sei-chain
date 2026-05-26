package receipt

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	dbconfig "github.com/sei-protocol/sei-chain/sei-db/config"
	seidbtypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("db", "ledger-db", "receipt")

// Sentinel errors for consistent error checking.
var (
	ErrNotFound               = errors.New("receipt not found")
	ErrNotConfigured          = errors.New("receipt store not configured")
	ErrRangeQueryNotSupported = errors.New("range query not supported by this backend")
	// ErrTxIndexDisabled indicates that a receipt-by-tx-hash lookup missed the
	// in-memory cache and cannot be served because the parquet backend's pebble
	// tx hash index is disabled. A full parquet scan would require reading every
	// file on disk and is intentionally not attempted. The error wraps
	// ErrNotFound so callers that treat "not found" as a null/nil result (e.g.
	// eth_getTransactionReceipt) continue to behave correctly; the wrapping
	// preserves the underlying reason for operators and tests via errors.Is.
	ErrTxIndexDisabled = fmt.Errorf("receipt tx hash index is disabled; parquet fallback scan is not allowed: %w", ErrNotFound)
)

// ReceiptStore exposes receipt-specific operations without leaking the StateStore interface.
type ReceiptStore interface {
	LatestVersion() int64
	SetLatestVersion(version int64) error
	SetEarliestVersion(version int64) error
	GetReceipt(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error)
	GetReceiptFromStore(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error)
	SetReceipts(ctx sdk.Context, receipts []ReceiptRecord) error
	// FilterLogs queries logs across a range of blocks.
	// For single-block queries, set fromBlock == toBlock.
	FilterLogs(ctx sdk.Context, fromBlock, toBlock uint64, crit filters.FilterCriteria) ([]*ethtypes.Log, error)
	Close() error
}

type ReceiptRecord struct {
	TxHash       common.Hash
	Receipt      *types.Receipt
	ReceiptBytes []byte // Optional pre-marshaled receipt (must match Receipt if set)
}

// ReceiptReadMetrics records cache hits, misses, and timing for cached receipt
// and log reads.
type ReceiptReadMetrics interface {
	ReportReceiptCacheHit()
	ReportReceiptCacheMiss()
	ReportLogFilterCacheHit()
	ReportLogFilterCacheMiss()
	RecordCacheFilterScanDuration(seconds float64)
	RecordCacheGetDuration(seconds float64)
}

type receiptStore struct {
	db          seidbtypes.StateStore
	storeKey    sdk.StoreKey
	stopPruning chan struct{}
	pruneWg     sync.WaitGroup
	closeOnce   sync.Once
}

const (
	receiptBackendPebble  = "pebble"
	receiptBackendParquet = "parquet"
)

func normalizeReceiptBackend(backend string) string { _ = "STUB: not implemented"; return "" }

func NewReceiptStore(config dbconfig.ReceiptStoreConfig, storeKey sdk.StoreKey) (ReceiptStore, error) {
	_ = "STUB: not implemented"
	return *new(ReceiptStore), nil
}

// NewReceiptStoreWithReadMetrics constructs a receipt store and optionally
// records cache hits, misses, and timings for cached receipt/log reads.
func NewReceiptStoreWithReadMetrics(
	config dbconfig.ReceiptStoreConfig,
	storeKey sdk.StoreKey,
	metrics ReceiptReadMetrics,
) (ReceiptStore, error) {
	_ = "STUB: not implemented"
	return *new(ReceiptStore), nil
}

// BackendTypeName returns the backend implementation name ("parquet" or "pebble") for testing.
// Returns "" if store is nil or the backend type is unknown.
func BackendTypeName(store ReceiptStore) string { _ = "STUB: not implemented"; return "" }

func newReceiptBackend(config dbconfig.ReceiptStoreConfig, storeKey sdk.StoreKey) (ReceiptStore, error) {
	_ = "STUB: not implemented"
	return *new(ReceiptStore), nil
}

func (s *receiptStore) LatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *receiptStore) SetLatestVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (s *receiptStore) SetEarliestVersion(version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *receiptStore) GetReceipt(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	// receipts are immutable, use latest version
	return nil, nil
}

// try persistent store

// try legacy store for older receipts

// Only used for testing.
func (s *receiptStore) GetReceiptFromStore(_ sdk.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	// receipts are immutable, use latest version
	return nil, nil
}

// try persistent store

func (s *receiptStore) SetReceipts(ctx sdk.Context, receipts []ReceiptRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// Genesis and some unit tests execute at block height 0. Async writes
// rely on a positive version to avoid regressions in the underlying
// state store metadata, so fall back to a synchronous apply in that case.

// for tests

// fallback to synchronous apply for stores that do not support async writes

// FilterLogs is not efficiently supported by the pebble backend since receipts
// are indexed by tx hash, not by block number. Returns ErrRangeQueryNotSupported.
// Callers should fall back to fetching receipts individually via GetReceipt.
func (s *receiptStore) FilterLogs(_ sdk.Context, _, _ uint64, _ filters.FilterCriteria) ([]*ethtypes.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *receiptStore) Close() error { _ = "STUB: not implemented"; return nil }

func recoverReceiptStore(changelogPath string, db seidbtypes.StateStore) error {
	_ = "STUB: not implemented"
	return nil
}

// Look backward to find where we should start replay from

// Fresh store (or no applied versions) - start from the first offset

// Replay from the offset where the version is larger than SS store latest version

// commit to state store

func startReceiptPruning(db seidbtypes.StateStore, keepRecent int64, pruneInterval int64, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// prune all versions up to and including the pruneVersion

// Generate a random percentage (between 0% and 100%) of the fixed interval as a delay

// Continue to next iteration

func getLogsForTx(receipt *types.Receipt, logStartIndex uint) []*ethtypes.Log {
	_ = "STUB: not implemented"
	return nil
}

func convertLog(l *types.Log, receipt *types.Receipt, logStartIndex uint) *ethtypes.Log {
	_ = "STUB: not implemented"
	return nil
}
