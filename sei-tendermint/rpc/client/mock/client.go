package mock

/*
package mock returns a Client implementation that
accepts various (mock) implementations of the various methods.

This implementation is useful for using in tests, when you don't
need a real server, but want a high-level of control about
the server response you want to mock (eg. error handling),
or if you just want to record the calls to verify in your tests.

For real clients, you probably want the "http" package.  If you
want to directly call a tendermint node in process, you can use the
"local" package.
*/

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/rpc/core"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// Client wraps arbitrary implementations of the various interfaces.
type Client struct {
	client.Client
	env *core.Environment
}

func New() Client { _ = "STUB: not implemented"; return *new(Client) }

var _ client.Client = Client{}

// Call is used by recorders to save a call and response.
// It can also be used to configure mock responses.
type Call struct {
	Name     string
	Args     interface{}
	Response interface{}
	Error    error
}

// GetResponse will generate the apporiate response for us, when
// using the Call struct to configure a Mock handler.
//
// When configuring a response, if only one of Response or Error is
// set then that will always be returned. If both are set, then
// we return Response if the Args match the set args, Error otherwise.
func (c Call) GetResponse(args interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	// handle the case with no response
	return nil, nil
}

// response without error

// have both, we must check args....

func (c Client) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) LagStatus(ctx context.Context) (*coretypes.ResultLagStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ABCIQueryWithOptions(
	ctx context.Context,
	path string,
	data bytes.HexBytes,
	opts client.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastTxCommit(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastTxAsync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastTxSync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) CheckTx(ctx context.Context, tx types.Tx) (*coretypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) NetInfo(ctx context.Context) (*coretypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ConsensusState(ctx context.Context) (*coretypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) DumpConsensusState(ctx context.Context) (*coretypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) ConsensusParams(ctx context.Context, height *int64) (*coretypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Health(ctx context.Context) (*coretypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Genesis(ctx context.Context) (*coretypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BlockByHash(ctx context.Context, hash bytes.HexBytes) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Commit(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) Validators(ctx context.Context, height *int64, page, perPage *int) (*coretypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Client) BroadcastEvidence(ctx context.Context, ev types.Evidence) (*coretypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
