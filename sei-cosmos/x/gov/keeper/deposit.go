package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

// GetDeposit gets the deposit of a specific depositor on a specific proposal
func (keeper Keeper) GetDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) (deposit types.Deposit, found bool) {
	_ = "STUB: not implemented"
	return *new(types.Deposit), false
}

// SetDeposit sets a Deposit to the gov store
func (keeper Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	_ = "STUB: not implemented"
	return
}

// GetAllDeposits returns all the deposits from the store
func (keeper Keeper) GetAllDeposits(ctx sdk.Context) (deposits types.Deposits) {
	_ = "STUB: not implemented"
	return *new(types.Deposits)
}

// GetDeposits returns all the deposits from a proposal
func (keeper Keeper) GetDeposits(ctx sdk.Context, proposalID uint64) (deposits types.Deposits) {
	_ = "STUB: not implemented"
	return *new(types.Deposits)
}

// DeleteDeposits deletes all the deposits on a specific proposal without refunding them
func (keeper Keeper) DeleteDeposits(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}

// IterateAllDeposits iterates over the all the stored deposits and performs a callback function
func (keeper Keeper) IterateAllDeposits(ctx sdk.Context, cb func(deposit types.Deposit) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// IterateDeposits iterates over the all the proposals deposits and performs a callback function
func (keeper Keeper) IterateDeposits(ctx sdk.Context, proposalID uint64, cb func(deposit types.Deposit) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// AddDeposit adds or updates a deposit of a specific depositor on a specific proposal
// Activates voting period when appropriate
func (keeper Keeper) AddDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress, depositAmount sdk.Coins) (bool, error) {
	_ = "STUB: not implemented"
	// Checks to see if proposal exists
	return false, nil
}

// Check if proposal is still depositable

// update the governance module's account coins pool

// Update proposal

// Check if deposit has provided sufficient total funds to transition the proposal into the voting period

// Add or update deposit object

// called when deposit has been added to a proposal, however the proposal may not be active

// RefundDeposits refunds and deletes all the deposits on a specific proposal
func (keeper Keeper) RefundDeposits(ctx sdk.Context, proposalID uint64) {
	_ = "STUB: not implemented"
	return
}
