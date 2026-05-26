package vtype

// PendingAccountWrite tracks field-level changes to an account that have not yet been committed.
// Each field has a value and a flag indicating whether it has been set. Only set fields are
// applied when merging into a base AccountData.
//
// It is legal to operate on a nil PendingAccountWrite. A nil PendingAccountWrite will always return 0s from getters,
// and will return a non-nil result when a setter is called.
type PendingAccountWrite struct {
	balance  *Balance
	nonce    uint64
	nonceSet bool
	codeHash *CodeHash
}

// NewPendingAccountWrite creates a new PendingAccountWrite with no fields set.
func NewPendingAccountWrite() *PendingAccountWrite { _ = "STUB: not implemented"; return nil }

// GetBalance returns the pending balance value, or nil if not set.
func (p *PendingAccountWrite) GetBalance() *Balance { _ = "STUB: not implemented"; return nil }

// IsBalanceSet reports whether the balance has been set in this pending write.
func (p *PendingAccountWrite) IsBalanceSet() bool { _ = "STUB: not implemented"; return false }

// GetNonce returns the pending nonce value.
func (p *PendingAccountWrite) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

// IsNonceSet reports whether the nonce has been set in this pending write.
func (p *PendingAccountWrite) IsNonceSet() bool { _ = "STUB: not implemented"; return false }

// GetCodeHash returns the pending code hash value, or nil if not set.
func (p *PendingAccountWrite) GetCodeHash() *CodeHash { _ = "STUB: not implemented"; return nil }

// IsCodeHashSet reports whether the code hash has been set in this pending write.
func (p *PendingAccountWrite) IsCodeHashSet() bool { _ = "STUB: not implemented"; return false }

// SetBalance marks the balance as changed. A nil balance is treated as all zeros.
// The pointer is stored directly; the caller must not modify the underlying array
// after calling SetBalance. Returns self.
func (p *PendingAccountWrite) SetBalance(balance *Balance) *PendingAccountWrite {
	_ = "STUB: not implemented"
	return nil
}

// SetNonce marks the nonce as changed. Returns self.
func (p *PendingAccountWrite) SetNonce(nonce uint64) *PendingAccountWrite {
	_ = "STUB: not implemented"
	return nil
}

// SetCodeHash marks the code hash as changed. A nil code hash is treated as all zeros.
// The pointer is stored directly; the caller must not modify the underlying array
// after calling SetCodeHash. Returns self.
func (p *PendingAccountWrite) SetCodeHash(codeHash *CodeHash) *PendingAccountWrite {
	_ = "STUB: not implemented"
	return nil
}

// Merge applies the pending field changes onto a copy of the base AccountData, updating the
// block height. Only fields that have been set via Set* methods are overwritten; all other
// fields are carried over from the base. The base is not modified. If a nil base is provided,
// the pending writes are applied to a new AccountData instantiated to all 0s.
func (p *PendingAccountWrite) Merge(base *AccountData, blockHeight int64) *AccountData {
	_ = "STUB: not implemented"
	return nil
}
