package keyring

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/hd"
)

// SignatureAlgo defines the interface for a keyring supported algorithm.
type SignatureAlgo interface {
	Name() hd.PubKeyType
	Derive() hd.DeriveFn
	Generate() hd.GenerateFn
}

// NewSigningAlgoFromString creates a supported SignatureAlgo.
func NewSigningAlgoFromString(str string, algoList SigningAlgoList) (SignatureAlgo, error) {
	_ = "STUB: not implemented"
	return *new(SignatureAlgo), nil
}

// SigningAlgoList is a slice of signature algorithms
type SigningAlgoList []SignatureAlgo

// Contains returns true if the SigningAlgoList the given SignatureAlgo.
func (sal SigningAlgoList) Contains(algo SignatureAlgo) bool {
	_ = "STUB: not implemented"
	return false
}

// String returns a comma separated string of the signature algorithm names in the list.
func (sal SigningAlgoList) String() string { _ = "STUB: not implemented"; return "" }
