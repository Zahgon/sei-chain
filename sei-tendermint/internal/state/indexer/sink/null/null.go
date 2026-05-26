package null

import (
	"context"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var _ indexer.EventSink = (*EventSink)(nil)

// EventSink implements a no-op indexer.
type EventSink struct{}

func NewEventSink() indexer.EventSink { _ = "STUB: not implemented"; return *new(indexer.EventSink) }

func (nes *EventSink) Type() indexer.EventSinkType {
	_ = "STUB: not implemented"
	return *new(indexer.EventSinkType)
}

func (nes *EventSink) IndexBlockEvents(bh types.EventDataNewBlockHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func (nes *EventSink) IndexTxEvents(results []*abci.TxResultV2) error {
	_ = "STUB: not implemented"
	return nil
}

func (nes *EventSink) SearchBlockEvents(ctx context.Context, q *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nes *EventSink) SearchTxEvents(ctx context.Context, q *query.Query) ([]*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nes *EventSink) GetTxByHash(hash []byte) (*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nes *EventSink) HasBlock(h int64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (nes *EventSink) Stop() error { _ = "STUB: not implemented"; return nil }
