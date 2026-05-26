package tx

import (
	signingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

type config struct {
	handler     signing.SignModeHandler
	decoder     sdk.TxDecoder
	encoder     sdk.TxEncoder
	jsonDecoder sdk.TxDecoder
	jsonEncoder sdk.TxEncoder
	protoCodec  codec.ProtoCodecMarshaler
}

// NewTxConfig returns a new protobuf TxConfig using the provided ProtoCodec and sign modes. The
// first enabled sign mode will become the default sign mode.
// NOTE: Use NewTxConfigWithHandler to provide a custom signing handler in case the sign mode
// is not supported by default (eg: SignMode_SIGN_MODE_EIP_191).
func NewTxConfig(protoCodec codec.ProtoCodecMarshaler, enabledSignModes []signingtypes.SignMode) client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}

// NewTxConfig returns a new protobuf TxConfig using the provided ProtoCodec and signing handler.
func NewTxConfigWithHandler(protoCodec codec.ProtoCodecMarshaler, handler signing.SignModeHandler) client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}

func (g config) NewTxBuilder() client.TxBuilder {
	_ = "STUB: not implemented"
	return *

	// WrapTxBuilder returns a builder from provided transaction
	new(client.TxBuilder)
}

func (g config) WrapTxBuilder(newTx sdk.Tx) (client.TxBuilder, error) {
	_ = "STUB: not implemented"
	return *new(client.TxBuilder), nil
}

func (g config) SignModeHandler() signing.SignModeHandler {
	_ = "STUB: not implemented"
	return *new(signing.SignModeHandler)
}

func (g config) TxEncoder() sdk.TxEncoder { _ = "STUB: not implemented"; return *new(sdk.TxEncoder) }

func (g config) TxDecoder() sdk.TxDecoder { _ = "STUB: not implemented"; return *new(sdk.TxDecoder) }

func (g config) ProtoCodec() codec.ProtoCodecMarshaler {
	_ = "STUB: not implemented"
	return *new(codec.ProtoCodecMarshaler)
}

func (g config) TxJSONEncoder() sdk.TxEncoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxEncoder)
}

func (g config) TxJSONDecoder() sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}
