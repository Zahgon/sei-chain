package evmrpc

import (
	"context"
	"time"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/eth/tracers/tracersutils"
	"github.com/ethereum/go-ethereum/rpc"
)

const profiledDefaultTraceTimeout = 5 * time.Second
const profiledDefaultTraceReexec = uint64(128)
const maxProfiledTraceWorkers = 16

func (api *DebugAPI) shouldUseProfiledBlockTrace(config *tracers.TraceConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (api *DebugAPI) profiledTraceBlockByNumber(ctx context.Context, number rpc.BlockNumber, config *tracers.TraceConfig) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) profiledTraceBlockByHash(ctx context.Context, hash gethcommon.Hash, config *tracers.TraceConfig) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) profiledTraceBlock(
	ctx context.Context,
	block *gethtypes.Block,
	metadata []tracersutils.TraceBlockMetadata,
	config *tracers.TraceConfig,
) ([]*tracers.TxTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) profiledTraceBlockSequential(
	ctx context.Context,
	block *gethtypes.Block,
	metadata []tracersutils.TraceBlockMetadata,
	config *tracers.TraceConfig,
	statedb vm.StateDB,
	blockCtx vm.BlockContext,
	signer gethtypes.Signer,
	blockHash gethcommon.Hash,
	results []*tracers.TxTraceResult,
) ([]*tracers.TxTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type profiledTxTraceTask struct {
	index   int
	statedb vm.StateDB
}

func (api *DebugAPI) profiledTraceBlockParallel(
	ctx context.Context,
	block *gethtypes.Block,
	metadata []tracersutils.TraceBlockMetadata,
	config *tracers.TraceConfig,
	statedb vm.StateDB,
	signer gethtypes.Signer,
	blockHash gethcommon.Hash,
	results []*tracers.TxTraceResult,
	threads int,
) ([]*tracers.TxTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fill error entries for txs that were never dispatched to workers,
// matching the sequential path's per-tx error semantics.

func (api *DebugAPI) profiledTraceTx(
	ctx context.Context,
	tx *gethtypes.Transaction,
	message *core.Message,
	txctx *tracers.Context,
	vmctx vm.BlockContext,
	statedb vm.StateDB,
	config *tracers.TraceConfig,
	precompiles vm.PrecompiledContracts,
	noFlush bool,
	phaseDurations *traceExecutionPhaseDurations,
) (value interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func profiledErrorTrace(err error, tx *gethtypes.Transaction, message *core.Message, txctx *tracers.Context, vmctx vm.BlockContext, config *tracers.TraceConfig) (value interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}
