package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

// create proto for new data value
// because we don't have a lastUpdate, we set it to 0

// Migrate3to4 migrates from version 3 to 4
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	// we need to migrate the miss counters to be stored as VotePenaltyCounter to introduce abstain count logic
	return nil
}

// previously the data was stored as uint64, now it is VotePenaltyCounter proto

// create proto for new data value
// because we don't have a lastUpdate, we set it to 0

// Migrate3to4 migrates from version 4 to 5
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	// we remove the prevotes from store in this migration
	return nil
}

func (m Migrator) Migrate5To6(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	// Do a one time backfill for success count in the vote penalty counter
	return nil
}

// previously the data was stored as uint64, now it is VotePenaltyCounter proto

// nolint:gosec
