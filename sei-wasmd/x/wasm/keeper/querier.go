package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var _ types.QueryServer = &grpcQuerier{}

type grpcQuerier struct {
	cdc           codec.Codec
	storeKey      sdk.StoreKey
	keeper        types.ViewKeeper
	queryGasLimit sdk.Gas
	paramsKeeper  types.ParamsKeeper
}

// NewGrpcQuerier constructor
func NewGrpcQuerier(cdc codec.Codec, storeKey sdk.StoreKey, keeper types.ViewKeeper, queryGasLimit sdk.Gas, paramsKeeper types.ParamsKeeper) *grpcQuerier {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

func (q grpcQuerier) ContractInfo(c context.Context, req *types.QueryContractInfoRequest) (*types.QueryContractInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q grpcQuerier) ContractHistory(c context.Context, req *types.QueryContractHistoryRequest) (*types.QueryContractHistoryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// redact

// ContractsByCode lists all smart contracts for a code id
func (q grpcQuerier) ContractsByCode(c context.Context, req *types.QueryContractsByCodeRequest) (*types.QueryContractsByCodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q grpcQuerier) AllContractState(c context.Context, req *types.QueryAllContractStateRequest) (*types.QueryAllContractStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q grpcQuerier) RawContractState(c context.Context, req *types.QueryRawContractStateRequest) (*types.QueryRawContractStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q grpcQuerier) SmartContractState(c context.Context, req *types.QuerySmartContractStateRequest) (rsp *types.QuerySmartContractStateResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get cosmos gas params

// set gas meter with appropriate multiplier

// recover from out-of-gas panic

func (q grpcQuerier) Code(c context.Context, req *types.QueryCodeRequest) (*types.QueryCodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q grpcQuerier) Codes(c context.Context, req *types.QueryCodesRequest) (*types.QueryCodesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryContractInfo(ctx sdk.Context, addr sdk.AccAddress, keeper types.ViewKeeper) (*types.QueryContractInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// redact the Created field (just used for sorting, not part of public API)

func queryCode(ctx sdk.Context, codeID uint64, keeper types.ViewKeeper) (*types.QueryCodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nil, nil leads to 404 in rest handler

func (q grpcQuerier) PinnedCodes(c context.Context, req *types.QueryPinnedCodesRequest) (*types.QueryPinnedCodesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
