package commands

import (
	"github.com/sei-protocol/seilog"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
)

var logger = seilog.NewLogger("tendermint", "cmd", "tendermint", "commands")

// InspectCmd constructs the command to start an inspect server.
func MakeInspectCommand(conf *config.Config) *cobra.Command { _ = "STUB: not implemented"; return nil }
