package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/gaskv"
)

/*
Context is an immutable object contains all information needed to
process a request.

It contains a context.Context object inside if you want to use that,
but please do not over-use it. We try to keep all data structured
and standard additions here would be better just to add to the Context struct
*/
type Context struct {
	ctx                context.Context
	ms                 MultiStore
	nextMs             MultiStore          // ms of the next height; only used in tracing
	nextStoreKeys      map[string]struct{} // store key names that should use nextMs
	header             tmproto.Header
	headerHash         tmbytes.HexBytes
	chainID            string
	txBytes            []byte
	txSum              [32]byte
	voteInfo           []abci.VoteInfo
	gasMeter           GasMeter
	gasEstimate        uint64
	occEnabled         bool
	blockGasMeter      GasMeter
	checkTx            bool
	recheckTx          bool // if recheckTx == true, then checkTx must also be true
	minGasPrice        DecCoins
	consParams         *tmproto.ConsensusParams
	eventManager       *EventManager
	evmEventManager    *EVMEventManager
	priority           int64         // The tx priority, only relevant in CheckTx
	hasPriority        bool          // Whether the tx has a priority set
	deliverTxCallback  func(Context) // callback to make at the end of DeliverTx.
	evmRequiredBalance *big.Int      // Required sender balance for this EVM tx, only relevant in CheckTx.

	// EVM properties
	evm                                 bool   // EVM transaction flag
	evmNonce                            uint64 // EVM Transaction nonce
	evmSenderAddress                    common.Address
	seiSenderAddress                    AccAddress
	evmTxHash                           string // EVM TX hash
	evmVmError                          string // EVM VM error during execution
	evmEntryViaWasmdPrecompile          bool   // EVM is entered via wasmd precompile directly
	evmPrecompileCalledFromDelegateCall bool   // EVM precompile is called from a delegate call

	messageIndex int // Used to track current message being processed
	txIndex      int

	closestUpgradeName string

	traceSpanContext context.Context

	isTracing    bool
	isSimulation bool
	storeTracer  gaskv.IStoreTracer
}

// Proposed rename, not done to avoid API breakage
type Request = Context

// Read-only accessors
func (c Context) Context() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (c Context) MultiStore() MultiStore { _ = "STUB: not implemented"; return *new(MultiStore) }

func (c Context) GigaMultiStore() GigaMultiStore {
	_ = "STUB: not implemented"
	return *new(GigaMultiStore)
}

func (c Context) BlockHeight() int64 { _ = "STUB: not implemented"; return 0 }

func (c Context) BlockTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c Context) ChainID() string { _ = "STUB: not implemented"; return "" }

func (c Context) TxBytes() []byte { _ = "STUB: not implemented"; return nil }

func (c Context) TxSum() [32]byte { _ = "STUB: not implemented"; return nil }

func (c Context) VoteInfos() []abci.VoteInfo { _ = "STUB: not implemented"; return nil }

func (c Context) GasMeter() GasMeter { _ = "STUB: not implemented"; return *new(GasMeter) }

func (c Context) GasEstimate() uint64 { _ = "STUB: not implemented"; return 0 }

func (c Context) IsCheckTx() bool { _ = "STUB: not implemented"; return false }

func (c Context) IsReCheckTx() bool { _ = "STUB: not implemented"; return false }

func (c Context) IsOCCEnabled() bool { _ = "STUB: not implemented"; return false }

func (c Context) MinGasPrices() DecCoins { _ = "STUB: not implemented"; return *new(DecCoins) }

func (c Context) EventManager() *EventManager { _ = "STUB: not implemented"; return nil }

func (c Context) EVMEventManager() *EVMEventManager { _ = "STUB: not implemented"; return nil }

func (c Context) Priority() int64 { _ = "STUB: not implemented"; return 0 }

func (c Context) EVMSenderAddress() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (c Context) EVMNonce() uint64 { _ = "STUB: not implemented"; return 0 }

func (c Context) SeiSenderAddress() AccAddress { _ = "STUB: not implemented"; return *new(AccAddress) }

func (c Context) EVMTxHash() string { _ = "STUB: not implemented"; return "" }

func (c Context) IsEVM() bool { _ = "STUB: not implemented"; return false }

func (c Context) EVMVMError() string { _ = "STUB: not implemented"; return "" }

func (c Context) EVMEntryViaWasmdPrecompile() bool { _ = "STUB: not implemented"; return false }

func (c Context) EVMPrecompileCalledFromDelegateCall() bool {
	_ = "STUB: not implemented"
	return false
}

func (c Context) DeliverTxCallback() func(Context) { _ = "STUB: not implemented"; return nil }

func (c Context) EVMRequiredBalance() *big.Int { _ = "STUB: not implemented"; return nil }

func (c Context) MessageIndex() int { _ = "STUB: not implemented"; return 0 }

func (c Context) TxIndex() int {
	_ = "STUB: not implemented"

	// clone the header before returning
	return 0
}

func (c Context) BlockHeader() tmproto.Header {
	_ = "STUB: not implemented"
	return *new(tmproto.Header)
}

func (c Context) TraceSpanContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c Context) IsTracing() bool {
	_ = "STUB: not implemented"

	// IsSimulation reports whether the context is being run in transaction
	// simulation mode.
	return false
}

func (c Context) IsSimulation() bool { _ = "STUB: not implemented"; return false }

func (c Context) StoreTracer() gaskv.IStoreTracer {
	_ = "STUB: not implemented"
	return *new(gaskv.IStoreTracer)
}

// WithPriority returns a Context with an updated tx priority.
func (c Context) WithPriority(p int64) Context { _ = "STUB: not implemented"; return *new(Context) }

// HasPriority returns true iff the priority is set for this Context even if it
// was set to zero.
func (c Context) HasPriority() bool { _ = "STUB: not implemented"; return false }

// HeaderHash returns a copy of the header hash obtained during abci.RequestBeginBlock
func (c Context) HeaderHash() tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

func (c Context) ConsensusParams() *tmproto.ConsensusParams { _ = "STUB: not implemented"; return nil }

// create a new context
func NewContext(ms MultiStore, header tmproto.Header, isCheckTx bool) Context {
	_ = "STUB: not implemented"
	// https://github.com/gogo/protobuf/issues/519
	return *new(Context)
}

// WithContext returns a Context with an updated context.Context.
func (c Context) WithContext(ctx context.Context) Context {
	_ = "STUB: not implemented"
	return *

	// WithMultiStore returns a Context with an updated MultiStore.
	new(Context)
}

func (c Context) WithMultiStore(ms MultiStore) Context {
	_ = "STUB: not implemented"
	return *

	// WithBlockHeader returns a Context with an updated tendermint block header in UTC time.
	new(Context)
}

func (c Context) WithBlockHeader(header tmproto.Header) Context {
	_ = "STUB: not implemented"
	// https://github.com/gogo/protobuf/issues/519
	return *new(Context)
}

// WithHeaderHash returns a Context with an updated tendermint block header hash.
func (c Context) WithHeaderHash(hash []byte) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithBlockTime returns a Context with an updated tendermint block header time in UTC time
func (c Context) WithBlockTime(newTime time.Time) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// https://github.com/gogo/protobuf/issues/519

// WithProposer returns a Context with an updated proposer consensus address.
func (c Context) WithProposer(addr ConsAddress) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithBlockHeight returns a Context with an updated block height.
func (c Context) WithBlockHeight(height int64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithChainID returns a Context with an updated chain identifier.
func (c Context) WithChainID(chainID string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithTxBytes returns a Context with an updated txBytes.
func (c Context) WithTxBytes(txBytes []byte) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithTxSum(txSum [32]byte) Context {
	_ = "STUB: not implemented"
	return *

	// WithVoteInfos returns a Context with an updated consensus VoteInfo.
	new(Context)
}

func (c Context) WithVoteInfos(voteInfo []abci.VoteInfo) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithGasMeter returns a Context with an updated transaction GasMeter.
func (c Context) WithGasMeter(meter GasMeter) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithGasEstimate returns a Context with an updated gas estimate.
func (c Context) WithGasEstimate(gasEstimate uint64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithIsCheckTx enables or disables CheckTx value for verifying transactions and returns an updated Context
func (c Context) WithIsCheckTx(isCheckTx bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithIsOCCEnabled enables or disables whether OCC is used as the concurrency algorithm
func (c Context) WithIsOCCEnabled(isOCCEnabled bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithIsRecheckTx called with true will also set true on checkTx in order to
// enforce the invariant that if recheckTx = true then checkTx = true as well.
func (c Context) WithIsReCheckTx(isRecheckTx bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithMinGasPrices returns a Context with an updated minimum gas price value
func (c Context) WithMinGasPrices(gasPrices DecCoins) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithConsensusParams returns a Context with an updated consensus params
func (c Context) WithConsensusParams(params *tmproto.ConsensusParams) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithEventManager returns a Context with an updated event manager
func (c Context) WithEventManager(em *EventManager) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithEvmEventManager(em *EVMEventManager) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithMessageIndex returns a Context with the current message index that's being processed
func (c Context) WithMessageIndex(messageIndex int) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithTxIndex returns a Context with the current transaction index that's being processed
func (c Context) WithTxIndex(txIndex int) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) WithTraceSpanContext(ctx context.Context) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithEVMSenderAddress(address common.Address) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithEVMNonce(nonce uint64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithSeiSenderAddress(address AccAddress) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithIsEVM(isEVM bool) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) WithEVMTxHash(txHash string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithEVMVMError(vmError string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithEVMEntryViaWasmdPrecompile(e bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithEVMPrecompileCalledFromDelegateCall(e bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithDeliverTxCallback(deliverTxCallback func(Context)) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithEVMRequiredBalance(evmRequiredBalance *big.Int) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) WithIsTracing(it bool) Context { _ = "STUB: not implemented"; return *new(Context) }

// WithIsSimulation sets the simulation flag on the context.
func (c Context) WithIsSimulation(isSimulation bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// WithTraceMode enables historical tracing behavior without allocating a KV
// store tracer. This keeps upgrade-aware tracing semantics for ordinary
// debug_trace* RPCs without paying the per-access StoreTracer overhead.
func (c Context) WithTraceMode(it bool) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) WithNextMs(ms MultiStore, nextStoreKeys []string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// TODO: remove???
func (c Context) IsZero() bool {
	_ = "STUB: not implemented"

	// WithValue is deprecated, provided for backwards compatibility
	// Please use
	//
	//	ctx = ctx.WithContext(context.WithValue(ctx.Context(), key, false))
	//
	// instead of
	//
	//	ctx = ctx.WithValue(key, false)
	return false
}

func (c Context) WithValue(key, value interface{}) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// Value is deprecated, provided for backwards compatibility
// Please use
//
//	ctx.Context().Value(key)
//
// instead of
//
//	ctx.Value(key)
func (c Context) Value(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// ----------------------------------------------------------------------------
// Store / Caching
// ----------------------------------------------------------------------------

// KVStore fetches a KVStore from the MultiStore.
func (c Context) KVStore(key StoreKey) KVStore { _ = "STUB: not implemented"; return *new(KVStore) }

func (c Context) GigaKVStore(key StoreKey) KVStore { _ = "STUB: not implemented"; return *new(KVStore) }

// TransientStore fetches a TransientStore from the MultiStore.
func (c Context) TransientStore(key StoreKey) KVStore {
	_ = "STUB: not implemented"
	return *new(KVStore)
}

// CacheContext returns a new Context with the multi-store cached and a new
// EventManager. The cached context is written to the context when writeCache
// is called.
func (c Context) CacheContext() (cc Context, writeCache func()) {
	_ = "STUB: not implemented"
	return *new(Context), nil
}

// ContextKey defines a type alias for a stdlib Context key.
type ContextKey string

// SdkContextKey is the key in the context.Context which holds the sdk.Context.
const SdkContextKey ContextKey = "sdk-context"

// WrapSDKContext returns a stdlib context.Context with the provided sdk.Context's internal
// context as a value. It is useful for passing an sdk.Context  through methods that take a
// stdlib context.Context parameter such as generated gRPC methods. To get the original
// sdk.Context back, call UnwrapSDKContext.
func WrapSDKContext(ctx Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// UnwrapSDKContext retrieves a Context from a context.Context instance
// attached with WrapSDKContext. It panics if a Context was not properly
// attached
func UnwrapSDKContext(ctx context.Context) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) ClosestUpgradeName() string { _ = "STUB: not implemented"; return "" }

func (c Context) WithClosestUpgradeName(name string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}
