package receipt

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	dbconfig "github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/parquet"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

// parquetReceiptStore wraps the parquet.Store and implements ReceiptStore.
type parquetReceiptStore struct {
	store       *parquet.Store
	storeKey    sdk.StoreKey
	txHashIndex TxHashIndex
	indexPruner *txHashIndexPruner
}

func newParquetReceiptStore(cfg dbconfig.ReceiptStoreConfig, storeKey sdk.StoreKey) (ReceiptStore, error) {
	_ = "STUB: not implemented"
	return *new(ReceiptStore), nil
}

func (s *parquetReceiptStore) LatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (s *parquetReceiptStore) SetLatestVersion(version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *parquetReceiptStore) SetEarliestVersion(version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *parquetReceiptStore) cacheRotateInterval() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *parquetReceiptStore) warmupReceipts() []ReceiptRecord {
	_ = "STUB: not implemented"
	return nil
}

func (s *parquetReceiptStore) GetReceipt(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *parquetReceiptStore) GetReceiptFromStore(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// indexedReceiptLookup uses the tx hash index to narrow the parquet search to
// a single file. When the index is disabled the lookup returns
// ErrTxIndexDisabled instead of performing a full parquet scan, which would
// be prohibitively expensive at production scale.
func (s *parquetReceiptStore) indexedReceiptLookup(ctx context.Context, txHash common.Hash) (*parquet.ReceiptResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *parquetReceiptStore) SetReceipts(ctx sdk.Context, receipts []ReceiptRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// Let the parquet store rotate on aligned boundaries even when the
// block has no EVM receipts. Without this the rotation invariant
// drifts whenever a boundary block happens to be empty, and the
// reader's file-pruning logic silently misses queries that fall
// into the over-sized open file.
//nolint:gosec // block heights fit within uint64

//nolint:gosec // block numbers won't exceed int64 max

// indexReceiptInputs batches tx hashes by block number and writes them to the
// tx hash index.
func (s *parquetReceiptStore) indexReceiptInputs(inputs []parquet.ReceiptInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *parquetReceiptStore) FilterLogs(ctx sdk.Context, fromBlock, toBlock uint64, crit filters.FilterCriteria) ([]*ethtypes.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *parquetReceiptStore) Close() error { _ = "STUB: not implemented"; return nil }

func (s *parquetReceiptStore) replayWAL() error { _ = "STUB: not implemented"; return nil }

// Collect tx hashes per block during replay so the index can be
// populated in a single batch after the parquet store is consistent.

// A boundary entry about to rotate makes every prior entry stale (those
// blocks are flushed into the file being closed). Advance dropOffset so
// the post-replay truncate removes them.

// Re-index replayed blocks so the tx hash index stays consistent
// with the parquet store after a crash/restart.

//nolint:gosec // block numbers won't exceed int64 max

func truncateReplayWAL(w interface{ TruncateBefore(offset uint64) error }, dropOffset uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func BuildParquetLogRecords(logs []*ethtypes.Log, blockHash common.Hash) []parquet.LogRecord {
	_ = "STUB: not implemented"
	return nil
}

func buildTopicsFromParquetLogResult(lr parquet.LogResult) []common.Hash {
	_ = "STUB: not implemented"
	return nil
}

func ExtractLogTopics(topics []common.Hash) ([]byte, []byte, []byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
