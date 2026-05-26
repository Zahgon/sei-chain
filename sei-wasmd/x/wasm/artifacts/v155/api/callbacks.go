package api

// Check https://akrennmair.github.io/golang-cgo-slides/ to learn
// how this embedded C code works.

/*
#include "bindings155.h"

// typedefs for _cgo functions (db)
typedef GoError155 (*read_db_fn)(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView key, UnmanagedVector *val, UnmanagedVector *errOut);
typedef GoError155 (*write_db_fn)(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView key, U8SliceView val, UnmanagedVector *errOut);
typedef GoError155 (*remove_db_fn)(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView key, UnmanagedVector *errOut);
typedef GoError155 (*scan_db_fn)(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView start, U8SliceView end, int32_t order, GoIter *out, UnmanagedVector *errOut);
// iterator
typedef GoError155 (*db_next)(iterator_t idx, gas_meter_t *gas_meter, uint64_t *used_gas, UnmanagedVector *key, UnmanagedVector *val, UnmanagedVector *errOut);
typedef GoError155 (*db_next_key)(iterator_t idx, gas_meter_t *gas_meter, uint64_t *used_gas, UnmanagedVector *key, UnmanagedVector *errOut);
typedef GoError155 (*db_next_value)(iterator_t idx, gas_meter_t *gas_meter, uint64_t *used_gas, UnmanagedVector *val, UnmanagedVector *errOut);
// and api
typedef GoError155 (*humanize_address_fn)(api_t *ptr, U8SliceView src, UnmanagedVector *dest, UnmanagedVector *errOut, uint64_t *used_gas);
typedef GoError155 (*canonicalize_address_fn)(api_t *ptr, U8SliceView src, UnmanagedVector *dest, UnmanagedVector *errOut, uint64_t *used_gas);
typedef GoError155 (*query_external_fn)(querier_t *ptr, uint64_t gas_limit, uint64_t *used_gas, U8SliceView request, UnmanagedVector *result, UnmanagedVector *errOut);

// forward declarations (db)
GoError155 cGet155_cgo(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView key, UnmanagedVector *val, UnmanagedVector *errOut);
GoError155 cSet155_cgo(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView key, U8SliceView val, UnmanagedVector *errOut);
GoError155 cDelete155_cgo(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView key, UnmanagedVector *errOut);
GoError155 cScan155_cgo(db_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, U8SliceView start, U8SliceView end, int32_t order, GoIter *out, UnmanagedVector *errOut);
// iterator
GoError155 cNext155_cgo(iterator_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, UnmanagedVector *key, UnmanagedVector *val, UnmanagedVector *errOut);
GoError155 cNextKey155_cgo(iterator_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, UnmanagedVector *key, UnmanagedVector *errOut);
GoError155 cNextValue155_cgo(iterator_t *ptr, gas_meter_t *gas_meter, uint64_t *used_gas, UnmanagedVector *val, UnmanagedVector *errOut);
// api
GoError155 cHumanAddress155_cgo(api_t *ptr, U8SliceView src, UnmanagedVector *dest, UnmanagedVector *errOut, uint64_t *used_gas);
GoError155 cCanonicalAddress155_cgo(api_t *ptr, U8SliceView src, UnmanagedVector *dest, UnmanagedVector *errOut, uint64_t *used_gas);
// and querier
GoError155 cQueryExternal155_cgo(querier_t *ptr, uint64_t gas_limit, uint64_t *used_gas, U8SliceView request, UnmanagedVector *result, UnmanagedVector *errOut);


*/
import "C"

import (
	"github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

// Note: we have to include all exports in the same file (at least since they both import bindings155.h),
// or get odd cgo build errors about duplicate definitions

func recoverPanic(ret *C.GoError155) { _ = "STUB: not implemented"; return }

// This is used to handle ErrorOutOfGas panics.
//
// What we do here is something that should not be done in the first place.
// "A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast
// on errors that shouldn’t occur during normal operation, or that we aren’t prepared to
// handle gracefully." says https://gobyexample.com/panic.
// And 'Ask yourself "when this happens, should the application immediately crash?" If yes,
// use a panic; otherwise, use an error.' says this popular answer on SO: https://stackoverflow.com/a/44505268.
// Oh, and "If you're already worrying about discriminating different kinds of panics, you've lost sight of the ball."
// (Rob Pike) from https://eli.thegreenplace.net/2018/on-the-uses-and-misuses-of-panics-in-go/
//
// We don't want to import Cosmos SDK and also cannot use interfaces to detect these
// error types (as they have no methods). So, let's just rely on the descriptive names.

// These three types are "thrown" (which is not a thing in Go 🙃) in panics from the gas module
// (https://github.com/cosmos/cosmos-sdk/blob/v0.45.4/store/types/gas.go):
// 1. ErrorOutOfGas
// 2. ErrorGasOverflow
// 3. ErrorNegativeGasConsumed
//
// In the baseapp, ErrorOutOfGas gets special treatment:
// - https://github.com/cosmos/cosmos-sdk/blob/v0.45.4/baseapp/baseapp.go#L607
// - https://github.com/cosmos/cosmos-sdk/blob/v0.45.4/baseapp/recovery.go#L50-L60
// This turns the panic into a regular error with a helpful error message.
//
// The other two gas related panic types indicate programming errors and are handled along
// with all other errors in https://github.com/cosmos/cosmos-sdk/blob/v0.45.4/baseapp/recovery.go#L66-L77.

// TODO: figure out how to pass the text in its `Descriptor` field through all the FFI

/****** DB ********/

var db_vtable = C.Db_vtable{
	read_db:   (C.read_db_fn)(C.cGet155_cgo),
	write_db:  (C.write_db_fn)(C.cSet155_cgo),
	remove_db: (C.remove_db_fn)(C.cDelete155_cgo),
	scan_db:   (C.scan_db_fn)(C.cScan155_cgo),
}

type DBState struct {
	Store types.KVStore
	// CallID is used to lookup the proper frame for iterators associated with this contract call (iterator.go)
	CallID uint64
}

// use this to create C.Db in two steps, so the pointer lives as long as the calling stack
//
//	state := buildDBState(kv, callID)
//	db := buildDB(&state, &gasMeter)
//	// then pass db into some FFI function
func buildDBState(kv types.KVStore, callID uint64) DBState {
	_ = "STUB: not implemented"
	return *new(DBState)
}

// contract: original pointer/struct referenced must live longer than C.Db struct
// since this is only used internally, we can verify the code that this is the case
func buildDB(state *DBState, gm *types.GasMeter) C.Db { _ = "STUB: not implemented"; return *new(C.Db) }

var iterator_vtable = C.Iterator_vtable{
	next:       (C.db_next)(C.cNext155_cgo),
	next_key:   (C.db_next_key)(C.cNextKey155_cgo),
	next_value: (C.db_next_value)(C.cNextValue155_cgo),
}

// An iterator including referenced objects is 117 bytes large (calculated using https://github.com/DmitriyVTitov/size).
// We limit the number of iterators per contract call ID here in order limit memory usage to 32768*117 = ~3.8 MB as a safety measure.
// In any reasonable contract, gas limits should hit sooner than that though.
const frameLenLimit = 32768

// contract: original pointer/struct referenced must live longer than C.Db struct
// since this is only used internally, we can verify the code that this is the case
func buildIterator(callID uint64, it types.Iterator) (C.iterator_t, error) {
	_ = "STUB: not implemented"
	return *new(C.iterator_t), nil
}

//export cGet155
func cGet155(ptr *C.db_t, gasMeter *C.gas_meter_t, usedGas *cu64, key C.U8SliceView, val *C.UnmanagedVector, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// we received an invalid pointer

// v will equal nil when the key is missing
// https://github.com/cosmos/cosmos-sdk/blob/1083fa948e347135861f88e07ec76b0314296832/store/types/store.go#L174

//export cSet155
func cSet155(ptr *C.db_t, gasMeter *C.gas_meter_t, usedGas *cu64, key C.U8SliceView, val C.U8SliceView, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// we received an invalid pointer

//export cDelete155
func cDelete155(ptr *C.db_t, gasMeter *C.gas_meter_t, usedGas *cu64, key C.U8SliceView, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// we received an invalid pointer

//export cScan155
func cScan155(ptr *C.db_t, gasMeter *C.gas_meter_t, usedGas *cu64, start C.U8SliceView, end C.U8SliceView, order ci32, out *C.GoIter, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// we received an invalid pointer

// Ascending

// Descending

// store the actual error message in the return buffer

//export cNext155
func cNext155(ref C.iterator_t, gasMeter *C.gas_meter_t, usedGas *cu64, key *C.UnmanagedVector, val *C.UnmanagedVector, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	// typical usage of iterator
	// 	for ; itr.Valid(); itr.Next() {
	// 		k, v := itr.Key(); itr.Value()
	// 		...
	// 	}
	return *new(C.GoError155)
}

// we received an invalid pointer

// end of iterator, return as no-op, nil key is considered end

// call Next at the end, upon creation we have first data loaded

// check iter.Error() ????

//export cNextKey155
func cNextKey155(ref C.iterator_t, gasMeter *C.gas_meter_t, usedGas *cu64, key *C.UnmanagedVector, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

//export cNextValue155
func cNextValue155(ref C.iterator_t, gasMeter *C.gas_meter_t, usedGas *cu64, value *C.UnmanagedVector, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// nextPart is a helper function that contains the shared code for key- and value-only iteration.
func nextPart(ref C.iterator_t, gasMeter *C.gas_meter_t, usedGas *cu64, output *C.UnmanagedVector, errOut *C.UnmanagedVector, valFn func(types.Iterator) []byte) (ret C.GoError155) {
	_ = "STUB: not implemented"
	// typical usage of iterator
	// 	for ; itr.Valid(); itr.Next() {
	// 		k, v := itr.Key(); itr.Value()
	// 		...
	// 	}
	return *new(C.GoError155)
}

// we received an invalid pointer

// end of iterator, return as no-op, nil `output` is considered end

// call Next at the end, upon creation we have first data loaded

// check iter.Error() ????

var api_vtable = C.GoApi_vtable{
	humanize_address:     (C.humanize_address_fn)(C.cHumanAddress155_cgo),
	canonicalize_address: (C.canonicalize_address_fn)(C.cCanonicalAddress155_cgo),
}

// contract: original pointer/struct referenced must live longer than C.GoApi struct
// since this is only used internally, we can verify the code that this is the case
func buildAPI(api *types.GoAPI) C.GoApi { _ = "STUB: not implemented"; return *new(C.GoApi) }

//export cHumanAddress155
func cHumanAddress155(ptr *C.api_t, src C.U8SliceView, dest *C.UnmanagedVector, errOut *C.UnmanagedVector, used_gas *cu64) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// store the actual error message in the return buffer

//export cCanonicalAddress155
func cCanonicalAddress155(ptr *C.api_t, src C.U8SliceView, dest *C.UnmanagedVector, errOut *C.UnmanagedVector, used_gas *cu64) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// store the actual error message in the return buffer

/****** Go Querier ********/

var querier_vtable = C.Querier_vtable{
	query_external: (C.query_external_fn)(C.cQueryExternal155_cgo),
}

// contract: original pointer/struct referenced must live longer than C.GoQuerier struct
// since this is only used internally, we can verify the code that this is the case
func buildQuerier(q *Querier) C.GoQuerier { _ = "STUB: not implemented"; return *new(C.GoQuerier) }

//export cQueryExternal155
func cQueryExternal155(ptr *C.querier_t, gasLimit cu64, usedGas *cu64, request C.U8SliceView, result *C.UnmanagedVector, errOut *C.UnmanagedVector) (ret C.GoError155) {
	_ = "STUB: not implemented"
	return *new(C.GoError155)
}

// we received an invalid pointer

// query the data

// serialize the response
