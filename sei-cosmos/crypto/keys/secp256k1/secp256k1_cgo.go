//go:build libsecp256k1_sdk
// +build libsecp256k1_sdk

package secp256k1

// Sign creates an ECDSA signature on curve Secp256k1, using SHA256 on the msg.
func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we do not need v  in r||s||v:

// VerifySignature validates the signature.
// The msg will be hashed prior to signature verification.
func (pubKey *PubKey) VerifySignature(msg []byte, sigStr []byte) bool {
	_ = "STUB: not implemented"
	return false
}
