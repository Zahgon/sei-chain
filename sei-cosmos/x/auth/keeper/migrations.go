package keeper

import (
	"github.com/gogo/protobuf/grpc"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper      AccountKeeper
	queryServer grpc.Server
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper AccountKeeper, queryServer grpc.Server) Migrator {
	_ = "STUB: not implemented"
	return *new(Migrator)
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

func (m Migrator) Migrate2to3(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }
