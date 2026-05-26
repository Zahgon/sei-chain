package query

import (
	"math"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	db "github.com/tendermint/tm-db"
)

// DefaultLimit is the default `limit` for queries
// if the `limit` is not supplied, paginate will use `DefaultLimit`
const DefaultLimit = 100

// MaxLimit is the maximum limit the paginate function can handle
// which equals the maximum value that can be stored in uint64
const MaxLimit = math.MaxUint64

// ParsePagination validate PageRequest and returns page number & limit.
func ParsePagination(pageReq *PageRequest) (page, limit int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// #nosec G115 -- overflow checked below
// #nosec G115 -- overflow checked below

// Paginate does pagination of all the results in the PrefixStore based on the
// provided PageRequest. onResult should be used to do actual unmarshaling.
func Paginate(
	prefixStore types.KVStore,
	pageRequest *PageRequest,
	onResult func(key []byte, value []byte) error,
) (*PageResponse, error) {
	_ = "STUB: not implemented"

	// if the PageRequest is nil, use default PageRequest
	return nil, nil
}

// count total results when the limit is zero/not supplied

func getIterator(prefixStore types.KVStore, start []byte, reverse bool) db.Iterator {
	_ = "STUB: not implemented"
	return *new(db.Iterator)
}
