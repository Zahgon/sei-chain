package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// MigrateTraces migrates the DenomTraces to the correct format, accounting for slashes in the BaseDenom.
func (m Migrator) MigrateTraces(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	// list of traces that must replace the old traces in store
	return nil
}

// check if the new way of splitting FullDenom
// is the same as the current DenomTrace.
// If it isn't then store the new DenomTrace in the list of new traces.

// The new form of parsing will result in a token denomination change.
// A bank migration is required. A panic should occur to prevent the
// chain from using corrupted state.

// replace the outdated traces with the new trace information

func equalTraces(dtA, dtB types.DenomTrace) bool { _ = "STUB: not implemented"; return false }
