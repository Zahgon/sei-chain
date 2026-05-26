package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

var _ types.QueryServer = Keeper{}

// ClientState implements the Query/ClientState gRPC method
func (q Keeper) ClientState(c context.Context, req *types.QueryClientStateRequest) (*types.QueryClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientStates implements the Query/ClientStates gRPC method
func (q Keeper) ClientStates(c context.Context, req *types.QueryClientStatesRequest) (*types.QueryClientStatesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only append to the list if it's accumulating

// ConsensusState implements the Query/ConsensusState gRPC method
func (q Keeper) ConsensusState(c context.Context, req *types.QueryConsensusStateRequest) (*types.QueryConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConsensusStates implements the Query/ConsensusStates gRPC method
func (q Keeper) ConsensusStates(c context.Context, req *types.QueryConsensusStatesRequest) (*types.QueryConsensusStatesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter any metadata stored under consensus state key

// ConsensusStateHeights implements the Query/ConsensusStateHeights gRPC method
func (q Keeper) ConsensusStateHeights(c context.Context, req *types.QueryConsensusStateHeightsRequest) (*types.QueryConsensusStateHeightsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter any metadata stored under consensus state key

// ClientStatus implements the Query/ClientStatus gRPC method
func (q Keeper) ClientStatus(c context.Context, req *types.QueryClientStatusRequest) (*types.QueryClientStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientParams implements the Query/ClientParams gRPC method
func (q Keeper) ClientParams(c context.Context, _ *types.QueryClientParamsRequest) (*types.QueryClientParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpgradedClientState implements the Query/UpgradedClientState gRPC method
func (q Keeper) UpgradedClientState(c context.Context, req *types.QueryUpgradedClientStateRequest) (*types.QueryUpgradedClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpgradedConsensusState implements the Query/UpgradedConsensusState gRPC method
func (q Keeper) UpgradedConsensusState(c context.Context, req *types.QueryUpgradedConsensusStateRequest) (*types.QueryUpgradedConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
