package protoutils

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// Conv is a pair of functions to encode and decode between a type and a Message.
type Conv[T any, P Message] struct {
	Encode func(T) P
	Decode func(P) (T, error)
}

func (c Conv[T, P]) Marshal(t T) []byte { _ = "STUB: not implemented"; return nil }

func (c Conv[T, P]) Unmarshal(bytes []byte) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// EncodeSlice encodes a slice of T into a slice of P.
func (c Conv[T, P]) EncodeSlice(t []T) []P { _ = "STUB: not implemented"; return nil }

// DecodeSlice decodes a slice of P into a slice of T.
func (c Conv[T, P]) DecodeSlice(p []P) ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

// EncodeOpt encodes utils.Option[T], mapping utils.None to utils.Zero[P]().
func (c Conv[T, P]) EncodeOpt(mv utils.Option[T]) P { _ = "STUB: not implemented"; return *new(P) }

// DecodeReq decodes a ProtoMessage into a T, returning an error if p is nil.
func (c Conv[T, P]) DecodeReq(p P) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// DecodeOpt decodes a ProtoMessage into a T, returning nil if p is nil.
func (c Conv[T, P]) DecodeOpt(p P) (utils.Option[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
