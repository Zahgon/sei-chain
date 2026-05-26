package http

import (
	"context"
	"net/http"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	rpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	jsonrpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/client"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

/*
HTTP is a Client implementation that communicates with a Tendermint node over
JSON RPC and WebSockets.

This is the main implementation you probably want to use in production code.
There are other implementations when calling the Tendermint node in-process
(Local), or when you want to mock out the server for test code (mock).

You can subscribe for any event published by Tendermint using Subscribe method.
Note delivery is best-effort. If you don't read events fast enough or network is
slow, Tendermint might cancel the subscription. The client will attempt to
resubscribe (you don't need to do anything). It will keep trying every second
indefinitely until successful.

Request batching is available for JSON RPC requests over HTTP, which conforms to
the JSON RPC specification (https://www.jsonrpc.org/specification#batch). See
the example for more details.

Example:

	c, err := New("http://192.168.1.10:26657")
	if err != nil {
		// handle error
	}

	// call Start/Stop if you're subscribing to events
	err = c.Start()
	if err != nil {
		// handle error
	}
	defer c.Stop()

	res, err := c.Status()
	if err != nil {
		// handle error
	}

	// handle result
*/
type HTTP struct {
	remote string
	rpc    *jsonrpcclient.Client

	*baseRPCClient
	*wsEvents
}

// BatchHTTP provides the same interface as `HTTP`, but allows for batching of
// requests (as per https://www.jsonrpc.org/specification#batch). Do not
// instantiate directly - rather use the HTTP.NewBatch() method to create an
// instance of this struct.
//
// Batching of HTTP requests is thread-safe in the sense that multiple
// goroutines can each create their own batches and send them using the same
// HTTP client. Multiple goroutines could also enqueue transactions in a single
// batch, but ordering of transactions in the batch cannot be guaranteed in such
// an example.
type BatchHTTP struct {
	rpcBatch *jsonrpcclient.RequestBatch
	*baseRPCClient
}

// rpcClient is an internal interface to which our RPC clients (batch and
// non-batch) must conform. Acts as an additional code-level sanity check to
// make sure the implementations stay coherent.
type rpcClient interface {
	rpcclient.ABCIClient
	rpcclient.HistoryClient
	rpcclient.NetworkClient
	rpcclient.SignClient
	rpcclient.StatusClient
}

// baseRPCClient implements the basic RPC method logic without the actual
// underlying RPC call functionality, which is provided by `caller`.
type baseRPCClient struct {
	caller jsonrpcclient.Caller
}

var (
	_ rpcClient = (*HTTP)(nil)
	_ rpcClient = (*BatchHTTP)(nil)
	_ rpcClient = (*baseRPCClient)(nil)
)

//-----------------------------------------------------------------------------
// HTTP

// New takes a remote endpoint in the form <protocol>://<host>:<port>. An error
// is returned on invalid remote.
func New(remote string) (*HTTP, error) { _ = "STUB: not implemented"; return nil, nil }

// NewWithTimeout does the same thing as New, except you can set a Timeout for
// http.Client. A Timeout of zero means no timeout.
func NewWithTimeout(remote string, t time.Duration) (*HTTP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWithClient constructs an RPC client using a custom HTTP client.
// An error is reported if c == nil or remote is an invalid address.
func NewWithClient(remote string, c *http.Client) (*HTTP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ rpcclient.Client = (*HTTP)(nil)

// Remote returns the remote network address in a string form.
func (c *HTTP) Remote() string {
	_ = "STUB: not implemented"

	// NewBatch creates a new batch client for this HTTP client.
	return ""
}

func (c *HTTP) NewBatch() *BatchHTTP { _ = "STUB: not implemented"; return nil }

//-----------------------------------------------------------------------------
// BatchHTTP

// Send is a convenience function for an HTTP batch that will trigger the
// compilation of the batched requests and send them off using the client as a
// single request. On success, this returns a list of the deserialized results
// from each request in the sent batch.
func (b *BatchHTTP) Send(ctx context.Context) ([]any, error) {
	_ = "STUB: not implemented"
	return nil,

		// Clear will empty out this batch of requests and return the number of requests
		// that were cleared out.
		nil
}

func (b *BatchHTTP) Clear() int { _ = "STUB: not implemented"; return 0 }

// Count returns the number of enqueued requests waiting to be sent.
func (b *BatchHTTP) Count() int { _ = "STUB: not implemented"; return 0 }

//-----------------------------------------------------------------------------
// baseRPCClient

func (c *baseRPCClient) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) LagStatus(ctx context.Context) (*coretypes.ResultLagStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ABCIQueryWithOptions(ctx context.Context, path string, data bytes.HexBytes, opts rpcclient.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastTxCommit(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastTxAsync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastTxSync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastTx(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) broadcastTX(ctx context.Context, route string, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) UnconfirmedTxs(ctx context.Context, page *int, perPage *int) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) NumUnconfirmedTxs(ctx context.Context) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) CheckTx(ctx context.Context, tx types.Tx) (*coretypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) NetInfo(ctx context.Context) (*coretypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) DumpConsensusState(ctx context.Context) (*coretypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ConsensusState(ctx context.Context) (*coretypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) ConsensusParams(ctx context.Context, height *int64) (*coretypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Events(ctx context.Context, req *coretypes.RequestEvents) (*coretypes.ResultEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Health(ctx context.Context) (*coretypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Genesis(ctx context.Context) (*coretypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) GenesisChunked(ctx context.Context, id uint) (*coretypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // id is a small genesis chunk index

func (c *baseRPCClient) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BlockByHash(ctx context.Context, hash bytes.HexBytes) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Header(ctx context.Context, height *int64) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) HeaderByHash(ctx context.Context, hash bytes.HexBytes) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Commit(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Tx(ctx context.Context, hash bytes.HexBytes, prove bool) (*coretypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) TxSearch(ctx context.Context, query string, prove bool, page, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BlockSearch(ctx context.Context, query string, page, perPage *int, orderBy string) (*coretypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) Validators(ctx context.Context, height *int64, page, perPage *int) (*coretypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseRPCClient) BroadcastEvidence(ctx context.Context, ev types.Evidence) (*coretypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
