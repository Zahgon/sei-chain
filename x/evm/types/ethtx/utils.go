package ethtx

import (
	"math/big"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Effective gas price is the smaller of base fee + tip limit vs total fee limit
func EffectiveGasPrice(baseFee, feeCap, tipCap *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// Convert a value with the provided converter and set it using the provided setter
func SetConvertIfPresent[U comparable, V any](orig U, converter func(U) V, setter func(V)) {
	_ = "STUB: not implemented"
	return
}

// validate a ethtypes.Transaction for sdk.Int overflow
func ValidateEthTx(tx *ethtypes.Transaction) error { _ = "STUB: not implemented"; return nil }

func DecodeSignature(sig []byte) (r, s, v *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
