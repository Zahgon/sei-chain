package commands

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/spf13/cobra"
)

const wasmDirName = "wasm"

// MakeResetCommand constructs a command that removes the database of
// the specified Tendermint core instance.
func MakeResetCommand(conf *config.Config) *cobra.Command { _ = "STUB: not implemented"; return nil }

// If home is empty, use conf.RootDir as a fallback

// ResetAll removes address book files plus all data, and resets the privValdiator data.
// Exported for extenal CLI usage
// XXX: this is unsafe and should only suitable for testnets.
func ResetAll(dbDir, privValKeyFile, privValStateFile string, keyType string, homeDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// recreate the dbDir since the privVal state needs to live there

// removeIfExists removes a path if it exists, logging the result.
func removeIfExists(path, label string) { _ = "STUB: not implemented"; return }

// ResetState removes all blocks, tendermint state, indexed transactions and evidence.
// It handles both the legacy flat layout and the new subdirectory layout.
func ResetState(dbDir string) error {
	_ = "STUB: not implemented"
	// Legacy paths (flat under data/)
	return nil
}

// New paths (subdirectory layout — all tendermint DBs under data/tendermint/)

// ResetFilePV loads the file private validator and resets the watermark to 0. If used on an existing network,
// this can cause the node to double sign.
// XXX: this is unsafe and should only suitable for testnets.
func ResetFilePV(privValKeyFile, privValStateFile string, keyType string) error {
	_ = "STUB: not implemented"
	return nil
}

// ResetPeerStore removes the peer store containing all information used by the tendermint networking layer.
// In the case of a reset, new peers will need to be set either via the config or through the discovery mechanism.
// It checks both legacy (data/peerstore.db) and new (data/tendermint/peerstore.db) locations.
func ResetPeerStore(dbDir string) error { _ = "STUB: not implemented"; return nil }

func MakeUnsafeResetAllCommand(conf *config.Config) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// Get the --home flag value from the command

// If home is empty, use conf.RootDir as a fallback
