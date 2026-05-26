package types

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Validate performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

// NOTE: this errors if supply for any given coin is zero

func getTotalSupply(genState *GenesisState) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

var OneUseiInWei sdk.Int = sdk.NewInt(1_000_000_000_000)

func SplitUseiWeiAmount(amt sdk.Int) (sdk.Int, sdk.Int) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), *new(sdk.Int)
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(params Params, balances []Balance, supply sdk.Coins, denomMetaData []Metadata, weiBalances []WeiBalance) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState returns a default bank module genesis state.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// GetGenesisStateFromAppState returns x/bank GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}
