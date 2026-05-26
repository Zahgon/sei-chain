package main

import (
	"context"
	"math/rand"
	"time"

	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// Load generates transactions against the network until the given context is
// canceled.
func Load(ctx context.Context, r *rand.Rand, testnet *e2e.Testnet) error {
	_ = "STUB: not implemented"
	// Since transactions are executed across all nodes in the network, we need
	// to reduce transaction load for larger networks to avoid using too much
	// CPU. This gives high-throughput small networks and low-throughput large ones.
	// This also limits the number of TCP connections, since each worker has
	// a connection to all nodes.
	return nil
}

// success counts per iteration

// Spawn job generator and processors.

// Montior transaction to ensure load propagates to the network
//
// This loop doesn't check or time out for stalls, since a stall here just
// aborts the load generator sooner and could obscure backpressure
// from the test harness, and there are other checks for
// stalls in the framework. Ideally we should monitor latency as a guide
// for when to give up, but we don't have a good way to track that yet.

// TODO perhaps allow test networks to
// declare required transaction rates, which
// might allow us to avoid the special case
// around 0 txs above.

// loadGenerate generates jobs until the context is canceled.
//
// The chTx has multiple consumers, thus the rate limiting of the load
// generation is primarily the result of backpressure from the
// broadcast transaction, though there is still some timer-based
// limiting.
func loadGenerate(ctx context.Context, r *rand.Rand, chTx chan<- types.Tx, txSize int, networkSize int) {
	_ = "STUB: not implemented"
	return
}

// Constrain the key space to avoid using too much
// space, while reduce the size of the data in the app.

// sleep for a bit before sending the
// next transaction.

func loadGenerateWaitTime(r *rand.Rand, size int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// loadProcess processes transactions
func loadProcess(ctx context.Context, testnet *e2e.Testnet, chTx <-chan types.Tx, chSuccess chan<- int) {
	_ = "STUB: not implemented"
	// Each worker gets its own client to each usable node, which
	// allows for some concurrency while still bounding it.
	return
}

// Construct a list of usable nodes for the creating
// load. Don't send load through seed nodes because
// they do not provide the RPC endpoints required to
// broadcast transaction.

// Put the clients in a ring so they can be used in a
// round-robin fashion.

// reset counter for the next iteration
