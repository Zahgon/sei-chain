package receipt

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	dbtypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// TxHashIndex maps transaction hashes to the block number that contains them.
// Implementations must be safe for concurrent use.
type TxHashIndex interface {
	// GetBlockNumber returns the block number for a given tx hash.
	// Returns (0, false, nil) when the hash is not indexed.
	GetBlockNumber(ctx context.Context, txHash common.Hash) (blockNumber uint64, ok bool, err error)

	// IndexBlock associates every tx hash in the slice with blockNumber.
	IndexBlock(ctx context.Context, blockNumber uint64, txHashes []common.Hash) error

	// PruneBefore removes all index entries for blocks strictly below blockNumber.
	PruneBefore(ctx context.Context, blockNumber uint64) error

	Close() error
}

// Key layout for the Pebble-backed index:
//
//	Primary:  'h' + txHash (32 bytes)  -> blockNumber (8 bytes big-endian)
//	Reverse:  'b' + blockNumber (8 BE) + txHash (32 bytes) -> empty
//
// The reverse mapping lets PruneBefore delete old entries efficiently
// using a range scan on the 'b' prefix.
const (
	txHashPrefix = 'h'
	blockPrefix  = 'b'
	txHashLen    = 32
	blockNumLen  = 8
)

func makeTxHashKey(txHash common.Hash) []byte { _ = "STUB: not implemented"; return nil }

func makeBlockPrefixKey(blockNumber uint64) []byte { _ = "STUB: not implemented"; return nil }

func makeBlockTxKey(blockNumber uint64, txHash common.Hash) []byte {
	_ = "STUB: not implemented"
	return nil
}

func encodeBlockNumber(blockNumber uint64) []byte { _ = "STUB: not implemented"; return nil }

func decodeBlockNumber(b []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// PebbleTxHashIndex is the first concrete implementation of TxHashIndex,
// backed by the shared sei-db Pebble KV wrapper.
type PebbleTxHashIndex struct {
	db        dbtypes.KeyValueDB
	closeOnce sync.Once
}

var _ TxHashIndex = (*PebbleTxHashIndex)(nil)

// NewPebbleTxHashIndex opens (or creates) a Pebble-backed tx hash index
// in the given directory.
func NewPebbleTxHashIndex(dir string) (*PebbleTxHashIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *PebbleTxHashIndex) GetBlockNumber(_ context.Context, txHash common.Hash) (uint64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (idx *PebbleTxHashIndex) IndexBlock(_ context.Context, blockNumber uint64, txHashes []common.Hash) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Avoid Sync on every block: fsync per commit would add large latency;
// Pebble still appends to the WAL without forcing a full sync each time.

func (idx *PebbleTxHashIndex) PruneBefore(_ context.Context, blockNumber uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Only delete the primary key if it still points to this block.
// An overwrite (same tx hash re-indexed at a newer block) would
// have updated the primary key but left the old reverse entry;
// blindly deleting the primary key would corrupt the newer mapping.

func (idx *PebbleTxHashIndex) Close() error { _ = "STUB: not implemented"; return nil }

// txHashIndexPruner runs periodic pruning on a TxHashIndex, aligned with
// parquet file pruning. It is started by the parquet receipt store when the
// tx hash index is enabled.
type txHashIndexPruner struct {
	index    TxHashIndex
	interval time.Duration
	stopCh   chan struct{}
	wg       sync.WaitGroup
	// latestBlock is called to determine the latest block for prune calculations.
	latestBlock func() int64
	keepRecent  int64
}

func newTxHashIndexPruner(index TxHashIndex, keepRecent, pruneIntervalSec int64, latestBlock func() int64) *txHashIndexPruner {
	_ = "STUB: not implemented"
	return nil
}

func (p *txHashIndexPruner) Start() { _ = "STUB: not implemented"; return }

func (p *txHashIndexPruner) Stop() { _ = "STUB: not implemented"; return }

// TxHashIndexDir returns the canonical subdirectory name for the Pebble
// tx-hash index within a receipt store DB directory.
func TxHashIndexDir(receiptDBDir string) string { _ = "STUB: not implemented"; return "" }
