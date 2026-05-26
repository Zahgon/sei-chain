package keeper

import (
	"github.com/sei-protocol/sei-chain/giga/deps/xbank/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/prefix"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var _ ViewKeeper = (*BaseViewKeeper)(nil)

// ViewKeeper defines a module interface that facilitates read only access to
// account balances.
type ViewKeeper interface {
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool

	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetWeiBalance(ctx sdk.Context, addr sdk.AccAddress) sdk.Int
}

// BaseViewKeeper implements a read only keeper implementation of ViewKeeper.
type BaseViewKeeper struct {
	cdc       codec.BinaryCodec
	storeKey  sdk.StoreKey
	ak        types.AccountKeeper
	cacheSize int

	// UseRegularStore when true causes GetKVStore to use ctx.KVStore instead of ctx.GigaKVStore.
	// This is for debugging/testing to isolate Giga executor logic from GigaKVStore layer.
	UseRegularStore bool
}

// NewBaseViewKeeper returns a new BaseViewKeeper.
func NewBaseViewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, ak types.AccountKeeper) BaseViewKeeper {
	_ = "STUB: not implemented"
	return *new(BaseViewKeeper)
}

// GetKVStore returns the appropriate KVStore based on the UseRegularStore flag.
// When UseRegularStore is true (for debugging/testing), returns regular KVStore.
// Otherwise returns GigaKVStore.
func (k BaseViewKeeper) GetKVStore(ctx sdk.Context) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

// HasBalance returns whether or not an account has at least amt balance.
func (k BaseViewKeeper) HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool {
	_ = "STUB: not implemented"
	return false
}

// GetBalance returns the balance of a specific denomination for a given account
// by address.
func (k BaseViewKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

// LockedCoins returns all the coins that are not spendable (i.e. locked) for an
// account by address. For standard accounts, the result will always be no coins.
// For vesting accounts, LockedCoins is delegated to the concrete vesting account
// type.
func (k BaseViewKeeper) LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// SpendableCoins returns the total balances of spendable coins for an account
// by address. If the account has no spendable coins, an empty Coins slice is
// returned.
func (k BaseViewKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// getAccountStore gets the account store of the given address.
func (k BaseViewKeeper) getAccountStore(ctx sdk.Context, addr sdk.AccAddress) prefix.Store {
	_ = "STUB: not implemented"
	return *new(prefix.Store)
}

func (k BaseViewKeeper) GetWeiBalance(ctx sdk.Context, addr sdk.AccAddress) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// should never happen
