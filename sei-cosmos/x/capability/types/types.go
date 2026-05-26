package types

// NewCapability returns a reference to a new Capability to be used as an
// actual capability.
func NewCapability(index uint64) *Capability { _ = "STUB: not implemented"; return nil }

// String returns the string representation of a Capability. The string contains
// the Capability's memory reference as the string is to be used in a composite
// key and to authenticate capabilities.
func (ck *Capability) String() string { _ = "STUB: not implemented"; return "" }

func NewOwner(module, name string) Owner { _ = "STUB: not implemented"; return *new(Owner) }

// Key returns a composite key for an Owner.
func (o Owner) Key() string { _ = "STUB: not implemented"; return "" }

func (o Owner) String() string { _ = "STUB: not implemented"; return "" }

func NewCapabilityOwners() *CapabilityOwners { _ = "STUB: not implemented"; return nil }

// Set attempts to add a given owner to the CapabilityOwners. If the owner
// already exists, an error will be returned. Set runs in O(log n) average time
// and O(n) in the worst case.
func (co *CapabilityOwners) Set(owner Owner) error { _ = "STUB: not implemented"; return nil }

// owner already exists at co.Owners[i]

// owner does not exist in the set of owners, so we insert at position i
// expand by 1 in amortized O(1) / O(n) worst case

// Remove removes a provided owner from the CapabilityOwners if it exists. If the
// owner does not exist, Remove is considered a no-op.
func (co *CapabilityOwners) Remove(owner Owner) { _ = "STUB: not implemented"; return }

// owner exists at co.Owners[i]

// Get returns (i, true) of the provided owner in the CapabilityOwners if the
// owner exists, where i indicates the owner's index in the set. Otherwise
// (i, false) where i indicates where in the set the owner should be added.
func (co *CapabilityOwners) Get(owner Owner) (int, bool) {
	_ = "STUB: not implemented"
	// find smallest index s.t. co.Owners[i] >= owner in O(log n) time
	return 0, false
}

// owner exists at co.Owners[i]
