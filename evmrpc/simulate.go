package evmrpc

import (
	"context"
	"math/big"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/eth/tracers/tracersutils"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/export"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/app/legacyabci"
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type CtxIsWasmdPrecompileCallKeyType string

const CtxIsWasmdPrecompileCallKey CtxIsWasmdPrecompileCallKeyType = "CtxIsWasmdPrecompileCallKey"

type SimulationAPI struct {
	backend        *Backend
	connectionType ConnectionType
	requestLimiter *semaphore.Weighted
}

func NewSimulationAPI(
	ctxProvider func(int64) sdk.Context,
	keeper *keeper.Keeper,
	beginBlockKeepers legacyabci.BeginBlockKeepers,
	txConfigProvider func(int64) client.TxConfig,
	tmClient client.LocalClient,
	config *SimulateConfig,
	app *baseapp.BaseApp,
	antehandler sdk.AnteHandler,
	connectionType ConnectionType,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
	watermarks *WatermarkManager,
) *SimulationAPI {
	_ = "STUB: not implemented"
	return nil
}

type AccessListResult struct {
	Accesslist *ethtypes.AccessList `json:"accessList"`
	Error      string               `json:"error,omitempty"`
	GasUsed    hexutil.Uint64       `json:"gasUsed"`
}

func (s *SimulationAPI) CreateAccessList(ctx context.Context, args export.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash) (result *AccessListResult, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SimulationAPI) EstimateGas(ctx context.Context, args export.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *export.StateOverride) (result hexutil.Uint64, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Uint64), nil
}

/* ---------- fail‑fast limiter ---------- */

func (s *SimulationAPI) EstimateGasAfterCalls(ctx context.Context, args export.TransactionArgs, calls []export.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *export.StateOverride) (result hexutil.Uint64, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Uint64), nil
}

/* ---------- fail‑fast limiter ---------- */

func (s *SimulationAPI) Call(ctx context.Context, args export.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *export.StateOverride, blockOverrides *export.BlockOverrides) (result hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

/* ---------- fail‑fast limiter ---------- */

// If the result contains a revert reason, try to unpack and return it.

func NewRevertError(result *core.ExecutionResult) *RevertError {
	_ = "STUB: not implemented"
	return nil
}

// RevertError is an API error that encompasses an EVM revertal with JSON error
// code and a binary data blob.
type RevertError struct {
	error
	reason string // revert reason hex encoded
}

// ErrorCode returns the JSON error code for a revertal.
// See: https://github.com/ethereum/wiki/wiki/JSON-RPC-Error-Codes-Improvement-Proposal
func (e *RevertError) ErrorCode() int {
	_ = "STUB: not implemented"

	// ErrorData returns the hex encoded revert reason.
	return 0
}

func (e *RevertError) ErrorData() interface{} { _ = "STUB: not implemented"; return nil }

type SimulateConfig struct {
	GasCap                       uint64
	EVMTimeout                   time.Duration
	MaxConcurrentSimulationCalls int
}

var _ tracers.Backend = (*Backend)(nil)

type Backend struct {
	*eth.EthAPIBackend
	ctxProvider        func(int64) sdk.Context
	traceCtxProvider   TraceContextProvider
	txConfigProvider   func(int64) client.TxConfig
	keeper             *keeper.Keeper
	tmClient           client.LocalClient
	config             *SimulateConfig
	app                *baseapp.BaseApp
	beginBlockKeepers  legacyabci.BeginBlockKeepers
	antehandler        sdk.AnteHandler
	globalBlockCache   BlockCache
	cacheCreationMutex *sync.Mutex
	watermarks         *WatermarkManager
}

type TraceContextProvider func(int64) (sdk.Context, func())

func NewBackend(
	ctxProvider func(int64) sdk.Context,
	keeper *keeper.Keeper,
	beginBlockKeepers legacyabci.BeginBlockKeepers,
	txConfigProvider func(int64) client.TxConfig,
	tmClient client.LocalClient,
	config *SimulateConfig,
	app *baseapp.BaseApp,
	antehandler sdk.AnteHandler,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
	watermarks *WatermarkManager,
) *Backend {
	_ = "STUB: not implemented"
	return nil
}

func defaultTraceContextProvider(ctxProvider func(int64) sdk.Context) TraceContextProvider {
	_ = "STUB: not implemented"
	return *new(TraceContextProvider)
}

func (b *Backend) isV65ActiveAtHeight(height int64) bool { _ = "STUB: not implemented"; return false }

func (b *Backend) SetTraceContextProvider(provider TraceContextProvider) {
	_ = "STUB: not implemented"
	return
}

func (b *Backend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (vm.StateDB, *ethtypes.Header, error) {
	_ = "STUB: not implemented"
	return *new(vm.StateDB), nil, nil
}

// no need to check version for latest block

func (b *Backend) GetTransaction(ctx context.Context, txHash common.Hash) (found bool, tx *ethtypes.Transaction, blockHash common.Hash, blockNumber uint64, index uint64, err error) {
	_ = "STUB: not implemented"
	return false, nil, *new(common.Hash), 0, 0, nil
}

// Use BlockID.Hash rather than Header.Hash(): under CometBFT they
// are equal, but under Autobahn the Block.Header returned by /block
// is sparse (the GigaRouter's translateGlobalBlock only populates
// ChainID/Height/Time), so Header.Hash() recomputes a Merkle root
// that doesn't match any stored value — and downstream
// debug_traceTransaction fails with "block not found by hash" when
// it tries to round-trip this value through BlockByHash.
// BlockID.Hash carries the actual block hash that the EVM receipt
// store recorded during FinalizeBlock: same on both engines,
// correct under both.

//nolint:gosec

func (b *Backend) ChainDb() ethdb.Database { _ = "STUB: not implemented"; return *new(ethdb.Database) }

func (b Backend) ConvertBlockNumber(bn rpc.BlockNumber) int64 { _ = "STUB: not implemented"; return 0 }

func (b Backend) BlockByNumber(ctx context.Context, bn rpc.BlockNumber) (*ethtypes.Block, []tracersutils.TraceBlockMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// AsTransaction may return nil if it fails to unpack the tx data.

func (b Backend) BlockByHash(ctx context.Context, hash common.Hash) (*ethtypes.Block, []tracersutils.TraceBlockMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (b *Backend) RPCGasCap() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *Backend) RPCEVMTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *Backend) chainConfigForHeight(height int64) *params.ChainConfig {
	_ = "STUB: not implemented"
	return nil
}

func (b *Backend) ChainConfig() *params.ChainConfig { _ = "STUB: not implemented"; return nil }

func (b *Backend) ChainConfigAtHeight(height int64) *params.ChainConfig {
	_ = "STUB: not implemented"
	return nil
}

func (b *Backend) GetPoolNonce(_ context.Context, addr common.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Backend) Engine() consensus.Engine {
	_ = "STUB: not implemented"
	return *new(consensus.Engine)
}

func (b *Backend) HeaderByNumber(ctx context.Context, bn rpc.BlockNumber) (*ethtypes.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) StateAtTransaction(ctx context.Context, block *ethtypes.Block, txIndex int, reexec uint64) (*ethtypes.Transaction, vm.BlockContext, vm.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(vm.BlockContext), *new(vm.StateDB), *new(tracers.StateReleaseFunc), nil
}

func (b *Backend) ReplayTransactionTillIndex(ctx context.Context, block *ethtypes.Block, txIndex int) (vm.StateDB, tmtypes.Txs, error) {
	_ = "STUB: not implemented"
	return *new(vm.StateDB), *new(tmtypes.Txs), nil
}

func (b *Backend) replayTransactionTillIndex(ctx context.Context, block *ethtypes.Block, txIndex int, ctxProvider TraceContextProvider) (vm.StateDB, tmtypes.Txs, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return *

	// Short circuit if it's genesis block.
	new(vm.StateDB), *new(tmtypes.Txs), *new(tracers.StateReleaseFunc), nil
}

func (b *Backend) StateAtBlock(ctx context.Context, block *ethtypes.Block, reexec uint64, base vm.StateDB, readOnly bool, preferDisk bool) (vm.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return *new(vm.StateDB), *new(tracers.StateReleaseFunc), nil
}

func (b *Backend) initializeBlock(ctx context.Context, block *ethtypes.Block, ctxProvider TraceContextProvider) (sdk.Context, *coretypes.ResultBlock, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	return *

	// get the parent block using block.parentHash
	new(sdk.Context), nil, *new(tracers.StateReleaseFunc), nil
}

// todo: load all

func (b *Backend) GetEVM(_ context.Context, msg *core.Message, stateDB vm.StateDB, h *ethtypes.Header, vmConfig *vm.Config, blockCtx *vm.BlockContext) *vm.EVM {
	_ = "STUB: not implemented"
	return nil
}

func (b *Backend) CurrentHeader() *ethtypes.Header { _ = "STUB: not implemented"; return nil }

func (b *Backend) SuggestGasTipCap(context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil,

		// getBlockByNumberOrHash resolves blockNrOrHash to a Tendermint ResultBlock in one RPC path
		// (by hash or by number, including latest). Callers pass the result to getHeader.
		nil
}

func (b *Backend) getBlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*coretypes.ResultBlock, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// fallbackToEthHeaderOnly builds a minimal header when the block cannot be loaded
// (e.g. CurrentHeader when Block RPC fails). BaseFee is overwritten by CurrentHeader afterward.
func (b *Backend) fallbackToEthHeaderOnly(height int64) *ethtypes.Header {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func (b *Backend) getHeader(ctx context.Context, tmBlock *coretypes.ResultBlock) *ethtypes.Header {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

//nolint:gosec

func (b *Backend) GetCustomPrecompiles(h int64) map[common.Address]vm.PrecompiledContract {
	_ = "STUB: not implemented"
	return nil
}

func (b *Backend) PrepareTx(statedb vm.StateDB, tx *ethtypes.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// skip ante if no signature is set

// PrepareTxNoFlush is like PrepareTx but uses ResetForTracer instead of
// CleanupForTracer, avoiding CacheMultiStore flushes. This is required in the
// parallel block trace path where copies of the statedb are concurrently read
// by worker goroutines; flushing would write to shared CacheMultiStore layers
// and cause data races.
func (b *Backend) PrepareTxNoFlush(statedb vm.StateDB, tx *ethtypes.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Backend) GetBlockContext(ctx context.Context, block *ethtypes.Block, statedb vm.StateDB, backend export.ChainContextBackend) (vm.BlockContext, error) {
	_ = "STUB: not implemented"
	return *new(vm.BlockContext), nil
}

func noSignatureSet(tx *ethtypes.Transaction) bool { _ = "STUB: not implemented"; return false }

type Engine struct {
	*ethash.Ethash
	ctxProvider func(int64) sdk.Context
	keeper      *keeper.Keeper
}

func (e *Engine) Author(*ethtypes.Header) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}
