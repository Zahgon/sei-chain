package conn

import (
	"context"
	"errors"
	"net/netip"
	"sync/atomic"
	"time"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/sei-protocol/seilog"
	"golang.org/x/time/rate"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

var logger = seilog.NewLogger("tendermint", "internal", "p2p", "conn")

// ChannelID is an arbitrary channel ID.
type ChannelID uint16

type ChannelDescriptor = ChannelDescriptorT[gogoproto.Message]

type ChannelDescriptorT[T gogoproto.Message] struct {
	ID       ChannelID
	Priority int

	MessageType T

	// PreDecode, if Some, runs on the raw wire bytes of an inbound message
	// before gogoproto.Unmarshal so a channel can enforce size or shape
	// invariants before decoding. It should be a cheap, bounded check.
	// Returning a non-nil error drops the message and evicts the sending peer.
	PreDecode utils.Option[func([]byte) error]

	// TODO: Remove once p2p refactor is complete.
	SendQueueCapacity   int
	RecvMessageCapacity int

	// RecvBufferCapacity defines the max buffer size of inbound messages for a
	// given p2p Channel queue.
	RecvBufferCapacity int

	// Human readable name of the channel, used in logging and
	// diagnostics.
	Name string
}

func (chDesc ChannelDescriptorT[T]) ToGeneric() ChannelDescriptor {
	_ = "STUB: not implemented"
	return *new(ChannelDescriptor)
}

func (chDesc ChannelDescriptorT[T]) withDefaults() ChannelDescriptorT[T] {
	_ = "STUB: not implemented"
	return nil
}

// 21MB

var errPongTimeout = errors.New("pong timeout")

type errBadEncoding struct{ error }
type errBadChannel struct{ error }

// mConnMessage passes MConnection messages through internal channels.
type mConnMessage struct {
	channelID ChannelID
	payload   []byte
}

/*
Each peer has one `MConnection` (multiplex connection) instance.

__multiplex__ *noun* a system or signal involving simultaneous transmission of
several messages along a single channel of communication.

Each `MConnection` handles message transmission on multiple abstract communication
`Channel`s.  Each channel has a globally unique byte id.
The byte id and the relative priorities of each `Channel` are configured upon
initialization of the connection.
*/
type MConnection struct {
	conn      Conn
	sendQueue utils.Watch[*sendQueue]
	recvPong  utils.Mutex[*utils.AtomicSend[bool]]
	recvCh    chan mConnMessage
	config    MConnConfig
}

// MConnConfig is a MConnection configuration.
type MConnConfig struct {
	SendRate                int64         // B/s
	RecvRate                int64         // B/s
	MaxPacketMsgPayloadSize int           // Maximum payload size
	FlushThrottle           time.Duration // Interval to flush writes (throttled)
	PingInterval            time.Duration // Interval to send pings
	PongTimeout             time.Duration // Time to wait for a pong
}

func (c *MConnConfig) getSendRateLimit() rate.Limit {
	_ = "STUB: not implemented"
	return *new(rate.Limit)
}

func (c *MConnConfig) getRecvRateLimit() rate.Limit {
	_ = "STUB: not implemented"
	return *new(rate.Limit)
}

// DefaultMConnConfig returns the default config.
func DefaultMConnConfig() MConnConfig {
	_ = "STUB: not implemented"

	// TODO(gprusak): RecvRate should be strictly larger than SendRate,
	// so that under maximal load the backpressure is at the sender.
	return *new(MConnConfig)
}

// 500KB/s
// 500KB/s
// mirrors MaxPacketMsgPayloadSize from config/config.go

type sendQueue struct {
	ping  bool
	pong  bool
	flush utils.Option[time.Time]
	// TODO(gprusak): restrict to channels that peer knows about
	channels map[ChannelID]*sendChannel
}

func newSendQueue(chDescs []*ChannelDescriptor) *sendQueue { _ = "STUB: not implemented"; return nil }

func (q *sendQueue) setFlush(t time.Time) { _ = "STUB: not implemented"; return }

// NewMConnection wraps net.Conn and creates multiplex connection with a config
func NewMConnection(
	conn Conn,
	chDescs []*ChannelDescriptor,
	config MConnConfig,
) *MConnection {
	_ = "STUB: not implemented"
	return nil
}

func (c *MConnection) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *MConnection) LocalAddr() netip.AddrPort {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort)
}
func (c *MConnection) RemoteAddr() netip.AddrPort {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort)
}
func (c *MConnection) Close() { _ = "STUB: not implemented"; return }

// String returns a safe, concise representation of the connection.
// This prevents the race caused by slog/fmt reflecting over mutable fields
// (such as recvPong) when MConnection is passed as a log value.
func (c *MConnection) String() string { _ = "STUB: not implemented"; return "" }

// Queues a message to be sent.
// WARNING: takes ownership of msgBytes
// TODO(gprusak): fix the ownership
func (c *MConnection) Send(ctx context.Context, chID ChannelID, msgBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Recv .
func (c *MConnection) Recv(ctx context.Context) (ChannelID, []byte, error) {
	_ = "STUB: not implemented"
	return *new(ChannelID), nil, nil
}

func (c *MConnection) recvPongSubscribe() utils.AtomicRecv[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (c *MConnection) pingRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"

	// Send ping.
	return nil
}

// Wait for pong.

// Sleep.

func (c *MConnection) statsRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Exponential decay of stats.
// TODO(gprusak): This is not atomic at all.

// popSendQueue pops a message from the send queue.
// Returns nil,nil if the connection should be flushed.
func (c *MConnection) popSendQueue(ctx context.Context) (*pb.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Choose a channel to create a PacketMsg from.
// The chosen channel will be the one whose recentlySent/priority is the least.

// It is flush time!

// sendRoutine polls for packets to send from channels.
func (c *MConnection) sendRoutine(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // burst size is bounded by config values; no overflow risk

// Marshalling is expected to always succeed.

// Here we ignore the fact that writing sized msg actually writes extra bytes to express size.

// recvRoutine receives messages and pushes them to recvCh.
// It also handles ping/pong messages.
func (c *MConnection) recvRoutine(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // burst size is bounded by config values; no overflow risk

// maxPacketMsgSize returns a maximum size of PacketMsg
func (c *MConnection) maxPacketMsgSize() uint64 { _ = "STUB: not implemented"; return 0 }

type sendChannel struct {
	desc         ChannelDescriptor
	recentlySent atomic.Uint64 // Exponential moving average.
	queue        utils.RingBuf[*[]byte]
}

func (ch *sendChannel) ratio() float32 { _ = "STUB: not implemented"; return 0 }

// Creates a new PacketMsg to send.
// Not goroutine-safe
func (ch *sendChannel) popMsg(maxPayload int) *pb.PacketMsg { _ = "STUB: not implemented"; return nil }

type recvChannel struct {
	desc ChannelDescriptor
	buf  []byte
}

func newRecvChannel(desc ChannelDescriptor) *recvChannel { _ = "STUB: not implemented"; return nil }

// Handles incoming PacketMsgs. It returns a message bytes if message is
// complete, which is owned by the caller and will not be modified.
// Not goroutine-safe
func (ch *recvChannel) pushMsg(packet *pb.PacketMsg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
