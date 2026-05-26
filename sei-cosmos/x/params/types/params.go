package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var ParamStoreKeyFeesParams = []byte("FeesParams")
var ParamStoreKeyCosmosGasParams = []byte("CosmosGasParams")

func NewFeesParams(minGasPrices sdk.DecCoins) FeesParams {
	_ = "STUB: not implemented"
	return *new(FeesParams)
}

// ParamTable for minting module.
func ParamKeyTable() KeyTable { _ = "STUB: not implemented"; return *new(KeyTable) }

func (fp *FeesParams) Validate() error { _ = "STUB: not implemented"; return nil }

func validateFeesParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

func NewCosmosGasParams(multiplierNumerator uint64, multiplierDenominator uint64) CosmosGasParams {
	_ = "STUB: not implemented"
	return *new(CosmosGasParams)
}

func (cg *CosmosGasParams) Validate() error { _ = "STUB: not implemented"; return nil }

func validateCosmosGasParams(i interface{}) error { _ = "STUB: not implemented"; return nil }
