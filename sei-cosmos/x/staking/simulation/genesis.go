package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

// Simulation parameter constants
const (
	unbondingTime     = "unbonding_time"
	maxValidators     = "max_validators"
	historicalEntries = "historical_entries"
)

// genUnbondingTime returns randomized UnbondingTime
func genUnbondingTime(r *rand.Rand) (ubdTime time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// genMaxValidators returns randomized MaxValidators
func genMaxValidators(r *rand.Rand) (maxValidators uint32) { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // Intn(250)+1 always returns a value in [1, 250], fits in uint32

// getHistEntries returns randomized HistoricalEntries between 0-100.
func getHistEntries(r *rand.Rand) uint32 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // DefaultHistoricalEntries is a small constant, result always fits in uint32

// RandomizedGenState generates a random GenesisState for staking
func RandomizedGenState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	// params
	return
}

// NOTE: the slashing module need to be defined after the staking module on the
// NewSimulationManager constructor for this to work

// validators & delegations
