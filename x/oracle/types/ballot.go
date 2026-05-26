package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NOTE: we don't need to implement proto interface on this file
//       these are not used in store or rpc response

// VoteForTally is a convenience wrapper to reduce redundant lookup cost
type VoteForTally struct {
	Denom        string
	ExchangeRate sdk.Dec
	Voter        sdk.ValAddress
	Power        int64
}

// NewVoteForTally returns a new VoteForTally instance
func NewVoteForTally(rate sdk.Dec, denom string, voter sdk.ValAddress, power int64) VoteForTally {
	_ = "STUB: not implemented"
	return *new(VoteForTally)
}

// ExchangeRateBallot is a convenience wrapper around a ExchangeRateVote slice
type ExchangeRateBallot []VoteForTally

// ToMap return organized exchange rate map by validator
func (pb ExchangeRateBallot) ToMap() map[string]sdk.Dec { _ = "STUB: not implemented"; return nil }

// ToCrossRate return cross_rate(base/exchange_rate) ballot
func (pb ExchangeRateBallot) ToCrossRate(bases map[string]sdk.Dec) (cb ExchangeRateBallot) {
	_ = "STUB: not implemented"
	return *new(ExchangeRateBallot)
}

// Quo will panic on overflow, so we wrap it in a defer/recover

// if overflow, set exchange rate to 0 and power to 0

// If we can't get reference Sei exchange rate, we just convert the vote as abstain vote

// ToCrossRateWithSort return cross_rate(base/exchange_rate) ballot
func (pb ExchangeRateBallot) ToCrossRateWithSort(bases map[string]sdk.Dec) (cb ExchangeRateBallot) {
	_ = "STUB: not implemented"
	return *new(ExchangeRateBallot)
}

// Power returns the total amount of voting power in the ballot
func (pb ExchangeRateBallot) Power() int64 { _ = "STUB: not implemented"; return 0 }

// WeightedMedian returns the median weighted by the power of the ExchangeRateVote.
// CONTRACT: ballot must be sorted
func (pb ExchangeRateBallot) WeightedMedian() sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// WeightedMedianWithAssertion returns the median weighted by the power of the ExchangeRateVote.
// CONTRACT: ballot must be sorted
func (pb ExchangeRateBallot) WeightedMedianWithAssertion() sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// StandardDeviation returns the standard deviation by the power of the ExchangeRateVote.
func (pb ExchangeRateBallot) StandardDeviation(median sdk.Dec) (standardDeviation sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Len implements sort.Interface
func (pb ExchangeRateBallot) Len() int {
	_ = "STUB: not implemented"

	// Less reports whether the element with
	// index i should sort before the element with index j.
	return 0
}

func (pb ExchangeRateBallot) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface.
func (pb ExchangeRateBallot) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Claim is an interface that directs its rewards to an attached bank account.
type Claim struct {
	Power     int64
	Weight    int64
	WinCount  int64
	DidVote   bool
	Recipient sdk.ValAddress
}

// NewClaim generates a Claim instance.
func NewClaim(power, weight, winCount int64, recipient sdk.ValAddress, didVote bool) Claim {
	_ = "STUB: not implemented"
	return *new(Claim)
}
