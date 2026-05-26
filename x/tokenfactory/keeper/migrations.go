package keeper

import (
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const KeySeparator = "|"

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	// Reset params after removing the denom creation fee param
	return nil
}

// We remove the denom creation fee whitelist in this migration

func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	// Set denom metadata for all denoms
	return nil
}

func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	// Add new params and set all to defaults
	return nil
}

func (m Migrator) SetMetadata(denomMetadata *banktypes.Metadata) { _ = "STUB: not implemented"; return }
