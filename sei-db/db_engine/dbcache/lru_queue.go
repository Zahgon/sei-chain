package dbcache

import (
	"container/list"
)

// Implements a queue-like abstraction with LRU semantics. Not thread safe.
type lruQueue struct {
	order     *list.List
	entries   map[string]*list.Element
	totalSize uint64
}

type lruQueueEntry struct {
	key  string
	size uint64
}

// Create a new LRU queue.
func newLRUQueue() *lruQueue { _ = "STUB: not implemented"; return nil }

// Add a new entry to the LRU queue. Can also be used to update an existing value with a new weight.
func (lru *lruQueue) Push(
	// the key in the cache that was recently interacted with
	key []byte,
	// the size of the key + value
	size uint64,
) {
	_ = "STUB: not implemented"
	return
}

// should be impossible

// Signal that an entry has been interacted with, moving it to the back of the queue
// (i.e. making it so it doesn't get popped soon).
func (lru *lruQueue) Touch(key []byte) { _ = "STUB: not implemented"; return }

// Returns the total size of all entries in the LRU queue.
func (lru *lruQueue) GetTotalSize() uint64 { _ = "STUB: not implemented"; return 0 }

// Returns a count of the number of entries in the LRU queue, where each entry counts for 1 regardless of size.
func (lru *lruQueue) GetCount() uint64 { _ = "STUB: not implemented"; return 0 }

// Pops a single element out of the queue. The element removed is the entry least recently passed to Update().
// Returns the key in string form to avoid copying the key an additional time.
// Panics if the queue is empty.
func (lru *lruQueue) PopLeastRecentlyUsed() string { _ = "STUB: not implemented"; return "" }

// should be impossible
