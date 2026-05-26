package tx

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	signingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

const aminoNonCriticalFieldsError = "protobuf transaction contains unknown non-critical fields. This is a transaction malleability issue and SIGN_MODE_LEGACY_AMINO_JSON cannot be used."

var _ signing.SignModeHandler = signModeLegacyAminoJSONHandler{}

// signModeLegacyAminoJSONHandler defines the SIGN_MODE_LEGACY_AMINO_JSON
// SignModeHandler.
type signModeLegacyAminoJSONHandler struct{}

func (s signModeLegacyAminoJSONHandler) DefaultMode() signingtypes.SignMode {
	_ = "STUB: not implemented"
	return *new(signingtypes.SignMode)
}

func (s signModeLegacyAminoJSONHandler) Modes() []signingtypes.SignMode {
	_ = "STUB: not implemented"
	return nil
}

func (s signModeLegacyAminoJSONHandler) GetSignBytes(mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
