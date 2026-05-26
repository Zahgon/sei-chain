package commands

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
)

// LightCmd constructs the base command called when invoked without any subcommands.
func MakeLightCommand(conf *config.Config) *cobra.Command { _ = "STUB: not implemented"; return nil }

// create a prefixed db on the chainID

// check to see if we can start from an existing state

//nolint:prealloc

// Initiate the light client. If the trusted store already has blocks in it, this
// will be used else we use the trusted options.

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435
// Note we don't need to adjust anything if the timeout is already unlimited.

// Error starting or closing listener:
