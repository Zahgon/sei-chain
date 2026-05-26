package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
)

// RandomGenesisDefaultSendParam computes randomized allow all send transfers param for the bank module
func RandomGenesisDefaultSendParam(r *rand.Rand) bool {
	_ = "STUB: not implemented"
	// 90% chance of transfers being enable or P(a) = 0.9 for success
	return false
}

// RandomGenesisSendParams randomized Parameters for the bank module
func RandomGenesisSendParams(r *rand.Rand) types.SendEnabledParams {
	_ = "STUB: not implemented"
	return *new(types.SendEnabledParams)
}

// 90% chance of transfers being DefaultSendEnabled=true or P(a) = 0.9 for success
// 50% of the time add an additional denom specific record (P(b) = 0.475 = 0.5 * 0.95)

// set send enabled 95% of the time

// overall probability of enabled for bond denom is 94.75% (P(a)+P(b) - P(a)*P(b))

// RandomGenesisBalances returns a slice of account balances. Each account has
// a balance of simState.InitialStake for sdk.DefaultBondDenom.
func RandomGenesisBalances(simState *module.SimulationState) []types.Balance {
	_ = "STUB: not implemented"
	return nil
}

// RandomizedGenState generates a random GenesisState for bank
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
