package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// ModuleCdc references the global interchain accounts module codec. Note, the codec
// should ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to interchain accounts and
// defined at the application level.
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// RegisterInterfaces registers the concrete InterchainAccount implementation against the associated
// x/auth AccountI and GenesisAccount interfaces
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

// SerializeCosmosTx serializes a slice of sdk.Msg's using the CosmosTx type. The sdk.Msg's are
// packed into Any's and inserted into the Messages field of a CosmosTx. The proto marshaled CosmosTx
// bytes are returned. Only the ProtoCodec is supported for serializing messages.
func SerializeCosmosTx(cdc codec.BinaryCodec, msgs []sdk.Msg) (bz []byte, err error) {
	_ = "STUB: not implemented"
	// only ProtoCodec is supported
	return nil, nil
}

// DeserializeCosmosTx unmarshals and unpacks a slice of transaction bytes
// into a slice of sdk.Msg's. Only the ProtoCodec is supported for message
// deserialization.
func DeserializeCosmosTx(cdc codec.BinaryCodec, data []byte) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	// only ProtoCodec is supported
	return nil, nil
}
