package simulation

import (
	"math/rand"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant"
)

// genFeeGrants returns a slice of randomly generated allowances.
func genFeeGrants(r *rand.Rand, accounts []simtypes.Account) []feegrant.Grant {
	_ = "STUB: not implemented"
	return nil
}

func generateRandomAllowances(granter, grantee sdk.AccAddress, r *rand.Rand) feegrant.Grant {
	_ = "STUB: not implemented"
	return *new(feegrant.Grant)
}

// RandomizedGenState generates a random GenesisState for feegrant
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
