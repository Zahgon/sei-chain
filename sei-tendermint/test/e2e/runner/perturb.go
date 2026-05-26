package main

import (
	"context"

	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
)

// Perturbs a running testnet.
func Perturb(ctx context.Context, testnet *e2e.Testnet) error {
	_ = "STUB: not implemented"
	return nil
	// first tick fires immediately; reset below
}

// give network some time to recover between each

// PerturbNode perturbs a node with a given perturbation, returning its status
// after recovering.
func PerturbNode(ctx context.Context, node *e2e.Node, perturbation e2e.Perturbation) (*rpctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Seed nodes do not have an RPC endpoint exposed so we cannot assert that
// the node recovered. All we can do is hope.
