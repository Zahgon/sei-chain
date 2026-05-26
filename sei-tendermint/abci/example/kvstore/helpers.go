package kvstore

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// RandVals returns a list of cnt validators for initializing
// the application. Note that the keys are deterministically
// derived from the index in the array, while the power is
// random (Change this if not desired)
func RandVals(cnt int) []types.ValidatorUpdate { _ = "STUB: not implemented"; return nil }

// Random value between [0, 2^16 - 1]
// nolint:gosec // G404: Use of weak random number generator
