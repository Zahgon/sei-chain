package multiversion

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// mvsMergeIterator merges a parent Iterator and a cache Iterator.
// The cache iterator may return nil keys to signal that an item
// had been deleted (but not deleted in the parent).
// If the cache iterator has the same key as the parent, the
// cache shadows (overrides) the parent.
type mvsMergeIterator struct {
	parent    types.Iterator
	cache     types.Iterator
	ascending bool
	ReadsetHandler
}

var _ types.Iterator = (*mvsMergeIterator)(nil)

func NewMVSMergeIterator(
	parent, cache types.Iterator,
	ascending bool,
	readsetHandler ReadsetHandler,
) *mvsMergeIterator {
	_ = "STUB: not implemented"
	return nil
}

// Domain implements Iterator.
// It returns the union of the iter.Parent doman, and the iter.Cache domain.
// If the domains are disjoint, this includes the domain in between them as well.
func (iter *mvsMergeIterator) Domain() (start, end []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Valid implements Iterator.
func (iter *mvsMergeIterator) Valid() bool { _ = "STUB: not implemented"; return false }

// Next implements Iterator
func (iter *mvsMergeIterator) Next() { _ = "STUB: not implemented"; return }

// If parent is invalid, get the next cache item.

// If cache is invalid, get the next parent item.

// Both are valid.  Compare keys.

// parent < cache

// parent == cache

// parent > cache

// Key implements Iterator
func (iter *mvsMergeIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// If parent is invalid, get the cache key.

// If cache is invalid, get the parent key.

// Both are valid.  Compare keys.

// parent < cache

// parent == cache

// parent > cache

// Value implements Iterator
func (iter *mvsMergeIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// If parent is invalid, get the cache value.

// If cache is invalid, get the parent value.

// add values read from parent to readset

// Both are valid.  Compare keys.

// parent < cache

// add values read from parent to readset

// parent >= cache

// Close implements Iterator
func (iter *mvsMergeIterator) Close() error { _ = "STUB: not implemented"; return nil }

// still want to close cache iterator regardless

// Error returns an error if the mvsMergeIterator is invalid defined by the
// Valid method.
func (iter *mvsMergeIterator) Error() error { _ = "STUB: not implemented"; return nil }

// If not valid, panics.
// NOTE: May have side-effect of iterating over cache.
func (iter *mvsMergeIterator) assertValid() { _ = "STUB: not implemented"; return }

// Like bytes.Compare but opposite if not ascending.
func (iter *mvsMergeIterator) compare(a, b []byte) int { _ = "STUB: not implemented"; return 0 }

// Skip all delete-items from the cache w/ `key < until`.  After this function,
// current cache item is a non-delete-item, or `until <= key`.
// If the current cache item is not a delete item, does nothing.
// If `until` is nil, there is no limit, and cache may end up invalid.
// CONTRACT: cache is valid.
func (iter *mvsMergeIterator) skipCacheDeletes(until []byte) { _ = "STUB: not implemented"; return }

// Fast forwards cache (or parent+cache in case of deleted items) until current
// item exists, or until iterator becomes invalid.
// Returns whether the iterator is valid.
func (iter *mvsMergeIterator) skipUntilExistsOrInvalid() bool {
	_ = "STUB: not implemented"

	// If parent is invalid, fast-forward cache.
	return false
}

// Parent is valid.

// Parent is valid, cache is valid.

// Compare parent and cache.

// parent < cache.

// parent == cache.
// Skip over if cache item is a delete.

// Cache is not a delete.

// cache exists.
// cache < parent
// Skip over if cache item is a delete.

// Cache is not a delete.

// cache exists.
