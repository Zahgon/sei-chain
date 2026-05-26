package tx

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// DefaultTxEncoder returns a default protobuf TxEncoder using the provided Marshaler
func DefaultTxEncoder() sdk.TxEncoder { _ = "STUB: not implemented"; return *new(sdk.TxEncoder) }

// DefaultJSONTxEncoder returns a default protobuf JSON TxEncoder using the provided Marshaler.
func DefaultJSONTxEncoder(cdc codec.ProtoCodecMarshaler) sdk.TxEncoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxEncoder)
}
