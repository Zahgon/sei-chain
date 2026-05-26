package simulation

// DONTCOVER

import (
	"math/rand"

	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

const (
	keyVotingParams          = "votingparams"
	keyDepositParams         = "depositparams"
	keyTallyParams           = "tallyparams"
	subkeyQuorum             = "quorum"
	subkeyThreshold          = "threshold"
	subkeyExpeditedThreshold = "expedited_threshold"
	subkeyVeto               = "veto"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange { _ = "STUB: not implemented"; return nil }
