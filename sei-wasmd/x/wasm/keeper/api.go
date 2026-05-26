package keeper

import (
	wasmvm "github.com/sei-protocol/sei-chain/sei-wasmvm"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

const (
	// DefaultGasCostHumanAddress is how moch SDK gas we charge to convert to a human address format
	DefaultGasCostHumanAddress = 5
	// DefaultGasCostCanonicalAddress is how moch SDK gas we charge to convert to a canonical address format
	DefaultGasCostCanonicalAddress = 4

	// DefaultDeserializationCostPerByte The formular should be `len(data) * deserializationCostPerByte`
	DefaultDeserializationCostPerByte = 1
)

var (
	costHumanize            = DefaultGasCostHumanAddress * DefaultGasMultiplier
	costCanonical           = DefaultGasCostCanonicalAddress * DefaultGasMultiplier
	costJSONDeserialization = wasmvmtypes.UFraction{
		Numerator:   DefaultDeserializationCostPerByte * DefaultGasMultiplier,
		Denominator: 1,
	}
)

func humanAddress(canon []byte) (string, uint64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func canonicalAddress(human string) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

var cosmwasmAPI = wasmvm.GoAPI{
	HumanAddress:     humanAddress,
	CanonicalAddress: canonicalAddress,
}
