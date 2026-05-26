package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// GetHistoricalInfo gets the historical info at a given height
func (k Keeper) GetHistoricalInfo(ctx sdk.Context, height int64) (types.HistoricalInfo, bool) {
	_ = "STUB: not implemented"
	return *new(types.HistoricalInfo), false
}

// SetHistoricalInfo sets the historical info at a given height
func (k Keeper) SetHistoricalInfo(ctx sdk.Context, height int64, hi *types.HistoricalInfo) {
	_ = "STUB: not implemented"
	return
}

// DeleteHistoricalInfo deletes the historical info at a given height
func (k Keeper) DeleteHistoricalInfo(ctx sdk.Context, height int64) {
	_ = "STUB: not implemented"
	return
}

// IterateHistoricalInfo provides an interator over all stored HistoricalInfo
//
//	objects. For each HistoricalInfo object, cb will be called. If the cb returns
//
// true, the iterator will close and stop.
func (k Keeper) IterateHistoricalInfo(ctx sdk.Context, cb func(types.HistoricalInfo) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllHistoricalInfo returns all stored HistoricalInfo objects.
func (k Keeper) GetAllHistoricalInfo(ctx sdk.Context) []types.HistoricalInfo {
	_ = "STUB: not implemented"
	return nil
}

// TrackHistoricalInfo saves the latest historical-info and deletes the oldest
// heights that are below pruning height
func (k Keeper) TrackHistoricalInfo(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// Prune store to ensure we only have parameter-defined historical entries.
// In most cases, this will involve removing a single historical entry.
// In the rare scenario when the historical entries gets reduced to a lower value k'
// from the original value k. k - k' entries must be deleted from the store.
// Since the entries to be deleted are always in a continuous range, we can iterate
// over the historical entries starting from the most recent version to be pruned
// and then return at the first empty entry.

// if there is no need to persist historicalInfo, return

// Create HistoricalInfo struct

// Set latest HistoricalInfo at current height
