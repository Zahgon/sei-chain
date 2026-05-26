package kv

import (
	"context"

	dbm "github.com/tendermint/tm-db"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	kvb "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer/block/kv"
	kvt "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer/tx/kv"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var _ indexer.EventSink = (*EventSink)(nil)

// The EventSink is an aggregator for redirecting the call path of the tx/block kvIndexer.
// For the implementation details please see the kv.go in the indexer/block and indexer/tx folder.
type EventSink struct {
	txi   *kvt.TxIndex
	bi    *kvb.BlockerIndexer
	store dbm.DB
}

func NewEventSink(store dbm.DB) indexer.EventSink {
	_ = "STUB: not implemented"
	return *new(indexer.EventSink)
}

func (kves *EventSink) Type() indexer.EventSinkType {
	_ = "STUB: not implemented"
	return *new(indexer.EventSinkType)
}

func (kves *EventSink) IndexBlockEvents(bh types.EventDataNewBlockHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func (kves *EventSink) IndexTxEvents(results []*abci.TxResultV2) error {
	_ = "STUB: not implemented"
	return nil
}

func (kves *EventSink) SearchBlockEvents(ctx context.Context, q *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kves *EventSink) SearchTxEvents(ctx context.Context, q *query.Query) ([]*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kves *EventSink) GetTxByHash(hash []byte) (*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kves *EventSink) HasBlock(h int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (kves *EventSink) Stop() error { _ = "STUB: not implemented"; return nil }
