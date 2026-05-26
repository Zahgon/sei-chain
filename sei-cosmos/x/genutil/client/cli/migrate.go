package cli

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil/types"
)

const flagGenesisTime = "genesis-time"

// Allow applications to extend and modify the migration process.
//
// Ref: https://github.com/cosmos/cosmos-sdk/issues/5041
var migrationMap = types.MigrationMap{}

// GetMigrationCallback returns a MigrationCallback for a given version.
func GetMigrationCallback(version string) types.MigrationCallback {
	_ = "STUB: not implemented"
	return *new(types.MigrationCallback)
}

// GetMigrationVersions get all migration version in a sorted slice.
func GetMigrationVersions() []string { _ = "STUB: not implemented"; return nil }

// MigrateGenesisCmd returns a command to execute genesis state migration.
func MigrateGenesisCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Since some default values are valid values, we just print to
// make sure the user didn't forget to update these values.

// TODO: handler error from migrationFunc call
