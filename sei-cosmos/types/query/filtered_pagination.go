package query

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// FilteredPaginate does pagination of all the results in the PrefixStore based on the
// provided PageRequest. onResult should be used to do actual unmarshaling and filter the results.
// If key is provided, the pagination uses the optimized querying.
// If offset is used, the pagination uses lazy filtering i.e., searches through all the records.
// The accumulate parameter represents if the response is valid based on the offset given.
// It will be false for the results (filtered) < offset  and true for `offset > accumulate <= end`.
// When accumulate is set to true the current result should be appended to the result set returned
// to the client.
func FilteredPaginate(
	prefixStore types.KVStore,
	pageRequest *PageRequest,
	onResult func(key []byte, value []byte, accumulate bool) (bool, error),
) (*PageResponse, error) {
	_ = "STUB: not implemented"

	// if the PageRequest is nil, use default PageRequest
	return nil, nil
}

// count total results when the limit is zero/not supplied

// GenericFilteredPaginate does pagination of all the results in the PrefixStore based on the
// provided PageRequest. `onResult` should be used to filter or transform the results.
// `c` is a constructor function that needs to return a new instance of the type T (this is to
// workaround some generic pitfalls in which we can't instantiate a T struct inside the function).
// If key is provided, the pagination uses the optimized querying.
// If offset is used, the pagination uses lazy filtering i.e., searches through all the records.
// The resulting slice (of type F) can be of a different type than the one being iterated through
// (type T), so it's possible to do any necessary transformation inside the onResult function.
func GenericFilteredPaginate[T codec.ProtoMarshaler, F codec.ProtoMarshaler](
	cdc codec.BinaryCodec,
	prefixStore types.KVStore,
	pageRequest *PageRequest,
	onResult func(key []byte, value T) (F, error),
	constructor func() T,
) ([]F, *PageResponse, error) {
	_ = "STUB: not implemented"
	// if the PageRequest is nil, use default PageRequest
	return nil, nil, nil
}

// count total results when the limit is zero/not supplied

// Previously this was the "accumulate" flag
