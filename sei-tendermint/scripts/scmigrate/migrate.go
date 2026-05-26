// Package scmigrate implements a migration for SeenCommit data
// between 0.34 and 0.35
//
// The Migrate implementation is idempotent and finds all seen commit
// records and deletes all *except* the record corresponding to the
// highest height.
package scmigrate

import (
	"context"

	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type toMigrate struct {
	key    []byte
	commit *types.Commit
}

const prefixSeenCommit = int64(3)

func makeKeyFromPrefix(ids ...int64) []byte { _ = "STUB: not implemented"; return nil }

func makeToMigrate(val []byte) (*types.Commit, error) { _ = "STUB: not implemented"; return nil, nil }

// theoretically we should error for all errors, but
// there's no reason to keep junk data in the
// database, and it makes testing easier.

func sortMigrations(scData []toMigrate) {
	_ = "STUB: not implemented"
	// put this in it's own function just to make it testable
	return
}

func getAllSeenCommits(ctx context.Context, db dbm.DB) ([]toMigrate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renameRecord(db dbm.DB, keep toMigrate) error { _ = "STUB: not implemented"; return nil }

// we already did this conversion

// This record's key has already been converted to the "new" format, we just
// now need to trim off the tail.

func deleteRecords(db dbm.DB, scData []toMigrate) error {
	_ = "STUB: not implemented"
	// delete all the remaining stale values in a single batch
	return nil
}

func Migrate(ctx context.Context, db dbm.DB) error { _ = "STUB: not implemented"; return nil }

// nothing to do

// Sort commits in decreasing order of height.

// Keep and rename the newest seen commit, delete the rest.
// In TM < v0.35 we kept a last-seen commit for each height; in v0.35 we
// retain only the latest.

// Remove any older seen commits. Prior to v0.35, we kept these records for
// all heights, but v0.35 keeps only the latest.
