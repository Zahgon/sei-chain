package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewAggregateExchangeRateVote creates a AggregateExchangeRateVote instance
func NewAggregateExchangeRateVote(exchangeRateTuples ExchangeRateTuples, voter sdk.ValAddress) AggregateExchangeRateVote {
	_ = "STUB: not implemented"
	return *new(AggregateExchangeRateVote)
}

// String implement stringify
func (v AggregateExchangeRateVote) String() string { _ = "STUB: not implemented"; return "" }

// NewExchangeRateTuple creates a ExchangeRateTuple instance
func NewExchangeRateTuple(denom string, exchangeRate sdk.Dec) ExchangeRateTuple {
	_ = "STUB: not implemented"
	return *new(ExchangeRateTuple)
}

// String implement stringify
func (v ExchangeRateTuple) String() string { _ = "STUB: not implemented"; return "" }

// ExchangeRateTuples - array of ExchangeRateTuple
type ExchangeRateTuples []ExchangeRateTuple

// String implements fmt.Stringer interface
func (tuples ExchangeRateTuples) String() string { _ = "STUB: not implemented"; return "" }

// ParseExchangeRateTuples ExchangeRateTuple parser
func ParseExchangeRateTuples(tuplesStr string) (ExchangeRateTuples, error) {
	_ = "STUB: not implemented"
	return *new(ExchangeRateTuples), nil
}

// String implement stringify
func (ex OracleExchangeRate) String() string { _ = "STUB: not implemented"; return "" }

// OracleExchangeRates - array of OracleExchangeRate
type DenomOracleExchangeRatePairs []DenomOracleExchangeRatePair

// String implements fmt.Stringer interface
func (rates DenomOracleExchangeRatePairs) String() string { _ = "STUB: not implemented"; return "" }

func NewDenomOracleExchangeRatePair(denom string, exchangeRate sdk.Dec, lastUpdate sdk.Int, lastUpdateTimestamp int64) DenomOracleExchangeRatePair {
	_ = "STUB: not implemented"
	return *new(DenomOracleExchangeRatePair)
}

// VotePenaltyCounter - array of VotePenaltyCounter
type VotePenaltyCounters []VotePenaltyCounter
