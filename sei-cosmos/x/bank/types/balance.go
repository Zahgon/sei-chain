package types

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/exported"
)

var _ exported.GenesisBalance = (*Balance)(nil)

// GetAddress returns the account address of the Balance object.
func (b Balance) GetAddress() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// GetCoins returns the account coins of the Balance object.
func (b Balance) GetCoins() sdk.Coins {
	_ = "STUB: not implemented"

	// Validate checks for address and coins correctness.
	return *new(sdk.Coins)
}

func (b Balance) Validate() error { _ = "STUB: not implemented"; return nil }

type balanceByAddress struct {
	addresses []sdk.AccAddress
	balances  []Balance
}

func (b balanceByAddress) Len() int           { _ = "STUB: not implemented"; return 0 }
func (b balanceByAddress) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (b balanceByAddress) Swap(i, j int) { _ = "STUB: not implemented"; return }

// SanitizeGenesisBalances sorts addresses and coin sets.
func SanitizeGenesisBalances(balances []Balance) []Balance {
	_ = "STUB: not implemented"
	// Given that this function sorts balances, using the standard library's
	// Quicksort based algorithms, we have algorithmic complexities of:
	// * Best case: O(nlogn)
	// * Worst case: O(n^2)
	// The comparator used MUST be cheap to use lest we incur expenses like we had
	// before whereby sdk.AccAddressFromBech32, which is a very expensive operation
	// compared n * n elements yet discarded computations each time, as per:
	//
	//	https://github.com/cosmos/cosmos-sdk/issues/7766#issuecomment-786671734
	return nil
}

// 1. Retrieve the address equivalents for each Balance's address.

// 2. Sort balances.

// GenesisBalancesIterator implements genesis account iteration.
type GenesisBalancesIterator struct{}

// IterateGenesisBalances iterates over all the genesis balances found in
// appGenesis and invokes a callback on each genesis account. If any call
// returns true, iteration stops.
func (GenesisBalancesIterator) IterateGenesisBalances(
	cdc codec.JSONCodec, appState map[string]json.RawMessage, cb func(exported.GenesisBalance) (stop bool),
) {
	_ = "STUB: not implemented"
	return
}
