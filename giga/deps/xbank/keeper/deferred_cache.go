package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/prefix"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type DeferredCache struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec
}

func NewDeferredCache(cdc codec.BinaryCodec, storeKey sdk.StoreKey) *DeferredCache {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeferredCache) getModuleTxIndexedStore(ctx sdk.Context, moduleAddr sdk.AccAddress, txIndex uint64) prefix.Store {
	_ = "STUB: not implemented"
	return *new(prefix.Store)
}

// GetBalance returns the balance of a specific denomination for a given module address and transaction index
func (d *DeferredCache) GetBalance(ctx sdk.Context, moduleAddr sdk.AccAddress, txIndex uint64, denom string) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

// setBalance sets the coin balance for a module and tx Index.
func (d *DeferredCache) setBalance(ctx sdk.Context, moduleAddr sdk.AccAddress, txIndex uint64, balance sdk.Coin) error {
	_ = "STUB: not implemented"
	return nil
}

// Bank invariants require to not store zero balances, so we follow the same pattern in deferred cache.

// upsertBalance updates or sets the coin balance for a module and tx combination keyed on balance denom.
func (d *DeferredCache) upsertBalance(ctx sdk.Context, moduleAddr sdk.AccAddress, txIndex uint64, balance sdk.Coin) error {
	_ = "STUB: not implemented"
	return nil
}

// UpsertBalances updates or sets the coin balances for a module and tx combination with the given coins.
func (d *DeferredCache) UpsertBalances(ctx sdk.Context, moduleAddr sdk.AccAddress, txIndex uint64, balances sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// iterate through coins and upsert

// IterateAccountBalances iterates over the balances of a single module for all tx indices and
// provides the token balance to a callback.
// Note that because there can be multiple tx indices per module,
// there can be multiple occurrences of the same denom in `balance`.
// If true is returned from the
// callback, iteration is halted.
func (d *DeferredCache) IterateDeferredBalances(ctx sdk.Context, cb func(moduleAddr sdk.AccAddress, balance sdk.Coin) bool) {
	_ = "STUB: not implemented"
	return
}

// Clear deletes all of the keys in the deferred cache
func (d *DeferredCache) Clear(ctx sdk.Context) { _ = "STUB: not implemented"; return }
