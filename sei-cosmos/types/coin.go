package types

import (
	"regexp"
	"sort"
	"sync"
)

//-----------------------------------------------------------------------------
// Coin

// NewCoin returns a new coin with a denomination and amount. It will panic if
// the amount is negative or if the denomination is invalid.
func NewCoin(denom string, amount Int) Coin { _ = "STUB: not implemented"; return *new(Coin) }

// NewInt64Coin returns a new coin with a denomination and amount. It will panic
// if the amount is negative.
func NewInt64Coin(denom string, amount int64) Coin { _ = "STUB: not implemented"; return *new(Coin) }

// String provides a human-readable representation of a coin
func (coin Coin) String() string { _ = "STUB: not implemented"; return "" }

// Validate returns an error if the Coin has a negative amount or if
// the denom is invalid.
func (coin Coin) Validate() error { _ = "STUB: not implemented"; return nil }

// IsValid returns true if the Coin has a non-negative amount and the denom is valid.
func (coin Coin) IsValid() bool { _ = "STUB: not implemented"; return false }

// IsZero returns if this represents no money
func (coin Coin) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsGTE returns true if they are the same type and the receiver is
// an equal or greater value
func (coin Coin) IsGTE(other Coin) bool { _ = "STUB: not implemented"; return false }

// IsLT returns true if they are the same type and the receiver is
// a smaller value
func (coin Coin) IsLT(other Coin) bool { _ = "STUB: not implemented"; return false }

// IsEqual returns true if the two sets of Coins have the same value
func (coin Coin) IsEqual(other Coin) bool { _ = "STUB: not implemented"; return false }

// Add adds amounts of two coins with same denom. If the coins differ in denom then
// it panics.
func (coin Coin) Add(coinB Coin) Coin { _ = "STUB: not implemented"; return *new(Coin) }

// AddAmount adds an amount to the Coin.
func (coin Coin) AddAmount(amount Int) Coin { _ = "STUB: not implemented"; return *new(Coin) }

// Sub subtracts amounts of two coins with same denom. If the coins differ in denom
// then it panics.
func (coin Coin) Sub(coinB Coin) Coin { _ = "STUB: not implemented"; return *new(Coin) }

func (coin Coin) SubUnsafe(coinB Coin) Coin { _ = "STUB: not implemented"; return *new(Coin) }

// SubAmount subtracts an amount from the Coin.
func (coin Coin) SubAmount(amount Int) Coin { _ = "STUB: not implemented"; return *new(Coin) }

// IsPositive returns true if coin amount is positive.
//
// TODO: Remove once unsigned integers are used.
func (coin Coin) IsPositive() bool { _ = "STUB: not implemented"; return false }

// IsNegative returns true if the coin amount is negative and false otherwise.
//
// TODO: Remove once unsigned integers are used.
func (coin Coin) IsNegative() bool { _ = "STUB: not implemented"; return false }

// IsNil returns true if the coin amount is nil and false otherwise.
func (coin Coin) IsNil() bool { _ = "STUB: not implemented"; return false }

//-----------------------------------------------------------------------------
// Coins

// Coins is a set of Coin, one per currency
type Coins []Coin

// NewCoins constructs a new coin set. The provided coins will be sanitized by removing
// zero coins and sorting the coin set. A panic will occur if the coin set is not valid.
func NewCoins(coins ...Coin) Coins { _ = "STUB: not implemented"; return *new(Coins) }

func sanitizeCoins(coins []Coin) Coins { _ = "STUB: not implemented"; return *new(Coins) }

type coinsJSON Coins

// MarshalJSON implements a custom JSON marshaller for the Coins type to allow
// nil Coins to be encoded as an empty array.
func (coins Coins) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (coins Coins) String() string { _ = "STUB: not implemented"; return "" }

// Build the string with a string builder

// Validate checks that the Coins are sorted, have positive amount, with a valid and unique
// denomination (i.e no duplicates). Otherwise, it returns an error.
func (coins Coins) Validate() error { _ = "STUB: not implemented"; return nil }

// check single coin case

// we compare each coin against the last denom

func (coins Coins) isSorted() bool { _ = "STUB: not implemented"; return false }

// IsValid calls Validate and returns true when the Coins are sorted, have positive amount, with a
// valid and unique denomination (i.e no duplicates).
func (coins Coins) IsValid() bool { _ = "STUB: not implemented"; return false }

// Add adds two sets of coins.
//
// e.g.
// {2A} + {A, 2B} = {3A, 2B}
// {2A} + {0B} = {2A}
//
// NOTE: Add operates under the invariant that coins are sorted by
// denominations.
//
// CONTRACT: Add will never return Coins where one Coin has a non-positive
// amount. In otherwords, IsValid will always return true.
// The function panics if `coins` or  `coinsB` are not sorted (ascending).
func (coins Coins) Add(coinsB ...Coin) Coins { _ = "STUB: not implemented"; return *new(Coins) }

// safeAdd will perform addition of two coins sets. If both coin sets are
// empty, then an empty set is returned. If only a single set is empty, the
// other set is returned. Otherwise, the coins are compared in order of their
// denomination and addition only occurs when the denominations match, otherwise
// the coin is simply added to the sum assuming it's not zero.
// The function panics if `coins` or  `coinsB` are not sorted (ascending).
func (coins Coins) safeAdd(coinsB Coins) Coins {
	_ = "STUB: not implemented"
	// probably the best way will be to make Coins and interface and hide the structure
	// definition (type alias)
	return *new(Coins)
}

// return nil coins if both sets are empty

// return set B (excluding zero coins) if set A is empty

// return set A (excluding zero coins) if set B is empty

// coin A denom < coin B denom

// coin A denom == coin B denom

// coin A denom > coin B denom

// DenomsSubsetOf returns true if receiver's denom set
// is subset of coinsB's denoms.
func (coins Coins) DenomsSubsetOf(coinsB Coins) bool {
	_ = "STUB: not implemented"
	// more denoms in B than in receiver
	return false
}

// Sub subtracts a set of coins from another.
//
// e.g.
// {2A, 3B} - {A} = {A, 3B}
// {2A} - {0B} = {2A}
// {A, B} - {A} = {B}
//
// CONTRACT: Sub will never return Coins where one Coin has a non-positive
// amount. In otherwords, IsValid will always return true.
func (coins Coins) Sub(coinsB Coins) Coins { _ = "STUB: not implemented"; return *new(Coins) }

// SafeSub performs the same arithmetic as Sub but returns a boolean if any
// negative coin amount was returned.
// The function panics if `coins` or  `coinsB` are not sorted (ascending).
func (coins Coins) SafeSub(coinsB Coins) (Coins, bool) {
	_ = "STUB: not implemented"
	return *new(Coins), false
}

// PartitionSigned separates the coins between those with positive values and those with negative values,
// and returns both positive and negative coins. This also discards any coins with zero values (not returned in either group).
func (coins Coins) PartitionSigned() (positives Coins, negatives Coins) {
	_ = "STUB: not implemented"
	return *new(Coins), *new(Coins)
}

// Max takes two valid Coins inputs and returns a valid Coins result
// where for every denom D, AmountOf(D) of the result is the maximum
// of AmountOf(D) of the inputs.  Note that the result might be not
// be equal to either input. For any valid Coins a, b, and c, the
// following are always true:
//
//	a.IsAllLTE(a.Max(b))
//	b.IsAllLTE(a.Max(b))
//	a.IsAllLTE(c) && b.IsAllLTE(c) == a.Max(b).IsAllLTE(c)
//	a.Add(b...).IsEqual(a.Min(b).Add(a.Max(b)...))
//
// E.g.
// {1A, 3B, 2C}.Max({4A, 2B, 2C} == {4A, 3B, 2C})
// {2A, 3B}.Max({1B, 4C}) == {2A, 3B, 4C}
// {1A, 2B}.Max({}) == {1A, 2B}
func (coins Coins) Max(coinsB Coins) Coins { _ = "STUB: not implemented"; return *new(Coins) }

// denom missing from coinsB

// same denom in both

// denom missing from coinsA

// Min takes two valid Coins inputs and returns a valid Coins result
// where for every denom D, AmountOf(D) of the result is the minimum
// of AmountOf(D) of the inputs.  Note that the result might be not
// be equal to either input. For any valid Coins a, b, and c, the
// following are always true:
//
//	a.Min(b).IsAllLTE(a)
//	a.Min(b).IsAllLTE(b)
//	c.IsAllLTE(a) && c.IsAllLTE(b) == c.IsAllLTE(a.Min(b))
//	a.Add(b...).IsEqual(a.Min(b).Add(a.Max(b)...))
//
// E.g.
// {1A, 3B, 2C}.Min({4A, 2B, 2C} == {1A, 2B, 2C})
// {2A, 3B}.Min({1B, 4C}) == {1B}
// {1A, 2B}.Min({3C}) == empty
//
// See also DecCoins.Intersect().
func (coins Coins) Min(coinsB Coins) Coins { _ = "STUB: not implemented"; return *new(Coins) }

// denom missing from coinsB

// same denom in both

// denom missing from coins

// IsAllGT returns true if for every denom in coinsB,
// the denom is present at a greater amount in coins.
func (coins Coins) IsAllGT(coinsB Coins) bool { _ = "STUB: not implemented"; return false }

// IsAllGTE returns false if for any denom in coinsB,
// the denom is present at a smaller amount in coins;
// else returns true.
func (coins Coins) IsAllGTE(coinsB Coins) bool { _ = "STUB: not implemented"; return false }

// IsAllLT returns True iff for every denom in coins, the denom is present at
// a smaller amount in coinsB.
func (coins Coins) IsAllLT(coinsB Coins) bool { _ = "STUB: not implemented"; return false }

// IsAllLTE returns true iff for every denom in coins, the denom is present at
// a smaller or equal amount in coinsB.
func (coins Coins) IsAllLTE(coinsB Coins) bool { _ = "STUB: not implemented"; return false }

// IsAnyGT returns true iff for any denom in coins, the denom is present at a
// greater amount in coinsB.
//
// e.g.
// {2A, 3B}.IsAnyGT{A} = true
// {2A, 3B}.IsAnyGT{5C} = false
// {}.IsAnyGT{5C} = false
// {2A, 3B}.IsAnyGT{} = false
func (coins Coins) IsAnyGT(coinsB Coins) bool { _ = "STUB: not implemented"; return false }

// IsAnyGTE returns true iff coins contains at least one denom that is present
// at a greater or equal amount in coinsB; it returns false otherwise.
//
// NOTE: IsAnyGTE operates under the invariant that both coin sets are sorted
// by denominations and there exists no zero coins.
func (coins Coins) IsAnyGTE(coinsB Coins) bool { _ = "STUB: not implemented"; return false }

// IsZero returns true if there are no coins or all coins are zero.
func (coins Coins) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsEqual returns true if the two sets of Coins have the same value
func (coins Coins) IsEqual(coinsB Coins) bool { _ = "STUB: not implemented"; return false }

// Empty returns true if there are no coins and false otherwise.
func (coins Coins) Empty() bool { _ = "STUB: not implemented"; return false }

// AmountOf returns the amount of a denom from coins
func (coins Coins) AmountOf(denom string) Int { _ = "STUB: not implemented"; return *new(Int) }

// NonZeroAmountsOf returns non-zero coins for provided denoms
func (coins Coins) NonZeroAmountsOf(denoms []string) (subset Coins) {
	_ = "STUB: not implemented"
	return *new(Coins)
}

// AmountOfNoDenomValidation returns the amount of a denom from coins
// without validating the denomination.
func (coins Coins) AmountOfNoDenomValidation(denom string) Int {
	_ = "STUB: not implemented"
	return *new(Int)
}

// GetDenomByIndex returns the Denom of the certain coin to make the findDup generic
func (coins Coins) GetDenomByIndex(i int) string { _ = "STUB: not implemented"; return "" }

// IsAllPositive returns true if there is at least one coin and all currencies
// have a positive value.
func (coins Coins) IsAllPositive() bool { _ = "STUB: not implemented"; return false }

// IsAnyNegative returns true if there is at least one coin whose amount
// is negative; returns false otherwise. It returns false if the coin set
// is empty too.
//
// TODO: Remove once unsigned integers are used.
func (coins Coins) IsAnyNegative() bool { _ = "STUB: not implemented"; return false }

// IsAnyNil returns true if there is at least one coin whose amount
// is nil; returns false otherwise. It returns false if the coin set
// is empty too.
func (coins Coins) IsAnyNil() bool { _ = "STUB: not implemented"; return false }

// negative returns a set of coins with all amount negative.
//
// TODO: Remove once unsigned integers are used.
func (coins Coins) negative() Coins { _ = "STUB: not implemented"; return *new(Coins) }

// removeZeroCoins removes all zero coins from the given coin set in-place.
func removeZeroCoins(coins Coins) Coins { _ = "STUB: not implemented"; return *new(Coins) }

//-----------------------------------------------------------------------------
// Sort interface

// Len implements sort.Interface for Coins
func (coins Coins) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface for Coins
	return 0
}

func (coins Coins) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface for Coins
func (coins Coins) Swap(i, j int) { _ = "STUB: not implemented"; return }

var _ sort.Interface = Coins{}

// Sort is a helper function to sort the set of coins in-place
func (coins Coins) Sort() Coins { _ = "STUB: not implemented"; return *new(Coins) }

//-----------------------------------------------------------------------------
// Parsing

var (
	// Denominations can be 3 ~ 128 characters long and support letters, followed by either
	// a letter, a number or a separator ('/').
	reDnmString = `[a-zA-Z][a-zA-Z0-9/-]{2,127}`
	reDecAmt    = `[[:digit:]]+(?:\.[[:digit:]]+)?|\.[[:digit:]]+`
	reSpc       = `[[:space:]]*`
	reDnm       *regexp.Regexp
	reDecCoin   *regexp.Regexp

	// denomRegexMu protects coinDenomRegex, reDnm, and reDecCoin from concurrent access
	denomRegexMu sync.RWMutex
)

func init() {
	SetCoinDenomRegex(DefaultCoinDenomRegex)
}

// DefaultCoinDenomRegex returns the default regex string
func DefaultCoinDenomRegex() string {
	_ = "STUB: not implemented"

	// coinDenomRegex returns the current regex string and can be overwritten for custom validation
	return ""
}

var coinDenomRegex = DefaultCoinDenomRegex

// SetCoinDenomRegex allows for coin's custom validation by overriding the regular
// expression string used for denom validation.
func SetCoinDenomRegex(reFn func() string) { _ = "STUB: not implemented"; return }

// ValidateDenom is the default validation function for Coin.Denom.
func ValidateDenom(denom string) error { _ = "STUB: not implemented"; return nil }

func mustValidateDenom(denom string) { _ = "STUB: not implemented"; return }

// ParseCoinNormalized parses and normalize a cli input for one coin type, returning errors if invalid or on an empty string
// as well.
// Expected format: "{amount}{denomination}"
func ParseCoinNormalized(coinStr string) (coin Coin, err error) {
	_ = "STUB: not implemented"
	return *new(Coin), nil
}

// ParseCoinsNormalized will parse out a list of coins separated by commas, and normalize them by converting to the smallest
// unit. If the parsing is successful, the provided coins will be sanitized by removing zero coins and sorting the coin
// set. Lastly a validation of the coin set is executed. If the check passes, ParseCoinsNormalized will return the
// sanitized coins.
// Otherwise, it will return an error.
// If an empty string is provided to ParseCoinsNormalized, it returns nil Coins.
// ParseCoinsNormalized supports decimal coins as inputs, and truncate them to int after converted to the smallest unit.
// Expected format: "{amount0}{denomination},...,{amountN}{denominationN}"
func ParseCoinsNormalized(coinStr string) (Coins, error) {
	_ = "STUB: not implemented"
	return *new(Coins), nil
}
