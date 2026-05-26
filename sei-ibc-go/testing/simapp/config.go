package simapp

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

// List of available flags for the simulator
var (
	FlagGenesisFileValue        string
	FlagParamsFileValue         string
	FlagExportParamsPathValue   string
	FlagExportParamsHeightValue int
	FlagExportStatePathValue    string
	FlagExportStatsPathValue    string
	FlagSeedValue               int64
	FlagInitialBlockHeightValue int
	FlagNumBlocksValue          int
	FlagBlockSizeValue          int
	FlagLeanValue               bool
	FlagCommitValue             bool
	FlagOnOperationValue        bool // TODO: Remove in favor of binary search for invariant violation
	FlagAllInvariantsValue      bool

	FlagEnabledValue     bool
	FlagVerboseValue     bool
	FlagPeriodValue      uint
	FlagGenesisTimeValue int64
)

// GetSimulatorFlags gets the values of all the available simulation flags
func GetSimulatorFlags() {
	_ = "STUB: not implemented"
	// config fields
	return
}

// simulation flags

// NewConfigFromFlags creates a simulation from the retrieved values of the flags.
func NewConfigFromFlags() simulation.Config {
	_ = "STUB: not implemented"
	return *new(simulation.Config)
}
