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

// Migrate1to2 migrates from version 1 to 2.
// This migration prunes:
// - migrates solo machine client state from protobuf definition v1 to v2
// - prunes solo machine consensus states
// - prunes expired tendermint consensus states
// - adds ProcessedHeight and Iteration keys for unexpired tendermint consensus states
func (m Migrator) Migrate1to2(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

func (m Migrator) Migrate2to3(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }
