// mux package provides a TCP connection multiplexer - it allows to run
// multiple reliable independent bidirectional streams over a single TCP connection:
// The data is sent in frames of bounded size in round robin fashion over all the streams
// (fairness). There is no head-of-line blocking: a sender is not allowed to send bytes,
// until peer allows it - a TCP-like buffer window is maintained: peer declares the
// maximal size of message it is willing to consume, and the number of messages it currently
// can buffer locally.
//
// Each mux stream has its own Kind number. Kind numbers are supposed to identify the stream-level communication
// protocol (for example, if you implement an RPC server on top of this multiplexer, each RPC will have its own Kind number).
//
// # LOW LEVEL PROTOCOL
//
// Multiplexer traffic consists of frames. Frame looks as follows:
// [header size (1B)] [header] [payload]
// Header is a binary protobuf message with size up to 255B (because size is sent as a single byte).
// This message is intentionally flat with a small number of fields (see mux.proto), so that it is encoded efficiently.
// In particular header contains PayloadSize field which indicates the size of the payload of the frame which is sent after the header.
// There are multiple frame types:
// * OPEN (opens a stream)
// * RESIZE (extends the window of the stream, allowing peer to send more messages)
// * MSG (actual payload of the stream)
// * CLOSE (closes the stream - this is the last frame of the stream)
package mux

import (
	"context"
	"errors"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/conn"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/mux/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

const handshakeMaxSize = 10 * 1024 // 10kB

var errUnknownStream = errors.New("frame for an unknown stream")
var errTooManyAccepts = errors.New("too many concurrent accepted streams")
var errFrameAfterClose = errors.New("received frame after CLOSE frame")
var errTooManyMsgs = errors.New("too many messages")
var errTooLargeMsg = errors.New("message too large")
var errUnknownKind = errors.New("unknown kind")
var errStreamKindMismatch = errors.New("stream kind mismatch")
var errAlreadyOpened = errors.New("stream already opened")
var errAlreadyClosed = errors.New("stream already closed")

type Config struct {
	// Maximal number of bytes in a frame (excluding header).
	FrameSize uint64
	// Limits on the number of concurrent streams of each kind.
	Kinds map[StreamKind]*StreamKindConfig
}

type StreamKindConfig struct {
	// Maximal number of concurrent outbound streams.
	MaxConnects uint64
	// Maximal number of concurrent inbound streams.
	MaxAccepts uint64
}

type handshake struct {
	Kinds map[StreamKind]*StreamKindConfig
}

var handshakeConv = protoutils.Conv[*handshake, *pb.Handshake]{
	Encode: func(h *handshake) *pb.Handshake {
		kinds := make([]*pb.StreamKindConfig, 0, len(h.Kinds))
		for kind, c := range h.Kinds {
			kinds = append(kinds, &pb.StreamKindConfig{
				Kind:        uint64(kind),
				MaxConnects: c.MaxConnects,
				MaxAccepts:  c.MaxAccepts,
			})
		}
		return &pb.Handshake{Kinds: kinds}
	},
	Decode: func(x *pb.Handshake) (*handshake, error) {
		kinds := map[StreamKind]*StreamKindConfig{}
		for _, pc := range x.Kinds {
			kinds[StreamKind(pc.Kind)] = &StreamKindConfig{
				MaxConnects: pc.MaxConnects,
				MaxAccepts:  pc.MaxAccepts,
			}
		}
		return &handshake{Kinds: kinds}, nil
	},
}

type frame struct {
	Header  *pb.Header
	Payload []byte
}

type kindState struct {
	connectsQueue chan *streamState
	acceptsQueue  chan *streamState
}

type runnerInner struct {
	nextID     streamID
	streams    map[streamID]*streamState
	acceptsSem map[StreamKind]uint64
}

// State of the running multiplexer.
type runner struct {
	mux   *Mux
	inner utils.RWMutex[*runnerInner]
}

func newRunner(mux *Mux) *runner { _ = "STUB: not implemented"; return nil }

// getOrAccept() gets the current state of the stream for the given header message.
// If the stream does not exist yet, it tries to create it as an accept (inbound) stream.
// In that case the inbound stream limit for the given kind is checked.
func (r *runner) getOrAccept(h *pb.Header) (*streamState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *runnerInner) newConnectStream(kind StreamKind) *streamState {
	_ = "STUB: not implemented"
	// Non-blocking since we just closed a connect Stream.
	return nil
}

func (r *runner) tryPrune(id streamID) { _ = "STUB: not implemented"; return }

// Check if the stream is fully closed.

// Delete stream state.

// Free the stream capacity.

// runSend handles the frame queue.
// The frames from all streams are interleaved in a round robin fashion.
// frames have bounded size to make sure that large messages do not slow down smaller ones.
// Stream priorities are not implemented (not needed).
// WARNING: it respects ctx only partially, because conn does not.
func (r *runner) runSend(ctx context.Context, conn conn.Conn) error {
	_ = "STUB: not implemented"

	// Collect frames in round robin over streams.
	return nil
}

// Send the frames

// Notify sender about local buffer capacity.

// TODO(gprusak): this is counterintuitive asymmetric behavior:
// * tryPrune in runRecv follows immediately remote close
// * tryPrune in runSend happens with a delay (local close happens in Stream.close() call).
// As a result runRecv might still be sending frames on behalf of a pruned stream.
// Ownership of r.inner.streams should be clearer.

// runRecv receives and processes the incoming frames sequentially.
func (r *runner) runRecv(ctx context.Context, conn conn.Conn) error {
	_ = "STUB: not implemented"

	// frame size is hard capped here at 255B.
	// Currently we have 7 varint fields (up to 77B)
	return nil
}

// Process the frame content in order: OPEN, RESIZE, MSG, CLOSE

// Read the payload.

// Run runs the multiplexer for the given connection.
// It closes the connection before return.
func (m *Mux) Run(ctx context.Context, conn conn.Conn) error { _ = "STUB: not implemented"; return nil }

// Handshake exchange.

//nolint:gosec // handshake size bounded by handshakeMaxSize

// Initialize runner with handshake data.

// Run the tasks.

// queue is a queue of frames to send, consumed by runSend.
type queue map[streamID]*frame

// Get returns the frame corresponding to the given stream id.
// If it doesn't exist, it initializes the frame first.
func (q queue) Get(id streamID) *frame { _ = "STUB: not implemented"; return nil }

// Pop removes a frame of the given stream from the queue.
// Panics if there is no frame for this id.
// If a frame is too large (payload larger than maxPayload) it splits
// the frame into 2 smaller ones and returns the first one.
func (q queue) Pop(id streamID, maxPayload uint64) *frame { _ = "STUB: not implemented"; return nil }

// Split the frame into first and second.

// Close and MsgEnd fields are left in the second frame.

// Clear the fields from the first frame.

type Mux struct {
	cfg   *Config
	kinds map[StreamKind]*kindState
	queue *utils.Watch[queue]
}

// NewMux constructs a new multipexer.
// Remember to spawn Mux.Run() afterwards.
func NewMux(cfg *Config) *Mux { _ = "STUB: not implemented"; return nil }

// Connect establishes a new stream of the given kind.
// Blocks until the number of concurrent connects falls below the allowed limit.
// Then it waits until peer accepts the connection.
// Remember to Close() the stream after use.
func (m *Mux) Connect(ctx context.Context, kind StreamKind, maxMsgSize uint64, window uint64) (*Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accept accepts an incoming stream of the given kind.
// Blocks until peer opens a connect stream.
// Remember to Close() the stream after use.
func (m *Mux) Accept(ctx context.Context, kind StreamKind, maxMsgSize uint64, window uint64) (*Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
