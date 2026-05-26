package evmrpc

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type traceTxConfig struct {
	client.TxConfig
	decoder sdk.TxDecoder
}

func (c traceTxConfig) TxDecoder() sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

type protoCodecProvider interface {
	ProtoCodec() codec.ProtoCodecMarshaler
}

func traceCompatTxConfig(txConfig client.TxConfig, v65ActiveAtHeight bool) client.TxConfig {
	_ = "STUB: not implemented"
	return *new(client.TxConfig)
}

func traceCompatTxConfigProvider(txConfigProvider func(int64) client.TxConfig, isV65ActiveAtHeight func(int64) bool) func(int64) client.TxConfig {
	_ = "STUB: not implemented"
	return nil
}

func traceCompatTxDecoder(txConfig client.TxConfig, v65ActiveAtHeight bool) sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}
