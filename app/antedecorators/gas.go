package antedecorators

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
)

func GetGasMeterSetter(pk paramskeeper.Keeper) func(bool, sdk.Context, uint64, sdk.Tx) sdk.Context {
	_ = "STUB: not implemented"
	return nil
}

// In simulation, still use multiplier but with infinite gas limit
