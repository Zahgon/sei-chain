package mux

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type streamID uint64
type StreamKind uint64

func (id streamID) isConnect() bool {
	_ = "STUB: not implemented"

	// The least significant bit of streamID decides whether the stream is
	// outbound (connect) or inbound (accept). When receiving streamID from peer
	// we need to convert it to local streamID.
	return false
}

func streamIDFromRemote(x uint64) streamID { _ = "STUB: not implemented"; return *new(streamID) }

type closeState struct {
	local  bool
	remote bool
}

type sendState struct {
	remoteOpened bool
	maxMsgSize   uint64
	bufBegin     uint64
	begin        uint64
	end          uint64
}

type recvState struct {
	opened     bool
	maxMsgSize uint64
	begin      uint64
	used       uint64
	end        uint64
	msgs       [][]byte
}

type streamStateInner struct {
	send   sendState
	recv   recvState
	closed closeState
}

type streamState struct {
	id    streamID
	kind  StreamKind
	inner utils.Watch[*streamStateInner]
}

func newStreamState(id streamID, kind StreamKind) *streamState {
	_ = "STUB: not implemented"
	return nil
}

func (s *streamState) RemoteOpen(maxMsgSize uint64) error { _ = "STUB: not implemented"; return nil }

// Do not allow remote open before we connect.

func (s *streamState) RemoteClose() error { _ = "STUB: not implemented"; return nil }

func (s *streamState) RemoteWindowEnd(windowEnd uint64) { _ = "STUB: not implemented"; return }

// RemotePayloadSize checks if there is place for the payload.
func (s *streamState) RemotePayloadSize(payloadSize uint64) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // recv.used is bounded by recv.end which is bounded by len(recv.msgs)

func (s *streamState) RemotePayload(payload []byte) { _ = "STUB: not implemented"; return }

//nolint:gosec // recv.used is bounded by recv.end which is bounded by len(recv.msgs)

func (s *streamState) RemoteMsgEnd() error { _ = "STUB: not implemented"; return nil }
