package oracle

import (
	"github.com/sei-protocol/sei-chain/x/oracle/keeper"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func MidBlocker(ctx sdk.Context, k keeper.Keeper) { _ = "STUB: not implemented"; return }

// TODO(PLT-336): remove once oracle_mid_blocker_duration_seconds verified

// Build claim map over all validators in active set

// Exclude not bonded validator

// Organize votes to ballot by denom
// NOTE: **Filter out inactive or jailed validators**
// NOTE: **Make abstain votes to have zero vote power**

// belowThresholdVoteMap has assets that failed to meet threshold

// make voteMap of Reference denom to calculate cross exchange rates

// Iterate through ballots and update exchange rates; drop if not enough votes have been achieved.

// Convert ballot to cross exchange rates

// Get weighted median of cross exchange rates

// if exchange rate is somehow 0, exclude it from ballot?

// skip this denom

// Transform into the original form base/quote

// Set the exchange rate, emit ABCI event

// TODO(PLT-336): remove once oracle_price_update_total verified

// in this case, all assets would be in the belowThresholdVoteMap

// perform tally for below threshold assets to calculate total win count

//---------------------------
// Do miss counting & slashing

// we require validator to have submitted in-range data
// for all assets to not be counted as a miss

// Increase miss counter

// Clear the ballot

// Update vote targets

func EndBlocker(ctx sdk.Context, k keeper.Keeper) { _ = "STUB: not implemented"; return }

// TODO(PLT-336): remove once oracle_end_blocker_duration_seconds verified

// Do slash who did miss voting over threshold and
// reset miss counters of all validators at the last block of slash window

// Compare vote targets and actives and remove excess feeds
