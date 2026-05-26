package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

func min(a int, b int) int { _ = "STUB: not implemented"; return 0 }

// SimulateParamChangeProposalContent returns random parameter change content.
// It will generate a ParameterChangeProposal object with anywhere between 1 and
// the total amount of defined parameters changes, all of which have random valid values.
func SimulateParamChangeProposalContent(paramChangePool []simulation.ParamChange) simulation.ContentSimulatorFn {
	_ = "STUB: not implemented"

	// Bound the maximum number of simultaneous parameter changes
	return *new(simulation.ContentSimulatorFn)
}

// perm here takes at most len(paramChangePool) calls to random

// add a new distinct parameter to the set of changes

// randomly generate some expedited proposal

// title
// description
// set of changes
// is expedited proposal or not
