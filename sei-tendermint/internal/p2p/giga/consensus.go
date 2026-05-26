package giga

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/rpc"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// Sends a consensus message to the peer whenever atomic watch is updated.
func sendUpdates[T interface {
	comparable
	types.ConsensusReq
}](
	ctx context.Context,
	client rpc.Client[API],
	w utils.AtomicRecv[utils.Option[T]],
) error {
	_ = "STUB: not implemented"
	return nil
}

const pingInterval = 10 * time.Second
const pingTimeout = 5 * time.Second

// sendPings periodically sends Ping messages.
func (x *Service) clientPing(ctx context.Context, client rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(gprusak): add random payload to actually verify roundtrip latency.

//

// Run sends newest consensus messages to the peer.
func (x *Service) clientConsensus(ctx context.Context, c rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}

// Send updates about new consensus messages.

// Ping implements pb.StreamAPIServer.
// Note that we use streaming RPC, because unary RPC apparently causes 10ms extra delay on avg (empirically tested).
func (x *Service) serverPing(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

// Consensus implements pb.StreaAPIServer.
func (x *Service) serverConsensus(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}
