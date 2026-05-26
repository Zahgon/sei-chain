package simulation

import (
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

// FindAccount find a specific address from an account list
func FindAccount(accs []simtypes.Account, address string) (simtypes.Account, bool) {
	_ = "STUB: not implemented"
	return *new(simtypes.Account), false
}
