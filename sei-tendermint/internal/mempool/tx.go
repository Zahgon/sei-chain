package mempool

import (
	"context"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/libs/clist"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// TxInfo are parameters that get passed when attempting to add a tx to the
// mempool.
type TxInfo struct {
	// SenderID is the internal peer ID used in the mempool to identify the
	// sender, storing two bytes with each transaction instead of 20 bytes for
	// the types.NodeID.
	SenderID uint16

	// SenderNodeID is the actual types.NodeID of the sender.
	SenderNodeID types.NodeID
}

type hashedTx struct {
	tx   types.Tx
	hash types.TxHash
}

func newHashedTx(tx types.Tx) hashedTx { _ = "STUB: not implemented"; return *new(hashedTx) }

func (ktx *hashedTx) Tx() types.Tx       { _ = "STUB: not implemented"; return *new(types.Tx) }
func (ktx *hashedTx) Hash() types.TxHash { _ = "STUB: not implemented"; return *new(types.TxHash) }
func (ktx *WrappedTx) Size() int         { _ = "STUB: not implemented"; return 0 }

// WrappedTx defines a wrapper around a raw transaction with additional metadata
// that is used for indexing.
type WrappedTx struct {
	// hashedTx represents the raw binary transaction data and its memoized hash.
	hashedTx

	// height defines the height at which the transaction was validated at
	height int64

	// gasWanted defines the amount of gas the transaction sender requires
	gasWanted int64

	// estimatedGas defines the amount of gas that the transaction is estimated to use
	estimatedGas int64

	// priority defines the transaction's priority as specified by the application
	// in the ResponseCheckTx response.
	priority int64

	// timestamp is the time at which the node first received the transaction from
	// a peer. It is used as a second dimension is prioritizing transactions when
	// two transactions have the same priority.
	timestamp time.Time

	// peers records a mapping of all peers that sent a given transaction
	peers map[uint16]struct{}

	// heapIndex defines the index of the item in the heap
	heapIndex int

	// gossipEl references the linked-list element in the gossip index
	gossipEl *clist.CElement[*WrappedTx]

	// removed marks the transaction as removed from the mempool. This is set
	// during RemoveTx and is needed due to the fact that a given existing
	// transaction in the mempool can be evicted when it is simultaneously having
	// a reCheckTx callback executed.
	removed bool

	// evm properties that aid in prioritization
	evm utils.Option[evmTx]
}

type evmTx struct {
	address    common.Address
	seiAddress []byte
	nonce      uint64
	// evmRequiredBalance is the sender balance threshold for this EVM tx to become ready.
	requiredBalance *big.Int
}

// IsBefore returns true if the WrappedTx is before the given WrappedTx
// this applies to EVM transactions only
func (wtx *WrappedTx) EVMNonce() uint64 { _ = "STUB: not implemented"; return 0 }

type txStoreInner struct {
	byHash    map[types.TxHash]*WrappedTx // primary index
	sizeBytes utils.AtomicSend[int64]
}

// TxStore implements a thread-safe mapping of valid transaction(s).
//
// NOTE:
//   - Concurrent read-only access to a *WrappedTx object is OK. However, mutative
//     access is not allowed. Regardless, it is not expected for the mempool to
//     need mutative access.
type TxStore struct {
	inner     utils.RWMutex[*txStoreInner]
	sizeBytes utils.AtomicRecv[int64]
}

func NewTxStore() *TxStore { _ = "STUB: not implemented"; return nil }

// Size returns the total number of transactions in the store.
func (txs *TxStore) Size() int { _ = "STUB: not implemented"; return 0 }

// AllTxsBytes returns the total size in bytes of all transactions in the store.
func (txs *TxStore) AllTxsBytes() int64 { _ = "STUB: not implemented"; return 0 }

// WaitForTxs waits until the store becomes non-empty.
func (txs *TxStore) WaitForTxs(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// GetAllTxs returns all the transactions currently in the store.
func (txs *TxStore) GetAllTxs() []*WrappedTx { _ = "STUB: not implemented"; return nil }

// GetOlderThan have older timestamp than minTime OR lower height than minHeight.
func (txs *TxStore) GetOlderThan(minTime utils.Option[time.Time], minHeight utils.Option[int64]) []*WrappedTx {
	_ = "STUB: not implemented"
	return nil
}

// GetTxByHash returns a *WrappedTx by the transaction's hash.
func (txs *TxStore) GetTxByHash(key types.TxHash) *WrappedTx { _ = "STUB: not implemented"; return nil }

func (txs *TxStore) IsTxRemovedByHash(txHash types.TxHash) bool {
	_ = "STUB: not implemented"
	return false
}

// IsTxRemoved returns true if a transaction by hash is marked as removed and
// false otherwise.
func (txs *TxStore) IsTxRemoved(wtx *WrappedTx) bool { _ = "STUB: not implemented"; return false }

// if this instance has already been marked, return true

// otherwise if the same hash exists, return its state

// otherwise we haven't seen this tx

// SetTx stores a *WrappedTx by its hash.
func (txs *TxStore) SetTx(wtx *WrappedTx) { _ = "STUB: not implemented"; return }

// RemoveTx removes a *WrappedTx from the transaction store. It deletes all
// indexes of the transaction.
func (txs *TxStore) RemoveTx(wtx *WrappedTx) { _ = "STUB: not implemented"; return }

// TxHasPeer returns true if a transaction by hash has a given peer ID and false
// otherwise. If the transaction does not exist, false is returned.
func (txs *TxStore) TxHasPeer(key types.TxHash, peerID uint16) bool {
	_ = "STUB: not implemented"
	return false
}

// GetOrSetPeerByTxHash looks up a WrappedTx by transaction hash and adds the
// given peerID to the WrappedTx's set of peers that sent us this transaction.
// We return true if we've already recorded the given peer for this transaction
// and false otherwise. If the transaction does not exist by hash, we return
// (nil, false).
func (txs *TxStore) GetOrSetPeerByTxHash(hash types.TxHash, peerID uint16) (*WrappedTx, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type PendingTxs struct {
	inner     utils.RWMutex[*pendingTxsInner]
	config    *Config
	sizeBytes atomic.Int64
}

type pendingTxsInner struct {
	txs []*WrappedTx
}

func NewPendingTxs(conf *Config) *PendingTxs { _ = "STUB: not implemented"; return nil }

func (p *PendingTxs) EvaluatePendingTransactions(
	evaluate func(*WrappedTx) abci.PendingTxCheckerResponse,
) (
	acceptedTxs []*WrappedTx,
	rejectedTxs []*WrappedTx,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Assumes the pending tx store is already write-locked.
func (p *PendingTxs) popTxsAtIndices(inner *pendingTxsInner, indices []int) {
	_ = "STUB: not implemented"
	return
}

func (p *PendingTxs) Insert(tx *WrappedTx) error { _ = "STUB: not implemented"; return nil }

func (p *PendingTxs) SizeBytes() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PendingTxs) Peek(max int) []*WrappedTx { _ = "STUB: not implemented"; return nil }

// priority is fifo

func (p *PendingTxs) Size() int { _ = "STUB: not implemented"; return 0 }

func (p *PendingTxs) PurgeExpired(blockHeight int64, now time.Time, cb func(wtx *WrappedTx)) {
	_ = "STUB: not implemented"
	return
}

// txs retains the ordering of insertion

// once found, we can break because these are ordered

// once found, we can break because these are ordered
