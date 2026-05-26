package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("x", "mint", "keeper")

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a v3 Migrator.
func NewMigrator(keeper Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

func (m Migrator) Migrate1to2(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

func (m Migrator) Migrate2to3(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

// Migrate Minter First

//nolint:gosec

// Migrate TokenReleaseSchedule

//nolint:gosec
