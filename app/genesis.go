package app

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	genesistypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/genesis"

	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
)

var DefaultGenesisConfig = genesistypes.GenesisImportConfig{
	StreamGenesisImport: false,
	GenesisStreamFile:   "",
}

const (
	flagGenesisStreamImport = "genesis.stream-import"
	flagGenesisImportFile   = "genesis.import-file"
)

func ReadGenesisImportConfig(opts servertypes.AppOptions) (genesistypes.GenesisImportConfig, error) {
	_ = "STUB: not implemented"
	return *
	// copy
	new(genesistypes.GenesisImportConfig), nil
}

// The genesis state of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	_ = "STUB: not implemented"
	return *new(GenesisState)
}

// Override distribution config to remove community tax
