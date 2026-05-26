package address

// MaxAddrLen is the maximum allowed length (in bytes) for an address.
const MaxAddrLen = 255

// LengthPrefix prefixes the address bytes with its length, this is used
// for example for variable-length components in store keys.
func LengthPrefix(bz []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MustLengthPrefix is LengthPrefix with panic on error.
func MustLengthPrefix(bz []byte) []byte { _ = "STUB: not implemented"; return nil }
