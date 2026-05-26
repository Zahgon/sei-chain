package epoch

import (
	"math/rand"

	seiappparams "github.com/sei-protocol/sei-chain/app/params"
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"
	"github.com/sei-protocol/sei-chain/testutil/sample"
	epochsimulation "github.com/sei-protocol/sei-chain/x/epoch/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = epochsimulation.FindAccount
	_ = seiappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	return
}

// this line is used by starport scaffolding # simapp/module/genesisState

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	_ = "STUB: not implemented"

	// RandomizedParams creates randomized  param changes for the simulator
	return nil
}

func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	_ = "STUB: not implemented"
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
	_ = "STUB: not implemented"

	// WeightedOperations returns the all the gov module operations with their respective weights.
	return
}

func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	_ = "STUB: not implemented"
	return nil
}

// this line is used by starport scaffolding # simapp/module/operation
