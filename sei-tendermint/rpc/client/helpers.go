package client

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// Waiter is informed of current height, decided whether to quit early
type Waiter func(delta int64) (abort error)

// DefaultWaitStrategy is the standard backoff algorithm,
// but you can plug in another one
func DefaultWaitStrategy(delta int64) (abort error) { _ = "STUB: not implemented"; return nil }

// estimate of wait time....
// wait half a second for the next block (in progress)
// plus one second for every full block

// Wait for height will poll status at reasonable intervals until
// the block at the given height is available.
//
// If waiter is nil, we use DefaultWaitStrategy, but you can also
// provide your own implementation
func WaitForHeight(ctx context.Context, c StatusClient, h int64, waiter Waiter) error {
	_ = "STUB: not implemented"
	return nil
}

// wait for the time, or abort early

// WaitForOneEvent waits for the first event matching the given query on c, or
// until ctx ends. It reports an error if ctx ends before a matching event is
// received.
func WaitForOneEvent(ctx context.Context, c EventsClient, query string) (types.EventData, error) {
	_ = "STUB: not implemented"
	return *new(types.EventData), nil
}

// duration doesn't matter, limited by ctx timeout

// continue polling until ctx expires
