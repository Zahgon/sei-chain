package gov

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/keeper"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "gov")

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) { _ = "STUB: not implemented"; return }

// delete inactive proposal from store and its deposits

// called when proposal become inactive

// fetch active proposals whose voting periods have ended (are passed the block time)

// If an expedited proposal fails, we do not want to update
// the deposit at this point since the proposal is converted to regular.
// As a result, the deposits are either deleted or refunded in all casses
// EXCEPT when an expedited proposal fails.

// The proposal handler may execute state mutating logic depending
// on the proposal content. If the handler fails, no state mutation
// is written and the error message is logged.

// The cached context is created with a new EventManager. However, since
// the proposal handler execution was successful, we want to track/keep
// any events emitted, so we re-emit to "merge" the events into the
// original Context's EventManager.

// write state to the underlying multi-store

// The proposal didn't pass after voting period ends

// When expedited proposal fails, it is converted to a regular proposal.
// As a result, the voting period is extended.
// Once the regular voting period expires again, the tally is repeated
// according to the regular proposal rules.

// When regular proposal fails, it is rejected and
// the proposal with that id is done forever.

// when proposal become active
