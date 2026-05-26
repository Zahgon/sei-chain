package utils

import (
	"sync"

	"github.com/gogo/protobuf/proto"
)

func ErrorAs[T error](err error) Option[T] { _ = "STUB: not implemented"; return nil }

// ProtoClone clones a proto.Message object.
func ProtoClone[T proto.Message](item T) T { _ = "STUB: not implemented"; return *new(T) }

// ProtoEqual compares two proto.Message objects.
func ProtoEqual[T proto.Message](a, b T) bool { _ = "STUB: not implemented"; return false }

// ProtoMessage is comparable proto.Message.
type ProtoMessage interface {
	comparable
	proto.Message
}

// ProtoConv is a pair of functions to encode and decode between a type and a ProtoMessage.
type ProtoConv[T any, P ProtoMessage] struct {
	Encode func(T) P
	Decode func(P) (T, error)
}

// EncodeSlice encodes a slice of T into a slice of P.
func (c ProtoConv[T, P]) EncodeSlice(t []T) []P { _ = "STUB: not implemented"; return nil }

// DecodeSlice decodes a slice of P into a slice of T.
func (c ProtoConv[T, P]) DecodeSlice(p []P) ([]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Slice constructs a slice.
// It is a syntax sugar for `[]T{v...}`, which avoids
// spelling out T. Not very useful if you need to spell
// out T to construct the elements: in that case
// you might prefer the []T{{...},{...}} syntax instead.
func Slice[T any](v ...T) []T {
	_ = "STUB: not implemented"

	// Alloc moves value to heap.
	return nil
}

func Alloc[T any](v T) *T {
	_ = "STUB: not implemented"

	// Zero returns a zero value of type T.
	return nil
}

func Zero[T any]() (zero T) {
	_ = "STUB: not implemented"

	// NoCopy may be added to structs which must not be copied
	// after the first use.
	//
	// See https://golang.org/issues/8005#issuecomment-190753527
	// for details.
	//
	// Note that it must not be embedded, otherwise Lock and Unlock methods
	// will be exported.
	return *new(T)
}

type NoCopy struct{}

// Lock implements sync.Locker.
func (*NoCopy) Lock() {
	_ = "STUB: not implemented"

	// Unlock implements sync.Locker.
	return
}

func (*NoCopy) Unlock() { _ = "STUB: not implemented"; return }

var _ sync.Locker = (*NoCopy)(nil)

// NoCompare may be added to structs which must not be used as
// map keys.
type NoCompare [0]func()

// EncodeOpt encodes Option[T], mapping None to Zero[P]().
func (c ProtoConv[T, P]) EncodeOpt(mv Option[T]) P { _ = "STUB: not implemented"; return *new(P) }

// DecodeReq decodes a ProtoMessage into a T, returning an error if p is nil.
func (c ProtoConv[T, P]) DecodeReq(p P) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// DecodeOpt decodes a ProtoMessage into a T, returning nil if p is nil.
func (c ProtoConv[T, P]) DecodeOpt(p P) (Option[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
