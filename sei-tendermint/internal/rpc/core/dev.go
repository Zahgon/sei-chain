package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// UnsafeFlushMempool removes all transactions from the mempool.
func (env *Environment) UnsafeFlushMempool(ctx context.Context) (*coretypes.ResultUnsafeFlushMempool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
