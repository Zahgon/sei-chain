package evmrpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type traceExecutionPhaseDurations struct {
	PrepareTxNanos   int64 `json:"prepareTxNanos"`
	ExecutionNanos   int64 `json:"executionNanos"`
	TraceResultNanos int64 `json:"traceResultNanos"`
}

type TraceTransactionProfilePhases struct {
	LookupTransactionNanos   int64 `json:"lookupTransactionNanos"`
	LoadBlockNanos           int64 `json:"loadBlockNanos"`
	ReplayHistoricalTxsNanos int64 `json:"replayHistoricalTxsNanos"`
	BuildBlockContextNanos   int64 `json:"buildBlockContextNanos"`
	traceExecutionPhaseDurations
}

type TraceTransactionProfile struct {
	TotalNanos              int64                         `json:"totalNanos"`
	HistoricalDBLookupNanos int64                         `json:"historicalDbLookupNanos"`
	OtherNanos              int64                         `json:"otherNanos"`
	Phases                  TraceTransactionProfilePhases `json:"phases"`
	Store                   *sdk.StoreTraceDump           `json:"store,omitempty"`
}

type TraceTransactionProfileResponse struct {
	Trace   interface{}             `json:"trace"`
	Profile TraceTransactionProfile `json:"profile"`
}

func (api *DebugAPI) TraceTransactionProfile(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

//nolint:gosec

func (api *DebugAPI) newProfileTracingBackend() *Backend { _ = "STUB: not implemented"; return nil }

func dumpStoreTrace(statedb vm.StateDB) *sdk.StoreTraceDump { _ = "STUB: not implemented"; return nil }

func historicalLookupNanos(storeDump *sdk.StoreTraceDump) int64 {
	_ = "STUB: not implemented"
	return 0
}
