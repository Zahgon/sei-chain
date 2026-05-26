package eventbus

import (
	"context"

	tmpubsub "github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub"
	tmquery "github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var DefaultBufferCapacity = 10000

// Subscription is a proxy interface for a pubsub Subscription.
type Subscription interface {
	ID() string
	Next(context.Context) (tmpubsub.Message, error)
}

// EventBus is a common bus for all events going through the system.
// It is a type-aware wrapper around an underlying pubsub server.
// All events should be published via the bus.
type EventBus struct {
	service.BaseService
	pubsub *tmpubsub.Server
}

// NewDefault returns a new event bus with default options.
func NewDefault() *EventBus { _ = "STUB: not implemented"; return nil }

func (b *EventBus) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *EventBus) OnStop() { _ = "STUB: not implemented"; return }

func (b *EventBus) NumClients() int { _ = "STUB: not implemented"; return 0 }

func (b *EventBus) NumClientSubscriptions(clientID string) int { _ = "STUB: not implemented"; return 0 }

func (b *EventBus) SubscribeWithArgs(ctx context.Context, args tmpubsub.SubscribeArgs) (Subscription, error) {
	_ = "STUB: not implemented"
	return *new(Subscription), nil
}

func (b *EventBus) Unsubscribe(ctx context.Context, args tmpubsub.UnsubscribeArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) UnsubscribeAll(ctx context.Context, subscriber string) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) Observe(ctx context.Context, observe func(tmpubsub.Message) error, queries ...*tmquery.Query) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) Publish(eventValue string, eventData types.EventData) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventNewBlock(data types.EventDataNewBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// add Tendermint-reserved new block event
// copy to a new destination explicitly

func (b *EventBus) PublishEventNewBlockHeader(data types.EventDataNewBlockHeader) error {
	_ = "STUB: not implemented"
	// no explicit deadline for publishing events
	return nil
}

// add Tendermint-reserved new block header event
// copy to a new destination explicitly

func (b *EventBus) PublishEventNewEvidence(evidence types.EventDataNewEvidence) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventVote(data types.EventDataVote) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventValidBlock(data types.EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventBlockSyncStatus(data types.EventDataBlockSyncStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventStateSyncStatus(data types.EventDataStateSyncStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// PublishEventTx publishes tx event with events from Result. Note it will add
// predefined keys (EventTypeKey, TxHashKey). Existing events with the same keys
// will be overwritten.
func (b *EventBus) PublishEventTx(data types.EventDataTx) error {
	_ = "STUB: not implemented"
	return nil
}

// add Tendermint-reserved events

func (b *EventBus) PublishEventNewRoundStep(data types.EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventTimeoutPropose(data types.EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventTimeoutWait(data types.EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventNewRound(data types.EventDataNewRound) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventCompleteProposal(data types.EventDataCompleteProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventPolka(data types.EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventRelock(data types.EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventLock(data types.EventDataRoundState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventValidatorSetUpdates(data types.EventDataValidatorSetUpdates) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *EventBus) PublishEventEvidenceValidated(evidence types.EventDataEvidenceValidated) error {
	_ = "STUB: not implemented"
	return nil
}
