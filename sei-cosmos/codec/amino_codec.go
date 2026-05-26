package codec

import (
	"github.com/gogo/protobuf/proto"
)

// AminoCodec defines a codec that utilizes Codec for both binary and JSON
// encoding.
type AminoCodec struct {
	*LegacyAmino
}

var _ Codec = &AminoCodec{}

// NewAminoCodec returns a reference to a new AminoCodec
func NewAminoCodec(codec *LegacyAmino) *AminoCodec { _ = "STUB: not implemented"; return nil }

// Marshal implements BinaryMarshaler.Marshal method.
func (ac *AminoCodec) Marshal(o ProtoMarshaler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustMarshal implements BinaryMarshaler.MustMarshal method.
func (ac *AminoCodec) MustMarshal(o ProtoMarshaler) []byte { _ = "STUB: not implemented"; return nil }

// MarshalLengthPrefixed implements BinaryMarshaler.MarshalLengthPrefixed method.
func (ac *AminoCodec) MarshalLengthPrefixed(o ProtoMarshaler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustMarshalLengthPrefixed implements BinaryMarshaler.MustMarshalLengthPrefixed method.
func (ac *AminoCodec) MustMarshalLengthPrefixed(o ProtoMarshaler) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Unmarshal implements BinaryMarshaler.Unmarshal method.
func (ac *AminoCodec) Unmarshal(bz []byte, ptr ProtoMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshal implements BinaryMarshaler.MustUnmarshal method.
func (ac *AminoCodec) MustUnmarshal(bz []byte, ptr ProtoMarshaler) {
	_ = "STUB: not implemented"
	return
}

// UnmarshalLengthPrefixed implements BinaryMarshaler.UnmarshalLengthPrefixed method.
func (ac *AminoCodec) UnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshalLengthPrefixed implements BinaryMarshaler.MustUnmarshalLengthPrefixed method.
func (ac *AminoCodec) MustUnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler) {
	_ = "STUB: not implemented"
	return
}

// MarshalJSON implements JSONCodec.MarshalJSON method,
// it marshals to JSON using legacy amino codec.
func (ac *AminoCodec) MarshalAsJSON(o proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustMarshalJSON implements JSONCodec.MustMarshalJSON method,
// it executes MarshalJSON except it panics upon failure.
func (ac *AminoCodec) MustMarshalJSON(o proto.Message) []byte {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements JSONCodec.UnmarshalJSON method,
// it unmarshals from JSON using legacy amino codec.
func (ac *AminoCodec) UnmarshalAsJSON(bz []byte, ptr proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshalJSON implements JSONCodec.MustUnmarshalJSON method,
// it executes UnmarshalJSON except it panics upon failure.
func (ac *AminoCodec) MustUnmarshalJSON(bz []byte, ptr proto.Message) {
	_ = "STUB: not implemented"
	return
}

// MarshalInterface is a convenience function for amino marshaling interfaces.
// The `i` must be an interface.
// NOTE: to marshal a concrete type, you should use Marshal instead
func (ac *AminoCodec) MarshalInterface(i proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalInterface is a convenience function for amino unmarshaling interfaces.
// `ptr` must be a pointer to an interface.
// NOTE: to unmarshal a concrete type, you should use Unmarshal instead
//
// Example:
//
//	var x MyInterface
//	err := cdc.UnmarshalInterface(bz, &x)
func (ac *AminoCodec) UnmarshalInterface(bz []byte, ptr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalInterfaceJSON is a convenience function for amino marshaling interfaces.
// The `i` must be an interface.
// NOTE: to marshal a concrete type, you should use MarshalJSON instead
func (ac *AminoCodec) MarshalInterfaceJSON(i proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalInterfaceJSON is a convenience function for amino unmarshaling interfaces.
// `ptr` must be a pointer to an interface.
// NOTE: to unmarshal a concrete type, you should use UnmarshalJSON instead
//
// Example:
//
//	var x MyInterface
//	err := cdc.UnmarshalInterfaceJSON(bz, &x)
func (ac *AminoCodec) UnmarshalInterfaceJSON(bz []byte, ptr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
