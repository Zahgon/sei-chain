package oracle

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

// SpammingPreventionDecorator will check if the transaction's gas is smaller than
// configured hard cap
type SpammingPreventionDecorator struct {
	oracleKeeper keeper.Keeper
}

// NewSpammingPreventionDecorator returns new spamming prevention decorator instance
func NewSpammingPreventionDecorator(oracleKeeper keeper.Keeper) SpammingPreventionDecorator {
	_ = "STUB: not implemented"
	return *new(SpammingPreventionDecorator)
}

// AnteHandle handles msg tax fee checking
func (spd SpammingPreventionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// CheckOracleSpamming check whether the msgs are spamming purpose or not
func (spd SpammingPreventionDecorator) CheckOracleSpamming(ctx sdk.Context, msgs []sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

type VoteAloneDecorator struct{}

func NewOracleVoteAloneDecorator() VoteAloneDecorator {
	_ = "STUB: not implemented"
	return *new(VoteAloneDecorator)
}

// AnteHandle handles msg tax fee checking
func (VoteAloneDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}
