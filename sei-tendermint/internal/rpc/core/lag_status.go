package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// LagStatus returns Tendermint lag status, if lag is over a certain threshold
func (env *Environment) LagStatus(ctx context.Context) (*coretypes.ResultLagStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate lag

// Return a response with error code to differentiate the lagging status by http response code
