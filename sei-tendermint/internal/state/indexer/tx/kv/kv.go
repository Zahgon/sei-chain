package kv

import (
	"context"

	dbm "github.com/tendermint/tm-db"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query/syntax"
	indexer "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
)

var _ indexer.TxIndexer = (*TxIndex)(nil)

// TxIndex is the simplest possible indexer
// It is backed by two kv stores:
// 1. txhash - result  (primary key)
// 2. event - txhash   (secondary key)
type TxIndex struct {
	store dbm.DB
}

// NewTxIndex creates new KV indexer.
func NewTxIndex(store dbm.DB) *TxIndex { _ = "STUB: not implemented"; return nil }

// Get gets transaction from the TxIndex storage and returns it or nil if the
// transaction is not found.
func (txi *TxIndex) Get(hash []byte) (*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Index indexes transactions using the given list of events. Each key
// that indexed from the tx's events is a composite of the event type and the
// respective attribute's key delimited by a "." (eg. "account.number").
// Any event with an empty type is not indexed.
func (txi *TxIndex) Index(results []*abci.TxResultV2) error { _ = "STUB: not implemented"; return nil }

// index tx by events

// index by height (always)

// index by hash (always)

func (txi *TxIndex) indexEvents(result *abci.TxResultV2, hash []byte, store dbm.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// only index events with a non-empty type

// index if `index: true` is set

// ensure event does not conflict with a reserved prefix key

// Search performs a search using the given query.
//
// It breaks the query into conditions (like "tx.height > 5"). For each
// condition, it queries the DB index. One special use cases here: (1) if
// "tx.hash" is found, it returns tx result for it (2) for range queries it is
// better for the client to provide both lower and upper bounds, so we are not
// performing a full scan. Results from querying indexes are then intersected
// and returned to the caller, in no particular order.
//
// Search will exit early and return any result fetched so far,
// when a message is received on the context chan.
func (txi *TxIndex) Search(ctx context.Context, q *query.Query) ([]*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get a list of conditions (like "tx.height > 5")

// if there is a hash condition, return the result immediately

// conditions to skip because they're handled before "everything else"

// extract ranges
// if both upper and lower bounds exist, it's better to get them in order not
// no iterate over kvs that are not within range.

// Ignore any remaining conditions if the first condition resulted
// in no matches (assuming implicit AND operand).

// if there is a height condition ("tx.height=3"), extract it

// for all other conditions

// Ignore any remaining conditions if the first condition resulted
// in no matches (assuming implicit AND operand).

// Potentially exit early.

func lookForHash(conditions []syntax.Condition) (hash []byte, ok bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// lookForHeight returns a height if there is an "height=X" condition.
func lookForHeight(conditions []syntax.Condition) (height int64) {
	_ = "STUB: not implemented"
	return 0
}

// match returns all matching txs by hash that meet a given condition and start
// key. An already filtered result (filteredHashes) is provided such that any
// non-intersecting matches are removed.
//
// NOTE: filteredHashes may be empty if no previous condition has matched.
func (txi *TxIndex) match(
	ctx context.Context,
	c syntax.Condition,
	startKeyBz []byte,
	filteredHashes map[string][]byte,
	firstRun bool,
) map[string][]byte {
	_ = "STUB: not implemented"
	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil
}

// Potentially exit early.

// XXX: can't use startKeyBz here because c.Operand is nil
// (e.g. "account.owner/<nil>/" won't match w/ a single row)

// Potentially exit early.

// XXX: startKey does not apply here.
// For example, if startKey = "account.owner/an/" and search query = "account.owner CONTAINS an"
// we can't iterate with prefix "account.owner/an/" because we might miss keys like "account.owner/Ulan/"

// Potentially exit early.

// Potentially exit early.

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHashes that were not found in this
// match (tmpHashes).

// Potentially exit early.

// matchRange returns all matching txs by hash that meet a given queryRange and
// start key. An already filtered result (filteredHashes) is provided such that
// any non-intersecting matches are removed.
//
// NOTE: filteredHashes may be empty if no previous condition has matched.
func (txi *TxIndex) matchRange(
	ctx context.Context,
	qr indexer.QueryRange,
	startKey []byte,
	filteredHashes map[string][]byte,
	firstRun bool,
) map[string][]byte {
	_ = "STUB: not implemented"
	// A previous match was attempted but resulted in no matches, so we return
	// no matches (assuming AND operand).
	return nil
}

// XXX: passing time in a ABCI Events is not yet implemented
// case time.Time:
// 	v := strconv.ParseInt(extractValueFromKey(it.Key()), 10, 64)
// 	if v == r.upperBound {
// 		break
// 	}

// Potentially exit early.

// Either:
//
// 1. Regardless if a previous match was attempted, which may have had
// results, but no match was found for the current condition, then we
// return no matches (assuming AND operand).
//
// 2. A previous match was not attempted, so we return all results.

// Remove/reduce matches in filteredHashes that were not found in this
// match (tmpHashes).

// Potentially exit early.

// ##########################  Keys  #############################
//
// The indexer has two types of kv stores:
// 1. txhash - result  (primary key)
// 2. event - txhash   (secondary key)
//
// The event key can be decomposed into 4 parts.
// 1. A composite key which can be any string.
// Usually something like "tx.height" or "account.owner"
// 2. A value. That corresponds to the key. In the above
// example the value could be "5" or "Ivan"
// 3. The height of the Tx that aligns with the key and value.
// 4. The index of the Tx that aligns with the key and value

// the hash/primary key
func primaryKey(hash []byte) []byte { _ = "STUB: not implemented"; return nil }

// The event/secondary key
func secondaryKey(compositeKey, value string, height int64, index uint32) []byte {
	_ = "STUB: not implemented"
	return nil
}

// parseValueFromKey parses an event key and extracts out the value, returning an error if one arises.
// This will also involve ensuring that the key has the correct format.
// CONTRACT: function doesn't check that the prefix is correct. This should have already been done by the iterator
func parseValueFromKey(key []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func keyFromEvent(compositeKey string, value string, result *abci.TxResultV2) []byte {
	_ = "STUB: not implemented"
	return nil
}

func KeyFromHeight(result *abci.TxResultV2) []byte { _ = "STUB: not implemented"; return nil }

// Prefixes: these represent an initial part of the key and are used by iterators to iterate over a small
// section of the kv store during searches.

func prefixFromCompositeKey(compositeKey string) []byte { _ = "STUB: not implemented"; return nil }

func prefixFromCompositeKeyAndValue(compositeKey, value string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// a small utility function for getting a keys prefix based on a condition and a height
func prefixForCondition(c syntax.Condition, height int64) []byte {
	_ = "STUB: not implemented"
	return nil
}
