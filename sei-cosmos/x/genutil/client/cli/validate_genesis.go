package cli

import (
	"encoding/json"

	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

const (
	chainUpgradeGuide = "https://docs.cosmos.network/master/migrations/chain-upgrade-guide-040.html"
	flagStreaming     = "streaming"
)

// ValidateGenesisCmd takes a genesis file, and makes sure that it is valid.
func ValidateGenesisCmd(mbm module.BasicManager) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// Load default if passed no args, otherwise load passed file

type AppState struct {
	Module string          `json:"module"`
	Data   json.RawMessage `json:"data"`
}

type ModuleState struct {
	AppState AppState `json:"app_state"`
}

func parseModule(jsonStr string) (*ModuleState, error) { _ = "STUB: not implemented"; return nil, nil }

func validateGenesisStream(mbm module.BasicManager, cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Load default if passed no args, otherwise load passed file

// determine module name or genesisDoc

// new module

// same module

// validateGenDoc reads a genesis file and validates that it is a correct
// Tendermint GenesisDoc. This function does not do any cosmos-related
// validation.
func validateGenDoc(importGenesisFile string) (*tmtypes.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
