package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// AccountBalance retrieves the account balance of an address
// rosetta requires us to fetch the block information too
func (on OnlineNetwork) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Block gets the transactions in the given block
func (on OnlineNetwork) Block(ctx context.Context, request *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// block identifier is assumed not to be nil as rosetta will do this check for us
// check if we have to query via hash or block number

// BlockTransaction gets the given transaction in the specified block, we do not need to check the block itself too
// due to the fact that tendermint achieves instant finality
func (on OnlineNetwork) BlockTransaction(ctx context.Context, request *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mempool fetches the transactions contained in the mempool
func (on OnlineNetwork) Mempool(ctx context.Context, _ *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MempoolTransaction fetches a single transaction in the mempool
// NOTE: it is not implemented yet
func (on OnlineNetwork) MempoolTransaction(ctx context.Context, request *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (on OnlineNetwork) NetworkList(_ context.Context, _ *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (on OnlineNetwork) NetworkOptions(_ context.Context, _ *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (on OnlineNetwork) NetworkStatus(ctx context.Context, _ *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}
