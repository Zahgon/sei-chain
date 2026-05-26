package bytes

import (
	"fmt"
)

// HexBytes is a wrapper around []byte that encodes data as hexadecimal strings
// for use in JSON.
type HexBytes []byte

// Marshal needed for protobuf compatibility
func (bz HexBytes) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"

	// Unmarshal needed for protobuf compatibility
	return nil, nil
}

func (bz *HexBytes) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText encodes a HexBytes value as hexadecimal digits.
// This method is used by json.Marshal.
func (bz HexBytes) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText handles decoding of HexBytes from JSON strings.
// This method is used by json.Unmarshal.
// It allows decoding of both hex and base64-encoded byte arrays.
func (bz *HexBytes) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// Bytes fulfills various interfaces in light-client, etc...
func (bz HexBytes) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (bz HexBytes) String() string { _ = "STUB: not implemented"; return "" }

// Format writes either address of 0th element in a slice in base 16 notation,
// with leading 0x (%p), or casts HexBytes to bytes and writes as hexadecimal
// string to s.
func (bz HexBytes) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Matches the hexbytes MarshalJSON of tendermint/tendermint. Overrides the
// default []byte Marshal implementation. This is basically the point of hex bytes.
func (bz HexBytes) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Matches the hexbytes UnmarshalJSON of tendermint/tendermint. Overrides the
// default []byte Marshal implementation. This is basically the point of hex bytes.
func (bz *HexBytes) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
