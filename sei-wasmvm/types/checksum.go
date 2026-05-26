package types

// Checksum represents a hash of the Wasm bytecode that serves as an ID. Must be generated from this library.
// The length of a checksum must always be ChecksumLen.
type Checksum []byte

func (cs Checksum) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cs *Checksum) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

const ChecksumLen = 32

// ForceNewChecksum creates a Checksum instance from a hex string.
// It panics in case the input is invalid.
func ForceNewChecksum(input string) Checksum { _ = "STUB: not implemented"; return *new(Checksum) }
