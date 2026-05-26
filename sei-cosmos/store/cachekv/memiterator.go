package cachekv

import (
	"sync"

	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// Iterates over iterKVCache items.
// if key is nil, means it was deleted.
// Implements Iterator.
type memIterator struct {
	types.Iterator

	lastKey []byte
	deleted *sync.Map
}

func newMemIterator(
	start, end []byte,
	items *dbm.MemDB,
	deleted *sync.Map,
	ascending bool,
) *memIterator {
	_ = "STUB: not implemented"
	return nil
}

func (mi *memIterator) Value() []byte {
	_ = "STUB: not implemented"

	// We need to handle the case where deleted is modified and includes our current key
	// We handle this by maintaining a lastKey object in the iterator.
	// If the current key is the same as the last key (and last key is not nil / the start)
	// then we are calling value on the same thing as last time.
	// Therefore we don't check the mi.deleted to see if this key is included in there.
	return nil
}
