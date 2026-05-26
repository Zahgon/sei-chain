package conn

import (
	"context"
	"errors"

	"github.com/gogo/protobuf/proto"
)

var errMsgTooLarge = errors.New("message too large")

// Writes size-prefixed proto message.
func WriteSizedMsg(ctx context.Context, conn Conn, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Reads size-prefixed proto message.
// It is slow, because size is encoded as varint, and therefore needs to be read byte at a time.
func ReadSizedMsg(ctx context.Context, conn Conn, maxSize uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshals length prefixed message.
// Length is encoded as varint.
func UnmarshalSizedProto(data []byte, msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // n is positive (checked above); len(data) is always non-negative

//nolint:gosec // size is bounded by len(data) which fits in int
