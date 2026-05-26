package consensus

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/libs/wal"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	tmcons "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/consensus"
)

const (
	// time.Time + max consensus msg size
	maxMsgSizeBytes = maxMsgSize + 24
)

type ErrBadSize struct{ error }

//--------------------------------------------------------
// types and functions for savings consensus messages

// MsgToProto takes a consensus message type and returns the proto defined
// consensus message.
func MsgToProto(msg Message) *tmcons.Message { _ = "STUB: not implemented"; return nil }

// MsgFromProto takes a consensus proto message and returns the native go type.
func MsgFromProto(msg *tmcons.Message) (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

// TimedWALMessage wraps WALMessage and adds Time for debugging purposes.
type TimedWALMessage struct {
	Time time.Time
	Msg  WALMessage
}

// EndHeightMessage marks the end of the given height inside WAL.
type EndHeightMessage struct {
	Height int64
}

type WALMessage struct{ any }

func NewWALMessage[T msgInfo | timeoutInfo | EndHeightMessage | types.EventDataRoundState](v T) WALMessage {
	_ = "STUB: not implemented"
	return *

	// WALtoProto takes a WAL message and return a proto walMessage and error.
	new(WALMessage)
}

func (msg WALMessage) toProto() *tmcons.WALMessage { _ = "STUB: not implemented"; return nil }

// walFromProto takes a proto wal message and return a consensus walMessage and
// error.
func walFromProto(msg *tmcons.WALMessage) (WALMessage, error) {
	_ = "STUB: not implemented"
	return *new(WALMessage), nil
}

// deny message based on possible overflow

//--------------------------------------------------------

// Write ahead logger writes msgs to disk before they are processed.
// Can be used for crash-recovery and deterministic replay.
type WAL struct{ inner utils.Mutex[*wal.Log] }

// OpenWAL opens WAL.
func OpenWAL(walFile string) (res *WAL, resErr error) { _ = "STUB: not implemented"; return nil, nil }

// For backward compatibility, we insert a marker in case WAL is empty.
// Current logic doesn't need it any more though.

// Sync flushes and fsync's the buffered entries to underlying files.
func (w *WAL) Sync() error { _ = "STUB: not implemented"; return nil }

// Close releases all underlying resources unconditionally.
// Other methods will return an error after calling Close.
func (w *WAL) Close() { _ = "STUB: not implemented"; return }

// Append appends an entry to the WAL.
// You need to call Sync afterwards to ensure entry is persisted on disk.
func (w *WAL) Append(msg WALMessage) error { _ = "STUB: not implemented"; return nil }

func walFromBytes(msgBytes []byte) (WALMessage, error) {
	_ = "STUB: not implemented"
	return *new(WALMessage), nil
}

// ReadLastHeightMsgs() - reads and returns all messages after the last EndHeightMessage marker.
// Returns height the messages belong to (i.e. EndHeightMessage.Height + 1)
// If WAL contains no marker, height 1 is assumed.
func (w *WAL) ReadLastHeightMsgs() (int64, []WALMessage, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Read files from the last, looking for the first EndHeightMessage marker.
