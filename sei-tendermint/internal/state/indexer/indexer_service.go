package indexer

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "internal", "state", "indexer")

// Service connects event bus, transaction and block indexers together in
// order to index transactions and blocks coming from the event bus.
type Service struct {
	service.BaseService

	eventSinks []EventSink
	eventBus   *eventbus.EventBus
	metrics    *Metrics

	currentBlock struct {
		header types.EventDataNewBlockHeader
		height int64
		batch  *Batch
	}
}

// NewService constructs a new indexer service from the given arguments.
func NewService(args ServiceArgs) *Service { _ = "STUB: not implemented"; return nil }

// publish publishes a pubsub message to the service. The service blocks until
// the message has been fully processed.
func (is *Service) publish(msg pubsub.Message) error {
	_ = "STUB: not implemented"
	// Indexing has three states. Initially, no block is in progress (WAIT) and
	// we expect a block header. Upon seeing a header, we are waiting for zero
	// or more transactions (GATHER). Once all the expected transactions have
	// been delivered (in some order), we are ready to index. After indexing a
	// block, we revert to the WAIT state for the next block.
	return nil
}

// WAIT: Start a new block.

// If the block does not expect any transactions, fall through and index
// it immediately.  This shouldn't happen, but this check ensures we do
// not get stuck if it does.

// GATHER: Accumulate a transaction into the current block's batch.

// This may have been the last transaction in the batch, so fall through
// to check whether it is time to index.

// INDEX: We have all the transactions we expect for the current block.

// return to the WAIT state for the next block

// OnStart implements part of service.Service. It registers an observer for the
// indexer if the underlying event sinks support indexing.
//
// TODO(creachadair): Can we get rid of the "enabled" check?
func (is *Service) OnStart(ctx context.Context) error {
	_ = "STUB: not implemented"
	// If the event sinks support indexing, register an observer to capture
	// block header data for the indexer.
	return nil
}

// OnStop implements service.Service by closing the event sinks.
func (is *Service) OnStop() { _ = "STUB: not implemented"; return }

// ServiceArgs are arguments for constructing a new indexer service.
type ServiceArgs struct {
	Sinks    []EventSink
	EventBus *eventbus.EventBus
	Metrics  *Metrics
}

// KVSinkEnabled returns the given eventSinks is containing KVEventSink.
func KVSinkEnabled(sinks []EventSink) bool { _ = "STUB: not implemented"; return false }

// IndexingEnabled returns the given eventSinks is supporting the indexing services.
func IndexingEnabled(sinks []EventSink) bool { _ = "STUB: not implemented"; return false }
