// Package pubsub implements an event dispatching server with a single publisher
// and multiple subscriber clients. Multiple goroutines can safely publish to a
// single Server instance.
//
// Clients register subscriptions with a query to select which messages they
// wish to receive. When messages are published, they are broadcast to all
// clients whose subscription query matches that message. Queries are
// constructed using the github.com/tendermint/tendermint/internal/pubsub/query
// package.
//
// Example:
//
//	q, err := query.New(`account.name='John'`)
//	if err != nil {
//	    return err
//	}
//	sub, err := pubsub.SubscribeWithArgs(ctx, pubsub.SubscribeArgs{
//	    ClientID: "johns-transactions",
//	    Query:    q,
//	})
//	if err != nil {
//	    return err
//	}
//
//	for {
//	    next, err := sub.Next(ctx)
//	    if err == pubsub.ErrTerminated {
//	       return err // terminated by publisher
//	    } else if err != nil {
//	       return err // timed out, client unsubscribed, etc.
//	    }
//	    process(next)
//	}
package pubsub

import (
	"context"
	"errors"
	"sync"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var (
	logger = seilog.NewLogger("tendermint", "internal", "pubsub")

	// ErrSubscriptionNotFound is returned when a client tries to unsubscribe
	// from not existing subscription.
	ErrSubscriptionNotFound = errors.New("subscription not found")

	// ErrAlreadySubscribed is returned when a client tries to subscribe twice or
	// more using the same query.
	ErrAlreadySubscribed = errors.New("already subscribed")

	// ErrServerStopped is returned when attempting to publish or subscribe to a
	// server that has been stopped.
	ErrServerStopped = errors.New("pubsub server is stopped")
)

// SubscribeArgs are the parameters to create a new subscription.
type SubscribeArgs struct {
	ClientID string       // Client ID
	Query    *query.Query // filter query for events (required)
	Limit    int          // subscription queue capacity limit (0 means 1)
	Quota    int          // subscription queue soft quota (0 uses Limit)
}

// UnsubscribeArgs are the parameters to remove a subscription.
// The subscriber ID must be populated, and at least one of the client ID or
// the registered query.
type UnsubscribeArgs struct {
	Subscriber string       // subscriber ID chosen by the client (required)
	ID         string       // subscription ID (assigned by the server)
	Query      *query.Query // the query registered with the subscription
}

// Validate returns nil if args are valid to identify a subscription to remove.
// Otherwise, it reports an error.
func (args UnsubscribeArgs) Validate() error { _ = "STUB: not implemented"; return nil }

// Server allows clients to subscribe/unsubscribe for messages, publishing
// messages with or without events, and manages internal state.
type Server struct {
	service.BaseService

	queue  chan item
	done   <-chan struct{} // closed when server should exit
	pubs   sync.RWMutex    // excl: shutdown; shared: active publisher
	exited chan struct{}   // server exited

	// All subscriptions currently known.
	// Lock exclusive to add, remove, or cancel subscriptions.
	// Lock shared to look up or publish to subscriptions.
	subs struct {
		sync.RWMutex
		index *subIndex

		// This function is called synchronously with each message published
		// before it is delivered to any other subscriber. This allows an index
		// to be persisted before any subscribers see the messages.
		observe func(Message) error
	}

	// TODO(creachadair): Rework the options so that this does not need to live
	// as a field. It is not otherwise needed.
	queueCap int
}

// Option sets a parameter for the server.
type Option func(*Server)

// NewServer returns a new server. See the commentary on the Option functions
// for a detailed description of how to configure buffering. If no options are
// provided, the resulting server's queue is unbuffered.
func NewServer(options ...Option) *Server { _ = "STUB: not implemented"; return nil }

// The queue receives items to be published.

// The index tracks subscriptions by ID and query terms.

// BufferCapacity allows you to specify capacity for publisher's queue.  This
// is the number of messages that can be published without blocking.  If no
// buffer is specified, publishing is synchronous with delivery.  This function
// will panic if cap < 0.
func BufferCapacity(cap int) Option { _ = "STUB: not implemented"; return *new(Option) }

// BufferCapacity returns capacity of the publication queue.
func (s *Server) BufferCapacity() int {
	_ = "STUB: not implemented"

	// Observe registers an observer function that will be called synchronously
	// with each published message matching any of the given queries, prior to it
	// being forwarded to any subscriber.  If no queries are specified, all
	// messages will be observed. An error is reported if an observer is already
	// registered.
	return 0
}

func (s *Server) Observe(ctx context.Context, observe func(Message) error, queries ...*query.Query) error {
	_ = "STUB: not implemented"
	return nil
}

// Compile the message filter.

// nothing to do for this message

// SubscribeWithArgs creates a subscription for the given arguments.  It is an
// error if the query is nil, a subscription already exists for the specified
// client ID and query, or if the capacity arguments are invalid.
func (s *Server) SubscribeWithArgs(ctx context.Context, args SubscribeArgs) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unsubscribe removes the subscription for the given client and/or query.  It
// returns ErrSubscriptionNotFound if no such subscription exists.
func (s *Server) Unsubscribe(ctx context.Context, args UnsubscribeArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(creachadair): Do we need to support unsubscription for an "empty"
// query?  I believe that case is not possible by the Query grammar, but we
// should make sure.
//
// Revisit this logic once we are able to remove indexing by query.

// UnsubscribeAll removes all subscriptions for the given client ID.
// It returns ErrSubscriptionNotFound if no subscriptions exist for that client.
func (s *Server) UnsubscribeAll(ctx context.Context, clientID string) error {
	_ = "STUB: not implemented"
	return nil
}

// NumClients returns the number of clients.
func (s *Server) NumClients() int { _ = "STUB: not implemented"; return 0 }

// NumClientSubscriptions returns the number of subscriptions the client has.
func (s *Server) NumClientSubscriptions(clientID string) int { _ = "STUB: not implemented"; return 0 }

// Publish publishes the given message. An error will be returned to the caller
// if the pubsub server has shut down.
func (s *Server) Publish(msg types.EventData) error { _ = "STUB: not implemented"; return nil }

// PublishWithEvents publishes the given message with the set of events. The set
// is matched with clients queries. If there is a match, the message is sent to
// the client.
func (s *Server) PublishWithEvents(msg types.EventData, events []abci.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// OnStop implements part of the Service interface. It is a no-op.
func (s *Server) OnStop() {
	_ = "STUB: not implemented"

	// Wait implements Service.Wait by blocking until the server has exited, then
	// yielding to the base service wait.
	return
}

func (s *Server) Wait() { _ = "STUB: not implemented"; return }

// OnStart implements Service.OnStart by starting the server.
func (s *Server) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) publish(data types.EventData, events []abci.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) run(ctx context.Context) {
	_ = "STUB: not implemented"
	// The server runs until ctx is canceled.
	return
}

// Shutdown monitor: When the context ends, wait for any active publish
// calls to exit, then close the queue to signal the sender to exit.

// Sender: Service the queue and forward messages to subscribers.

// Terminate all subscribers before exit.

// removeSubs cancels and removes all the subscriptions in evict with the given
// error. The caller must hold the s.subs lock.
func (s *Server) removeSubs(evict subInfoSet, reason error) { _ = "STUB: not implemented"; return }

// send delivers the given message to all matching subscribers.  An error in
// query matching stops transmission and is returned.
func (s *Server) send(data types.EventData, events []abci.Event) error {
	_ = "STUB: not implemented"
	// At exit, evict any subscriptions that were too slow.
	return nil
}

// N.B. Order is important here. We must acquire and defer the lock release
// AFTER deferring the eviction cleanup: The cleanup must happen after the
// reader lock has released, or it will deadlock.

// If an observer is defined, give it control of the message before
// attempting to deliver it to any matching subscribers. If the observer
// fails, the message will not be forwarded.

// Publish the events to the subscriber's queue. If this fails, e.g.,
// because the queue is over capacity or out of quota, evict the
// subscription from the index.
