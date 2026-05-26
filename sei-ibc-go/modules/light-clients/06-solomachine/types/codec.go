package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// RegisterInterfaces register the ibc channel submodule interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

func UnmarshalSignatureData(cdc codec.BinaryCodec, data []byte) (signing.SignatureData, error) {
	_ = "STUB: not implemented"
	return *new(signing.SignatureData), nil
}

// UnmarshalDataByType attempts to unmarshal the data to the specified type. An error is
// return if it fails.
func UnmarshalDataByType(cdc codec.BinaryCodec, dataType DataType, data []byte) (Data, error) {
	_ = "STUB: not implemented"
	return *new(Data), nil
}

// unpack any

// unpack any
