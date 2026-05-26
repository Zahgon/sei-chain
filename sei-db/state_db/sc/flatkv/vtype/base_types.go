package vtype

const (
	AddressLen  = 20
	CodeHashLen = 32
	NonceLen    = 8
	SlotLen     = 32
	BalanceLen  = 32
)

// Address is an EVM address (20 bytes).
type Address [AddressLen]byte

// CodeHash is a contract code hash (32 bytes).
type CodeHash [CodeHashLen]byte

// Slot is a storage slot key (32 bytes).
type Slot [SlotLen]byte

// Balance is an EVM balance (32 bytes, big-endian uint256).
type Balance [BalanceLen]byte

// ParseNonce parses a nonce value from a byte slice.
func ParseNonce(b []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// ParseCodeHash parses a codehash value from a byte slice.
func ParseCodeHash(b []byte) (*CodeHash, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseBalance parses a balance value from a byte slice.
func ParseBalance(b []byte) (*Balance, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseStorageValue parses a storage value from a byte slice.
func ParseStorageValue(b []byte) (*[32]byte, error) { _ = "STUB: not implemented"; return nil, nil }
