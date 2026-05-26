package address

import (
	"crypto/sha256"
)

// Len is the length of base addresses
const Len = sha256.Size

// Addressable represents any type from which we can derive an address.
type Addressable interface {
	Address() []byte
}

// Hash creates a new address from address type and key
func Hash(typ string, key []byte) []byte { _ = "STUB: not implemented"; return nil }

// the error always nil, it's here only to satisfy the io.Writer interface

// Compose creates a new address based on sub addresses.
func Compose(typ string, subAddresses []Addressable) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Module is a specialized version of a composed address for modules. Each module account
// is constructed from a module name and module account key.
func Module(moduleName string, key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Derive derives a new address from the main `address` and a derivation `key`.
func Derive(address []byte, key []byte) []byte { _ = "STUB: not implemented"; return nil }
