package hashable

import (
	"crypto/sha256"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Hashable interface {
	proto.Message
	IsHashable()
}

// Hash is a SHA-256 hash.
type Hash[T Hashable] [sha256.Size]byte

// ParseHash parses a Hash from bytes.
func ParseHash[T Hashable](raw []byte) (Hash[T], error) { _ = "STUB: not implemented"; return nil, nil }

// ToHash hashes a Hashable proto object.
func ToHash[T Hashable](a T) Hash[T] { _ = "STUB: not implemented"; return nil }

// MarshalCanonical returns the canonical protobuf encoding of msg according to
// the custom Tendermint hashing/signing rules described in canonical.go.
// The output is deterministic and suitable for hashing and signing.
func MarshalCanonical[T Hashable](msg T) []byte { _ = "STUB: not implemented"; return nil }

type builder []byte

func (b builder) Tag(num protowire.Number, typ protowire.Type) builder {
	_ = "STUB: not implemented"
	return *new(builder)
}

func (b builder) Varint(v uint64) builder    { _ = "STUB: not implemented"; return *new(builder) }
func (b builder) Fixed32(v uint32) builder   { _ = "STUB: not implemented"; return *new(builder) }
func (b builder) Fixed64(v uint64) builder   { _ = "STUB: not implemented"; return *new(builder) }
func (b builder) Bytes(bytes []byte) builder { _ = "STUB: not implemented"; return *new(builder) }
func (b builder) String(s string) builder    { _ = "STUB: not implemented"; return *new(builder) }

func (b builder) Message(msg protoreflect.Message) builder {
	_ = "STUB: not implemented"
	// NOTE: we ignore unknown fields - we are unable to encode them canonically.
	// NOTE: we can sort fields on init if needed (in the generated files).
	return *new(builder)
}

func (b builder) List(num protoreflect.FieldNumber, kind protoreflect.Kind, list protoreflect.List) builder {
	_ = "STUB: not implemented"
	return *new(builder)
}

// We pack only lists longer than 1 for backward compatibility of optional -> repeated changes.

func (b builder) Singular(num protoreflect.FieldNumber, kind protoreflect.Kind, value protoreflect.Value) builder {
	_ = "STUB: not implemented"
	return *new(builder)
}

func (b builder) Value(kind protoreflect.Kind, value protoreflect.Value) builder {
	_ = "STUB: not implemented"
	return *new(builder)
}

//nolint:gosec // protobuf enum values fit in uint64

//nolint:gosec // intentional truncation to 32-bit per protobuf wire format

//nolint:gosec // reinterpret signed as unsigned per protobuf varint encoding

//nolint:gosec // intentional truncation to 32-bit per protobuf zigzag encoding

//nolint:gosec // intentional truncation to 32-bit per protobuf wire format

//nolint:gosec // intentional truncation to 32-bit per protobuf fixed32

//nolint:gosec // intentional truncation to 32-bit per protobuf sfixed32

//nolint:gosec // reinterpret signed as unsigned per protobuf fixed64 encoding

func isPackable(kind protoreflect.Kind) bool { _ = "STUB: not implemented"; return false }

func sortedFields(fields protoreflect.FieldDescriptors) []protoreflect.FieldDescriptor {
	_ = "STUB: not implemented"
	return nil
}
