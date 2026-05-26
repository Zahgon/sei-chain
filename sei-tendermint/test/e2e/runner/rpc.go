package main

import (
	"context"

	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// waitForHeight waits for the network to reach a certain height (or above),
// returning the block at the height seen. Errors if the network is not making
// progress at all.
// If height == 0, the initial height of the test network is used as the target.
func waitForHeight(ctx context.Context, testnet *e2e.Testnet, height int64) (*types.Block, *types.BlockID, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// skip nodes that have reached the target height

// skip nodes that don't have state or haven't started yet

// cache the clients

// the node has achieved the target height!

// add this node to the set of target
// height nodes

// if not all of the nodes that we
// have clients for have reached the
// target height, keep trying.

// All nodes are at or above the target height. Now fetch the block for that target height
// and return it. We loop again through all clients because some may have pruning set but
// at least two of them should be archive nodes.

// waitForNode waits for a node to become available and catch up to the given block height.
func waitForNode(ctx context.Context, node *e2e.Node, height int64) (*rpctypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	// If the node is the light client or seed note, we do not check for the last height.
	// The light client and seed note can be behind the full node and validator
	return nil, nil
}

// if there was a problem with the request in
// the previous recreate the client to ensure
// reconnection

// If the node is the light client, it is not essential to wait for it to catch up, but we must return status info

// getLatestBlock returns the last block that all active nodes in the network have
// agreed upon i.e. the earlist of each nodes latest block
func getLatestBlock(ctx context.Context, testnet *e2e.Testnet) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip nodes that don't have state or haven't started yet
