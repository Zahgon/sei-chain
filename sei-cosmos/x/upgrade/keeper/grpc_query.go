package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/types"
)

var _ types.QueryServer = Keeper{}

// CurrentPlan implements the Query/CurrentPlan gRPC method
func (k Keeper) CurrentPlan(c context.Context, req *types.QueryCurrentPlanRequest) (*types.QueryCurrentPlanResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppliedPlan implements the Query/AppliedPlan gRPC method
func (k Keeper) AppliedPlan(c context.Context, req *types.QueryAppliedPlanRequest) (*types.QueryAppliedPlanResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpgradedConsensusState implements the Query/UpgradedConsensusState gRPC method
// nolint: staticcheck
func (k Keeper) UpgradedConsensusState(c context.Context, req *types.QueryUpgradedConsensusStateRequest) (*types.QueryUpgradedConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ModuleVersions implements the Query/QueryModuleVersions gRPC method
func (k Keeper) ModuleVersions(c context.Context, req *types.QueryModuleVersionsRequest) (*types.QueryModuleVersionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// check if a specific module was requested
}

// return the requested module

// module requested, but not found

// if no module requested return all module versions from state
