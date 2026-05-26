package null

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var _ indexer.BlockIndexer = (*BlockerIndexer)(nil)

// TxIndex implements a no-op block indexer.
type BlockerIndexer struct{}

func (idx *BlockerIndexer) Has(height int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (idx *BlockerIndexer) Index(types.EventDataNewBlockHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *BlockerIndexer) Search(ctx context.Context, q *query.Query) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
