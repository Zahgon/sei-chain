package types

import (
	"github.com/gogo/protobuf/jsonpb"

	amino "github.com/tendermint/go-amino"
)

type anyCompat struct {
	aminoBz []byte
	jsonBz  []byte
	err     error
}

var Debug = true

func anyCompatError(errType string, x interface{}) error { _ = "STUB: not implemented"; return nil }

func (any Any) MarshalAmino() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (any *Any) UnmarshalAmino(bz []byte) error { _ = "STUB: not implemented"; return nil }

func (any *Any) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (any *Any) UnmarshalJSON(bz []byte) error { _ = "STUB: not implemented"; return nil }

// AminoUnpacker is an AnyUnpacker provided for backwards compatibility with
// amino for the binary un-marshaling phase
type AminoUnpacker struct {
	Cdc *amino.Codec
}

var _ AnyUnpacker = AminoUnpacker{}

func (a AminoUnpacker) UnpackAny(any *Any, iface interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// this is necessary for tests that use reflect.DeepEqual and compare
// proto vs amino marshaled values

// AminoUnpacker is an AnyUnpacker provided for backwards compatibility with
// amino for the binary marshaling phase
type AminoPacker struct {
	Cdc *amino.Codec
}

var _ AnyUnpacker = AminoPacker{}

func (a AminoPacker) UnpackAny(any *Any, _ interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// AminoUnpacker is an AnyUnpacker provided for backwards compatibility with
// amino for the JSON marshaling phase
type AminoJSONUnpacker struct {
	Cdc *amino.Codec
}

var _ AnyUnpacker = AminoJSONUnpacker{}

func (a AminoJSONUnpacker) UnpackAny(any *Any, iface interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// this is necessary for tests that use reflect.DeepEqual and compare
// proto vs amino marshaled values

// AminoUnpacker is an AnyUnpacker provided for backwards compatibility with
// amino for the JSON un-marshaling phase
type AminoJSONPacker struct {
	Cdc *amino.Codec
}

var _ AnyUnpacker = AminoJSONPacker{}

func (a AminoJSONPacker) UnpackAny(any *Any, _ interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// ProtoJSONPacker is an AnyUnpacker provided for compatibility with jsonpb
type ProtoJSONPacker struct {
	JSONPBMarshaler *jsonpb.Marshaler
}

var _ AnyUnpacker = ProtoJSONPacker{}

func (a ProtoJSONPacker) UnpackAny(any *Any, _ interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
