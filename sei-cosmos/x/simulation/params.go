package simulation

import (
	"encoding/json"
	"math/rand"

	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

const (
	// Minimum time per block
	minTimePerBlock int64 = 10000 / 2

	// Maximum time per block
	maxTimePerBlock int64 = 10000
)

// TODO: explain transitional matrix usage
var (
	// Currently there are 3 different liveness types,
	// fully online, spotty connection, offline.
	defaultLivenessTransitionMatrix, _ = CreateTransitionMatrix([][]int{
		{90, 20, 1},
		{10, 50, 5},
		{0, 10, 1000},
	})

	// 3 states: rand in range [0, 4*provided blocksize],
	// rand in range [0, 2 * provided blocksize], 0
	defaultBlockSizeTransitionMatrix, _ = CreateTransitionMatrix([][]int{
		{85, 5, 0},
		{15, 92, 1},
		{0, 3, 99},
	})
)

// Params define the parameters necessary for running the simulations
type Params struct {
	pastEvidenceFraction      float64
	numKeys                   int
	evidenceFraction          float64
	initialLivenessWeightings []int
	livenessTransitionMatrix  simulation.TransitionMatrix
	blockSizeTransitionMatrix simulation.TransitionMatrix
}

func (p Params) PastEvidenceFraction() float64 { _ = "STUB: not implemented"; return 0 }

func (p Params) NumKeys() int { _ = "STUB: not implemented"; return 0 }

func (p Params) EvidenceFraction() float64 { _ = "STUB: not implemented"; return 0 }

func (p Params) InitialLivenessWeightings() []int { _ = "STUB: not implemented"; return nil }

func (p Params) LivenessTransitionMatrix() simulation.TransitionMatrix {
	_ = "STUB: not implemented"
	return *new(simulation.TransitionMatrix)
}

func (p Params) BlockSizeTransitionMatrix() simulation.TransitionMatrix {
	_ = "STUB: not implemented"
	return *new(simulation.TransitionMatrix)
}

// RandomParams returns random simulation parameters
func RandomParams(r *rand.Rand) Params { _ = "STUB: not implemented"; return *new(Params) }

// number of accounts created for the simulation

// Param change proposals

// ParamChange defines the object used for simulating parameter change proposals
type ParamChange struct {
	subspace string
	key      string
	simValue simulation.SimValFn
}

func (spc ParamChange) Subspace() string { _ = "STUB: not implemented"; return "" }

func (spc ParamChange) Key() string { _ = "STUB: not implemented"; return "" }

func (spc ParamChange) SimValue() simulation.SimValFn {
	_ = "STUB: not implemented"
	return *

	// NewSimParamChange creates a new ParamChange instance
	new(simulation.SimValFn)
}

func NewSimParamChange(subspace, key string, simVal simulation.SimValFn) simulation.ParamChange {
	_ = "STUB: not implemented"
	return *new(simulation.ParamChange)
}

// ComposedKey creates a new composed key for the param change proposal
func (spc ParamChange) ComposedKey() string { _ = "STUB: not implemented"; return "" }

// Proposal Contents

// WeightedProposalContent defines a common struct for proposal contents defined by
// external modules (i.e outside gov)
type WeightedProposalContent struct {
	appParamsKey       string                        // key used to retrieve the value of the weight from the simulation application params
	defaultWeight      int                           // default weight
	contentSimulatorFn simulation.ContentSimulatorFn // content simulator function
}

func NewWeightedProposalContent(appParamsKey string, defaultWeight int, contentSimulatorFn simulation.ContentSimulatorFn) simulation.WeightedProposalContent {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedProposalContent)
}

func (w WeightedProposalContent) AppParamsKey() string { _ = "STUB: not implemented"; return "" }

func (w WeightedProposalContent) DefaultWeight() int { _ = "STUB: not implemented"; return 0 }

func (w WeightedProposalContent) ContentSimulatorFn() simulation.ContentSimulatorFn {
	_ = "STUB: not implemented"
	return *new(simulation.ContentSimulatorFn)
}

// Param change proposals

// randomConsensusParams returns random simulation consensus parameters, it extracts the Evidence from the Staking genesis state.
func randomConsensusParams(r *rand.Rand, appState json.RawMessage, cdc codec.JSONCodec) *tmproto.ConsensusParams {
	_ = "STUB: not implemented"
	return nil
}
