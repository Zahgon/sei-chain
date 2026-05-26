package evmrpc

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/receipt"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

var errNoHeightSource = errors.New("unable to determine height information")

// ErrBlockHeightNotYetAvailable is returned when a concrete block height is above the
// node's safe latest watermark. eth_getBlockByNumber maps this to result null (Ethereum spec).
var ErrBlockHeightNotYetAvailable = errors.New("block height not yet available")

// WatermarkManager coordinates access to block, state, and receipt stores to
// determine queryable block heights for RPC consumers. It ensures read-side
// requests only target heights where all backing data sources are fully
// synchronized.
type WatermarkManager struct {
	tmClient     client.LocalClient
	ctxProvider  func(int64) sdk.Context
	stateStore   types.StateStore
	receiptStore receipt.ReceiptStore
}

func NewWatermarkManager(
	tmClient client.LocalClient,
	ctxProvider func(int64) sdk.Context,
	stateStore types.StateStore,
	receiptStore receipt.ReceiptStore,
) *WatermarkManager {
	_ = "STUB: not implemented"
	return nil
}

// Watermarks returns the earliest block height, earliest state height, and
// latest height that are safe to serve. Earliest heights are inclusive.
func (m *WatermarkManager) Watermarks(ctx context.Context) (int64, int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

// Tendermint heights govern both block availability and the latest safe height.

// State store heights (historical state DB) may lag behind block pruning.

// Receipt store version participates only in the latest watermark.

// LatestHeight returns the inclusive latest height guaranteed to have complete
// data.
func (m *WatermarkManager) LatestHeight(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EarliestHeight returns the earliest height that remains fully queryable.
func (m *WatermarkManager) EarliestHeight(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EarliestStateHeight returns the earliest height with state availability.
func (m *WatermarkManager) EarliestStateHeight(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ResolveHeight normalizes a requested block identifier into a concrete height.
// If the resolved height sits outside the tracked watermarks, the method returns
// an error explaining whether it is too old (pruned) or too new (not yet
// available).
func (m *WatermarkManager) ResolveHeight(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EnsureBlockHeightAvailable verifies that the provided block height falls within
// the computed watermarks.
func (m *WatermarkManager) EnsureBlockHeightAvailable(ctx context.Context, height int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *WatermarkManager) ensureWithinWatermarks(height, earliest, latest int64) error {
	_ = "STUB: not implemented"
	return nil
}

func blockByNumberRespectingWatermarks(
	ctx context.Context,
	client client.LocalClient,
	wm *WatermarkManager,
	heightPtr *int64,
	maxRetries int,
) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blockByHashRespectingWatermarks(
	ctx context.Context,
	client client.LocalClient,
	wm *WatermarkManager,
	hash []byte,
	maxRetries int,
) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *WatermarkManager) fetchTendermintWatermarks(ctx context.Context) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (m *WatermarkManager) fetchStateStoreWatermarks() (int64, int64, bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}
