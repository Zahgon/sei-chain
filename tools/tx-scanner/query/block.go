package query

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client/grpc/tmservice"
)

// GetLatestBlock query the latest block data
func GetLatestBlock() (*tmservice.GetLatestBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockByHeight query the block data at height
func GetBlockByHeight(height int64) (*tmservice.GetBlockByHeightResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
