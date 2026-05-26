package codec

import (
	"io"

	amino "github.com/tendermint/go-amino"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// LegacyAmino defines a wrapper for an Amino codec that properly
// handles protobuf types with Any's. Deprecated.
type LegacyAmino struct {
	Amino *amino.Codec
}

func (cdc *LegacyAmino) Seal() { _ = "STUB: not implemented"; return }

func NewLegacyAmino() *LegacyAmino { _ = "STUB: not implemented"; return nil }

// RegisterEvidences registers Tendermint evidence types with the provided Amino
// codec.
func RegisterEvidences(cdc *LegacyAmino) { _ = "STUB: not implemented"; return }

// MarshalJSONIndent provides a utility for indented JSON encoding of an object
// via an Amino codec. It returns an error if it cannot serialize or indent as
// JSON.
func MarshalJSONIndent(cdc *LegacyAmino, obj interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustMarshalJSONIndent executes MarshalJSONIndent except it panics upon failure.
func MustMarshalJSONIndent(cdc *LegacyAmino, obj interface{}) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (cdc *LegacyAmino) marshalAnys(o interface{}) error { _ = "STUB: not implemented"; return nil }

func (cdc *LegacyAmino) unmarshalAnys(o interface{}) error { _ = "STUB: not implemented"; return nil }

func (cdc *LegacyAmino) jsonMarshalAnys(o interface{}) error { _ = "STUB: not implemented"; return nil }

func (cdc *LegacyAmino) jsonUnmarshalAnys(o interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cdc *LegacyAmino) Marshal(o interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdc *LegacyAmino) MustMarshal(o interface{}) []byte { _ = "STUB: not implemented"; return nil }

func (cdc *LegacyAmino) MarshalLengthPrefixed(o interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdc *LegacyAmino) MustMarshalLengthPrefixed(o interface{}) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (cdc *LegacyAmino) Unmarshal(bz []byte, ptr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cdc *LegacyAmino) MustUnmarshal(bz []byte, ptr interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cdc *LegacyAmino) UnmarshalLengthPrefixed(bz []byte, ptr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cdc *LegacyAmino) MustUnmarshalLengthPrefixed(bz []byte, ptr interface{}) {
	_ = "STUB: not implemented"
	return
}

// MarshalAsJSON implements codec.Codec interface
func (cdc *LegacyAmino) MarshalAsJSON(o interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdc *LegacyAmino) MustMarshalJSON(o interface{}) []byte {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalAsJSON implements codec.Codec interface
func (cdc *LegacyAmino) UnmarshalAsJSON(bz []byte, ptr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cdc *LegacyAmino) MustUnmarshalJSON(bz []byte, ptr interface{}) {
	_ = "STUB: not implemented"
	return
}

func (*LegacyAmino) UnpackAny(*types.Any, interface{}) error { _ = "STUB: not implemented"; return nil }

func (cdc *LegacyAmino) RegisterInterface(ptr interface{}, iopts *amino.InterfaceOptions) {
	_ = "STUB: not implemented"
	return
}

func (cdc *LegacyAmino) RegisterConcrete(o interface{}, name string, copts *amino.ConcreteOptions) {
	_ = "STUB: not implemented"
	return
}

func (cdc *LegacyAmino) MarshalJSONIndent(o interface{}, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdc *LegacyAmino) PrintTypes(out io.Writer) error { _ = "STUB: not implemented"; return nil }
