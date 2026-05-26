package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// ABCIQuery queries the application for some information.
// More: https://docs.tendermint.com/master/rpc/#/ABCI/abci_query
func (env *Environment) ABCIQuery(ctx context.Context, req *coretypes.RequestABCIQuery) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ABCIInfo gets some info about the application.
// More: https://docs.tendermint.com/master/rpc/#/ABCI/abci_info
func (env *Environment) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
