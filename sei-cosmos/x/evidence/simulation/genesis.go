package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/exported"
)

// Simulation parameter constants
const evidence = "evidence"

// GenEvidences returns an empty slice of evidences.
func GenEvidences(_ *rand.Rand, _ []simtypes.Account) []exported.Evidence {
	_ = "STUB: not implemented"
	return nil
}

// RandomizedGenState generates a random GenesisState for evidence
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
