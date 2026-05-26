package pubsub

import (
	"context"
	"errors"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/libs/queue"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var (
	// ErrUnsubscribed is returned by Next when the client has unsubscribed.
	ErrUnsubscribed = errors.New("subscription removed by client")

	// ErrTerminated is returned by Next when the subscription was terminated by
	// the publisher.
	ErrTerminated = errors.New("subscription terminated by publisher")
)

// A Subscription represents a client subscription for a particular query.
type Subscription struct {
	id      string
	queue   *queue.Queue // open until the subscription ends
	stopErr error        // after queue is closed, the reason why
}

// newSubscription returns a new subscription with the given queue capacity.
func newSubscription(quota, limit int) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Next blocks until a message is available, ctx ends, or the subscription
// ends.  Next returns ErrUnsubscribed if s was unsubscribed, ErrTerminated if
// s was terminated by the publisher, or a context error if ctx ended without a
// message being available.
func (s *Subscription) Next(ctx context.Context) (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

// ID returns the unique subscription identifier for s.
func (s *Subscription) ID() string {
	_ = "STUB: not implemented"

	// publish transmits msg to the subscriber. It reports a queue error if the
	// queue cannot accept any further messages.
	return ""
}

func (s *Subscription) publish(msg Message) error { _ = "STUB: not implemented"; return nil }

// stop terminates the subscription with the given error reason.
func (s *Subscription) stop(err error) { _ = "STUB: not implemented"; return }

// Message glues data and events together.
type Message struct {
	subID  string
	data   types.EventData
	events []abci.Event
}

// SubscriptionID returns the unique identifier for the subscription
// that produced this message.
func (msg Message) SubscriptionID() string {
	_ = "STUB: not implemented"

	// Data returns an original data published.
	return ""
}

func (msg Message) Data() types.EventData { _ = "STUB: not implemented"; return *new(types.EventData) }

func (msg Message) LegacyData() types.LegacyEventData {
	_ = "STUB: not implemented"
	return *new(types.LegacyEventData)
}

// Events returns events, which matched the client's query.
func (msg Message) Events() []abci.Event { _ = "STUB: not implemented"; return nil }
