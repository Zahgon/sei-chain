package simulation

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

// TransitionMatrix is _almost_ a left stochastic matrix.  It is technically
// not one due to not normalizing the column values.  In the future, if we want
// to find the steady state distribution, it will be quite easy to normalize
// these values to get a stochastic matrix.  Floats aren't currently used as
// the default due to non-determinism across architectures
type TransitionMatrix struct {
	weights [][]int
	// total in each column
	totals []int
	n      int
}

// CreateTransitionMatrix creates a transition matrix from the provided weights.
// TODO: Provide example usage
func CreateTransitionMatrix(weights [][]int) (simulation.TransitionMatrix, error) {
	_ = "STUB: not implemented"
	return *new(simulation.TransitionMatrix), nil
}

// NextState returns the next state randomly chosen using r, and the weightings
// provided in the transition matrix.
func (t TransitionMatrix) NextState(r *rand.Rand, i int) int { _ = "STUB: not implemented"; return 0 }

// This line should never get executed

// GetMemberOfInitialState takes an initial array of weights, of size n.
// It returns a weighted random number in [0,n).
func GetMemberOfInitialState(r *rand.Rand, weights []int) int { _ = "STUB: not implemented"; return 0 }

// This line should never get executed
