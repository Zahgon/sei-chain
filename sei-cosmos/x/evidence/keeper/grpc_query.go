package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/types"
)

var _ types.QueryServer = Keeper{}

// Evidence implements the Query/Evidence gRPC method
func (k Keeper) Evidence(c context.Context, req *types.QueryEvidenceRequest) (*types.QueryEvidenceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllEvidence implements the Query/AllEvidence gRPC method
func (k Keeper) AllEvidence(c context.Context, req *types.QueryAllEvidenceRequest) (*types.QueryAllEvidenceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
