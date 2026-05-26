package keeper // noalias

import (
	"math/rand"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// does a certain by-power index record exist
func ValidatorByPowerIndexExists(ctx sdk.Context, keeper Keeper, power []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// update validator for testing
func TestingUpdateValidator(keeper Keeper, ctx sdk.Context, validator types.Validator, apply bool) types.Validator {
	_ = "STUB: not implemented"
	return *new(types.Validator)
}

// Remove any existing power key for validator.

// RandomValidator returns a random validator given access to the keeper and ctx
func RandomValidator(r *rand.Rand, keeper Keeper, ctx sdk.Context) (val types.Validator, ok bool) {
	_ = "STUB: not implemented"
	return *new(types.Validator), false
}
