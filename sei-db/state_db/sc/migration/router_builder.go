package migration

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/memiavl"
)

// Builds a router for the given migration write mode. A router is responsible for splitting
// reads/writes between the memiavl and flatkv backends.
func BuildRouter(
	ctx context.Context,
	writeMode config.WriteMode,
	memIAVL *memiavl.CommitStore,
	flatKV flatkv.Store,
	// If this router will be doing data migration, this is the number of keys to migrate in each batch.
	migrationBatchSize int,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

/* Data flow: MemiavlOnly (0)

                       ┌─────────────┐                                  ┌─────────┐
──all-modules────────▶ │ passthrough │ ──────────all-modules──────────▶ │ memIAVL │
                       └─────────────┘                                  └─────────┘
*/

// Build a router for handling write mode MemiavlOnly. Operates on a schema at migration version 0.
func buildMemiavlOnlyRouter(
	memIAVL *memiavl.CommitStore,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

/* Data flow: MigrateEVM (0 -> 1)

                       ┌──────────────┐                                  ┌─────────┐
──all-modules────────▶ │ moduleRouter │ ──everything-except-evm/───────▶ │ memIAVL │
                       └──────────────┘                                  └─────────┘
                              │                                               ▲
                             evm/                                             │
                              │       ┌──────un-migrated-keys─────────────────┘
                              │       │
                              ▼       │
                       ┌──────────────────┐                              ┌────────┐
                       │ migrationManager │ ────────migrated-keys──────▶ │ flatKV │
                       └──────────────────┘                              └────────┘
*/

// Build a router for handling write mode MigrateEVM. Migrates from version 0 to version 1.
func buildMigrateEVMRouter(
	ctx context.Context,
	memIAVL *memiavl.CommitStore,
	flatKV flatkv.Store,
	migrationBatchSize int,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

// Manages migration and routing for keys in the evm/ module.

/* Data flow: EVMMigrated (1)

                       ┌──────────────┐                                  ┌─────────┐
──all-modules────────▶ │ moduleRouter │ ──everything-except-evm/───────▶ │ memIAVL │
                       └──────────────┘                                  └─────────┘
                              │
                              │
                              │
                              │
                              │
                              │                                          ┌────────┐
                              └────────────evm/────────────────────────▶ │ flatKV │
                                                                         └────────┘
*/

// Build a router for handling write mode EVMMigrated. Operates on a schema at migration version 1.
func buildEVMMigratedRouter(
	memIAVL *memiavl.CommitStore,
	flatKV flatkv.Store,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

/* Data flow: MigrateAllButBank (1 -> 2)

                       ┌──────────────┐                                  ┌─────────┐
──all-modules────────▶ │ moduleRouter │ ──────────────bank/────────────▶ │ memIAVL │
                       └──────────────┘                                  └─────────┘
                        │     │                                               ▲
                        │   all but                                           │
                        │   bank/ and evm/    ┌──────un-migrated-keys─────────┘
                        │     │               │
                        │     ▼               │
                        │   ┌──────────────────┐                         ┌────────┐
                        │   │ migrationManager │ ───migrated-keys──────▶ │ flatKV │
                        │   └──────────────────┘                         └────────┘
                        │                                                    ▲
                        │                                                    │
                        └────────────────────────────evm/────────────────────┘
*/

// Build a router for handling write mode MigrateAllButBank. Migrates from version 1 to version 2.
func buildMigrateAllButBankRouter(
	ctx context.Context,
	memIAVL *memiavl.CommitStore,
	flatKV flatkv.Store,
	migrationBatchSize int,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

// Manages migration and routing for all keys except evm/ (already migrated) and bank/ (not migrating yet)

/* Data flow: AllMigratedButBank (2)

                       ┌──────────────┐                                  ┌─────────┐
──all-modules────────▶ │ moduleRouter │ ───bank/───────────────────────▶ │ memIAVL │
                       └──────────────┘                                  └─────────┘
                              │
                              │
                              │
                              │
                              │
                              │                                          ┌────────┐
                              └────────────all─but─bank/───────────────▶ │ flatKV │
                                                                         └────────┘
*/

// Build a router for handling write mode AllMigratedButBank. Operates on a schema at migration version 2.
func buildAllMigratedButBankRouter(
	memIAVL *memiavl.CommitStore,
	flatKV flatkv.Store,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

/* Data flow: MigrateBank (2 -> 3)

                       ┌──────────────┐                                  ┌─────────┐
──all-modules────────▶ │ moduleRouter │                                  │ memIAVL │
                       └──────────────┘                                  └─────────┘
                        │     │                                               ▲
                        │   bank/             ┌──────un-migrated-keys─────────┘
                        │     │               │
                        │     ▼               │
                        │   ┌──────────────────┐                         ┌────────┐
                        │   │ migrationManager │ ───migrated-keys──────▶ │ flatKV │
                        │   └──────────────────┘                         └────────┘
                        │                                                    ▲
                        │                                                    │
                        └───────────────────all─but─bank/────────────────────┘
*/

// Build a router for handling write mode MigrateBank. Migrates from version 2 to version 3.
func buildMigrateBankRouter(
	ctx context.Context,
	memIAVL *memiavl.CommitStore,
	flatKV flatkv.Store,
	migrationBatchSize int,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

// Manages migration and routing for keys in the bank/ module (the
// final module remaining in memiavl; every other module already
// lives in flatkv from prior migrations).

/* Data flow: FlatKVOnly (3)

                       ┌─────────────┐                                  ┌────────┐
──all-modules────────▶ │ passthrough │ ──────────all-modules──────────▶ │ flatKV │
                       └─────────────┘                                  └────────┘
*/

// Build a router for handling write mode FlatKVOnly. Operates on a schema at migration version 3.
func buildFlatKVOnlyRouter(
	flatKV flatkv.Store,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

// iteration not supported by flatkv
// proof building not supported by flatkv

/* Data flow: dual write (test only)

                       ┌──────────────┐                                  ┌─────────┐
──all-modules────────▶ │ moduleRouter │ ──everything-except-evm/───────▶ │ memIAVL │
                       └──────────────┘                                  └─────────┘
                              │                                               ▲
                             evm/                                             │
                              │       ┌──────evm/─reads-and-writes────────────┘
                              │       │
                              ▼       │
                       ┌───────────────────┐                             ┌────────┐
                       │ dual write router │ ───────evm/-writes────────▶ │ flatKV │
                       └───────────────────┘                             └────────┘
*/

// Build a test-only dual-write router.
//
// CRITICAL: this is a test-only router and should never be deployed to production machines.
func buildTestOnlyDualWriteRouter(
	memIAVL *memiavl.CommitStore,
	flatKV flatkv.Store,
) (Router, error) {
	_ = "STUB: not implemented"
	return *new(Router), nil
}

// Sends evm/ traffic to both memIAVL and flatKV.
// Note that a TestOnlyDualWriteRouter ignores module names; it's only job is to duplicate traffic.
// The routes given to the dual write router do not specify modules for this reason.

// Build a function capable of reading data from memiavl.
func buildMemIAVLReader(memIAVL *memiavl.CommitStore) DBReader {
	_ = "STUB: not implemented"
	return *new(DBReader)
}

// Build a function capable of writing data to memiavl.
func buildMemIAVLWriter(memIAVL *memiavl.CommitStore) DBWriter {
	_ = "STUB: not implemented"
	return *new(DBWriter)
}

// Build a function capable of getting an iterator over a range of keys in a memiavl store.
func buildMemIAVLIteratorBuilder(memIAVL *memiavl.CommitStore) DBIteratorBuilder {
	_ = "STUB: not implemented"
	return *new(DBIteratorBuilder)
}

// Build a function capable of building a proof of the value for a key in a memiavl store.
func buildMemIAVLProofBuilder(memIAVL *memiavl.CommitStore) DBProofBuilder {
	_ = "STUB: not implemented"
	return *new(DBProofBuilder)
}

// Build a function capable of reading data from flatkv.
func buildFlatKVReader(flatKV flatkv.Store) DBReader {
	_ = "STUB: not implemented"
	return *new(DBReader)
}

// Build a function capable of writing data to flatkv.
func buildFlatKVWriter(flatKV flatkv.Store) DBWriter {
	_ = "STUB: not implemented"
	return *new(DBWriter)
}

// Build a route to a memiavl store for the given module names.
func routeToMemIAVL(memIAVL *memiavl.CommitStore, moduleNames ...string) (*Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build a route to a flatkv store for the given module names.
func routeToFlatKV(flatKV flatkv.Store, moduleNames ...string) (*Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// iteration not supported
// proof building not supported
