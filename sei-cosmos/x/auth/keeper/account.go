package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

// NewAccountWithAddress implements AccountKeeperI.
func (ak AccountKeeper) NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	_ = "STUB: not implemented"
	return *new(types.AccountI)
}

// NewAccount sets the next account number to a given account interface
func (ak AccountKeeper) NewAccount(ctx sdk.Context, acc types.AccountI) types.AccountI {
	_ = "STUB: not implemented"
	return *new(types.AccountI)
}

// HasAccount implements AccountKeeperI.
func (ak AccountKeeper) HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

// GetAccount implements AccountKeeperI.
func (ak AccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	_ = "STUB: not implemented"
	return *new(types.AccountI)
}

// GetAllAccounts returns all accounts in the accountKeeper.
func (ak AccountKeeper) GetAllAccounts(ctx sdk.Context) (accounts []types.AccountI) {
	_ = "STUB: not implemented"
	return nil
}

// SetAccount implements AccountKeeperI.
func (ak AccountKeeper) SetAccount(ctx sdk.Context, acc types.AccountI) {
	_ = "STUB: not implemented"
	return
}

// RemoveAccount removes an account for the account mapper store.
// NOTE: this will cause supply invariant violation if called
func (ak AccountKeeper) RemoveAccount(ctx sdk.Context, acc types.AccountI) {
	_ = "STUB: not implemented"
	return
}

// IterateAccounts iterates over all the stored accounts and performs a callback function.
// Stops iteration when callback returns true.
func (ak AccountKeeper) IterateAccounts(ctx sdk.Context, cb func(account types.AccountI) (stop bool)) {
	_ = "STUB: not implemented"
	return
}
