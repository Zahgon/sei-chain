package types

import (
	"sort"
)

// ----------------------------------------------------------------------------
// Decimal Coin

// NewDecCoin creates a new DecCoin instance from an Int.
func NewDecCoin(denom string, amount Int) DecCoin { _ = "STUB: not implemented"; return *new(DecCoin) }

// NewDecCoinFromDec creates a new DecCoin instance from a Dec.
func NewDecCoinFromDec(denom string, amount Dec) DecCoin {
	_ = "STUB: not implemented"
	return *new(DecCoin)
}

// NewDecCoinFromCoin creates a new DecCoin from a Coin.
func NewDecCoinFromCoin(coin Coin) DecCoin { _ = "STUB: not implemented"; return *new(DecCoin) }

// NewInt64DecCoin returns a new DecCoin with a denomination and amount. It will
// panic if the amount is negative or denom is invalid.
func NewInt64DecCoin(denom string, amount int64) DecCoin {
	_ = "STUB: not implemented"
	return *new(DecCoin)
}

// IsZero returns if the DecCoin amount is zero.
func (coin DecCoin) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsGTE returns true if they are the same type and the receiver is
// an equal or greater value.
func (coin DecCoin) IsGTE(other DecCoin) bool { _ = "STUB: not implemented"; return false }

// IsLT returns true if they are the same type and the receiver is
// a smaller value.
func (coin DecCoin) IsLT(other DecCoin) bool { _ = "STUB: not implemented"; return false }

// IsEqual returns true if the two sets of Coins have the same value.
func (coin DecCoin) IsEqual(other DecCoin) bool { _ = "STUB: not implemented"; return false }

// Add adds amounts of two decimal coins with same denom.
func (coin DecCoin) Add(coinB DecCoin) DecCoin { _ = "STUB: not implemented"; return *new(DecCoin) }

// Sub subtracts amounts of two decimal coins with same denom.
func (coin DecCoin) Sub(coinB DecCoin) DecCoin { _ = "STUB: not implemented"; return *new(DecCoin) }

// TruncateDecimal returns a Coin with a truncated decimal and a DecCoin for the
// change. Note, the change may be zero.
func (coin DecCoin) TruncateDecimal() (Coin, DecCoin) {
	_ = "STUB: not implemented"
	return *new(Coin), *new(DecCoin)
}

// IsPositive returns true if coin amount is positive.
//
// TODO: Remove once unsigned integers are used.
func (coin DecCoin) IsPositive() bool { _ = "STUB: not implemented"; return false }

// IsNegative returns true if the coin amount is negative and false otherwise.
//
// TODO: Remove once unsigned integers are used.
func (coin DecCoin) IsNegative() bool { _ = "STUB: not implemented"; return false }

// String implements the Stringer interface for DecCoin. It returns a
// human-readable representation of a decimal coin.
func (coin DecCoin) String() string { _ = "STUB: not implemented"; return "" }

// Validate returns an error if the DecCoin has a negative amount or if the denom is invalid.
func (coin DecCoin) Validate() error { _ = "STUB: not implemented"; return nil }

// IsValid returns true if the DecCoin has a non-negative amount and the denom is valid.
func (coin DecCoin) IsValid() bool { _ = "STUB: not implemented"; return false }

// ----------------------------------------------------------------------------
// Decimal Coins

// DecCoins defines a slice of coins with decimal values
type DecCoins []DecCoin

// NewDecCoins constructs a new coin set with with decimal values
// from DecCoins. The provided coins will be sanitized by removing
// zero coins and sorting the coin set. A panic will occur if the coin set is not valid.
func NewDecCoins(decCoins ...DecCoin) DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

func sanitizeDecCoins(decCoins []DecCoin) DecCoins {
	_ = "STUB: not implemented"
	// remove zeroes
	return *new(DecCoins)
}

// NewDecCoinsFromCoins constructs a new coin set with decimal values
// from regular Coins.
func NewDecCoinsFromCoins(coins ...Coin) DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

// String implements the Stringer interface for DecCoins. It returns a
// human-readable representation of decimal coins.
func (coins DecCoins) String() string { _ = "STUB: not implemented"; return "" }

// TruncateDecimal returns the coins with truncated decimals and returns the
// change. Note, it will not return any zero-amount coins in either the truncated or
// change coins.
func (coins DecCoins) TruncateDecimal() (truncatedCoins Coins, changeCoins DecCoins) {
	_ = "STUB: not implemented"
	return *new(Coins), *new(DecCoins)
}

// Add adds two sets of DecCoins.
//
// NOTE: Add operates under the invariant that coins are sorted by
// denominations.
//
// CONTRACT: Add will never return Coins where one Coin has a non-positive
// amount. In otherwords, IsValid will always return true.
func (coins DecCoins) Add(coinsB ...DecCoin) DecCoins {
	_ = "STUB: not implemented"
	return *new(DecCoins)
}

// safeAdd will perform addition of two DecCoins sets. If both coin sets are
// empty, then an empty set is returned. If only a single set is empty, the
// other set is returned. Otherwise, the coins are compared in order of their
// denomination and addition only occurs when the denominations match, otherwise
// the coin is simply added to the sum assuming it's not zero.
func (coins DecCoins) safeAdd(coinsB DecCoins) DecCoins {
	_ = "STUB: not implemented"
	return *new(DecCoins)
}

// return nil coins if both sets are empty

// return set B (excluding zero coins) if set A is empty

// return set A (excluding zero coins) if set B is empty

// coin A denom < coin B denom

// coin A denom == coin B denom

// coin A denom > coin B denom

// negative returns a set of coins with all amount negative.
func (coins DecCoins) negative() DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

// Sub subtracts a set of DecCoins from another (adds the inverse).
func (coins DecCoins) Sub(coinsB DecCoins) DecCoins {
	_ = "STUB: not implemented"
	return *new(DecCoins)
}

// SafeSub performs the same arithmetic as Sub but returns a boolean if any
// negative coin amount was returned.
func (coins DecCoins) SafeSub(coinsB DecCoins) (DecCoins, bool) {
	_ = "STUB: not implemented"
	return *new(DecCoins), false
}

// Intersect will return a new set of coins which contains the minimum DecCoin
// for common denoms found in both `coins` and `coinsB`. For denoms not common
// to both `coins` and `coinsB` the minimum is considered to be 0, thus they
// are not added to the final set. In other words, trim any denom amount from
// coin which exceeds that of coinB, such that (coin.Intersect(coinB)).IsLTE(coinB).
// See also Coins.Min().
func (coins DecCoins) Intersect(coinsB DecCoins) DecCoins {
	_ = "STUB: not implemented"
	return *new(DecCoins)
}

// UnionMax will return a new set of coins which contains the maximum DecCoin
// for denoms found in both `coins` and `coinsB`.
func (coins DecCoins) UnionMax(coinsB DecCoins) DecCoins {
	_ = "STUB: not implemented"
	return *new(DecCoins)
}

// Convert the map back to a slice.

// GetDenomByIndex returns the Denom to make the findDup generic
func (coins DecCoins) GetDenomByIndex(i int) string { _ = "STUB: not implemented"; return "" }

// IsAnyNegative returns true if there is at least one coin whose amount
// is negative; returns false otherwise. It returns false if the DecCoins set
// is empty too.
//
// TODO: Remove once unsigned integers are used.
func (coins DecCoins) IsAnyNegative() bool { _ = "STUB: not implemented"; return false }

// MulDec multiplies all the coins by a decimal.
//
// CONTRACT: No zero coins will be returned.
func (coins DecCoins) MulDec(d Dec) DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

// MulDecTruncate multiplies all the decimal coins by a decimal, truncating. It
// panics if d is zero.
//
// CONTRACT: No zero coins will be returned.
func (coins DecCoins) MulDecTruncate(d Dec) DecCoins {
	_ = "STUB: not implemented"
	return *new(DecCoins)
}

// QuoDec divides all the decimal coins by a decimal. It panics if d is zero.
//
// CONTRACT: No zero coins will be returned.
func (coins DecCoins) QuoDec(d Dec) DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

// QuoDecTruncate divides all the decimal coins by a decimal, truncating. It
// panics if d is zero.
//
// CONTRACT: No zero coins will be returned.
func (coins DecCoins) QuoDecTruncate(d Dec) DecCoins {
	_ = "STUB: not implemented"
	return *new(DecCoins)
}

// Empty returns true if there are no coins and false otherwise.
func (coins DecCoins) Empty() bool { _ = "STUB: not implemented"; return false }

// AmountOf returns the amount of a denom from deccoins
func (coins DecCoins) AmountOf(denom string) Dec { _ = "STUB: not implemented"; return *new(Dec) }

// IsEqual returns true if the two sets of DecCoins have the same value.
func (coins DecCoins) IsEqual(coinsB DecCoins) bool { _ = "STUB: not implemented"; return false }

// IsZero returns whether all coins are zero
func (coins DecCoins) IsZero() bool { _ = "STUB: not implemented"; return false }

// Validate checks that the DecCoins are sorted, have positive amount, with a valid and unique
// denomination (i.e no duplicates). Otherwise, it returns an error.
func (coins DecCoins) Validate() error { _ = "STUB: not implemented"; return nil }

// check single coin case

// we compare each coin against the last denom

// IsValid calls Validate and returns true when the DecCoins are sorted, have positive amount, with a
// valid and unique denomination (i.e no duplicates).
func (coins DecCoins) IsValid() bool { _ = "STUB: not implemented"; return false }

// IsAllPositive returns true if there is at least one coin and all currencies
// have a positive value.
//
// TODO: Remove once unsigned integers are used.
func (coins DecCoins) IsAllPositive() bool { _ = "STUB: not implemented"; return false }

func removeZeroDecCoins(coins DecCoins) DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

//-----------------------------------------------------------------------------
// Sorting

var _ sort.Interface = DecCoins{}

// Len implements sort.Interface for DecCoins
func (coins DecCoins) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface for DecCoins
	return 0
}

func (coins DecCoins) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface for DecCoins
func (coins DecCoins) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Sort is a helper function to sort the set of decimal coins in-place.
func (coins DecCoins) Sort() DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

// ----------------------------------------------------------------------------
// Parsing

// ParseDecCoin parses a decimal coin from a string, returning an error if
// invalid. An empty string is considered invalid.
func ParseDecCoin(coinStr string) (coin DecCoin, err error) {
	_ = "STUB: not implemented"
	return *new(DecCoin), nil
}

// ParseDecCoins will parse out a list of decimal coins separated by commas. If the parsing is successuful,
// the provided coins will be sanitized by removing zero coins and sorting the coin set. Lastly
// a validation of the coin set is executed. If the check passes, ParseDecCoins will return the sanitized coins.
// Otherwise it will return an error.
// If an empty string is provided to ParseDecCoins, it returns nil Coins.
// Expected format: "{amount0}{denomination},...,{amountN}{denominationN}"
func ParseDecCoins(coinsStr string) (DecCoins, error) {
	_ = "STUB: not implemented"
	return *new(DecCoins), nil
}
