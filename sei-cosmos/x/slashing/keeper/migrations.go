package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "slashing", "keeper")

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	_ = "STUB: not implemented"

	// Migrate2to3 migrates from version 2 to 3.
	return nil
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

// Note that we close the iterator twice. 2 iterators cannot be open at the same time due to mutex on the storage
// This close within defer is a safety net, while the close() after iteration is to close the iterator before opening
// a new one.

// need to use the key to extract validator cons addr
// last 8 bytes are the index
// remove the store prefix + length prefix

//nolint:gosec // index represents a block index, stored as uint64
// load legacy signing info type

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

// use previous height to calculate index offset

// need to turn this into a bool array
