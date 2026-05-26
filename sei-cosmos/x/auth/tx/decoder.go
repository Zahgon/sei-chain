package tx

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
)

// DefaultTxDecoder returns a default protobuf TxDecoder using the provided Marshaler.
func DefaultTxDecoder(cdc codec.ProtoCodecMarshaler) sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

// DefaultTxDecoderWithoutBodyBloatRejection returns a protobuf TxDecoder that
// preserves pre-v6.5 decode behavior for historical tooling. Do not use this for
// mempool, CheckTx, or DeliverTx paths.
func DefaultTxDecoderWithoutBodyBloatRejection(cdc codec.ProtoCodecMarshaler) sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

func defaultTxDecoder(cdc codec.ProtoCodecMarshaler, rejectBodyBloat bool) sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

// Make sure txBytes follow ADR-027.

// reject all unknown proto fields in the root TxRaw

// allow non-critical unknown fields in TxBody

// reject all unknown proto fields in AuthInfo

// DefaultJSONTxDecoder returns a default protobuf JSON TxDecoder using the provided Marshaler.
func DefaultJSONTxDecoder(cdc codec.ProtoCodecMarshaler) sdk.TxDecoder {
	_ = "STUB: not implemented"
	return *new(sdk.TxDecoder)
}

// rejectNonADR027TxRaw rejects txBytes that do not follow ADR-027. This is NOT
// a generic ADR-027 checker, it only applies decoding TxRaw. Specifically, it
// only checks that:
// - field numbers are in ascending order (1, 2, and potentially multiple 3s),
// - and varints are as short as possible.
// All other ADR-027 edge cases (e.g. default values) are not applicable with
// TxRaw.
func rejectNonADR027TxRaw(txBytes []byte) error {
	_ = "STUB: not implemented"
	// Make sure all fields are ordered in ascending order with this variable.
	return nil
}

// TxRaw only has bytes fields.

// Make sure fields are ordered in ascending order.

// All 3 fields of TxRaw have wireType == 2, so their next component
// is a varint, so we can safely call ConsumeVarint here.
// Byte structure: <varint of bytes length><bytes sequence>
// Inner  fields are verified in `DefaultTxDecoder`

// We make sure that this varint is as short as possible.

// Skip over the bytes that store fieldNumber and wireType bytes.

// rejectBloatedBody rejects tx bodies where the raw wire encoding is larger
// than the canonical re-marshal of the decoded struct. This catches protobuf-level
// bloat (e.g. padded sdk.Int fields, oversized Any.Value) that UnpackAny would
// otherwise silently canonicalize away before validation runs.
func rejectBloatedBody(rawBodyBytes []byte, body *tx.TxBody) error {
	_ = "STUB: not implemented"
	return nil
}

// varintMinLength returns the minimum number of bytes necessary to encode an
// uint using varint encoding.
func varintMinLength(n uint64) int {
	_ = "STUB: not implemented"

	// Note: 1<<N == 2**N.
	return 0
}
