package tests

import (
	"context"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/mock"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type MockClient struct {
	mock.Client
	blocks           [][][]byte
	txResults        [][]*abci.ExecTxResult
	consParamUpdates []*tmproto.ConsensusParams
	events           [][]abci.Event

	mockedBlockResults        map[int64]*coretypes.ResultBlock
	mockedBlockByHashResults  map[string]*coretypes.ResultBlock
	mockedBlockResultsResults map[int64]*coretypes.ResultBlockResults
	mockedValidators          map[int64]*coretypes.ResultValidators
	mockedGenesis             *coretypes.ResultGenesis
}

func (c *MockClient) EvmNextPendingNonce(_ common.Address) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *MockClient) EvmProxy(common.Address) (*url.URL, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *MockClient) Block(_ context.Context, h *int64) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) BlockByHash(_ context.Context, hash tmbytes.HexBytes) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) getBlock(i int64) *coretypes.ResultBlock {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockClient) Header(_ context.Context, h *int64) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) Genesis(context.Context) (*coretypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) BlockResults(_ context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) Status(context.Context) (*coretypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) recordBlockResult(txResults []*abci.ExecTxResult, consParamUpdates *tmproto.ConsensusParams, events []abci.Event) {
	_ = "STUB: not implemented"
	return
}

func (c *MockClient) Validators(ctx context.Context, height *int64, page, perPage *int) (*coretypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) Events(_ context.Context, req *coretypes.RequestEvents) (*coretypes.ResultEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MockClient) UnconfirmedTxs(context.Context, *int, *int) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mockHash(height int64, prefix int64) tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

//nolint:gosec
//nolint:gosec

func mockBlockHeader(height int64) *tmtypes.Header { _ = "STUB: not implemented"; return nil }
