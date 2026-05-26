package giga

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/rpc"
)

func (x *Service) serverStreamLaneProposals(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) serverStreamLaneVotes(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) serverStreamAppVotes(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) serverStreamAppQCs(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) serverStreamCommitQCs(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) clientStreamLaneProposals(ctx context.Context, c rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(gprusak): dissemination of LaneProposals is the main source of bandwidth consumption.
// * to keep low latency, we need to push the lane proposals (streaming is required)
// * to avoid wasting bandwidth, we should set req.FirstBlockNumber (for that we need to authenticate validator in handshake)
// * the current implementation assumes a fully connected network - with a different topology we will need to be smarter.

// Sanity check, checking that the producer only sends their own proposals.
// TODO(gprusak): authenticate the peer to be able to do this check.
/*if got, want := proposal.Msg().Block().Header().Lane(), c.cfg.GetKey(); got != want {
	return fmt.Errorf("producer = %q, want %q", got, want)
}*/

func (x *Service) clientStreamLaneVotes(ctx context.Context, c rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) clientStreamCommitQCs(ctx context.Context, c rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) clientStreamAppVotes(ctx context.Context, c rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) clientStreamAppQCs(ctx context.Context, c rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}
