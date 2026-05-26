package tx

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/ante"
	authsigning "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// wrapper is a wrapper around the tx.Tx proto.Message which retain the raw
// body and auth_info bytes.
type wrapper struct {
	tx *tx.Tx

	// bodyBz represents the protobuf encoding of TxBody. This should be encoding
	// from the client using TxRaw if the tx was decoded from the wire
	bodyBz []byte

	// authInfoBz represents the protobuf encoding of TxBody. This should be encoding
	// from the client using TxRaw if the tx was decoded from the wire
	authInfoBz []byte

	txBodyHasUnknownNonCriticals bool
}

var (
	_ authsigning.Tx             = &wrapper{}
	_ client.TxBuilder           = &wrapper{}
	_ ante.HasExtensionOptionsTx = &wrapper{}
	_ ExtensionOptionsTxBuilder  = &wrapper{}
)

// ExtensionOptionsTxBuilder defines a TxBuilder that can also set extensions.
type ExtensionOptionsTxBuilder interface {
	client.TxBuilder

	SetExtensionOptions(...*codectypes.Any)
	SetNonCriticalExtensionOptions(...*codectypes.Any)
}

func newBuilder() *wrapper { _ = "STUB: not implemented"; return nil }

func (w *wrapper) GetMsgs() []sdk.Msg { _ = "STUB: not implemented"; return nil }

func (w *wrapper) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (w *wrapper) GetGasEstimate() uint64 { _ = "STUB: not implemented"; return 0 }

func (w *wrapper) getBodyBytes() []byte { _ = "STUB: not implemented"; return nil }

// if bodyBz is empty, then marshal the body. bodyBz will generally
// be set to nil whenever SetBody is called so the result of calling
// this method should always return the correct bytes. Note that after
// decoding bodyBz is derived from TxRaw so that it matches what was
// transmitted over the wire

func (w *wrapper) getAuthInfoBytes() []byte { _ = "STUB: not implemented"; return nil }

// if authInfoBz is empty, then marshal the body. authInfoBz will generally
// be set to nil whenever SetAuthInfo is called so the result of calling
// this method should always return the correct bytes. Note that after
// decoding authInfoBz is derived from TxRaw so that it matches what was
// transmitted over the wire

func (w *wrapper) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (w *wrapper) GetPubKeys() ([]cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: it is okay to leave this nil if there is no PubKey in the SignerInfo.
// PubKey's can be left unset in SignerInfo.

func (w *wrapper) GetGas() uint64 { _ = "STUB: not implemented"; return 0 }

func (w *wrapper) GetFee() sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

func (w *wrapper) FeePayer() sdk.AccAddress { _ = "STUB: not implemented"; return *new(sdk.AccAddress) }

// use first signer as default if no payer specified

func (w *wrapper) FeeGranter() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (w *wrapper) GetMemo() string { _ = "STUB: not implemented"; return "" }

// GetTimeoutHeight returns the transaction's timeout height (if set).
func (w *wrapper) GetTimeoutHeight() uint64 { _ = "STUB: not implemented"; return 0 }

func (w *wrapper) GetSignaturesV2() ([]signing.SignatureV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handle nil signatures (in case of simulation)

func (w *wrapper) SetMsgs(msgs ...sdk.Msg) error { _ = "STUB: not implemented"; return nil }

// set bodyBz to nil because the cached bodyBz no longer matches tx.Body

// SetTimeoutHeight sets the transaction's height timeout.
func (w *wrapper) SetTimeoutHeight(height uint64) { _ = "STUB: not implemented"; return }

// set bodyBz to nil because the cached bodyBz no longer matches tx.Body

func (w *wrapper) SetMemo(memo string) { _ = "STUB: not implemented"; return }

// set bodyBz to nil because the cached bodyBz no longer matches tx.Body

func (w *wrapper) SetGasLimit(limit uint64) { _ = "STUB: not implemented"; return }

// set authInfoBz to nil because the cached authInfoBz no longer matches tx.AuthInfo

func (w *wrapper) SetGasEstimate(estimate uint64) { _ = "STUB: not implemented"; return }

func (w *wrapper) SetFeeAmount(coins sdk.Coins) { _ = "STUB: not implemented"; return }

// set authInfoBz to nil because the cached authInfoBz no longer matches tx.AuthInfo

func (w *wrapper) SetFeePayer(feePayer sdk.AccAddress) { _ = "STUB: not implemented"; return }

// set authInfoBz to nil because the cached authInfoBz no longer matches tx.AuthInfo

func (w *wrapper) SetFeeGranter(feeGranter sdk.AccAddress) { _ = "STUB: not implemented"; return }

// set authInfoBz to nil because the cached authInfoBz no longer matches tx.AuthInfo

func (w *wrapper) SetSignatures(signatures ...signing.SignatureV2) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrapper) setSignerInfos(infos []*tx.SignerInfo) { _ = "STUB: not implemented"; return }

// set authInfoBz to nil because the cached authInfoBz no longer matches tx.AuthInfo

func (w *wrapper) setSignatures(sigs [][]byte) { _ = "STUB: not implemented"; return }

func (w *wrapper) GetTx() authsigning.Tx { _ = "STUB: not implemented"; return *new(authsigning.Tx) }

func (w *wrapper) GetProtoTx() *tx.Tx {
	_ = "STUB: not implemented"

	// Deprecated: AsAny extracts proto Tx and wraps it into Any.
	// NOTE: You should probably use `GetProtoTx` if you want to serialize the transaction.
	return nil
}

func (w *wrapper) AsAny() *codectypes.Any { _ = "STUB: not implemented"; return nil }

// WrapTx creates a TxBuilder wrapper around a tx.Tx proto message.
func WrapTx(protoTx *tx.Tx) client.TxBuilder {
	_ = "STUB: not implemented"
	return *new(client.TxBuilder)
}

func (w *wrapper) GetExtensionOptions() []*codectypes.Any { _ = "STUB: not implemented"; return nil }

func (w *wrapper) GetNonCriticalExtensionOptions() []*codectypes.Any {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrapper) SetExtensionOptions(extOpts ...*codectypes.Any) {
	_ = "STUB: not implemented"
	return
}

func (w *wrapper) SetNonCriticalExtensionOptions(extOpts ...*codectypes.Any) {
	_ = "STUB: not implemented"
	return
}
