package json

import (
	"encoding/json"
	"io"
	"reflect"
	"time"
)

var (
	timeType            = reflect.TypeOf(time.Time{})
	jsonMarshalerType   = reflect.TypeOf(new(json.Marshaler)).Elem()
	jsonUnmarshalerType = reflect.TypeOf(new(json.Unmarshaler)).Elem()
)

// Marshal marshals the value as JSON, using Amino-compatible JSON encoding (strings for
// 64-bit numbers, and type wrappers for registered types).
func Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalIndent marshals the value as JSON, using the given prefix and indentation.
func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encode(w io.Writer, v any) error {
	_ = "STUB: not implemented"
	// Bare nil values can't be reflected, so we must handle them here.
	return nil
}

func encodeReflect(w io.Writer, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// Recursively dereference if pointer.

// Convert times to UTC.

// If the value implements json.Marshaler, defer to stdlib directly. Since we've already
// dereferenced, we try implementations with both value receiver and pointer receiver. We must
// do this after the time normalization above, and thus after dereferencing.

// Complex types must be recursively encoded.

// 64-bit integers are emitted as strings, to avoid precision problems with e.g.
// Javascript which uses 64-bit floats (having 53-bit precision).

// For everything else, defer to the stdlib encoding/json encoder

func encodeReflectList(w io.Writer, rv reflect.Value) error {
	_ = "STUB: not implemented"
	// Emit nil slices as null.
	return nil
}

// Encode byte slices as base64 with the stdlib encoder.

// Stdlib does not base64-encode byte arrays, only slices, so we copy to slice.

// Anything else we recursively encode ourselves.

func encodeReflectMap(w io.Writer, rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

// nil maps are not emitted as nil, to retain Amino compatibility.

func encodeReflectStruct(w io.Writer, rv reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeStdlib(w io.Writer, v any) error {
	_ = "STUB: not implemented"
	// Doesn't stream the output because that adds a newline, as per:
	// https://golang.org/pkg/encoding/json/#Encoder.Encode
	return nil
}

func writeStr(w io.Writer, s string) error { _ = "STUB: not implemented"; return nil }
