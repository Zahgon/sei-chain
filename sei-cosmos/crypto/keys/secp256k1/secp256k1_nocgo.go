//go:build !libsecp256k1_sdk
// +build !libsecp256k1_sdk

package secp256k1

// Sign creates an ECDSA signature on curve Secp256k1, using SHA256 on the msg.
// The returned signature will be of the form R || S (in lower-S form).
func (pk *PrivKey) Sign(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SignCompact returns [recovery_id][R][S] - note: newer btcec versions only return 1 value
// true=compressed pubkey
// Remove the recovery id and return the R||S bytes

// VerifySignature verifies a signature of the form R || S.
// It uses the standard btcec/v2 signature verification approach.
func (pubKey *PubKey) VerifySignature(msg []byte, sigStr []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Check for overflow: SetByteSlice returns true if the value >= curve order N

// Enforce low-S: reject S > N/2 to prevent signature malleability

// Use single hash to match the signing process
