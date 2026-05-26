package simulation

// DONTCOVER

import (
	"math/rand"

	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

const (
	keyCommunityTax        = "communitytax"
	keyBaseProposerReward  = "baseproposerreward"
	keyBonusProposerReward = "bonusproposerreward"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange { _ = "STUB: not implemented"; return nil }
