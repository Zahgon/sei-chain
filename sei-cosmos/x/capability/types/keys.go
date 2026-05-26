package types

const (
	// ModuleName defines the module name
	ModuleName = "capability"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

var (
	// KeyIndex defines the key that stores the current globally unique capability
	// index.
	KeyIndex = []byte("index")

	// KeyPrefixIndexCapability defines a key prefix that stores index to capability
	// name mappings.
	KeyPrefixIndexCapability = []byte("capability_index")

	// KeyMemInitialized defines the key that stores the initialized flag in the memory store
	KeyMemInitialized = []byte("mem_initialized")
)

// RevCapabilityKey returns a reverse lookup key for a given module and capability
// name.
func RevCapabilityKey(module, name string) []byte { _ = "STUB: not implemented"; return nil }

// FwdCapabilityKey returns a forward lookup key for a given module and capability
// reference.
func FwdCapabilityKey(module string, cap *Capability) []byte {
	_ = "STUB: not implemented"
	// encode the key to a fixed length to avoid breaking consensus state machine
	// it's a hacky backport of https://github.com/cosmos/cosmos-sdk/pull/11737
	// the length 10 is picked so it's backward compatible on common architectures.
	return nil
}

// IndexToKey returns bytes to be used as a key for a given capability index.
func IndexToKey(index uint64) []byte { _ = "STUB: not implemented"; return nil }

// IndexFromKey returns an index from a call to IndexToKey for a given capability
// index.
func IndexFromKey(key []byte) uint64 { _ = "STUB: not implemented"; return 0 }
