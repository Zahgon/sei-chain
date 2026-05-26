package types

// KVStorePrefixIteratorPaginated returns iterator over items in the selected page.
// Items iterated and skipped in ascending order.
func KVStorePrefixIteratorPaginated(kvs KVStore, prefix []byte, page, limit uint) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

// KVStoreReversePrefixIteratorPaginated returns iterator over items in the selected page.
// Items iterated and skipped in descending order.
func KVStoreReversePrefixIteratorPaginated(kvs KVStore, prefix []byte, page, limit uint) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

// PaginatedIterator is a wrapper around Iterator that iterates over values starting for given page and limit.
type PaginatedIterator struct {
	Iterator

	page, limit uint // provided during initialization
	iterated    uint // incremented in a call to Next

}

func (pi *PaginatedIterator) skip() { _ = "STUB: not implemented"; return }

// Next will panic after limit is reached.
func (pi *PaginatedIterator) Next() { _ = "STUB: not implemented"; return }

// Valid if below limit and underlying iterator is valid.
func (pi *PaginatedIterator) Valid() bool { _ = "STUB: not implemented"; return false }
