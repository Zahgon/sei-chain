package main

import (
	"context"

	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
)

// Wait waits for a number of blocks to be produced, and for all nodes to catch
// up with it.
func Wait(ctx context.Context, testnet *e2e.Testnet, blocks int64) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitUntil waits until a given height has been reached.
func WaitUntil(ctx context.Context, testnet *e2e.Testnet, height int64) error {
	_ = "STUB: not implemented"
	return nil
}
