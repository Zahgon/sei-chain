package null

import (
	"context"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
)

var _ indexer.TxIndexer = (*TxIndex)(nil)

// TxIndex acts as a /dev/null.
type TxIndex struct{}

// Get on a TxIndex is disabled and panics when invoked.
func (txi *TxIndex) Get(hash []byte) (*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddBatch is a noop and always returns nil.
func (txi *TxIndex) AddBatch(batch *indexer.Batch) error {
	_ = "STUB: not implemented"

	// Index is a noop and always returns nil.
	return nil
}

func (txi *TxIndex) Index(results []*abci.TxResultV2) error { _ = "STUB: not implemented"; return nil }

func (txi *TxIndex) Search(ctx context.Context, q *query.Query) ([]*abci.TxResultV2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
