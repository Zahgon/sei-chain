// Package jsontypes supports decoding for interface types whose concrete
// implementations need to be stored as JSON. To do this, concrete values are
// packaged in wrapper objects having the form:
//
//	{
//	  "type": "<type-tag>",
//	  "value": <json-encoding-of-value>
//	}
//
// This package provides a registry for type tag strings and functions to
// encode and decode wrapper objects.
package jsontypes

import (
	"encoding/json"
	"reflect"
)

// The Tagged interface must be implemented by a type in order to register it
// with the jsontypes package. The TypeTag method returns a string label that
// is used to distinguish objects of that type.
type Tagged interface {
	TypeTag() string
}

// registry records the mapping from type tags to value types.
var registry = struct {
	types map[string]reflect.Type
}{types: make(map[string]reflect.Type)}

// register adds v to the type registry. It reports an error if the tag
// returned by v is already registered.
func register(v Tagged) error { _ = "STUB: not implemented"; return nil }

// MustRegister adds v to the type registry. It will panic if the tag returned
// by v is already registered. This function is meant for use during program
// initialization.
func MustRegister(v Tagged) { _ = "STUB: not implemented"; return }

type wrapper struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

// Marshal marshals a JSON wrapper object containing v. If v == nil, Marshal
// returns the JSON "null" value without error.
func Marshal(v Tagged) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal unmarshals a JSON wrapper object into v. It reports an error if
// the data do not encode a valid wrapper object, if the wrapper's type tag is
// not registered with jsontypes, or if the resulting value is not compatible
// with the type of v.
func Unmarshal(data []byte, v any) error {
	_ = "STUB: not implemented"
	// Verify that the target is some kind of pointer.
	return nil
}

// ok: registered type is directly assignable to the target

// ok: registered type is a pointer to a value assignable to the target

// we need a pointer to unmarshal

// isNull reports true if data is empty or is the JSON "null" value.
func isNull(data []byte) bool { _ = "STUB: not implemented"; return false }
