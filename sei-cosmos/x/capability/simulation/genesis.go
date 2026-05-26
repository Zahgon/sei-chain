package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

// Simulation parameter constants
const index = "index"

// GenIndex returns a random global index between 1-1000
func GenIndex(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // Int63n returns non-negative values

// RandomizedGenState generates a random GenesisState for capability
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
