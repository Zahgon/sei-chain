package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/types"
)

var logger = seilog.NewLogger("cosmos", "x", "evidence", "keeper")

// HandleEquivocationEvidence implements an equivocation evidence handler. Assuming the
// evidence is valid, the validator committing the misbehavior will be slashed,
// jailed and tombstoned. Once tombstoned, the validator will not be able to
// recover. Note, the evidence contains the block time and height at the time of
// the equivocation.
//
// The evidence is considered invalid if:
// - the evidence is too old
// - the validator is unbonded or does not exist
// - the signing info does not exist (will panic)
// - is already tombstoned
//
// TODO: Some of the invalid constraints listed above may need to be reconsidered
// in the case of a lunatic attack.
func (k Keeper) HandleEquivocationEvidence(ctx sdk.Context, evidence *types.Equivocation) {
	_ = "STUB: not implemented"
	return
}

// Ignore evidence that cannot be handled.
//
// NOTE: We used to panic with:
// `panic(fmt.Sprintf("Validator consensus-address %v not found", consAddr))`,
// but this couples the expectations of the app to both Tendermint and
// the simulator.  Both are expected to provide the full range of
// allowable but none of the disallowed evidence types.  Instead of
// getting this coordination right, it is easier to relax the
// constraints and ignore evidence that cannot be handled.

// calculate the age of the evidence

// Reject evidence if the double-sign is too old. Evidence is considered stale
// if the difference in time and number of blocks is greater than the allowed
// parameters defined.

// Defensive: Simulation doesn't take unbonding periods into account, and
// Tendermint might break this assumption at some point.

// ignore if the validator is already tombstoned

// We need to retrieve the stake distribution which signed the block, so we
// subtract ValidatorUpdateDelay from the evidence height.
// Note, that this *can* result in a negative "distributionHeight", up to
// -ValidatorUpdateDelay, i.e. at the end of the
// pre-genesis block (none) = at the beginning of the genesis block.
// That's fine since this is just used to filter unbonding delegations & redelegations.

// Slash validator. The `power` is the int64 power of the validator as provided
// to/by Tendermint. This value is validator.Tokens as sent to Tendermint via
// ABCI, and now received as evidence. The fraction is passed in to separately
// to slash unbonding and rebonding delegations.

// Jail the validator if not already jailed. This will begin unbonding the
// validator if not already unbonding (tombstoned).
