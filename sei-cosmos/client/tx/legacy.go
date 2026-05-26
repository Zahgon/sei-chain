package tx

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// ConvertTxToStdTx converts a transaction to the legacy StdTx format
func ConvertTxToStdTx(codec *codec.LegacyAmino, tx signing.Tx) (legacytx.StdTx, error) {
	_ = "STUB: not implemented"
	return *new(legacytx.StdTx), nil
}

// CopyTx copies a Tx to a new TxBuilder, allowing conversion between
// different transaction formats. If ignoreSignatureError is true, copying will continue
// tx even if the signature cannot be set in the target builder resulting in an unsigned tx.
func CopyTx(tx signing.Tx, builder client.TxBuilder, ignoreSignatureError bool) error {
	_ = "STUB: not implemented"
	return nil
}

// we call SetSignatures() agan with no args to clear any signatures in case the
// previous call to SetSignatures() had any partial side-effects
