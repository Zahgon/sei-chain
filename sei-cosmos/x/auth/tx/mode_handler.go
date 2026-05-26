package tx

import (
	signingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// DefaultSignModes are the default sign modes enabled for protobuf transactions.
var DefaultSignModes = []signingtypes.SignMode{
	signingtypes.SignMode_SIGN_MODE_DIRECT,
	signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
}

// makeSignModeHandler returns the default protobuf SignModeHandler supporting
// SIGN_MODE_DIRECT and SIGN_MODE_LEGACY_AMINO_JSON.
func makeSignModeHandler(modes []signingtypes.SignMode) signing.SignModeHandler {
	_ = "STUB: not implemented"
	return *new(signing.SignModeHandler)
}
