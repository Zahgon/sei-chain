package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// Tx allows you to query the transaction results. `nil` could mean the
// transaction is in the mempool, invalidated, or was not sent in the first
// place.
// More: https://docs.tendermint.com/master/rpc/#/Info/tx
func (env *Environment) Tx(ctx context.Context, req *coretypes.RequestTx) (*coretypes.ResultTx, error) {
	_ = "STUB: not implemented"
	// if index is disabled, return error
	return nil, nil
}

// TxSearch allows you to query for multiple transactions results. It returns a
// list of transactions (maximum ?per_page entries) and the total count.
// More: https://docs.tendermint.com/master/rpc/#/Info/tx_search
func (env *Environment) TxSearch(ctx context.Context, req *coretypes.RequestTxSearch) (*coretypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort results (must be done before pagination)

// paginate results
