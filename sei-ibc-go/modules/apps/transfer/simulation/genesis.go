package simulation

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

// Simulation parameter constants
const port = "port_id"

// RadomEnabled randomized send or receive enabled param with 75% prob of being true.
func RadomEnabled(r *rand.Rand) bool { _ = "STUB: not implemented"; return false }

// RandomizedGenState generates a random GenesisState for transfer.
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
