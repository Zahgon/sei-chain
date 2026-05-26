package protoutils

import (
	"google.golang.org/protobuf/proto"
)

// Message is comparable proto.Message.
type Message interface {
	comparable
	proto.Message
}

// Constructs an empty message.
func New[T Message]() T { _ = "STUB: not implemented"; return *new(T) }

// Computes the size of the message encoding.
func Size[T Message](t T) int { _ = "STUB: not implemented"; return 0 }

func Marshal[T Message](t T) []byte {
	_ = "STUB: not implemented"
	// Marshalling messages is always expected to succeed.
	return nil
}

func Unmarshal[T Message](bytes []byte) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// Clone clones a proto.Message object.
func Clone[T Message](item T) T { _ = "STUB: not implemented"; return *new(T) }

// Equal compares two Message objects.
func Equal[T Message](a, b T) bool { _ = "STUB: not implemented"; return false }
