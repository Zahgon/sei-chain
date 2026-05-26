package simulation

import (
	"math/rand"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

// genGrant returns a slice of authorization grants.
func genGrant(r *rand.Rand, accounts []simtypes.Account) []authz.GrantAuthorization {
	_ = "STUB: not implemented"
	return nil
}

func generateRandomGrant(r *rand.Rand) *codectypes.Any { _ = "STUB: not implemented"; return nil }

func newAnyAuthorization(a authz.Authorization) *codectypes.Any {
	_ = "STUB: not implemented"
	return nil
}

// RandomizedGenState generates a random GenesisState for authz.
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
