package core

import (
	"context"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// EvmProxy returns the EVM RPC URL of the autobahn validator that owns the
// sender shard. If the sender maps to the local validator, or if no EVM RPC
// endpoint is configured for the shard owner, it returns (nil, false).
func (env *Environment) EvmProxy(sender common.Address) (*url.URL, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//-----------------------------------------------------------------------------
// NOTE: tx should be signed, but this is only checked at the app level (not by Tendermint!)

// BroadcastTxAsync returns right away, with no response. Does not wait for
// CheckTx nor DeliverTx results.
// More:
// https://docs.tendermint.com/master/rpc/#/Tx/broadcast_tx_async
// Deprecated and should be removed in 0.37
func (env *Environment) BroadcastTxAsync(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated and should be remove in 0.37
func (env *Environment) BroadcastTxSync(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BroadcastTx returns with the response from CheckTx. Does not wait for
// DeliverTx result.
// More: https://docs.tendermint.com/master/rpc/#/Tx/broadcast_tx_sync
func (env *Environment) BroadcastTx(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BroadcastTxCommit returns with the responses from CheckTx and DeliverTx.
// More: https://docs.tendermint.com/master/rpc/#/Tx/broadcast_tx_commit
func (env *Environment) BroadcastTxCommit(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint: gosec

// UnconfirmedTxs gets unconfirmed transactions from the mempool in order of priority
// More: https://docs.tendermint.com/master/rpc/#/Info/unconfirmed_txs
func (env *Environment) UnconfirmedTxs(ctx context.Context, req *coretypes.RequestUnconfirmedTxs) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NumUnconfirmedTxs gets number of unconfirmed transactions.
// More: https://docs.tendermint.com/master/rpc/#/Info/num_unconfirmed_txs
func (env *Environment) NumUnconfirmedTxs(ctx context.Context) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckTx checks the transaction without executing it. The transaction won't
// be added to the mempool either.
// More: https://docs.tendermint.com/master/rpc/#/Tx/check_tx
func (env *Environment) CheckTx(ctx context.Context, req *coretypes.RequestCheckTx) (*coretypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
