package migrations

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

// MigrateSstoreGas updates the SeiSstoreSetGasEip2200 parameter to the default value.
func MigrateSstoreGas(ctx sdk.Context, k *keeper.Keeper) error {
	_ = "STUB: not implemented"
	return nil
}
