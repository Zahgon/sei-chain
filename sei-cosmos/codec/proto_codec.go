package codec

import (
	"github.com/gogo/protobuf/proto"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// ProtoCodecMarshaler defines an interface for codecs that utilize Protobuf for both
// binary and JSON encoding.
type ProtoCodecMarshaler interface {
	Codec
	InterfaceRegistry() types.InterfaceRegistry
}

// ProtoCodec defines a codec that utilizes Protobuf for both binary and JSON
// encoding.
type ProtoCodec struct {
	interfaceRegistry types.InterfaceRegistry
}

var _ Codec = &ProtoCodec{}
var _ ProtoCodecMarshaler = &ProtoCodec{}

// NewProtoCodec returns a reference to a new ProtoCodec
func NewProtoCodec(interfaceRegistry types.InterfaceRegistry) *ProtoCodec {
	_ = "STUB: not implemented"
	return nil
}

// Marshal implements BinaryMarshaler.Marshal method.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.MarshalInterface
func (pc *ProtoCodec) Marshal(o ProtoMarshaler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MustMarshal implements BinaryMarshaler.MustMarshal method.
		// NOTE: this function must be used with a concrete type which
		// implements proto.Message. For interface please use the codec.MarshalInterface
		nil
}

func (pc *ProtoCodec) MustMarshal(o ProtoMarshaler) []byte { _ = "STUB: not implemented"; return nil }

// MarshalLengthPrefixed implements BinaryMarshaler.MarshalLengthPrefixed method.
func (pc *ProtoCodec) MarshalLengthPrefixed(o ProtoMarshaler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // Size() returns the serialized size which is always non-negative

// MustMarshalLengthPrefixed implements BinaryMarshaler.MustMarshalLengthPrefixed method.
func (pc *ProtoCodec) MustMarshalLengthPrefixed(o ProtoMarshaler) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Unmarshal implements BinaryMarshaler.Unmarshal method.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterface
func (pc *ProtoCodec) Unmarshal(bz []byte, ptr ProtoMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshal implements BinaryMarshaler.MustUnmarshal method.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterface
func (pc *ProtoCodec) MustUnmarshal(bz []byte, ptr ProtoMarshaler) {
	_ = "STUB: not implemented"
	return
}

// UnmarshalLengthPrefixed implements BinaryMarshaler.UnmarshalLengthPrefixed method.
func (pc *ProtoCodec) UnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // remaining is validated non-negative above

//nolint:gosec // remaining is validated non-negative above

// MustUnmarshalLengthPrefixed implements BinaryMarshaler.MustUnmarshalLengthPrefixed method.
func (pc *ProtoCodec) MustUnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler) {
	_ = "STUB: not implemented"
	return
}

// MarshalJSON implements JSONCodec.MarshalJSON method,
// it marshals to JSON using proto codec.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.MarshalInterfaceJSON
func (pc *ProtoCodec) MarshalAsJSON(o proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustMarshalJSON implements JSONCodec.MustMarshalJSON method,
// it executes MarshalJSON except it panics upon failure.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.MarshalInterfaceJSON
func (pc *ProtoCodec) MustMarshalJSON(o proto.Message) []byte {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements JSONCodec.UnmarshalJSON method,
// it unmarshals from JSON using proto codec.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterfaceJSON
func (pc *ProtoCodec) UnmarshalAsJSON(bz []byte, ptr proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshalJSON implements JSONCodec.MustUnmarshalJSON method,
// it executes UnmarshalJSON except it panics upon failure.
// NOTE: this function must be used with a concrete type which
// implements proto.Message. For interface please use the codec.UnmarshalInterfaceJSON
func (pc *ProtoCodec) MustUnmarshalJSON(bz []byte, ptr proto.Message) {
	_ = "STUB: not implemented"
	return
}

// MarshalInterface is a convenience function for proto marshalling interfaces. It packs
// the provided value, which must be an interface, in an Any and then marshals it to bytes.
// NOTE: to marshal a concrete type, you should use Marshal instead
func (pc *ProtoCodec) MarshalInterface(i proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalInterface is a convenience function for proto unmarshaling interfaces. It
// unmarshals an Any from bz bytes and then unpacks it to the `ptr`, which must
// be a pointer to a non empty interface with registered implementations.
// NOTE: to unmarshal a concrete type, you should use Unmarshal instead
//
// Example:
//
//	var x MyInterface
//	err := cdc.UnmarshalInterface(bz, &x)
func (pc *ProtoCodec) UnmarshalInterface(bz []byte, ptr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalInterfaceJSON is a convenience function for proto marshalling interfaces. It
// packs the provided value in an Any and then marshals it to bytes.
// NOTE: to marshal a concrete type, you should use MarshalJSON instead
func (pc *ProtoCodec) MarshalInterfaceJSON(x proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalInterfaceJSON is a convenience function for proto unmarshaling interfaces.
// It unmarshals an Any from bz bytes and then unpacks it to the `iface`, which must
// be a pointer to a non empty interface, implementing proto.Message with registered implementations.
// NOTE: to unmarshal a concrete type, you should use UnmarshalJSON instead
//
// Example:
//
//	var x MyInterface  // must implement proto.Message
//	err := cdc.UnmarshalInterfaceJSON(&x, bz)
func (pc *ProtoCodec) UnmarshalInterfaceJSON(bz []byte, iface interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackAny implements AnyUnpacker.UnpackAny method,
// it unpacks the value in any to the interface pointer passed in as
// iface.
func (pc *ProtoCodec) UnpackAny(any *types.Any, iface interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// InterfaceRegistry returns InterfaceRegistry
func (pc *ProtoCodec) InterfaceRegistry() types.InterfaceRegistry {
	_ = "STUB: not implemented"
	return *new(types.InterfaceRegistry)
}

func assertNotNil(i interface{}) error { _ = "STUB: not implemented"; return nil }
