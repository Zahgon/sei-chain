package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
)

var _ types.QueryServer = Keeper{}

// DenomTrace implements the Query/DenomTrace gRPC method
func (q Keeper) DenomTrace(c context.Context, req *types.QueryDenomTraceRequest) (*types.QueryDenomTraceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DenomTraces implements the Query/DenomTraces gRPC method
func (q Keeper) DenomTraces(c context.Context, req *types.QueryDenomTracesRequest) (*types.QueryDenomTracesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Params implements the Query/Params gRPC method
func (q Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DenomHash implements the Query/DenomHash gRPC method
func (q Keeper) DenomHash(c context.Context, req *types.QueryDenomHashRequest) (*types.QueryDenomHashResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert given request trace path to DenomTrace struct to confirm the path in a valid denom trace format

// EscrowAddress implements the EscrowAddress gRPC method
func (q Keeper) EscrowAddress(c context.Context, req *types.QueryEscrowAddressRequest) (*types.QueryEscrowAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
