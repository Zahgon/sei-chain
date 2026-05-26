package simulation

import (
	"math/rand"
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// shamelessly copied from
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang#31832326

// RandStringOfLength generates a random string of a particular length
func RandStringOfLength(r *rand.Rand, n int) string {
	_ = "STUB: not implemented"

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	return ""
}

//nolint:gosec // letterIdxMask ensures idx is small and non-negative

//nolint:gosec // intentional zero-alloc []byte to string conversion, b is not modified after

// RandPositiveInt get a rand positive sdk.Int
func RandPositiveInt(r *rand.Rand, max sdk.Int) (sdk.Int, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), nil
}

// RandomAmount generates a random amount
// Note: The range of RandomAmount includes max, and is, in fact, biased to return max as well as 0.
func RandomAmount(r *rand.Rand, max sdk.Int) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// randInt = big.NewInt(0)

// NOTE: there are 10 total cases.
// up to max - 1

// RandomDecAmount generates a random decimal amount
// Note: The range of RandomDecAmount includes max, and is, in fact, biased to return max as well as 0.
func RandomDecAmount(r *rand.Rand, max sdk.Dec) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// randInt = big.NewInt(0)

// the underlying big int with all precision bits.
// NOTE: there are 10 total cases.

// RandTimestamp generates a random timestamp
func RandTimestamp(r *rand.Rand) time.Time {
	_ = "STUB: not implemented"
	// json.Marshal breaks for timestamps greater with year greater than 9999
	return *new(time.Time)
}

// RandIntBetween returns a random int between two numbers inclusively.
func RandIntBetween(r *rand.Rand, min, max int) int { _ = "STUB: not implemented"; return 0 }

// returns random subset of the provided coins
// will return at least one coin unless coins argument is empty or malformed
// i.e. 0 amt in coins
func RandSubsetCoins(r *rand.Rand, coins sdk.Coins) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// make sure at least one coin added

// malformed coin. 0 amt in coins

// skip denom that we already chose earlier

// coin flip if multiple coins
// if there is single coin then return random amount of it

// ignore errors and try another denom

// DeriveRand derives a new Rand deterministically from another random source.
// Unlike rand.New(rand.NewSource(seed)), the result is "more random"
// depending on the source and state of r.
//
// NOTE: not crypto safe.
func DeriveRand(r *rand.Rand) *rand.Rand {
	_ = "STUB: not implemented"
	// TODO what's a good number?  Too large is too slow.
	return nil
}

type multiSource []rand.Source

func (ms multiSource) Int63() (r int64) { _ = "STUB: not implemented"; return 0 }

func (ms multiSource) Seed(seed int64) { _ = "STUB: not implemented"; return }
