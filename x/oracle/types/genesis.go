package types

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params, rates []ExchangeRateTuple,
	feederDelegations []FeederDelegation, penaltyCounters []PenaltyCounter,
	aggregateExchangeRateVotes []AggregateExchangeRateVote,
	priceSnapshots []PriceSnapshot,
) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState - default GenesisState used by columbus-2
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// ValidateGenesis validates the oracle genesis state
func ValidateGenesis(data *GenesisState) error { _ = "STUB: not implemented"; return nil }

// GetGenesisStateFromAppState returns x/oracle GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}
