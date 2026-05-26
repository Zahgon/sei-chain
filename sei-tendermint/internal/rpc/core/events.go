package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventlog"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventlog/cursor"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

const (
	// Buffer on the Tendermint (server) side to allow some slowness in clients.
	subBufferSize = 100

	// maxQueryLength is the maximum length of a query string that will be
	// accepted. This is just a safety check to avoid outlandish queries.
	maxQueryLength = 512
)

// Subscribe for events via WebSocket.
// More: https://docs.tendermint.com/master/rpc/#/Websocket/subscribe
func (env *Environment) Subscribe(ctx context.Context, req *coretypes.RequestSubscribe) (*coretypes.ResultSubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture the current ID, since it can change in the future.

// The subscription was removed by the client.

// The subscription was terminated by the publisher.

// We have a message to deliver to the client.

// Unsubscribe from events via WebSocket.
// More: https://docs.tendermint.com/master/rpc/#/Websocket/unsubscribe
func (env *Environment) Unsubscribe(ctx context.Context, req *coretypes.RequestUnsubscribe) (*coretypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnsubscribeAll from all events via WebSocket.
// More: https://docs.tendermint.com/master/rpc/#/Websocket/unsubscribe_all
func (env *Environment) UnsubscribeAll(ctx context.Context) (*coretypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Events applies a query to the event log. If an event log is not enabled,
// Events reports an error. Otherwise, it filters the current contents of the
// log to return matching events.
//
// Events returns up to maxItems of the newest eligible event items. An item is
// eligible if it is older than before (or before is zero), it is newer than
// after (or after is zero), and its data matches the filter. A nil filter
// matches all event data.
//
// If before is zero and no eligible event items are available, Events waits
// for up to waitTime for a matching item to become available. The wait is
// terminated early if ctx ends.
//
// If maxItems ≤ 0, a default positive number of events is chosen. The values
// of maxItems and waitTime may be capped to sensible internal maxima without
// reporting an error to the caller.
func (env *Environment) Events(ctx context.Context, req *coretypes.RequestEvents) (*coretypes.ResultEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse and validate parameters.

// N.B. We accept up to one item more than requested, so we can tell how
// to set the "more" flag in the response.

// Long poll. The loop here is because new items may not match the query,
// and we want to keep waiting until we have relevant results (or time out).

// Don't report a timeout as a request failure.

// Quick poll, return only what is already available.

func cursorString(c cursor.Cursor) string { _ = "STUB: not implemented"; return "" }

func cursorInRange(c, before, after cursor.Cursor) bool { _ = "STUB: not implemented"; return false }

func marshalItems(items []*eventlog.Item) ([]*coretypes.EventItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
