package kv

import (
	"context"

	dbm "github.com/tendermint/tm-db"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query/syntax"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var _ indexer.BlockIndexer = (*BlockerIndexer)(nil)

// BlockerIndexer implements a block indexer, indexing FinalizeBlock
// events with an underlying KV store. Block events are indexed by their height,
// such that matching search criteria returns the respective block height(s).
type BlockerIndexer struct {
	store dbm.DB
}

func New(store dbm.DB) *BlockerIndexer { _ = "STUB: not implemented"; return nil }

// Has returns true if the given height has been indexed. An error is returned
// upon database query failure.
func (idx *BlockerIndexer) Has(height int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Index indexes FinalizeBlock events for a given block by its height.
// The following is indexed:
//
// primary key: encode(block.height | height) => encode(height)
// FinalizeBlock events: encode(eventType.eventAttr|eventValue|height|finalize_block) => encode(height)
func (idx *BlockerIndexer) Index(bh types.EventDataNewBlockHeader) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. index by height

// 2. index FinalizeBlock events

// Search performs a query for block heights that match a given FinalizeBlock
// The given query can match against zero or more block heights. In the case
// of height queries, i.e. block.height=H, if the height is indexed, that height
// alone will be returned. An error and nil slice is returned. Otherwise, a
// non-nil slice and nil error is returned.
func (idx *BlockerIndexer) Search(ctx context.Context, q *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If there is an exact height query, return the result immediately
// (if it exists).

// conditions to skip because they're handled before "everything else"

// Extract ranges. If both upper and lower bounds exist, it's better to get
// them in order as to not iterate over kvs that are not within range.

// Ignore any remaining conditions if the first condition resulted in no
// matches (assuming implicit AND operand).

// for all other conditions

// Ignore any remaining conditions if the first condition resulted in no
// matches (assuming implicit AND operand).

// fetch matching heights

// matchRange returns all matching block heights that match a given QueryRange
// and start key. An already filtered result (filteredHeights) is provided such
// that any non-intersecting matches are removed.
//
// NOTE: The provided filteredHeights may be empty if no previous condition has
// matched.
func (idx *BlockerIndexer) matchRange(
	ctx context.Context,
	qr indexer.QueryRange,
	startKey []byte,
	filteredHeights map[string][]byte,
	firstRun bool,
) (map[string][]byte, error) {
	_ = "STUB: not implemented"

	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil, nil
}

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHashes that were not found in this
// match (tmpHashes).

// match returns all matching heights that meet a given query condition and start
// key. An already filtered result (filteredHeights) is provided such that any
// non-intersecting matches are removed.
//
// NOTE: The provided filteredHeights may be empty if no previous condition has
// matched.
func (idx *BlockerIndexer) match(
	ctx context.Context,
	c syntax.Condition,
	startKeyBz []byte,
	filteredHeights map[string][]byte,
	firstRun bool,
) (map[string][]byte, error) {
	_ = "STUB: not implemented"

	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil, nil
}

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHeights that were not found in this
// match (tmpHeights).

func (idx *BlockerIndexer) indexEvents(batch dbm.Batch, events []abci.Event, typ string, height int64) error {
	_ = "STUB: not implemented"
	return nil
}

// only index events with a non-empty type

// index iff the event specified index:true and it's not a reserved event
