package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// NetInfo returns network info.
// More: https://docs.tendermint.com/master/rpc/#/Info/net_info
func (env *Environment) NetInfo(ctx context.Context) (*coretypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Genesis returns genesis file.
// More: https://docs.tendermint.com/master/rpc/#/Info/genesis
func (env *Environment) Genesis(ctx context.Context) (*coretypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (env *Environment) GenesisChunked(ctx context.Context, req *coretypes.RequestGenesisChunked) (*coretypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
