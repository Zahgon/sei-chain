package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
)

var _ types.QueryServer = Keeper{}

// Connection implements the Query/Connection gRPC method
func (q Keeper) Connection(c context.Context, req *types.QueryConnectionRequest) (*types.QueryConnectionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Connections implements the Query/Connections gRPC method
func (q Keeper) Connections(c context.Context, req *types.QueryConnectionsRequest) (*types.QueryConnectionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientConnections implements the Query/ClientConnections gRPC method
func (q Keeper) ClientConnections(c context.Context, req *types.QueryClientConnectionsRequest) (*types.QueryClientConnectionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionClientState implements the Query/ConnectionClientState gRPC method
func (q Keeper) ConnectionClientState(c context.Context, req *types.QueryConnectionClientStateRequest) (*types.QueryConnectionClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionConsensusState implements the Query/ConnectionConsensusState gRPC method
func (q Keeper) ConnectionConsensusState(c context.Context, req *types.QueryConnectionConsensusStateRequest) (*types.QueryConnectionConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
