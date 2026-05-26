package simapp

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	dbm "github.com/tendermint/tm-db"
)

// SetupSimulation creates the config, db (levelDB), temporary directory and logger for
// the simulation tests. If `FlagEnabledValue` is false it skips the current test.
// Returns error on an invalid db instantiation or temp dir creation.
func SetupSimulation(dirPrefix, dbName string) (simtypes.Config, dbm.DB, string, bool, error) {
	_ = "STUB: not implemented"
	return *new(simtypes.Config), *new(dbm.DB), "", false, nil
}

// CheckExportSimulation exports the app state and simulation parameters to JSON
// if the export paths are defined.
func CheckExportSimulation(
	app App, config simtypes.Config, params simtypes.Params,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintStats prints the corresponding statistics from the app DB.
func PrintStats(db dbm.DB) { _ = "STUB: not implemented"; return }

// GetSimulationLog unmarshals the KVPair's Value to the corresponding type based on the
// each's module store key and the prefix bytes of the KVPair's key.
func GetSimulationLog(storeName string, sdr sdk.StoreDecoderRegistry, kvAs, kvBs []kv.Pair) (log string) {
	_ = "STUB: not implemented"
	return ""
}

// skip if the value doesn't have any bytes
