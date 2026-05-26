package genutil

import (
	"encoding/json"
	"time"

	cfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// ExportGenesisFile creates and writes the genesis configuration to disk. An
// error is returned if building or writing the configuration to file fails.
func ExportGenesisFile(genDoc *tmtypes.GenesisDoc, genFile string) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesisFileWithTime creates and writes the genesis configuration to disk.
// An error is returned if building or writing the configuration to file fails.
func ExportGenesisFileWithTime(
	genFile, chainID string, validators []tmtypes.GenesisValidator,
	appState json.RawMessage, genTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// InitializeNodeValidatorFiles creates private validator and p2p configuration files.
func InitializeNodeValidatorFiles(config *cfg.Config) (nodeID string, valPubKey cryptotypes.PubKey, err error) {
	_ = "STUB: not implemented"
	return "", *new(cryptotypes.PubKey), nil
}

// InitializeNodeValidatorFiles creates private validator and p2p configuration files using the given mnemonic.
// If no valid mnemonic is given, a random one will be used instead.
func InitializeNodeValidatorFilesFromMnemonic(config *cfg.Config, mnemonic string) (nodeID string, valPubKey cryptotypes.PubKey, err error) {
	_ = "STUB: not implemented"
	return "", *new(cryptotypes.PubKey), nil
}
