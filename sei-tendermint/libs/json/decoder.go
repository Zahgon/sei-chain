package json

import (
	"reflect"
)

// Unmarshal unmarshals JSON into the given value, using Amino-compatible JSON encoding (strings
// for 64-bit numbers, and type wrappers for registered types).
func Unmarshal(bz []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func decode(bz []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func decodeReflect(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Handle null for slices, interfaces, and pointers

// Dereference-and-construct pointers, to handle nested pointers.

// Times must be UTC and end with Z

// If value implements json.Umarshaler, call it.

// Decode complex types recursively.

// For 64-bit integers, unwrap expected string and defer to stdlib for integer decoding.

// Anything else we defer to the stdlib.

func decodeReflectList(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Decode base64-encoded bytes using stdlib decoder, via byte slice for arrays.

// Decode anything else into a raw JSON slice, and decode values recursively.

// arrays of wrong size

// Replace empty slices with nil slices, for Amino compatibility

func decodeReflectMap(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Decode into a raw JSON map, using string keys.

// Recursively decode values.

func decodeReflectStruct(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Decode raw JSON values into a string-keyed map.

func decodeStdlib(bz []byte, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Make sure we are unmarshaling into a pointer.
