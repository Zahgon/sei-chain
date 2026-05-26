package p2p

import (
	"context"

	"github.com/gogo/protobuf/proto"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/conn"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type ChannelID = conn.ChannelID

type ChannelDescriptor[T proto.Message] = conn.ChannelDescriptorT[T]

// sendMsg is a message to be sent to a peer.
type sendMsg struct {
	Message   proto.Message // message payload
	ChannelID ChannelID     // channel id
}

// RecvMsg is a message received from a peer.
type RecvMsg[T proto.Message] struct {
	Message T            // message payload
	From    types.NodeID // sender
}

// PeerError is a peer error reported via Channel.Error.
//
// FIXME: This currently just disconnects the peer, which is too simplistic.
// For example, some errors should be logged, some should cause disconnects,
// and some should ban the peer.
//
// FIXME: This should probably be replaced by a more general PeerBehavior
// concept that can mark good and bad behavior and contributes to peer scoring.
// It should possibly also allow reactors to request explicit actions, e.g.
// disconnection or banning, in addition to doing this based on aggregates.
type PeerError struct {
	NodeID types.NodeID
	Err    error
	Fatal  bool
}

func (pe PeerError) Error() string { _ = "STUB: not implemented"; return "" }
func (pe PeerError) Unwrap() error {
	_ = "STUB: not implemented"

	// channel is a bidirectional channel to exchange Protobuf messages with peers.
	return nil
}

type channel struct {
	desc      conn.ChannelDescriptor
	recvQueue *Queue[RecvMsg[proto.Message]] // inbound messages (peers to reactors)
}

type Channel[T proto.Message] struct {
	*channel
	router *Router
}

// NewChannel creates a new channel. It is primarily for internal and test
// use, reactors should use Router.OpenChannel().
func newChannel(desc conn.ChannelDescriptor) *channel { _ = "STUB: not implemented"; return nil }

// TODO(gprusak): get rid of this random cap*cap value once we understand
// what the sizes per channel really should be.

func (ch *Channel[T]) send(msg T, queues ...*Queue[sendMsg]) { _ = "STUB: not implemented"; return }

func (ch *Channel[T]) Send(msg T, to types.NodeID) { _ = "STUB: not implemented"; return }

// reactor tried to send a message across a channel that the
// peer doesn't have available. This is a known issue due to
// how peer subscriptions work:
// https://github.com/tendermint/tendermint/issues/6598

// Broadcasts msg to all peers on the channel.
func (ch *Channel[T]) Broadcast(msg T) { _ = "STUB: not implemented"; return }

func (ch *Channel[T]) String() string { _ = "STUB: not implemented"; return "" }

func (ch *Channel[T]) ReceiveLen() int { _ = "STUB: not implemented"; return 0 }

// Recv Receives the next message from the channel.
func (ch *Channel[T]) Recv(ctx context.Context) (RecvMsg[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
