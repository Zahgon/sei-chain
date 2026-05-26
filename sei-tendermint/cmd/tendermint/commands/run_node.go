package commands

import (
	"github.com/spf13/cobra"

	cfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
)

var (
	genesisHash []byte
)

// AddNodeFlags exposes some common configuration options from conf in the flag
// set for cmd. This is a convenience for commands embedding a Tendermint node.
func AddNodeFlags(cmd *cobra.Command, conf *cfg.Config) {
	_ = "STUB: not implemented"
	// bind flags
	return
}

// mode flags

// priv val flags

// node flags

// abci flags

// rpc flags

// p2p flags

// consensus flags

func mustMarkDeprecated(cmd *cobra.Command, name, message string) {
	_ = "STUB: not implemented"
	return
}

func addDBFlags(cmd *cobra.Command, conf *cfg.Config) { _ = "STUB: not implemented"; return }
