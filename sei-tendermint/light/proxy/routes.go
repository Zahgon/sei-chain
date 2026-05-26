package proxy

import (
	"context"

	lrpc "github.com/sei-protocol/sei-chain/sei-tendermint/light/rpc"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// proxyService wraps a light RPC client to export the RPC service interfaces.
// The interfaces are implemented by delegating to the underlying node via the
// specified client.
type proxyService struct {
	Client *lrpc.Client
}

func (p proxyService) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) ABCIQuery(ctx context.Context, req *coretypes.RequestABCIQuery) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Block(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BlockByHash(ctx context.Context, req *coretypes.RequestBlockByHash) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BlockResults(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BlockSearch(ctx context.Context, req *coretypes.RequestBlockSearch) (*coretypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BlockchainInfo(ctx context.Context, req *coretypes.RequestBlockchainInfo) (*coretypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BroadcastEvidence(ctx context.Context, req *coretypes.RequestBroadcastEvidence) (*coretypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BroadcastTxAsync(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BroadcastTx(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BroadcastTxCommit(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) BroadcastTxSync(ctx context.Context, req *coretypes.RequestBroadcastTx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) CheckTx(ctx context.Context, req *coretypes.RequestCheckTx) (*coretypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Commit(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) ConsensusParams(ctx context.Context, req *coretypes.RequestConsensusParams) (*coretypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) DumpConsensusState(ctx context.Context) (*coretypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Events(ctx context.Context, req *coretypes.RequestEvents) (*coretypes.ResultEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Genesis(ctx context.Context) (*coretypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) GenesisChunked(ctx context.Context, req *coretypes.RequestGenesisChunked) (*coretypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // chunk index is validated upstream by the RPC layer; no negative values expected

func (p proxyService) GetConsensusState(ctx context.Context) (*coretypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Header(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) HeaderByHash(ctx context.Context, req *coretypes.RequestBlockByHash) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Health(ctx context.Context) (*coretypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) NetInfo(ctx context.Context) (*coretypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) NumUnconfirmedTxs(ctx context.Context) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) LagStatus(ctx context.Context) (*coretypes.ResultLagStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Subscribe(ctx context.Context, req *coretypes.RequestSubscribe) (*coretypes.ResultSubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Tx(ctx context.Context, req *coretypes.RequestTx) (*coretypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) TxSearch(ctx context.Context, req *coretypes.RequestTxSearch) (*coretypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) UnconfirmedTxs(ctx context.Context, req *coretypes.RequestUnconfirmedTxs) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Unsubscribe(ctx context.Context, req *coretypes.RequestUnsubscribe) (*coretypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) UnsubscribeAll(ctx context.Context) (*coretypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p proxyService) Validators(ctx context.Context, req *coretypes.RequestValidators) (*coretypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
