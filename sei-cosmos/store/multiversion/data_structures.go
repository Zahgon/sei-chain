package multiversion

import (
	"sync"

	"github.com/google/btree"
)

const (
	multiVersionBTreeDegree = 2
)

type MultiVersionValue interface {
	GetLatest() (value MultiVersionValueItem, found bool)
	GetLatestNonEstimate() (value MultiVersionValueItem, found bool)
	GetLatestBeforeIndex(index int) (value MultiVersionValueItem, found bool)
	Set(index int, incarnation int, value []byte)
	SetEstimate(index int, incarnation int)
	Delete(index int, incarnation int)
	Remove(index int)
}

type MultiVersionValueItem interface {
	IsDeleted() bool
	IsEstimate() bool
	Value() []byte
	Incarnation() int
	Index() int
}

type multiVersionItem struct {
	valueTree *btree.BTree // contains versions values written to this key
	mtx       sync.RWMutex // manages read + write accesses
}

var _ MultiVersionValue = (*multiVersionItem)(nil)

func NewMultiVersionItem() *multiVersionItem { _ = "STUB: not implemented"; return nil }

// GetLatest returns the latest written value to the btree, and returns a boolean indicating whether it was found.
func (item *multiVersionItem) GetLatest() (MultiVersionValueItem, bool) {
	_ = "STUB: not implemented"
	return *new(MultiVersionValueItem), false
}

// GetLatestNonEstimate returns the latest written value that isn't an ESTIMATE and returns a boolean indicating whether it was found.
// This can be used when we want to write finalized values, since ESTIMATEs can be considered to be irrelevant at that point
func (item *multiVersionItem) GetLatestNonEstimate() (MultiVersionValueItem, bool) {
	_ = "STUB: not implemented"
	return *new(MultiVersionValueItem), false
}

// only return if non-estimate

// if estimate, continue

// else we want to return

// GetLatest returns the latest written value to the btree prior to the index passed in, and returns a boolean indicating whether it was found.
//
// A `nil` value along with `found=true` indicates a deletion that has occurred and the underlying parent store doesn't need to be hit.
func (item *multiVersionItem) GetLatestBeforeIndex(index int) (MultiVersionValueItem, bool) {
	_ = "STUB: not implemented"
	return *new(MultiVersionValueItem), false
}

// we want to find the value at the index that is LESS than the current index

// start from pivot which contains our current index, and return on first item we hit.
// This will ensure we get the latest indexed value relative to our current index

func (item *multiVersionItem) Set(index int, incarnation int, value []byte) {
	_ = "STUB: not implemented"
	return
}

func (item *multiVersionItem) Delete(index int, incarnation int) { _ = "STUB: not implemented"; return }

func (item *multiVersionItem) Remove(index int) { _ = "STUB: not implemented"; return }

func (item *multiVersionItem) SetEstimate(index int, incarnation int) {
	_ = "STUB: not implemented"
	return
}

type valueItem struct {
	index       int
	incarnation int
	value       []byte
	estimate    bool
}

var _ MultiVersionValueItem = (*valueItem)(nil)

// Index implements MultiVersionValueItem.
func (v *valueItem) Index() int {
	_ = "STUB: not implemented"

	// Incarnation implements MultiVersionValueItem.
	return 0
}

func (v *valueItem) Incarnation() int { _ = "STUB: not implemented"; return 0 }

// IsDeleted implements MultiVersionValueItem.
func (v *valueItem) IsDeleted() bool { _ = "STUB: not implemented"; return false }

// IsEstimate implements MultiVersionValueItem.
func (v *valueItem) IsEstimate() bool {
	_ = "STUB: not implemented"

	// Value implements MultiVersionValueItem.
	return false
}

func (v *valueItem) Value() []byte {
	_ = "STUB: not implemented"

	// implement Less for btree.Item for valueItem
	return nil
}

func (i *valueItem) Less(other btree.Item) bool { _ = "STUB: not implemented"; return false }

func NewValueItem(index int, incarnation int, value []byte) *valueItem {
	_ = "STUB: not implemented"
	return nil
}

func NewEstimateItem(index int, incarnation int) *valueItem { _ = "STUB: not implemented"; return nil }

func NewDeletedItem(index int, incarnation int) *valueItem { _ = "STUB: not implemented"; return nil }
