package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"

	crgtypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/rosetta/lib/types"
)

// NewOffline instantiates the instance of an offline network
// whilst the offline network does not support the DataAPI,
// it supports a subset of the construction API.
func NewOffline(network *types.NetworkIdentifier, client crgtypes.Client) (crgtypes.API, error) {
	_ = "STUB: not implemented"
	return *new(crgtypes.API), nil
}

// OfflineNetwork implements an offline data API
// which is basically a data API that constantly
// returns errors, because it cannot be used if offline
type OfflineNetwork struct {
	OnlineNetwork
}

// Implement DataAPI in offline mode, which means no method is available
func (o OfflineNetwork) AccountBalance(_ context.Context, _ *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o OfflineNetwork) Block(_ context.Context, _ *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o OfflineNetwork) BlockTransaction(_ context.Context, _ *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o OfflineNetwork) Mempool(_ context.Context, _ *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o OfflineNetwork) MempoolTransaction(_ context.Context, _ *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o OfflineNetwork) NetworkStatus(_ context.Context, _ *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o OfflineNetwork) ConstructionSubmit(_ context.Context, _ *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o OfflineNetwork) ConstructionMetadata(_ context.Context, _ *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}
