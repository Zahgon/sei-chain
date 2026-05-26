package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("x", "oracle", "keeper")

// SlashAndResetCounters do slash any operator who over criteria & clear all operators miss counter to zero
func (k Keeper) SlashAndResetCounters(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// Calculate valid vote rate; (totalVotes - (MissCounter + AbstainCounter))/totalVotes
// this accounts for changes in vote period within a window, and will take the overall success rate
// as opposed to the one expected based on the number of vote period expected based on the ending slash window or vote period

//nolint:gosec
//nolint:gosec

// Penalize the validator whose the valid vote rate is smaller than min threshold

// TODO(PLT-336): remove once oracle_validator_slashed_total verified
