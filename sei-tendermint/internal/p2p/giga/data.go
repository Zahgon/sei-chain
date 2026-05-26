package giga

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/rpc"
)

func (s *Service) clientStreamFullCommitQCs(ctx context.Context, client rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add DoS protection (i.e. that only useful state.Data() has been actually sent).

// MaxConcurrentBlockFetches is the maximum number of blocks that client fetches concurrently.
const MaxConcurrentBlockFetches = 100

// BlockFetchTimeout after which the block fetch RPC is considered failed and needs to be retried.
const BlockFetchTimeout = 2 * time.Second

type req struct {
	n    types.GlobalBlockNumber
	done chan struct{}
}

func (s *Service) clientGetBlock(ctx context.Context, client rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) runBlockFetcher(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Wait for the QC.

func (s *Service) serverStreamFullCommitQCs(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't send the same QC twice.

func (x *Service) serverGetBlock(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}
