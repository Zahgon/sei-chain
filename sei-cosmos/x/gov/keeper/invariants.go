package keeper

// DONTCOVER

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// RegisterInvariants registers all governance invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper, bk types.BankKeeper) {
	_ = "STUB: not implemented"
	return
}

// AllInvariants runs all invariants of the governance module
func AllInvariants(keeper Keeper, bk types.BankKeeper) sdk.Invariant {
	_ = "STUB: not implemented"
	return *new(sdk.Invariant)
}

// ModuleAccountInvariant checks that the module account coins reflects the sum of
// deposit amounts held on store
func ModuleAccountInvariant(keeper Keeper, bk types.BankKeeper) sdk.Invariant {
	_ = "STUB: not implemented"
	return *new(sdk.Invariant)
}
