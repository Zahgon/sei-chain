package legacytx

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types/multisig"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	signingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// stdTxSignModeHandler is a SignModeHandler that handles SIGN_MODE_LEGACY_AMINO_JSON
type stdTxSignModeHandler struct{}

func NewStdTxSignModeHandler() signing.SignModeHandler {
	_ = "STUB: not implemented"
	return *new(signing.SignModeHandler)
}

// assert interface implementation
var _ signing.SignModeHandler = stdTxSignModeHandler{}

// DefaultMode implements SignModeHandler.DefaultMode
func (h stdTxSignModeHandler) DefaultMode() signingtypes.SignMode {
	_ = "STUB: not implemented"
	return *new(signingtypes.SignMode)
}

// Modes implements SignModeHandler.Modes
func (stdTxSignModeHandler) Modes() []signingtypes.SignMode { _ = "STUB: not implemented"; return nil }

// DefaultMode implements SignModeHandler.GetSignBytes
func (stdTxSignModeHandler) GetSignBytes(mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignatureDataToAminoSignature converts a SignatureData to amino-encoded signature bytes.
// Only SIGN_MODE_LEGACY_AMINO_JSON is supported.
func SignatureDataToAminoSignature(cdc *codec.LegacyAmino, data signingtypes.SignatureData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MultiSignatureDataToAminoMultisignature converts a MultiSignatureData to an AminoMultisignature.
// Only SIGN_MODE_LEGACY_AMINO_JSON is supported.
func MultiSignatureDataToAminoMultisignature(cdc *codec.LegacyAmino, mSig *signingtypes.MultiSignatureData) (multisig.AminoMultisignature, error) {
	_ = "STUB: not implemented"
	return *new(multisig.AminoMultisignature), nil
}
