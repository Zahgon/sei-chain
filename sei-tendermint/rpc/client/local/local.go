package local

import (
	"context"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub"
	rpccore "github.com/sei-protocol/sei-chain/sei-tendermint/internal/rpc/core"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	rpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "rpc", "client", "local")

/*
Local is a Client implementation that directly executes the rpc
functions on a given node, without going through HTTP or GRPC.

This implementation is useful for:

* Running tests against a node in-process without the overhead
of going through an http server
* Communication between an ABCI app and Tendermint core when they
are compiled in process.

For real clients, you probably want to use client.HTTP.  For more
powerful control during testing, you probably want the "client/mock" package.

You can subscribe for any event published by Tendermint using Subscribe method.
Note delivery is best-effort. If you don't read events fast enough, Tendermint
might cancel the subscription. The client will attempt to resubscribe (you
don't need to do anything). It will keep trying indefinitely with exponential
backoff (10ms -> 20ms -> 40ms) until successful.
*/
type Local struct {
	*eventbus.EventBus
	*rpccore.Environment
}

// NodeService describes the portion of the node interface that the
// local RPC client constructor needs to build a local client.
type NodeService interface {
	service.Service
	RPCEnvironment() *rpccore.Environment
	EventBus() *eventbus.EventBus
}

// New configures a client that calls the Node directly.
func New(node NodeService) (*Local, error) { _ = "STUB: not implemented"; return nil, nil }

var _ rpcclient.Client = (*Local)(nil)

func (c *Local) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) ABCIQueryWithOptions(ctx context.Context, path string, data bytes.HexBytes, opts rpcclient.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastTxCommit(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastTx(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastTxAsync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastTxSync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) UnconfirmedTxs(ctx context.Context, page, perPage *int) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) CheckTx(ctx context.Context, tx types.Tx) (*coretypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) EvmNextPendingNonce(addr common.Address) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *Local) EvmProxy(sender common.Address) (*url.URL, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Local) ConsensusState(ctx context.Context) (*coretypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) ConsensusParams(ctx context.Context, height *int64) (*coretypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) GenesisChunked(ctx context.Context, id uint) (*coretypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // id is a small genesis chunk index

func (c *Local) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BlockByHash(ctx context.Context, hash bytes.HexBytes) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Header(ctx context.Context, height *int64) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) HeaderByHash(ctx context.Context, hash bytes.HexBytes) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Commit(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Validators(ctx context.Context, height *int64, page, perPage *int) (*coretypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Tx(ctx context.Context, hash bytes.HexBytes, prove bool) (*coretypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) TxSearch(ctx context.Context, queryString string, prove bool, page, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BlockSearch(ctx context.Context, queryString string, page, perPage *int, orderBy string) (*coretypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) BroadcastEvidence(ctx context.Context, ev types.Evidence) (*coretypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) Subscribe(ctx context.Context, subscriber, queryString string, capacity ...int) (<-chan coretypes.ResultEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Local) eventsRoutine(ctx context.Context, sub eventbus.Subscription, subArgs pubsub.SubscribeArgs, outc chan<- coretypes.ResultEvent) {
	_ = "STUB: not implemented"
	return
}

// client unsubscribed

// client terminated

// Try to resubscribe with exponential backoff.
func (c *Local) resubscribe(ctx context.Context, subArgs pubsub.SubscribeArgs) eventbus.Subscription {
	_ = "STUB: not implemented"
	return *new(eventbus.Subscription)
}

//nolint:gosec // attempts is a small non-negative counter

func (c *Local) Unsubscribe(ctx context.Context, subscriber, queryString string) error {
	_ = "STUB: not implemented"
	return nil
}

// if this isn't a valid query it might be an ID, so
// we'll try that. It'll turn into an error when we
// try to unsubscribe. Eventually, perhaps, we'll want
// to change the interface to only allow
// unsubscription by ID, but that's a larger change.

func (c *Local) UnsubscribeAll(ctx context.Context, subscriber string) error {
	_ = "STUB: not implemented"
	return nil
}
