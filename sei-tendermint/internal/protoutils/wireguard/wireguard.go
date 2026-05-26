// Package wireguard runs bounded checks on raw protobuf wire bytes before
// the message is handed to Unmarshal. A caller registers a Schema describing
// which fields to descend into and which repeated fields to cap; Scan then
// walks the bytes once and rejects any payload that violates the rules.
//
// The intended use is as a channel/stream PreDecode hook: any size or shape
// invariant that must be enforced before decoding goes here, expressed
// declaratively as a schema next to the channel definition.
package wireguard

import (
	"google.golang.org/protobuf/encoding/protowire"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// Number re-exports protowire.Number so callers can build Schemas without
// also importing google.golang.org/protobuf/encoding/protowire directly.
type Number = protowire.Number

// Schema describes the validation applied to a single proto message type.
// Rules are keyed by proto field number; nesting is expressed by setting
// Rule.Nested to a child Schema. Schemas are immutable after construction
// and safe for concurrent use.
type Schema struct {
	Rules map[Number]Rule
}

// Rule is the validation applied to one field of a Schema's parent message.
// Nested and MaxCount compose: a field can both descend into a child Schema
// and cap its own occurrence count.
type Rule struct {
	// Nested, if Some, is applied to the contents of this length-delimited
	// field. Use for descending through wrapper layers on the way to a cap.
	Nested utils.Option[*Schema]
	// MaxCount, if non-zero, caps how many times this field may appear in the
	// scanned payload. The count is accumulated globally across the whole
	// Scan call — every match of this (Schema, field) pair increments one
	// shared counter, not a fresh counter per parent instance.
	MaxCount int
}

// Scan walks bz once, applying schema. Returns nil on success, an error on
// malformed wire bytes or a rule violation. A nil schema is a no-op.
func Scan(bz []byte, schema *Schema) error { _ = "STUB: not implemented"; return nil }

// Scan is the method form of the package-level Scan. It's the shape a
// ChannelDescriptor's PreDecode hook expects, so the generated SchemaForX
// values can be wired in directly without a wrapping closure.
func (s *Schema) Scan(bz []byte) error {
	_ = "STUB: not implemented"

	// counterKey scopes a MaxCount accumulator by (Schema, field number) so the
	// same Schema reached from multiple paths shares one counter, while two
	// unrelated Schemas that happen to use the same field number don't collide.
	return nil
}

type counterKey struct {
	schema *Schema
	num    Number
}

func scan(bz []byte, schema *Schema, counts map[counterKey]int) error {
	_ = "STUB: not implemented"
	return nil
}

// MustFieldNum reads the protobuf field number declared on T's field whose
// proto `name=` tag matches protoName. It panics if the field is missing or
// the tag is malformed, since both indicate a divergence between caller code
// and the regenerated proto bindings — a silent miscompare is worse than a
// loud startup failure for code that wires up Schemas at init.
//
// Reflection runs against the *struct type* via reflect.TypeFor[T](); no
// runtime instance is examined. Repeated fields, optional fields, and oneof
// variants all generate a struct field in the proto bindings regardless of
// whether any message instance populates them, so an empty / nil value at
// runtime is irrelevant here.
//
// To remove a proto field that a Schema currently references: first delete
// the MustFieldNum call and the Schema Rule that uses it, then regenerate
// proto with the field gone. Doing it in the other order panics at init.
func MustFieldNum[T any](protoName string) Number { _ = "STUB: not implemented"; return *new(Number) }

//nolint:gosec // ParseInt with bitSize=32 bounds num to int32 range
