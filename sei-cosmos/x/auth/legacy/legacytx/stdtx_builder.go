package legacytx

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	authsigning "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// StdTxBuilder wraps StdTx to implement to the context.TxBuilder interface.
// Note that this type just exists for backwards compatibility with amino StdTx
// and will not work for protobuf transactions.
type StdTxBuilder struct {
	StdTx
	cdc *codec.LegacyAmino
}

// ensure interface implementation
var _ client.TxBuilder = &StdTxBuilder{}

// GetTx implements TxBuilder.GetTx
func (s *StdTxBuilder) GetTx() authsigning.Tx {
	_ = "STUB: not implemented"

	// SetMsgs implements TxBuilder.SetMsgs
	return *new(authsigning.Tx)
}

func (s *StdTxBuilder) SetMsgs(msgs ...sdk.Msg) error { _ = "STUB: not implemented"; return nil }

// SetSignatures implements TxBuilder.SetSignatures.
func (s *StdTxBuilder) SetSignatures(signatures ...signing.SignatureV2) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *StdTxBuilder) SetFeeAmount(amount sdk.Coins) { _ = "STUB: not implemented"; return }

func (s *StdTxBuilder) SetGasLimit(limit uint64) { _ = "STUB: not implemented"; return }

func (s *StdTxBuilder) SetGasEstimate(estimate uint64) { _ = "STUB: not implemented"; return }

// SetMemo implements TxBuilder.SetMemo
func (s *StdTxBuilder) SetMemo(memo string) {
	_ = "STUB: not implemented"

	// SetTimeoutHeight sets the transaction's height timeout.
	return
}

func (s *StdTxBuilder) SetTimeoutHeight(height uint64) { _ = "STUB: not implemented"; return }

// SetFeeGranter does nothing for stdtx
func (s *StdTxBuilder) SetFeeGranter(_ sdk.AccAddress) {
	_ = "STUB: not implemented"

	// StdTxConfig is a context.TxConfig for StdTx
	return
}

type StdTxConfig struct {
	Cdc *codec.LegacyAmino
}

var _ client.TxConfig = StdTxConfig{}

// NewTxBuilder implements TxConfig.NewTxBuilder
func (s StdTxConfig) NewTxBuilder() client.TxBuilder {
	_ = "STUB: not implemented"
	return *new(client.TxBuilder)
}

// WrapTxBuilder returns a StdTxBuilder from provided transaction
func (s StdTxConfig) WrapTxBuilder(newTx sdk.Tx) (client.TxBuilder, error) {
	_ = "STUB: not implemented"
	return *new(client.TxBuilder), nil
}

// MarshalTx implements TxConfig.MarshalTx
func (s StdTxConfig) TxEncoder() sdk.TxEncoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxEncoder)
}

func (s StdTxConfig) TxDecoder() sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

func (s StdTxConfig) TxJSONEncoder() sdk.TxEncoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxEncoder)
}

func (s StdTxConfig) TxJSONDecoder() sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

func (s StdTxConfig) MarshalSignatureJSON(sigs []signing.SignatureV2) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s StdTxConfig) UnmarshalSignatureJSON(bz []byte) ([]signing.SignatureV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s StdTxConfig) SignModeHandler() authsigning.SignModeHandler {
	_ = "STUB: not implemented"
	return *new(authsigning.SignModeHandler)
}

// SignatureV2ToStdSignature converts a SignatureV2 to a StdSignature
// [Deprecated]
func SignatureV2ToStdSignature(cdc *codec.LegacyAmino, sig signing.SignatureV2) (StdSignature, error) {
	_ = "STUB: not implemented"
	return *new(StdSignature), nil
}

// Unmarshaler is a generic type for Unmarshal functions
type Unmarshaler func(bytes []byte, ptr interface{}) error

func mkDecoder(unmarshaler Unmarshaler) sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

// StdTx.Msg is an interface. The concrete types
// are registered by MakeTxCodec

// DefaultTxEncoder logic for standard transaction encoding
func DefaultTxEncoder(cdc *codec.LegacyAmino) sdk.TxEncoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxEncoder)
}
