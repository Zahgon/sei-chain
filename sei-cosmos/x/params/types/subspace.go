package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// StoreKey is the string store key for the param store
	StoreKey = "params"

	// TStoreKey is the string store key for the param transient store
	TStoreKey = "transient_params"
)

// Individual parameter store for each keeper
// Transient store persists for a block, so we use it for
// recording whether the parameter has been changed or not
type Subspace struct {
	cdc          codec.BinaryCodec
	legacyAmino  *codec.LegacyAmino
	key          sdk.StoreKey // []byte -> []byte, stores parameter
	tkey         sdk.StoreKey // []byte -> bool, stores parameter change
	name         []byte
	suffixedName []byte
	table        KeyTable
}

// NewSubspace constructs a store with namestore
func NewSubspace(cdc codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key sdk.StoreKey, tkey sdk.StoreKey, name string) Subspace {
	_ = "STUB: not implemented"
	return *new(Subspace)
}

// HasKeyTable returns if the Subspace has a KeyTable registered.
func (s Subspace) HasKeyTable() bool { _ = "STUB: not implemented"; return false }

// WithKeyTable initializes KeyTable and returns modified Subspace
func (s Subspace) WithKeyTable(table KeyTable) Subspace {
	_ = "STUB: not implemented"
	return *new(Subspace)
}

// Allocate additional capacity for Subspace.name
// So we don't have to allocate extra space each time appending to the key

// Returns a KVStore identical with ctx.KVStore(s.key).Prefix()
func (s Subspace) kvStore(ctx sdk.Context) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

// Returns a transient store for modification
func (s Subspace) transientStore(ctx sdk.Context) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

// Validate attempts to validate a parameter value by its key. If the key is not
// registered or if the validation of the value fails, an error is returned.
func (s Subspace) Validate(ctx sdk.Context, key []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Get queries for a parameter by key from the Subspace's KVStore and sets the
// value to the provided pointer. If the value does not exist, it will panic.
func (s Subspace) Get(ctx sdk.Context, key []byte, ptr interface{}) {
	_ = "STUB: not implemented"
	return
}

// GetIfExists queries for a parameter by key from the Subspace's KVStore and
// sets the value to the provided pointer. If the value does not exist, it will
// perform a no-op.
func (s Subspace) GetIfExists(ctx sdk.Context, key []byte, ptr interface{}) {
	_ = "STUB: not implemented"
	return
}

// GetRaw queries for the raw values bytes for a parameter by key.
func (s Subspace) GetRaw(ctx sdk.Context, key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Has returns if a parameter key exists or not in the Subspace's KVStore.
func (s Subspace) Has(ctx sdk.Context, key []byte) bool { _ = "STUB: not implemented"; return false }

// Modified returns true if the parameter key is set in the Subspace's transient
// KVStore.
func (s Subspace) Modified(ctx sdk.Context, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// checkType verifies that the provided key and value are comptable and registered.
func (s Subspace) checkType(key []byte, value interface{}) { _ = "STUB: not implemented"; return }

// Set stores a value for given a parameter key assuming the parameter type has
// been registered. It will panic if the parameter type has not been registered
// or if the value cannot be encoded. A change record is also set in the Subspace's
// transient KVStore to mark the parameter as modified.
func (s Subspace) Set(ctx sdk.Context, key []byte, value interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s Subspace) SetRaw(ctx sdk.Context, key []byte, value []byte) {
	_ = "STUB: not implemented"
	return
}

// Update stores an updated raw value for a given parameter key assuming the
// parameter type has been registered. It will panic if the parameter type has
// not been registered or if the value cannot be encoded. An error is returned
// if the raw value is not compatible with the registered type for the parameter
// key or if the new value is invalid as determined by the registered type's
// validation function.
func (s Subspace) Update(ctx sdk.Context, key, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// destValue contains the dereferenced value of dest so validation function do
// not have to operate on pointers.

// GetParamSet iterates through each ParamSetPair where for each pair, it will
// retrieve the value and set it to the corresponding value pointer provided
// in the ParamSetPair by calling Subspace#Get.
func (s Subspace) GetParamSet(ctx sdk.Context, ps ParamSet) { _ = "STUB: not implemented"; return }

// GetParamSetIfExists iterates through each ParamSetPair where for each pair, it will
// retrieve the value and set it to the corresponding value pointer provided
// in the ParamSetPair by calling Subspace#GetIfExists.
func (s Subspace) GetParamSetIfExists(ctx sdk.Context, ps ParamSet) {
	_ = "STUB: not implemented"
	return
}

// SetParamSet iterates through each ParamSetPair and sets the value with the
// corresponding parameter key in the Subspace's KVStore.
func (s Subspace) SetParamSet(ctx sdk.Context, ps ParamSet) { _ = "STUB: not implemented"; return }

// pair.Field is a pointer to the field, so indirecting the ptr.
// go-amino automatically handles it but just for sure,
// since SetStruct is meant to be used in InitGenesis
// so this method will not be called frequently

// Name returns the name of the Subspace.
func (s Subspace) Name() string { _ = "STUB: not implemented"; return "" }

// Wrapper of Subspace, provides immutable functions only
type ReadOnlySubspace struct {
	s Subspace
}

// Get delegates a read-only Get call to the Subspace.
func (ros ReadOnlySubspace) Get(ctx sdk.Context, key []byte, ptr interface{}) {
	_ = "STUB: not implemented"
	return
}

// GetRaw delegates a read-only GetRaw call to the Subspace.
func (ros ReadOnlySubspace) GetRaw(ctx sdk.Context, key []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Has delegates a read-only Has call to the Subspace.
func (ros ReadOnlySubspace) Has(ctx sdk.Context, key []byte) bool {
	_ = "STUB: not implemented"
	return false

	// Modified delegates a read-only Modified call to the Subspace.
}

func (ros ReadOnlySubspace) Modified(ctx sdk.Context, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Name delegates a read-only Name call to the Subspace.
func (ros ReadOnlySubspace) Name() string { _ = "STUB: not implemented"; return "" }
