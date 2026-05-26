package keeper

import (
	protoio "github.com/gogo/protobuf/io"
	snapshot "github.com/sei-protocol/sei-chain/sei-cosmos/snapshots/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var (
	logger = seilog.NewLogger("wasmd", "x", "wasm", "keeper")

	_ snapshot.ExtensionSnapshotter = &WasmSnapshotter{}
)

// SnapshotFormat format 1 is just gzipped wasm byte code for each item payload. No protobuf envelope, no metadata.
const SnapshotFormat = 1

type WasmSnapshotter struct {
	wasm *Keeper
	cms  sdk.MultiStore
}

func NewWasmSnapshotter(cms sdk.MultiStore, wasm *Keeper) *WasmSnapshotter {
	_ = "STUB: not implemented"
	return nil
}

func (ws *WasmSnapshotter) SnapshotName() string { _ = "STUB: not implemented"; return "" }

func (ws *WasmSnapshotter) SnapshotFormat() uint32 { _ = "STUB: not implemented"; return 0 }

func (ws *WasmSnapshotter) SupportedFormats() []uint32 {
	_ = "STUB: not implemented"
	// If we support older formats, add them here and handle them in Restore
	return nil
}

func (ws *WasmSnapshotter) Snapshot(height uint64, protoWriter protoio.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// Many code ids may point to the same code hash... only sync it once

// if seenBefore, just skip this one and move to the next

// load code and abort on error

func (ws *WasmSnapshotter) Restore(
	height uint64, format uint32, protoReader protoio.Reader,
) (snapshot.SnapshotItem, error) {
	_ = "STUB: not implemented"
	return *new(snapshot.SnapshotItem), nil
}

func restoreV1(ctx sdk.Context, k *Keeper, compressedCode []byte, num int) error {
	_ = "STUB: not implemented"
	// #nosec G115 -- MaxWasmSize is a constant and always non-negative
	return nil
}

// FIXME: check which codeIDs the checksum matches??

func finalizeV1(ctx sdk.Context, k *Keeper) error { _ = "STUB: not implemented"; return nil }

func (ws *WasmSnapshotter) processAllItems(
	height uint64,
	protoReader protoio.Reader,
	cb func(sdk.Context, *Keeper, []byte, int) error,
	finalize func(sdk.Context, *Keeper) error,
) (snapshot.SnapshotItem, error) {
	_ = "STUB: not implemented"
	return *new(snapshot.SnapshotItem), nil
}

// #nosec G115 -- height is bounds checked above

// keep the last item here... if we break, it will either be empty (if we hit io.EOF)
// or contain the last item (if we hit payload == nil)

// if it is not another ExtensionPayload message, then it is not for us.
// we should return it an let the manager handle this one
