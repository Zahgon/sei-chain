package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
)

type (
	PruningOptions = types.PruningOptions
)

type (
	Store                     = types.Store
	Committer                 = types.Committer
	CommitStore               = types.CommitStore
	Queryable                 = types.Queryable
	MultiStore                = types.MultiStore
	CacheMultiStore           = types.CacheMultiStore
	CommitMultiStore          = types.CommitMultiStore
	MultiStorePersistentCache = types.MultiStorePersistentCache
	GigaMultiStore            = types.GigaMultiStore
	KVStore                   = types.KVStore
	Iterator                  = types.Iterator
)

// StoreDecoderRegistry defines each of the modules store decoders. Used for ImportExport
// simulation.
type StoreDecoderRegistry map[string]func(kvA, kvB kv.Pair) string

// Iterator over all the keys with a certain prefix in ascending order
func KVStorePrefixIterator(kvs KVStore, prefix []byte) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

// Iterator over all the keys with a certain prefix in descending order.
func KVStoreReversePrefixIterator(kvs KVStore, prefix []byte) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

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

// DiffKVStores compares two KVstores and returns all the key/value pairs
// that differ from one another. It also skips value comparison for a set of provided prefixes
func DiffKVStores(a KVStore, b KVStore, prefixesToSkip [][]byte) (kvAs, kvBs []kv.Pair) {
	_ = "STUB: not implemented"
	return nil, nil
}

type (
	CacheKVStore  = types.CacheKVStore
	CommitKVStore = types.CommitKVStore
	CacheWrap     = types.CacheWrap
	CacheWrapper  = types.CacheWrapper
	CommitID      = types.CommitID
)

type StoreType = types.StoreType

const (
	StoreTypeMulti     = types.StoreTypeMulti
	StoreTypeDB        = types.StoreTypeDB
	StoreTypeIAVL      = types.StoreTypeIAVL
	StoreTypeTransient = types.StoreTypeTransient
	StoreTypeMemory    = types.StoreTypeMemory
)

type (
	StoreKey          = types.StoreKey
	CapabilityKey     = types.CapabilityKey
	KVStoreKey        = types.KVStoreKey
	TransientStoreKey = types.TransientStoreKey
	MemoryStoreKey    = types.MemoryStoreKey
)

// assertNoCommonPrefix will panic if there are two keys: k1 and k2 in keys, such that
// k1 is a prefix of k2
func assertNoPrefix(keys []string) { _ = "STUB: not implemented"; return }

// NewKVStoreKey returns a new pointer to a KVStoreKey.
func NewKVStoreKey(name string) *KVStoreKey { _ = "STUB: not implemented"; return nil }

// NewKVStoreKeys returns a map of new  pointers to KVStoreKey's.
// The function will panic if there is a potential conflict in names (see `assertNoPrefix`
// function for more details).
func NewKVStoreKeys(names ...string) map[string]*KVStoreKey { _ = "STUB: not implemented"; return nil }

// Constructs new TransientStoreKey
// Must return a pointer according to the ocap principle
func NewTransientStoreKey(name string) *TransientStoreKey { _ = "STUB: not implemented"; return nil }

// NewTransientStoreKeys constructs a new map of TransientStoreKey's
// Must return pointers according to the ocap principle
// The function will panic if there is a potential conflict in names (see `assertNoPrefix`
// function for more details).
func NewTransientStoreKeys(names ...string) map[string]*TransientStoreKey {
	_ = "STUB: not implemented"
	return nil
}

// NewMemoryStoreKeys constructs a new map matching store key names to their
// respective MemoryStoreKey references.
// The function will panic if there is a potential conflict in names (see `assertNoPrefix`
// function for more details).
func NewMemoryStoreKeys(names ...string) map[string]*MemoryStoreKey {
	_ = "STUB: not implemented"
	return nil
}

// PrefixEndBytes returns the []byte that would end a
// range query for all []byte with a certain prefix
// Deals with last byte of prefix being FF without overflowing
func PrefixEndBytes(prefix []byte) []byte { _ = "STUB: not implemented"; return nil }

// InclusiveEndBytes returns the []byte that would end a
// range query such that the input would be included
func InclusiveEndBytes(inclusiveBytes []byte) (exclusiveBytes []byte) {
	_ = "STUB: not implemented"
	return nil
}

//----------------------------------------

// key-value result for iterator queries
type KVPair = types.KVPair

//----------------------------------------

// TraceContext contains TraceKVStore context data. It will be written with
// every trace operation.
type TraceContext = types.TraceContext

// --------------------------------------

type (
	Gas       = types.Gas
	GasMeter  = types.GasMeter
	GasConfig = types.GasConfig
)

func NewGasMeter(limit Gas, multiplierNumerator uint64, multiplierDenominator uint64) GasMeter {
	_ = "STUB: not implemented"
	return *new(GasMeter)
}

type (
	ErrorOutOfGas    = types.ErrorOutOfGas
	ErrorGasOverflow = types.ErrorGasOverflow
)

func NewInfiniteGasMeter(multiplierNumerator uint64, multiplierDenominator uint64) GasMeter {
	_ = "STUB: not implemented"
	return *new(GasMeter)
}

// Helpers for setting gas meter with parent ctx multiplier
func NewGasMeterWithMultiplier(ctx Context, limit uint64) GasMeter {
	_ = "STUB: not implemented"
	return *new(GasMeter)
}

func NewInfiniteGasMeterWithMultiplier(ctx Context) GasMeter {
	_ = "STUB: not implemented"
	return *new(GasMeter)
}
