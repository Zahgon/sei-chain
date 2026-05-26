package mempool

import (
	"context"
	"crypto/sha256"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/libs/clist"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/libs/reservoir"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "internal", "mempool")

// ErrTxInCache is returned to the client if we saw tx earlier.
var ErrTxInCache = errors.New("tx already exists in cache")

// ErrTxTooLarge defines an error when a transaction is too big to be sent to peers.
var ErrTxTooLarge = errors.New("tx too large")

// Using SHA-256 truncated to 128 bits as the cache key: At 2K tx/sec, the
// collision probability is effectively zero (≈10^-29 for 120K keys in a minute,
// still negligible over years). If reduced 3× smaller (~43 bits), collisions
// become probable within a day and guaranteed over longer periods.
//
// For the purposes of the LRU cache key both sizes are sufficiently secure. For
// now. 128 bits is a safe balance between performance and collision probability
// and we may revisit later.
const maxCacheKeySize = sha256.Size / 2

// MinTxsPerBlock is how many txs we will attempt to have in a block if there's still space.
// MinGasEVMTx is the minimum the gas estimate can be for an EVM tx to be considered valid.
const (
	MinTxsToPeek = 10
	MinGasEVMTx  = 21000
)

type Config struct {
	// Maximum number of transactions in the mempool
	Size int

	// Limit the total size of all txs in the mempool.
	// This only accounts for raw transactions (e.g. given 1MB transactions and
	// max-txs-bytes=5MB, mempool will only accept 5 transactions).
	MaxTxsBytes int64

	// Size of the cache (used to filter transactions we saw earlier) in transactions
	CacheSize int

	// Size of the duplicate cache used to track duplicate txs
	DuplicateTxsCacheSize int

	// Do not remove invalid transactions from the cache (default: false)
	// Set to true if it's not possible for any invalid transaction to become
	// valid again in the future.
	KeepInvalidTxsInCache bool

	// Maximum size of a single transaction
	// NOTE: the max size of a tx transmitted over the network is {max-tx-bytes}.
	MaxTxBytes int

	// TTLDuration, if non-zero, defines the maximum amount of time a transaction
	// can exist for in the mempool.
	//
	// Note, if TTLNumBlocks is also defined, a transaction will be removed if it
	// has existed in the mempool at least TTLNumBlocks number of blocks or if it's
	// insertion time into the mempool is beyond TTLDuration.
	TTLDuration time.Duration

	// TTLNumBlocks, if non-zero, defines the maximum number of blocks a transaction
	// can exist for in the mempool.
	//
	// Note, if TTLDuration is also defined, a transaction will be removed if it
	// has existed in the mempool at least TTLNumBlocks number of blocks or if
	// it's insertion time into the mempool is beyond TTLDuration.
	TTLNumBlocks int64

	// TxNotifyThreshold, if non-zero, defines the minimum number of transactions
	// needed to trigger a notification in mempool's Tx notifier
	TxNotifyThreshold uint64

	// Maximum number of transactions in the pending set
	PendingSize int

	// Limit the total size of all txs in the pending set.
	MaxPendingTxsBytes int64

	RemoveExpiredTxsFromQueue bool

	// DropPriorityThreshold defines the percentage of transactions with the lowest
	// priority hint (expressed as a float in the range [0.0, 1.0]) that will be
	// dropped from the mempool once the configured utilisation threshold is reached.
	//
	// The default value of 0.1 means that the lowest 10% of transactions by
	// priority will be dropped when the mempool utilisation exceeds the
	// DropUtilisationThreshold.
	//
	// See DropUtilisationThreshold.
	DropPriorityThreshold float64

	// DropUtilisationThreshold defines the mempool utilisation level (expressed as
	// a percentage in the range [0.0, 1.0]) above which transactions will be
	// selectively dropped based on their priority hint.
	//
	// For example, if this parameter is set to 0.8, then once the mempool reaches
	// 80% capacity, transactions with priority hints below DropPriorityThreshold
	// percentile will be dropped to make room for new transactions.
	DropUtilisationThreshold float64

	// DropPriorityReservoirSize defines the size of the reservoir for keeping track
	// of the distribution of transaction priorities in the mempool.
	//
	// This is used to determine the priority threshold below which transactions will
	// be dropped when the mempool utilisation exceeds DropUtilisationThreshold.
	//
	// The reservoir is a statistically representative sample of transaction
	// priorities in the mempool, and is used to estimate the priority distribution
	// without needing to store all transaction priorities.
	//
	// A larger reservoir size will yield a more accurate estimate of the priority
	// distribution, but will consume more memory.
	//
	// The default value of 10,240 is a reasonable compromise between accuracy and
	// memory usage for most use cases. It takes approximately 80KB of memory storing
	// int64 transaction priorities.
	//
	// See DropUtilisationThreshold and DropPriorityThreshold.
	DropPriorityReservoirSize int `mapstructure:"drop-priority-reservoir-size"`
}

func DefaultConfig() *Config {
	_ = "STUB: not implemented"

	// Each signature verification takes .5ms, Size reduced until we implement
	// ABCI Recheck
	return nil
}

// 1GB

// 1MB
// prevent stale txs from filling mempool
// remove txs after 10 blocks

// 1GB

type evmAddrNonce struct {
	Address common.Address
	Nonce   uint64
}

// TxMempool defines a prioritized mempool data structure used by the v1 mempool
// reactor. It keeps a thread-safe priority queue of transactions that is used
// when a block proposer constructs a block and a thread-safe linked-list that
// is used to gossip transactions to peers in a FIFO manner.
type TxMempool struct {
	metrics *Metrics
	config  *Config
	app     *proxy.Proxy

	// txsAvailable fires once for each height when the mempool is not empty
	txsAvailable         chan struct{}
	notifiedTxsAvailable atomic.Bool

	// height defines the last block height process during Update()
	height int64

	// cache defines a fixed-size cache of already seen transactions as this
	// reduces pressure on the proxyApp.
	cache TxCache

	// blockFailedTxs tracks tx hashes that have previously failed during
	// block execution. Used to prevent infinite re-entry of txs that
	// consistently fail before fee charging in DeliverTx.
	blockFailedTxs TxCache

	// A TTL cache which keeps all txs that we have seen before over the TTL window.
	// Currently, this can be used for tracking whether checkTx is always serving the same tx or not.
	duplicateTxsCache utils.Option[*DuplicateTxCache]

	// txStore defines the main storage of valid transactions. Indexes are built
	// on top of this store.
	txStore *TxStore

	// gossipIndex defines the gossiping index of valid transactions via a
	// thread-safe linked-list. We also use the gossip index as a cursor for
	// rechecking transactions already in the mempool.
	gossipIndex *clist.CList[*WrappedTx]

	// recheckCursor and recheckEnd are used as cursors based on the gossip index
	// to recheck transactions that are already in the mempool. Iteration is not
	// thread-safe and transaction may be mutated in serial order.
	//
	// XXX/TODO: It might be somewhat of a codesmell to use the gossip index for
	// iterator and cursor management when rechecking transactions. If the gossip
	// index changes or is removed in a future refactor, this will have to be
	// refactored. Instead, we should consider just keeping a slice of a snapshot
	// of the mempool's current transactions during Update and an integer cursor
	// into that slice. This, however, requires additional O(n) space complexity.
	recheckCursor *clist.CElement[*WrappedTx] // next expected response
	recheckEnd    *clist.CElement[*WrappedTx] // re-checking stops here

	// priorityIndex defines the priority index of valid transactions via a
	// thread-safe priority queue.
	priorityIndex *TxPriorityQueue

	// pendingTxs stores transactions that are not valid yet but might become valid
	// once nonce ordering or sender balance catches up.
	pendingTxs *PendingTxs

	byAddrNonce utils.Mutex[map[evmAddrNonce]*WrappedTx]

	// A read/write lock is used to safe guard updates, insertions and deletions
	// from the mempool. A read-lock is implicitly acquired when executing CheckTx,
	// however, a caller must explicitly grab a write-lock via Lock when updating
	// the mempool via Update().
	mtx                  sync.RWMutex
	txConstraintsFetcher TxConstraintsFetcher

	priorityReservoir *reservoir.Sampler[int64]
}

func NewTxMempool(
	cfg *Config,
	app *proxy.Proxy,
	metrics *Metrics,
	txConstraintsFetcher TxConstraintsFetcher,
) *TxMempool {
	_ = "STUB: not implemented"
	return nil
}

// Use non-deterministic RNG

func (txmp *TxMempool) Config() *Config { _ = "STUB: not implemented"; return nil }

func (txmp *TxMempool) App() *proxy.Proxy { _ = "STUB: not implemented"; return nil }

func (txmp *TxMempool) EvmNextPendingNonce(addr common.Address) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (txmp *TxMempool) addNonce(wtx *WrappedTx) { _ = "STUB: not implemented"; return }

func (txmp *TxMempool) removeNonce(wtx *WrappedTx) { _ = "STUB: not implemented"; return }

func (txmp *TxMempool) TxStore() *TxStore {
	_ = "STUB: not implemented"

	// Lock obtains a write-lock on the mempool. A caller must be sure to explicitly
	// release the lock when finished.
	return nil
}

func (txmp *TxMempool) Lock() {
	_ = "STUB: not implemented"

	// Unlock releases a write-lock on the mempool.
	return
}

func (txmp *TxMempool) Unlock() {
	_ = "STUB: not implemented"

	// Size returns the number of valid transactions in the mempool. It is
	// thread-safe.
	return
}

func (txmp *TxMempool) Size() int { _ = "STUB: not implemented"; return 0 }

func (txmp *TxMempool) utilisation() float64 { _ = "STUB: not implemented"; return 0 }

func (txmp *TxMempool) NumTxsNotPending() int { _ = "STUB: not implemented"; return 0 }

func (txmp *TxMempool) BytesNotPending() int64 { _ = "STUB: not implemented"; return 0 }

func (txmp *TxMempool) TotalTxsBytesSize() int64 { _ = "STUB: not implemented"; return 0 }

// PendingSize returns the number of pending transactions in the mempool.
func (txmp *TxMempool) PendingSize() int        { _ = "STUB: not implemented"; return 0 }
func (txmp *TxMempool) PendingSizeBytes() int64 { _ = "STUB: not implemented"; return 0 }

// SizeBytes return the total sum in bytes of all the valid transactions in the
// mempool. It is thread-safe.
func (txmp *TxMempool) SizeBytes() int64 { _ = "STUB: not implemented"; return 0 }

// WaitForNextTx waits until the next transaction is available for gossip.
// Returns the next valid transaction to gossip.
func (txmp *TxMempool) WaitForNextTx(ctx context.Context) (*clist.CElement[*WrappedTx], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TxsAvailable returns a channel which fires once for every height, and only
// when transactions are available in the mempool. It is thread-safe.
func (txmp *TxMempool) TxsAvailable() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (txmp *TxMempool) checkResponseState(wtx *WrappedTx) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckTx executes the ABCI CheckTx method for a given transaction.
// It acquires a read-lock and attempts to execute the application's
// CheckTx ABCI method synchronously. We return an error if any of
// the following happen:
//
//   - The CheckTx execution fails.
//   - The transaction already exists in the cache and we've already received the
//     transaction from the peer. Otherwise, if it solely exists in the cache, we
//     return nil.
//   - The transaction size exceeds the maximum transaction size as defined by the
//     configuration provided to the mempool.
//   - The transaction fails the consensus-derived mempool checks.
//   - The app fails, e.g. the buffer is full.
//
// If the mempool is full, we still execute CheckTx and attempt to find a lower
// priority transaction to evict. If such a transaction exists, we remove the
// lower priority transaction and add the new one with higher priority.
//
// NOTE:
// - The applications' CheckTx implementation may panic.
// - The caller is not to explicitly require any locks for executing CheckTx.
func (txmp *TxMempool) CheckTx(ctx context.Context, tx types.Tx, txInfo TxInfo) (*abci.ResponseCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reject low priority transactions when the mempool is more than
// DropUtilisationThreshold full.

// We add the transaction to the mempool's cache and if the
// transaction is already present in the cache, i.e. false is returned, then we
// check if we've seen this transaction and error if we have.

// Check TTL cache to see if we've recently processed this transaction
// Only execute TTL cache logic if we're using a real TTL cache (not NOP)

// only add new transaction if checkTx passes and is not pending

// otherwise add to pending txs store

// TODO: eviction strategy for pending transactions

func (txmp *TxMempool) isInMempool(txHash types.TxHash) bool {
	_ = "STUB: not implemented"
	return false
}

func (txmp *TxMempool) HasTx(txHash types.TxHash) bool { _ = "STUB: not implemented"; return false }

func (txmp *TxMempool) GetTxsForHashes(txHashes []types.TxHash) types.Txs {
	_ = "STUB: not implemented"
	return *new(types.Txs)
}

func (txmp *TxMempool) SafeGetTxsForHashes(txHashes []types.TxHash) (types.Txs, []types.TxHash) {
	_ = "STUB: not implemented"
	return *new(types.Txs), nil
}

// Flush empties the mempool. It acquires a read-lock, fetches all the
// transactions currently in the transaction store and removes each transaction
// from the store and all indexes and finally resets the cache.
//
// NOTE:
// - Flushing the mempool may leave the mempool in an inconsistent state.
func (txmp *TxMempool) Flush() { _ = "STUB: not implemented"; return }

// ReapMaxBytesMaxGas returns a list of transactions within the provided size
// and gas constraints. The returned list starts with EVM transactions (in priority order),
// followed by non-EVM transactions (in priority order).
// There are 4 types of constraints.
//  1. maxBytes - stops pulling txs from mempool once maxBytes is hit.
//  2. maxGasWanted - stops pulling txs from mempool once total gas wanted exceeds maxGasWanted.
//     Can be set to -1 to be ignored.
//  3. maxGasEstimated - similar to maxGasWanted but will use the estimated gas used for EVM txs
//     while still using gas wanted for cosmos txs. Can be set to -1 to be ignored.
//
// NOTE:
//   - Transactions returned are not removed from the mempool transaction
//     store or indexes.
func (txmp *TxMempool) ReapMaxBytesMaxGas(maxBytes, maxGasWanted, maxGasEstimated int64) types.Txs {
	_ = "STUB: not implemented"
	return *new(types.Txs)
}

type ReapLimits struct {
	MaxTxs          utils.Option[uint64]
	MaxBytes        utils.Option[int64]
	MaxGasWanted    utils.Option[int64]
	MaxGasEstimated utils.Option[int64]
}

// ReapMaxTxsBytesMaxGas returns a list of transactions within the provided tx,
// byte, and gas constraints together with the total estimated gas for the
// returned transactions.
//
// NOTE: Gas limits are enforced using int64 running totals. If those totals
// overflow, gas limit enforcement no longer works correctly. This preserves the
// historical behavior for backward compatibility.
func (txmp *TxMempool) reapTxs(l ReapLimits) (types.Txs, int64) {
	_ = "STUB: not implemented"
	return *new(types.Txs), 0
}

//nolint:gosec // NumTxsNotPending returns non-negative value
// do not reap anything if threshold is not met

// bytes limit is a hard stop

// if the tx doesn't have a gas estimate, fallback to gas wanted

// prospective totals

// skip this unfit-by-gas tx once and attempt to pull up to 10 smaller ones

// include tx and update totals

// RemoveTxs removes the provided transactions from the mempool if present.
func (txmp *TxMempool) PopTxs(l ReapLimits) (types.Txs, int64) {
	_ = "STUB: not implemented"
	return *new(types.Txs), 0
}

// ReapMaxTxs returns a list of transactions within the provided number of
// transactions bound. Transaction are retrieved in priority order.
//
// NOTE:
//   - Transactions returned are not removed from the mempool transaction
//     store or indexes.
func (txmp *TxMempool) ReapMaxTxs(max int) types.Txs {
	_ = "STUB: not implemented"
	return *new(types.Txs)
}

// retrieve more from pending txs

// Update iterates over all the transactions provided by the block producer,
// removes them from the cache (if applicable), and removes
// the transactions from the main transaction store and associated indexes.
// If there are transactions remaining in the mempool, we initiate a
// re-CheckTx for them (if applicable), otherwise, we notify the caller more
// transactions are available.
//
// WARNING: callers should almost always pass recheck=false. recheck=true
// re-runs CheckTx on every tx still in the mempool after each block, and
// handleRecheckResult treats a "now pending" response as terminal: it
// evicts the tx and async-re-CheckTx-es it, which lands it back in
// pendingTxs. For chains whose antehandler returns pending for any
// ahead-of-nonce EVM tx (Sei), this evicts perfectly-valid queued txs.
//
// Example. txA (nonce 3), txB (nonce 2), txC (nonce 1) on the same sender.
//
//  1. txA, txB, txC are submitted in this order.
//  2. txA and txB enter pendingTxs (their nonce is ahead of the sender's
//     expected nonce at CheckTx time so the EVM antehandler marks them
//     pending). txC enters the priority index (its nonce matches expected).
//  3. Block 1 reaps and mines txC. The sender's expected nonce becomes 2.
//  4. handlePendingTransactions promotes txA and txB into the priority
//     index. The per-sender evmQueue is now [txB (head), txA (tail)].
//
// From step 5 onwards the recheck flag matters:
//
// recheck=false (correct):
//
//  5. updateReCheckTxs is skipped. The priority index keeps txB and txA.
//  6. Block 2 reaps the whole evmQueue. Both txB and txA mine.
//
// All 3 txs mine in 2 blocks, regardless of how out-of-order they arrived.
//
// recheck=true (broken):
//
//  5. updateReCheckTxs re-runs CheckTx on each tx in the priority index:
//     - txB: nonce 2 == expected 2 → not pending → stays.
//     - txA: nonce 3  > expected 2 → pending again. handleRecheckResult
//     evicts it and async-re-CheckTx-es it, which lands it back in
//     pendingTxs.
//  6. Block 2 reaps txB only (txA is no longer in the priority index).
//     handlePendingTransactions re-promotes txA. txA's nonce now matches
//     expected, so it survives the recheck this time.
//  7. Block 3 mines txA.
//
// All 3 txs take 3 blocks. With many out-of-order sequential nonces from
// one sender, this stalls the chain to 1-tx-per-block-per-sender throughput.
//
// CometBFT's default for ConsensusParams.ABCI.RecheckTx is false. Recheck
// primarily defended against state-dependent invalidation that modern
// chains catch in ProcessProposal/DeliverTx anyway.
//
// NOTE:
// - The caller must explicitly acquire a write-lock.
func (txmp *TxMempool) Update(
	ctx context.Context,
	blockHeight int64,
	blockTxs types.Txs,
	execTxResult []*abci.ExecTxResult,
	txConstraintsFetcher TxConstraintsFetcher,
	recheck bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// add the valid committed transaction to the cache (if missing)

// First block failure: allow one retry

// Subsequent failures: leave in cache to prevent infinite re-entry

// remove the committed transaction from the transaction store and indexes

// remove any tx that has the same nonce (because the committed tx
// may be from block proposal and is never in the local mempool)

// If there any uncommitted transactions left in the mempool, we either
// initiate re-CheckTx per remaining transaction or notify that remaining
// transactions are left.

// addNewTransaction is invoked for a new unique transaction after CheckTx
// has been executed by the ABCI application for the first time on that transaction.
// CheckTx can be called again for the same transaction later when re-checking;
// however, this function will not be called. A recheck after a block is committed
// goes to handleRecheckResult.
//
// addNewTransaction runs after the ABCI application executes CheckTx.
// It runs the consensus-derived post-check for the current state snapshot.
// If the post-check reports an error, the transaction is rejected. Otherwise,
// we attempt to insert the transaction into the mempool. CheckTx response codes
// are filtered earlier in CheckTx.
//
// When inserting a transaction, we first check if there is sufficient capacity.
// If there is, the transaction is added to the txStore and all indexes.
// Otherwise, if the mempool is full, we attempt to find a lower priority transaction
// to evict in place of the new incoming transaction. If no such transaction exists,
// the new incoming transaction is rejected.
//
// NOTE:
// - An explicit lock is NOT required.
func (txmp *TxMempool) addNewTransaction(wtx *WrappedTx) error {
	_ = "STUB: not implemented"
	// Update transaction priority reservoir with the true Tx priority
	// as determined by the application.
	//
	// NOTE: This is done before potentially rejecting the transaction due to
	// mempool being full. This is to ensure that the reservoir contains a
	// representative sample of all transactions that have been processed by
	// CheckTx.
	//
	// However, this is NOT done if the tx is pending, since a spammer could
	// throw off the correct priority percentiles otherwise.
	//
	// We do not use the priority hint here as it may be misleading and
	// inaccurate. The true priority as determined by the application is the
	// most accurate.
	return nil
}

// ignore bad transactions

// No room for the new incoming transaction so we just remove it from
// the cache.

// evict an existing transaction(s)
//
// NOTE:
// - The transaction, toEvict, can be removed while a concurrent
//   reCheckTx callback is being executed for the same transaction.

// handleRecheckResult handles the responses from ABCI CheckTx calls issued
// during the recheck phase of a block Update.  It removes any transactions
// invalidated by the application.
//
// The caller must hold a mempool write-lock (via Lock()) and when
// executing Update(), if the mempool is non-empty and Recheck is
// enabled, then all remaining transactions will be rechecked via
// CheckTx. The order transactions are rechecked must be the same as
// the order in which this callback is called.
//
// This method is NOT executed for the initial CheckTx on a new transaction;
// that case is handled by addNewTransaction instead.
func (txmp *TxMempool) handleRecheckResult(tx types.Tx, res *abci.ResponseCheckTxV2) {
	_ = "STUB: not implemented"
	return
}

// Search through the remaining list of tx to recheck for a transaction that matches
// the one we received from the ABCI application.

// we reached the end of the recheckTx list without finding a tx
// matching the one we received from the ABCI application.
// Return without processing any tx.

// Only evaluate transactions that have not been removed. This can happen
// if an existing transaction is evicted during CheckTx and while this
// callback is being executed for the same evicted transaction.

// we will treat a transaction that turns pending in a recheck as invalid and evict it

// move reCheckTx cursor to next element

// updateReCheckTxs updates the recheck cursors using the gossipIndex. For
// each transaction, it executes CheckTx. The global callback defined on
// the app will be executed for each transaction after CheckTx is
// executed.
//
// NOTE:
// - The caller must have a write-lock when executing updateReCheckTxs.
func (txmp *TxMempool) updateReCheckTxs(ctx context.Context) { _ = "STUB: not implemented"; return }

// Only execute CheckTx if the transaction is not marked as removed which
// could happen if the transaction was evicted.

// no need in retrying since the tx will be rechecked after the next block

// canAddTx returns an error if we cannot insert the provided *WrappedTx into
// the mempool due to mempool configured constraints. If it returns nil,
// the transaction can be inserted into the mempool.
func (txmp *TxMempool) canAddTx(wtx *WrappedTx) error { _ = "STUB: not implemented"; return nil }

func (txmp *TxMempool) canAddPendingTx(wtx *WrappedTx) error { _ = "STUB: not implemented"; return nil }

func (txmp *TxMempool) insertTx(wtx *WrappedTx) bool { _ = "STUB: not implemented"; return false }

// Insert the transaction into the gossip index and mark the reference to the
// linked-list element, which will be needed at a later point when the
// transaction is removed.

func (txmp *TxMempool) removeTx(wtx *WrappedTx, removeFromCache bool, shouldReenqueue bool, updatePriorityIndex bool) {
	_ = "STUB: not implemented"
	return
}

// Remove the transaction from the gossip index and cleanup the linked-list
// element so it can be garbage collected.

func (txmp *TxMempool) expire(blockHeight int64, wtx *WrappedTx) { _ = "STUB: not implemented"; return }

func (txmp *TxMempool) logExpiredTx(blockHeight int64, wtx *WrappedTx) {
	_ = "STUB: not implemented"
	// defensive check
	return
}

// purgeExpiredTxs removes all transactions that have exceeded their respective
// height- and/or time-based TTLs from their respective indexes. Every expired
// transaction will be removed from the mempool, but preserved in the cache (except for pending txs).
//
// NOTE: purgeExpiredTxs must only be called during TxMempool#Update in which
// the caller has a write-lock on the mempool and so we can safely iterate over
// the height and time based indexes.
func (txmp *TxMempool) purgeExpiredTxs(blockHeight int64) { _ = "STUB: not implemented"; return }

// remove pending txs that have expired

func (txmp *TxMempool) notifyTxsAvailable() { _ = "STUB: not implemented"; return }

// channel cap is 1, so this will send once

func (txmp *TxMempool) isPending(wtx *WrappedTx) bool { _ = "STUB: not implemented"; return false }

func (txmp *TxMempool) handlePendingTransactions() { _ = "STUB: not implemented"; return }

// Run executes mempool background tasks.
func (txmp *TxMempool) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO(gprusak): instead of actively updating stats,
// TxMempool should implement prometheus.Collector.
