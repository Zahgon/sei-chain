package http

import (
	"context"
	"sync"
	"time"

	rpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	jsonrpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/client"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "rpc", "client", "http")

// wsEvents is a wrapper around WSClient, which implements SubscriptionClient.
type wsEvents struct {
	ws *jsonrpcclient.WSClient

	mtx           sync.RWMutex
	subscriptions map[string]*wsSubscription
}

type wsSubscription struct {
	res   chan coretypes.ResultEvent
	id    string
	query string
}

var _ rpcclient.SubscriptionClient = (*wsEvents)(nil)

func newWsEvents(remote string) (*wsEvents, error) { _ = "STUB: not implemented"; return nil, nil }

// resubscribe immediately

// Start starts the websocket client and the event loop.
func (w *wsEvents) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop shuts down the websocket client.
func (w *wsEvents) Stop() error {
	_ = "STUB: not implemented"

	// Subscribe implements SubscriptionClient by using WSClient to subscribe given
	// subscriber to query. By default, it returns a channel with cap=1. Error is
	// returned if it fails to subscribe.
	//
	// When reading from the channel, keep in mind there's a single events loop, so
	// if you don't read events for this subscription fast enough, other
	// subscriptions will slow down in effect.
	//
	// The channel is never closed to prevent clients from seeing an erroneous
	// event.
	//
	// It returns an error if wsEvents is not running.
	return nil
}

func (w *wsEvents) Subscribe(ctx context.Context, subscriber, query string,
	outCapacity ...int) (out <-chan coretypes.ResultEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// subscriber param is ignored because Tendermint will override it with
// remote IP anyway.

// Unsubscribe implements SubscriptionClient by using WSClient to unsubscribe
// given subscriber from query.
//
// It returns an error if wsEvents is not running.
func (w *wsEvents) Unsubscribe(ctx context.Context, subscriber, query string) error {
	_ = "STUB: not implemented"
	return nil
}

// UnsubscribeAll implements SubscriptionClient by using WSClient to
// unsubscribe given subscriber from all the queries.
//
// It returns an error if wsEvents is not running.
func (w *wsEvents) UnsubscribeAll(ctx context.Context, subscriber string) error {
	_ = "STUB: not implemented"
	return nil
}

// After being reconnected, it is necessary to redo subscription to server
// otherwise no data will be automatically received.
func (w *wsEvents) redoSubscriptionsAfter(d time.Duration) { _ = "STUB: not implemented"; return }

func isErrAlreadySubscribed(err error) bool { _ = "STUB: not implemented"; return false }

func (w *wsEvents) eventListener(ctx context.Context) { _ = "STUB: not implemented"; return }

// Error can be ErrAlreadySubscribed or max client (subscriptions per
// client) reached or Tendermint exited.
// We can ignore ErrAlreadySubscribed, but need to retry in other
// cases.

// Resubscribe after 1 second to give Tendermint time to restart (if
// crashed).
