package unknownproto

import (
	"reflect"
	"sync"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"google.golang.org/protobuf/encoding/protowire"
)

const bit11NonCritical = 1 << 10

// MaxProtobufNestingDepth defines the maximum allowed nesting depth for protobuf messages
// to prevent stack overflow attacks. This matches similar limits in other protobuf implementations.
const MaxProtobufNestingDepth = 100

type descriptorIface interface {
	Descriptor() ([]byte, []int)
}

// RejectUnknownFieldsStrict rejects any bytes bz with an error that has unknown fields for the provided proto.Message type.
// This function traverses inside of messages nested via google.protobuf.Any. It does not do any deserialization of the proto.Message.
// An AnyResolver must be provided for traversing inside google.protobuf.Any's.
func RejectUnknownFieldsStrict(bz []byte, msg proto.Message, resolver jsonpb.AnyResolver) error {
	_ = "STUB: not implemented"
	return nil
}

// RejectUnknownFields rejects any bytes bz with an error that has unknown fields for the provided proto.Message type with an
// option to allow non-critical fields (specified as those fields with bit 11) to pass through. In either case, the
// hasUnknownNonCriticals will be set to true if non-critical fields were encountered during traversal. This flag can be
// used to treat a message with non-critical field different in different security contexts (such as transaction signing).
// This function traverses inside of messages nested via google.protobuf.Any. It does not do any deserialization of the proto.Message.
// An AnyResolver must be provided for traversing inside google.protobuf.Any's.
func RejectUnknownFields(bz []byte, msg proto.Message, allowUnknownNonCriticals bool, resolver jsonpb.AnyResolver) (hasUnknownNonCriticals bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// rejectUnknownFieldsWithDepth is the internal implementation that tracks recursion depth
func rejectUnknownFieldsWithDepth(bz []byte, msg proto.Message, allowUnknownNonCriticals bool, resolver jsonpb.AnyResolver, depth int) (hasUnknownNonCriticals bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

//nolint:gosec // protobuf field numbers are within int32 range by spec

// Assert that the wireTypes match.

//nolint:gosec // checked by wire type conversion

// The tag is critical, so report it.

// Skip over the bytes that store fieldNumber and wireType bytes.

// An unknown but non-critical field or just a scalar type (aka *INT and BYTES like).

// At this point only TYPE_STRING is expected to be unregistered, since FieldDescriptorProto.IsScalar() returns false for
// TYPE_BYTES and TYPE_STRING as per
// https://github.com/gogo/protobuf/blob/5628607bb4c51c3157aacc3a50f0ab707582b805/protoc-gen-gogo/descriptor/descriptor.go#L95-L118

// Let's recursively traverse and typecheck the field.

// consume length prefix of nested message

// Firstly typecheck types.Any to ensure nothing snuck in.

// And finally we can extract the TypeURL containing the protoMessageName.

var protoMessageForTypeNameMu sync.RWMutex
var protoMessageForTypeNameCache = make(map[string]proto.Message)

// protoMessageForTypeName takes in a fully qualified name e.g. testdata.TestVersionFD1
// and returns a corresponding empty protobuf message that serves the prototype for typechecking.
func protoMessageForTypeName(protoMessageName string) (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}

// Now cache it.

// checks is a mapping of protowire.Type to supported descriptor.FieldDescriptorProto_Type.
// it is implemented this way so as to have constant time lookups and avoid the overhead
// from O(n) walking of switch. The change to using this mapping boosts throughput by about 200%.
var checks = [...]map[descriptor.FieldDescriptorProto_Type]bool{
	// "0	Varint: int32, int64, uint32, uint64, sint32, sint64, bool, enum"
	0: {
		descriptor.FieldDescriptorProto_TYPE_INT32:  true,
		descriptor.FieldDescriptorProto_TYPE_INT64:  true,
		descriptor.FieldDescriptorProto_TYPE_UINT32: true,
		descriptor.FieldDescriptorProto_TYPE_UINT64: true,
		descriptor.FieldDescriptorProto_TYPE_SINT32: true,
		descriptor.FieldDescriptorProto_TYPE_SINT64: true,
		descriptor.FieldDescriptorProto_TYPE_BOOL:   true,
		descriptor.FieldDescriptorProto_TYPE_ENUM:   true,
	},

	// "1	64-bit:	fixed64, sfixed64, double"
	1: {
		descriptor.FieldDescriptorProto_TYPE_FIXED64:  true,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: true,
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:   true,
	},

	// "2	Length-delimited: string, bytes, embedded messages, packed repeated fields"
	2: {
		descriptor.FieldDescriptorProto_TYPE_STRING:  true,
		descriptor.FieldDescriptorProto_TYPE_BYTES:   true,
		descriptor.FieldDescriptorProto_TYPE_MESSAGE: true,
		// The following types can be packed repeated.
		// ref: "Only repeated fields of primitive numeric types (types which use the varint, 32-bit, or 64-bit wire types) can be declared "packed"."
		// ref: https://developers.google.com/protocol-buffers/docs/encoding#packed
		descriptor.FieldDescriptorProto_TYPE_INT32:    true,
		descriptor.FieldDescriptorProto_TYPE_INT64:    true,
		descriptor.FieldDescriptorProto_TYPE_UINT32:   true,
		descriptor.FieldDescriptorProto_TYPE_UINT64:   true,
		descriptor.FieldDescriptorProto_TYPE_SINT32:   true,
		descriptor.FieldDescriptorProto_TYPE_SINT64:   true,
		descriptor.FieldDescriptorProto_TYPE_BOOL:     true,
		descriptor.FieldDescriptorProto_TYPE_ENUM:     true,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:  true,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: true,
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:   true,
	},

	// "3	Start group:	groups (deprecated)"
	3: {
		descriptor.FieldDescriptorProto_TYPE_GROUP: true,
	},

	// "4	End group:	groups (deprecated)"
	4: {
		descriptor.FieldDescriptorProto_TYPE_GROUP: true,
	},

	// "5	32-bit:	fixed32, sfixed32, float"
	5: {
		descriptor.FieldDescriptorProto_TYPE_FIXED32:  true,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32: true,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:    true,
	},
}

// canEncodeType returns true if the wireType is suitable for encoding the descriptor type.
// See https://developers.google.com/protocol-buffers/docs/encoding#structure.
func canEncodeType(wireType protowire.Type, descType descriptor.FieldDescriptorProto_Type) bool {
	_ = "STUB: not implemented"
	return false
}

// errMismatchedWireType describes a mismatch between
// expected and got wireTypes for a specific tag number.
type errMismatchedWireType struct {
	Type         string
	GotWireType  protowire.Type
	WantWireType protowire.Type
	TagNum       protowire.Number
}

// String implements fmt.Stringer.
func (mwt *errMismatchedWireType) String() string { _ = "STUB: not implemented"; return "" }

// Error implements the error interface.
func (mwt *errMismatchedWireType) Error() string { _ = "STUB: not implemented"; return "" }

var _ error = (*errMismatchedWireType)(nil)

func wireTypeToString(wt protowire.Type) string { _ = "STUB: not implemented"; return "" }

// errUnknownField represents an error indicating that we encountered
// a field that isn't available in the target proto.Message.
type errUnknownField struct {
	Type     string
	TagNum   protowire.Number
	WireType protowire.Type
}

// String implements fmt.Stringer.
func (twt *errUnknownField) String() string { _ = "STUB: not implemented"; return "" }

// Error implements the error interface.
func (twt *errUnknownField) Error() string { _ = "STUB: not implemented"; return "" }

var _ error = (*errUnknownField)(nil)

var (
	protoFileToDesc   = make(map[string]*descriptor.FileDescriptorProto)
	protoFileToDescMu sync.RWMutex
)

func unnestDesc(mdescs []*descriptor.DescriptorProto, indices []int) *descriptor.DescriptorProto {
	_ = "STUB: not implemented"
	return nil
}

// Invoking descriptor.ForMessage(proto.Message.(Descriptor).Descriptor()) is incredibly slow
// for every single message, thus the need for a hand-rolled custom version that's performant and cacheable.
func extractFileDescMessageDesc(desc descriptorIface) (*descriptor.FileDescriptorProto, *descriptor.DescriptorProto, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Time to gunzip the content of the FileDescriptor and then proto unmarshal them.

// Now cache the FileDescriptor.

// Unnest the type if necessary.

type descriptorMatch struct {
	cache map[int32]*descriptor.FieldDescriptorProto
	desc  *descriptor.DescriptorProto
}

var descprotoCacheMu sync.RWMutex
var descprotoCache = make(map[reflect.Type]*descriptorMatch)

// getDescriptorInfo retrieves the mapping of field numbers to their respective field descriptors.
func getDescriptorInfo(desc descriptorIface, msg proto.Message) (map[int32]*descriptor.FieldDescriptorProto, *descriptor.DescriptorProto, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Now compute and cache the index.

// DefaultAnyResolver is a default implementation of AnyResolver which uses
// the default encoding of type URLs as specified by the protobuf specification.
type DefaultAnyResolver struct{}

var _ jsonpb.AnyResolver = DefaultAnyResolver{}

// Resolve is the AnyResolver.Resolve method.
func (d DefaultAnyResolver) Resolve(typeURL string) (proto.Message, error) {
	_ = "STUB: not implemented"
	// Only the part of typeURL after the last slash is relevant.
	return *new(proto.Message), nil
}
