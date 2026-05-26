package app

import (
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// ExportAppStateAndValidators exports the state of the application for a genesis
// file.
func (app *WasmApp) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string,
) (servertypes.ExportedApp, error) {
	_ = "STUB: not implemented"
	// as if they could withdraw from the start of the next block
	return *new(servertypes.ExportedApp), nil
}

// We export at last height + 1, because that's the height at which
// Tendermint will start InitChain.

// prepare for fresh start at zero height
// NOTE zero height genesis is a temporary feature which will be deprecated
//
//	in favour of export at a block height
func (app *WasmApp) prepForZeroHeightGenesis(ctx sdk.Context, jailAllowedAddrs []string) {
	_ = "STUB: not implemented"
	return
}

// check if there is a allowed address list

/* Just to be safe, assert the invariants on current state. */

/* Handle fee distribution state. */

// withdraw all validator commission

//nolint:errcheck

// withdraw all delegator rewards

//nolint:errcheck

// clear validator slash events

// clear validator historical rewards

// set context height to zero

// reinitialize all validators

// donate any unwithdrawn outstanding reward fraction tokens to the community pool

// reinitialize all delegations

// reset context height

/* Handle staking state. */

// iterate through redelegations, reset creation height

// iterate through unbonding delegations, reset creation height

// Iterate through validators by power descending, reset bond heights, and
// update bond intra-tx counters.

/* Handle slashing state. */

// reset start height on signing infos
