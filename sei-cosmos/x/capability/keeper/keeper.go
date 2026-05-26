package keeper

import (
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "capability", "keeper")

type (
	// Keeper defines the capability module's keeper. It is responsible for provisioning,
	// tracking, and authenticating capabilities at runtime. During application
	// initialization, the keeper can be hooked up to modules through unique function
	// references so that it can identify the calling module when later invoked.
	//
	// When the initial state is loaded from disk, the keeper allows the ability to
	// create new capability keys for all previously allocated capability identifiers
	// (allocated during execution of past transactions and assigned to particular modes),
	// and keep them in a memory-only store while the chain is running.
	//
	// The keeper allows the ability to create scoped sub-keepers which are tied to
	// a single specific module.
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		capMap        *sync.Map
		scopedModules map[string]struct{}
		sealed        bool
	}

	// ScopedKeeper defines a scoped sub-keeper which is tied to a single specific
	// module provisioned by the capability keeper. Scoped keepers must be created
	// at application initialization and passed to modules, which can then use them
	// to claim capabilities they receive and retrieve capabilities which they own
	// by name, in addition to creating new capabilities & authenticating capabilities
	// passed by other modules.
	ScopedKeeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		capMap   *sync.Map
		module   string
	}
)

// NewKeeper constructs a new CapabilityKeeper instance and initializes maps
// for capability map and scopedModules map.
func NewKeeper(cdc codec.BinaryCodec, storeKey, memKey sdk.StoreKey) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// internCapability returns the canonical *Capability for the given index,
// allocating one on first sight. The returned pointer is stable for the
// lifetime of the keeper, so the pointer-derived forward-lookup key in the
// in-memory store stays consistent for a given index. Addresses the
// long-standing TODO referenced in GetCapability (cosmos-sdk#7805).
func internCapability(m *sync.Map, index uint64) *types.Capability {
	_ = "STUB: not implemented"
	return nil
}

// ScopeToModule attempts to create and return a ScopedKeeper for a given module
// by name. It will panic if the keeper is already sealed or if the module name
// already has a ScopedKeeper.
func (k *Keeper) ScopeToModule(moduleName string) ScopedKeeper {
	_ = "STUB: not implemented"
	return *new(ScopedKeeper)
}

// Seal seals the keeper to prevent further modules from creating a scoped keeper.
// Seal may be called during app initialization for applications that do not wish to create scoped keepers dynamically.
func (k *Keeper) Seal() { _ = "STUB: not implemented"; return }

// InitMemStore will assure that the module store is a memory store (it will panic if it's not)
// and willl initialize it. The function is safe to be called multiple times.
// InitMemStore must be called every time the app starts before the keeper is used (so
// `BeginBlock` or `InitChain` - whichever is first). We need access to the store so we
// can't initialize it in a constructor.
func (k *Keeper) InitMemStore(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// check if memory store has not been initialized yet by checking if initialized flag is nil.

// initialize the in-memory store for all persisted capabilities

// set the initialized flag so we don't rerun initialization logic

// IsInitialized returns true if the keeper is properly initialized, and false otherwise.
func (k *Keeper) IsInitialized(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// InitializeIndex sets the index to one (or greater) in InitChain according
// to the GenesisState. It must only be called once.
// It will panic if the provided index is 0, or if the index is already set.
func (k Keeper) InitializeIndex(ctx sdk.Context, index uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// set the global index to the passed index

// GetLatestIndex returns the latest index of the CapabilityKeeper
func (k Keeper) GetLatestIndex(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

// SetOwners set the capability owners to the store
func (k Keeper) SetOwners(ctx sdk.Context, index uint64, owners types.CapabilityOwners) {
	_ = "STUB: not implemented"
	return
}

// set owners in persistent store

// GetOwners returns the capability owners with a given index.
func (k Keeper) GetOwners(ctx sdk.Context, index uint64) (types.CapabilityOwners, bool) {
	_ = "STUB: not implemented"
	return *new(types.CapabilityOwners), false
}

// get owners for index from persistent store

// InitializeCapability takes in an index and an owners array. It creates the capability in memory
// and sets the fwd and reverse keys for each owner in the memstore.
// It is used during initialization from genesis.
func (k Keeper) InitializeCapability(ctx sdk.Context, index uint64, owners types.CapabilityOwners) {
	_ = "STUB: not implemented"
	return
}

// Set the forward mapping between the module and capability tuple and the
// capability name in the memKVStore

// Set the reverse mapping between the module and capability name and the
// index in the in-memory store. Since marshalling and unmarshalling into a store
// will change memory address of capability, we simply store index as value here
// and retrieve the in-memory pointer to the capability from our map

// NewCapability attempts to create a new capability with a given name. If the
// capability already exists in the in-memory store, an error will be returned.
// Otherwise, a new capability is created with the current global unique index.
// The newly created capability has the scoped module name and capability name
// tuple set as the initial owner. Finally, the global index is incremented along
// with forward and reverse indexes set in the in-memory store.
//
// Note, namespacing is completely local, which is safe since records are prefixed
// with the module name and no two ScopedKeeper can have the same module name.
func (sk ScopedKeeper) NewCapability(ctx sdk.Context, name string) (*types.Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create new capability with the current global index

// update capability owner set

// increment global index

// Set the forward mapping between the module and capability tuple and the
// capability name in the memKVStore

// Set the reverse mapping between the module and capability name and the
// index in the in-memory store. Since marshalling and unmarshalling into a store
// will change memory address of capability, we simply store index as value here
// and retrieve the in-memory pointer to the capability from our map

// AuthenticateCapability attempts to authenticate a given capability and name
// from a caller. It allows for a caller to check that a capability does in fact
// correspond to a particular name. The scoped keeper will lookup the capability
// from the internal in-memory store and check against the provided name. It returns
// true upon success and false upon failure.
//
// Note, the capability's forward mapping is indexed by a string which should
// contain its unique memory reference.
func (sk ScopedKeeper) AuthenticateCapability(ctx sdk.Context, c *types.Capability, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// ClaimCapability attempts to claim a given Capability. The provided name and
// the scoped module's name tuple are treated as the owner. It will attempt
// to add the owner to the persistent set of capability owners for the capability
// index. If the owner already exists, it will return an error. Otherwise, it will
// also set a forward and reverse index for the capability and capability name.
func (sk ScopedKeeper) ClaimCapability(ctx sdk.Context, c *types.Capability, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// update capability owner set

// Set the forward mapping between the module and capability tuple and the
// capability name in the memKVStore

// Set the reverse mapping between the module and capability name and the
// index in the in-memory store. Since marshalling and unmarshalling into a store
// will change memory address of capability, we simply store index as value here
// and retrieve the in-memory pointer to the capability from our map

// ReleaseCapability allows a scoped module to release a capability which it had
// previously claimed or created. After releasing the capability, if no more
// owners exist, the capability will be globally removed.
func (sk ScopedKeeper) ReleaseCapability(ctx sdk.Context, c *types.Capability) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete the forward mapping between the module and capability tuple and the
// capability name in the memKVStore

// Delete the reverse mapping between the module and capability name and the
// index in the in-memory store.

// remove owner

// remove capability owner set

// The persistent owners store and the memstore reverse-lookup deleted
// above are authoritative for capability existence; leaving the index
// interned in capMap is harmless. See cosmos-sdk#7805.

// update capability owner set

// GetCapability allows a module to fetch a capability which it previously claimed
// by name. The module is not allowed to retrieve capabilities which it does not
// own.
func (sk ScopedKeeper) GetCapability(ctx sdk.Context, name string) (*types.Capability, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetCapabilityName allows a module to retrieve the name under which it stored a given
// capability given the capability
func (sk ScopedKeeper) GetCapabilityName(ctx sdk.Context, c *types.Capability) string {
	_ = "STUB: not implemented"
	return ""
}

// GetOwners all the Owners that own the capability associated with the name this ScopedKeeper uses
// to refer to the capability
func (sk ScopedKeeper) GetOwners(ctx sdk.Context, name string) (*types.CapabilityOwners, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// LookupModules returns all the module owners for a given capability
// as a string array and the capability itself.
// The method returns an error if either the capability or the owners cannot be
// retreived from the memstore.
func (sk ScopedKeeper) LookupModules(ctx sdk.Context, name string) ([]string, *types.Capability, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (sk ScopedKeeper) addOwner(ctx sdk.Context, c *types.Capability, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// update capability owner set

func (sk ScopedKeeper) getOwners(ctx sdk.Context, c *types.Capability) *types.CapabilityOwners {
	_ = "STUB: not implemented"
	return nil
}
