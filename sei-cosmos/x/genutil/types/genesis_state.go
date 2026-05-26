package types

import (
	"encoding/json"

	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(genTxs []json.RawMessage) *GenesisState {
	_ = "STUB: not implemented"
	// Ensure genTxs is never nil, https://github.com/cosmos/cosmos-sdk/issues/5086
	return nil
}

// DefaultGenesisState returns the genutil module's default genesis state.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// NewGenesisStateFromTx creates a new GenesisState object
// from auth transactions
func NewGenesisStateFromTx(txJSONEncoder sdk.TxEncoder, genTxs []sdk.Tx) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// GetGenesisStateFromAppState gets the genutil genesis state from the expected app state
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// SetGenesisStateInAppState sets the genutil genesis state within the expected app state
func SetGenesisStateInAppState(
	cdc codec.JSONCodec, appState map[string]json.RawMessage, genesisState *GenesisState,
) map[string]json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// GenesisStateFromGenDoc creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromGenDoc(genDoc tmtypes.GenesisDoc) (genesisState map[string]json.RawMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenesisStateFromGenFile creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromGenFile(genFile string) (genesisState map[string]json.RawMessage, genDoc *tmtypes.GenesisDoc, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ValidateGenesis validates GenTx transactions
func ValidateGenesis(genesisState *GenesisState, txJSONDecoder sdk.TxDecoder) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: abstract back to staking
