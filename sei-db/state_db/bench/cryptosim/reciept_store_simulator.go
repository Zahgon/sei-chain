package cryptosim

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	crand "github.com/sei-protocol/sei-chain/sei-db/common/rand"
	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/receipt"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
)

const (
	// Must be larger than cacheWindow * TransactionsPerBlock so that the
	// oldest ring entries have aged past the cache window, enabling duckdb-
	// only reads to find targets.  With the default cache window of ~1000
	// blocks and 1024 txns/block the minimum is ~1.02M; 3M gives comfortable
	// headroom for both cache and duckdb read modes.
	defaultTxHashRingSize = 3_000_000
)

// txHashEntry stores a written tx hash along with its block and contract address,
// used by reader goroutines to generate realistic log filter queries.
type txHashEntry struct {
	txHash          common.Hash
	blockNumber     uint64
	contractAddress common.Address
}

// txHashRing is a fixed-size ring buffer of recently written tx hashes.
// Writers call Push from the main loop; readers call RandomEntry from goroutines.
type txHashRing struct {
	mu      sync.RWMutex
	entries []txHashEntry
	size    int
	head    int
	count   int
}

func newTxHashRing(size int) *txHashRing { _ = "STUB: not implemented"; return nil }

// Push appends a tx hash entry to the ring, overwriting the oldest entry when full.
func (r *txHashRing) Push(txHash common.Hash, blockNumber uint64, contractAddress common.Address) {
	_ = "STUB: not implemented"
	return
}

// RandomEntry returns a random entry from the ring, using CannedRandom to
// avoid potential rand.Rand hotspots under high-concurrency benchmarks.
func (r *txHashRing) RandomEntry(crand *crand.CannedRandom) *txHashEntry {
	_ = "STUB: not implemented"
	return nil
}

const maxRingSampleAttempts = 100

// RandomEntryInBlockRange samples a random entry whose blockNumber falls within
// [minBlock, maxBlock]. Returns nil if no matching entry is found after a
// bounded number of attempts.
func (r *txHashRing) RandomEntryInBlockRange(crand *crand.CannedRandom, minBlock, maxBlock uint64) *txHashEntry {
	_ = "STUB: not implemented"
	return nil
}

// A simulated receipt store with concurrent reads, writes, and pruning
// backed by the production receipt.ReceiptStore (parquet + ledger cache).
type RecieptStoreSimulator struct {
	ctx    context.Context
	cancel context.CancelFunc

	config *CryptoSimConfig

	recieptsChan chan *block

	store                    receipt.ReceiptStore
	crand                    *crand.CannedRandom
	txRing                   *txHashRing
	metrics                  *CryptosimMetrics
	receiptCacheWindowBlocks uint64
}

// Creates a new receipt store simulator backed by the production ReceiptStore
// (parquet backend + ledger cache), with optional concurrent reader goroutines.
//
// The caller must supply a CannedRandom instance (typically via Clone) that
// shares the same (seed, bufferSize) as the block builder so that
// SyntheticTxHash reproduces the hashes the write path stored.
//
// Receipt-by-hash reads reconstruct tx hashes on the fly via SyntheticTxHash
// (no storage needed). Log filter reads use the ring buffer to sample contract
// addresses written by the write path. See SyntheticTxHash in receipt.go for details.
func NewRecieptStoreSimulator(
	ctx context.Context,
	config *CryptoSimConfig,
	recieptsChan chan *block,
	metrics *CryptosimMetrics,
	crand *crand.CannedRandom,
) (*RecieptStoreSimulator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nil StoreKey is safe: the parquet write path never touches the legacy KV store.
// Cryptosim passes its metrics into the cache wrapper so cache hits/misses are
// measured at the only layer that can distinguish them reliably.

func (r *RecieptStoreSimulator) mainLoop() { _ = "STUB: not implemented"; return }

// Processes a block of receipts using the production ReceiptStore.SetReceipts path,
// then populates the ring buffer with contract addresses for log filter reads.
func (r *RecieptStoreSimulator) processBlock(blk *block) { _ = "STUB: not implemented"; return }

//nolint:gosec

//nolint:gosec

//nolint:gosec

// startReceiptReaders launches dedicated goroutines for receipt-by-hash lookups.
func (r *RecieptStoreSimulator) startReceiptReaders() { _ = "STUB: not implemented"; return }

// startLogFilterReaders launches dedicated goroutines for log filter (eth_getLogs) queries.
func (r *RecieptStoreSimulator) startLogFilterReaders() { _ = "STUB: not implemented"; return }

func (r *RecieptStoreSimulator) tickerLoop(readsPerSecond int, crand *crand.CannedRandom, fn func(*crand.CannedRandom)) {
	_ = "STUB: not implemented"
	return
}

// executeReceiptRead samples a tx hash from the ring and queries GetReceipt.
//
// ReceiptReadMode controls which blocks are targeted:
//   - "cache": only blocks within the cache window (guaranteed cache hit).
//   - "duckdb": only blocks older than the cache window (guaranteed cache miss).
func (r *RecieptStoreSimulator) executeReceiptRead(crand *crand.CannedRandom) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

// executeLogFilterRead simulates an eth_getLogs query filtering by contract address
// over a configurable block range. Contract addresses come from the ring buffer.
//
// ReceiptLogFilterReadMode controls which blocks are targeted:
//   - "cache": block range falls entirely within the cache window (DuckDB skipped).
//   - "duckdb": block range falls entirely before the cache window (cache miss).
func (r *RecieptStoreSimulator) executeLogFilterRead(crand *crand.CannedRandom) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // Config validation guarantees a positive block range before this conversion.

//nolint:gosec

//nolint:gosec

//nolint:gosec
//nolint:gosec

//nolint:gosec

// convertLogsForTx converts evmtypes.Log entries to ethtypes.Log entries.
// Mirrors receipt.getLogsForTx.
func convertLogsForTx(rcpt *evmtypes.Receipt, logStartIndex uint) []*ethtypes.Log {
	_ = "STUB: not implemented"
	return nil
}

// convertLogEntry converts a single evmtypes.Log to an ethtypes.Log.
// Mirrors receipt.convertLog.
func convertLogEntry(l *evmtypes.Log, rcpt *evmtypes.Receipt, logStartIndex uint) *ethtypes.Log {
	_ = "STUB: not implemented"
	return nil
}

// mapTopics converts hex-encoded topic strings to common.Hash values.
func mapTopics(topics []string) []common.Hash { _ = "STUB: not implemented"; return nil }
