package commands

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
)

func MakeKeyMigrateCommand(conf *config.Config) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// allow database info to be overridden via cli

func RunDatabaseMigration(ctx context.Context, conf *config.Config) error {
	_ = "STUB: not implemented"
	return nil

	// this is ordered to put
	// the more ephemeral tables first to
	// reduce the possibility of the
	// ephemeral data overwriting later data
}
