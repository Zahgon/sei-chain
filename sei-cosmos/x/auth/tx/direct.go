package tx

import (
	signingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// signModeDirectHandler defines the SIGN_MODE_DIRECT SignModeHandler
type signModeDirectHandler struct{}

var _ signing.SignModeHandler = signModeDirectHandler{}

// DefaultMode implements SignModeHandler.DefaultMode
func (signModeDirectHandler) DefaultMode() signingtypes.SignMode {
	_ = "STUB: not implemented"
	return *new(signingtypes.SignMode)
}

// Modes implements SignModeHandler.Modes
func (signModeDirectHandler) Modes() []signingtypes.SignMode { _ = "STUB: not implemented"; return nil }

// GetSignBytes implements SignModeHandler.GetSignBytes
func (signModeDirectHandler) GetSignBytes(mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DirectSignBytes returns the SIGN_MODE_DIRECT sign bytes for the provided TxBody bytes, AuthInfo bytes, chain ID,
// account number and sequence.
func DirectSignBytes(bodyBytes, authInfoBytes []byte, chainID string, accnum uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
