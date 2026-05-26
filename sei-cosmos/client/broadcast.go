package client

import (
	"context"

	"github.com/pkg/errors"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
)

var ErrTxInCache = errors.New("tx already exists in cache")

type ErrTxTooLarge struct {
	Max    int
	Actual int
}

func (e ErrTxTooLarge) Error() string { _ = "STUB: not implemented"; return "" }

type ErrMempoolIsFull struct {
	NumTxs      int
	MaxTxs      int
	TxsBytes    int64
	MaxTxsBytes int64
}

func (e ErrMempoolIsFull) Error() string { _ = "STUB: not implemented"; return "" }

// BroadcastTx broadcasts a transactions either synchronously or asynchronously. The result of the broadcast is parsed into
// an intermediate structure which is logged if the context has a logger
// defined.
func BroadcastTx(ctx context.Context, node Client, broadcastMode string, txBytes []byte) (*sdk.TxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckTendermintError checks if the error returned from BroadcastTx is a
// Tendermint error that is returned before the tx is submitted due to
// precondition checks that failed. If an Tendermint error is detected, this
// function returns the correct code back in TxResponse.
//
// TODO: Avoid brittle string matching in favor of error matching. This requires
// a change to Tendermint's RPCError type to allow retrieval or matching against
// a concrete error type.
func CheckTendermintError(err error, tx tmtypes.Tx) *sdk.TxResponse {
	_ = "STUB: not implemented"
	return nil
}

// TxServiceBroadcast is a helper function to broadcast a Tx with the correct gRPC types
// from the tx service. Calls `clientCtx.BroadcastTx` under the hood.
func TxServiceBroadcast(grpcCtx context.Context, node Client, req *tx.BroadcastTxRequest) (*tx.BroadcastTxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// normalizeBroadcastMode converts a broadcast mode into a normalized string
// to be passed into the clientCtx.
func normalizeBroadcastMode(mode tx.BroadcastMode) string { _ = "STUB: not implemented"; return "" }
