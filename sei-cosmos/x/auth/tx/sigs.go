package tx

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// SignatureDataToModeInfoAndSig converts a SignatureData to a ModeInfo and raw bytes signature
func SignatureDataToModeInfoAndSig(data signing.SignatureData) (*tx.ModeInfo, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ModeInfoAndSigToSignatureData converts a ModeInfo and raw bytes signature to a SignatureData or returns
// an error
func ModeInfoAndSigToSignatureData(modeInfo *tx.ModeInfo, sig []byte) (signing.SignatureData, error) {
	_ = "STUB: not implemented"
	return *new(signing.SignatureData), nil
}

// decodeMultisignatures safely decodes the the raw bytes as a MultiSignature protobuf message
func decodeMultisignatures(bz []byte) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NOTE: it is import to reject multi-signatures that contain unrecognized fields because this is an exploitable
// malleability in the protobuf message. Basically an attacker could bloat a MultiSignature message with unknown
// fields, thus bloating the transaction and causing it to fail.

func (g config) MarshalSignatureJSON(sigs []signing.SignatureV2) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g config) UnmarshalSignatureJSON(bz []byte) ([]signing.SignatureV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
