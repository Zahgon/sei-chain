package antedecorators

import (
	"math"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	OraclePriority       = math.MaxInt64 - 100
	EVMAssociatePriority = math.MaxInt64 - 101
	// This is the max priority a non oracle or associate tx can take
	MaxPriority = math.MaxInt64 - 1000
)

type PriorityDecorator struct{}

func NewPriorityDecorator() PriorityDecorator {
	_ = "STUB: not implemented"
	return *new(PriorityDecorator)
}

func intMin(a, b int64) int64 { _ = "STUB: not implemented"; return 0 }

// Assigns higher priority to certain types of transactions including oracle
func (pd PriorityDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	// Cap priority
	// Use higher priorities for tiers including oracle tx's
	return *new(sdk.Context), nil
}

func isOracleTx(tx sdk.Tx) bool { _ = "STUB: not implemented"; return false }

// empty TX isn't oracle
