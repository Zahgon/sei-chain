package pubsub

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// An item to be published to subscribers.
type item struct {
	Data   types.EventData
	Events []abci.Event
}

// A subInfo value records a single subscription.
type subInfo struct {
	clientID string        // chosen by the client
	query    *query.Query  // chosen by the client
	subID    string        // assigned at registration
	sub      *Subscription // receives published events
}

// A subInfoSet is an unordered set of subscription info records.
type subInfoSet map[*subInfo]struct{}

func (s subInfoSet) contains(si *subInfo) bool { _ = "STUB: not implemented"; return false }
func (s subInfoSet) add(si *subInfo)           { _ = "STUB: not implemented"; return }
func (s subInfoSet) remove(si *subInfo)        { _ = "STUB: not implemented"; return }

// withQuery returns the subset of s whose query string matches qs.
func (s subInfoSet) withQuery(qs string) subInfoSet {
	_ = "STUB: not implemented"
	return *new(subInfoSet)
}

// A subIndex is an indexed collection of subscription info records.
// The index is not safe for concurrent use without external synchronization.
type subIndex struct {
	all      subInfoSet            // all subscriptions
	byClient map[string]subInfoSet // per-client subscriptions
	byQuery  map[string]subInfoSet // per-query subscriptions

	// TODO(creachadair): We allow indexing by query to support existing use by
	// the RPC service methods for event streaming. Fix up those methods not to
	// require this, and then remove indexing by query.
}

// newSubIndex constructs a new, empty subscription index.
func newSubIndex() *subIndex { _ = "STUB: not implemented"; return nil }

// findClients returns the set of subscriptions for the given client ID, or nil.
func (idx *subIndex) findClientID(id string) subInfoSet {
	_ = "STUB: not implemented"
	return *

	// findQuery returns the set of subscriptions on the given query string, or nil.
	new(subInfoSet)
}

func (idx *subIndex) findQuery(qs string) subInfoSet {
	_ = "STUB: not implemented"
	return *

	// contains reports whether idx contains any subscription matching the given
	// client ID and query pair.
	new(subInfoSet)
}

func (idx *subIndex) contains(clientID, query string) bool { _ = "STUB: not implemented"; return false }

// add adds si to the index, replacing any previous entry with the same terms.
// It is the caller's responsibility to check for duplicates before adding.
// See also the contains method.
func (idx *subIndex) add(si *subInfo) { _ = "STUB: not implemented"; return }

// removeAll removes all the elements of s from the index.
func (idx *subIndex) removeAll(s subInfoSet) { _ = "STUB: not implemented"; return }
