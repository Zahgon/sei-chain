package mempool

import (
	"container/heap"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

var _ heap.Interface = (*TxPriorityQueue)(nil)

// TxPriorityQueue defines a thread-safe priority queue for valid transactions.
type TxPriorityQueue struct {
	mtx sync.RWMutex
	txs []*WrappedTx // priority heap
	// invariant 1: no duplicate nonce in the same queue
	// invariant 2: no nonce gap in the same queue
	// invariant 3: head of the queue must be in heap
	evmQueue map[common.Address][]*WrappedTx // indexed by sender address, sorted by nonce
}

func insertToEVMQueue(queue []*WrappedTx, tx *WrappedTx, i int) []*WrappedTx {
	_ = "STUB: not implemented"
	// Make room for new value and add it
	return nil
}

// binarySearch finds the index at which nonce should be inserted in queue and
// whether an exact nonce match already exists.
func binarySearch(queue []*WrappedTx, nonce uint64) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func NewTxPriorityQueue() *TxPriorityQueue { _ = "STUB: not implemented"; return nil }

func (pq *TxPriorityQueue) TxByAddrNonce(addr common.Address, nonce uint64) (*WrappedTx, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (pq *TxPriorityQueue) txByAddrNonceUnsafe(addr common.Address, nonce uint64) (*WrappedTx, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (pq *TxPriorityQueue) tryReplacementUnsafe(tx *WrappedTx) (replaced *WrappedTx, shouldDrop bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// tx should be dropped since it's dominated by an existing tx

// should replace
// replace heap if applicable

// need to be in the heap since it has the same nonce

// replace queue item in-place

// GetEvictableTxs attempts to find and return a list of *WrappedTx than can be
// evicted to make room for another *WrappedTx with higher priority. If no such
// list of *WrappedTx exists, nil will be returned. The returned list of *WrappedTx
// indicate that these transactions can be removed due to them being of lower
// priority and that their total sum in size allows room for the incoming
// transaction according to the mempool's configured limits.
func (pq *TxPriorityQueue) GetEvictableTxs(priority, txSize, totalSize, cap int64) []*WrappedTx {
	_ = "STUB: not implemented"
	return nil
}

// Loop over all transactions in ascending priority order evaluating those
// that are only of less priority than the provided argument. We continue
// evaluating transactions until there is sufficient capacity for the new
// transaction (size) as defined by txSize.

// requires read lock
func (pq *TxPriorityQueue) numQueuedUnsafe() int { _ = "STUB: not implemented"; return 0 }

// first items in queue are also in heap, subtract one

// NumTxs returns the number of transactions in the priority queue. It is
// thread safe.
func (pq *TxPriorityQueue) NumTxs() int { _ = "STUB: not implemented"; return 0 }

func (pq *TxPriorityQueue) removeQueuedEvmTxUnsafe(tx *WrappedTx) (removedIdx int) {
	_ = "STUB: not implemented"
	return 0
}

func (pq *TxPriorityQueue) findTxIndexUnsafe(tx *WrappedTx) (int, bool) {
	_ = "STUB: not implemented"
	// safety check for race situation where heapIndex is out of range of txs
	return 0, false
}

// heap index isn't trustable here, so attempt to find it

// RemoveTx removes a specific transaction from the priority queue.
func (pq *TxPriorityQueue) RemoveTx(tx *WrappedTx, shouldReenqueue bool) (toBeReenqueued []*WrappedTx) {
	_ = "STUB: not implemented"
	return nil
}

func (pq *TxPriorityQueue) pushTxUnsafe(tx *WrappedTx) { _ = "STUB: not implemented"; return }

// if there aren't other waiting txs, init and return

// this item is on the heap at the moment

// the queue's first item (and ONLY the first item) must be on the heap
// if this tx is before the first item, then we need to remove the first
// item from the heap

// PushTx adds a valid transaction to the priority queue. It is thread safe.
func (pq *TxPriorityQueue) PushTx(tx *WrappedTx) (*WrappedTx, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// tx was not inserted, and nothing was replaced

// tx replaced an existing transaction

// tx was not inserted yet, so insert it

func (pq *TxPriorityQueue) popTxUnsafe() *WrappedTx { _ = "STUB: not implemented"; return nil }

// remove the first item from the heap

// this situation is primarily for a test case that inserts nils

// non-evm transactions do not have txs waiting on a nonce

// evm transactions can have txs waiting on this nonce
// if there are any, we should replace the heap with the next nonce
// for the address

// remove the first item from the evmQueue

// if there is a next item, now it can be added to the heap

// PopTx removes the top priority transaction from the queue. It is thread safe.
func (pq *TxPriorityQueue) PopTx() *WrappedTx { _ = "STUB: not implemented"; return nil }

// dequeue up to `max` transactions and reenqueue while locked
func (pq *TxPriorityQueue) ForEachTx(handler func(wtx *WrappedTx) bool) {
	_ = "STUB: not implemented"
	return
}

// dequeue up to `max` transactions and reenqueue while locked
// TODO: use ForEachTx instead
func (pq *TxPriorityQueue) PeekTxs(max int) []*WrappedTx { _ = "STUB: not implemented"; return nil }

// Push implements the Heap interface.
//
// NOTE: A caller should never call Push. Use PushTx instead.
func (pq *TxPriorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

// Pop implements the Heap interface.
//
// NOTE: A caller should never call Pop. Use PopTx instead.
func (pq *TxPriorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// avoid memory leak
// for safety

// Len implements the Heap interface.
//
// NOTE: A caller should never call Len. Use NumTxs instead.
func (pq *TxPriorityQueue) Len() int {
	_ = "STUB: not implemented"

	// Less implements the Heap interface. It returns true if the transaction at
	// position i in the queue is of less priority than the transaction at position j.
	return 0
}

func (pq *TxPriorityQueue) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// If there exists two transactions with the same priority, consider the one
	// that we saw the earliest as the higher priority transaction.
	return false
}

// We want Pop to give us the highest, not lowest, priority so we use greater
// than here.

// Swap implements the Heap interface. It swaps two transactions in the queue.
func (pq *TxPriorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func setHeapIndex(tx *WrappedTx, i int) {
	_ = "STUB: not implemented"
	// a removed tx can be nil
	return
}
