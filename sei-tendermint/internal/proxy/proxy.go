package proxy

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/metrics"

	"github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// Proxy wraps an ABCI application and records ABCI method timings.
type Proxy struct {
	app     types.Application
	metrics *Metrics
}

// New creates a proxied application interface around the provided ABCI application.
func New(app types.Application, metrics *Metrics) *Proxy { _ = "STUB: not implemented"; return nil }

func (app *Proxy) InitChain(ctx context.Context, req *types.RequestInitChain) (*types.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) ProcessProposal(ctx context.Context, req *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) FinalizeBlock(ctx context.Context, req *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) GetTxPriorityHint(ctx context.Context, req *types.RequestGetTxPriorityHintV2) (*types.ResponseGetTxPriorityHint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) EvmNonce(addr common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

func (app *Proxy) EvmBalance(addr common.Address, seiAddr []byte) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (app *Proxy) Commit(ctx context.Context) (*types.ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) CheckTxSafe(ctx context.Context, req *types.RequestCheckTxV2) (res *types.ResponseCheckTxV2, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) Info(ctx context.Context, req *types.RequestInfo) (*types.ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) Query(ctx context.Context, req *types.RequestQuery) (*types.ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) GetValidators() []types.ValidatorUpdate { _ = "STUB: not implemented"; return nil }

func (app *Proxy) ListSnapshots(ctx context.Context, req *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) OfferSnapshot(ctx context.Context, req *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) LoadSnapshotChunk(ctx context.Context, req *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (app *Proxy) ApplySnapshotChunk(ctx context.Context, req *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// addTimeSample returns a function that, when called, adds an observation to m.
// The observation added to m is the number of seconds ellapsed since addTimeSample
// was initially called. addTimeSample is meant to be called in a defer to calculate
// the amount of time a function takes to complete.
func addTimeSample(m metrics.Histogram) func() { _ = "STUB: not implemented"; return nil }
