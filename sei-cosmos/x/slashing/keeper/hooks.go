package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
)

func (k Keeper) AfterValidatorBonded(ctx sdk.Context, address sdk.ConsAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	// Update the signing info start height or create a new signing info
	return
}

// AfterValidatorCreated adds the address-pubkey relation when a validator is created.
func (k Keeper) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// AfterValidatorRemoved deletes the address-pubkey relation when a validator is removed,
func (k Keeper) AfterValidatorRemoved(ctx sdk.Context, address sdk.ConsAddress) {
	_ = "STUB: not implemented"
	return
}

// Hooks wrapper struct for slashing keeper
type Hooks struct {
	k Keeper
}

var _ types.StakingHooks = Hooks{}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	_ = "STUB: not implemented"

	// Implements sdk.ValidatorHooks
	return *new(Hooks)
}

func (h Hooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// Implements sdk.ValidatorHooks
func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

// Implements sdk.ValidatorHooks
func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}

func (h Hooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) BeforeDelegationCreated(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) AfterDelegationModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
	_ = "STUB: not implemented"
	return
}
func (h Hooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) {
	_ = "STUB: not implemented"
	return
}
