package mux

import (
	"context"
	"errors"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

var errRemoteClosed = errors.New("remote closed")
var errClosed = errors.New("closed")

type Stream struct {
	state *streamState
	queue *utils.Watch[queue]
}

func (s *Stream) maxSendMsgSize() uint64 { _ = "STUB: not implemented"; return 0 }

// open() opens the recv end of the Stream. Permits the peer to send "window" messages, up to maxMsgSize bytes each.
// Up to maxMsgSize*window bytes will be cached locally during the life of this Stream.
// Whenever you call Recv, you specify whether window should grow (i.e. whether to report that the messages
// have been consumed and he can send more).
func (s *Stream) open(ctx context.Context, maxMsgSize uint64, window uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Send sends a message to peer. Blocks until:
// * peer has permitted to send them a message (i.e. there is space in their local buffer)
// * the previous message has been sent by the multiplexer (at most 1 message per Stream is cached at all times)
// Returns an error if Close() was called already.
// Returns an error if the message is too large (exceeds maxMsgSize declared by the peer).
func (s *Stream) Send(ctx context.Context, msg []byte) error { _ = "STUB: not implemented"; return nil }

// Wait until the local buffer is empty && remote buffer has capacity.

// Will we never be able to send...

// ...or we can send now.

// We check msg size AFTER waiting because maxMsgSize could be set AFTER we wait.

// Push msg to the queue.

func (s *Stream) close(inner *streamStateInner) { _ = "STUB: not implemented"; return }

// Close sends a final CLOSE flag to the peer.
// All subsequent Send calls will fail.
// Recv calls will no longer be able to free buffer space.
// NOTE: we may consider separating Close into SendClose and RecvClose,
// to make send and recv parts of the stream entirely independent.
func (s *Stream) Close() { _ = "STUB: not implemented"; return }

// Recv receives a message from peer. Blocks until message is available OR
// until peer has closed their end of the Stream.
// If freeBuffer is set, it permits the peer to send more messages (since local buffer was freed).
func (s *Stream) Recv(ctx context.Context, freeBuffer bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A message is available or peer closed the Stream.

// Free buffer if requested AND the stream was not closed locally.
