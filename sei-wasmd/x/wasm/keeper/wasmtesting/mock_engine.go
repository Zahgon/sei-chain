package wasmtesting

import (
	wasmvm "github.com/sei-protocol/sei-chain/sei-wasmvm"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var _ types.WasmerEngine = &MockWasmer{}

// MockWasmer implements types.WasmerEngine for testing purpose. One or multiple messages can be stubbed.
// Without a stub function a panic is thrown.
type MockWasmer struct {
	CreateFn            func(codeID wasmvm.WasmCode) (wasmvm.Checksum, error)
	AnalyzeCodeFn       func(codeID wasmvm.Checksum) (*wasmvmtypes.AnalysisReport, error)
	InstantiateFn       func(codeID wasmvm.Checksum, env wasmvmtypes.Env, info wasmvmtypes.MessageInfo, initMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error)
	ExecuteFn           func(codeID wasmvm.Checksum, env wasmvmtypes.Env, info wasmvmtypes.MessageInfo, executeMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error)
	QueryFn             func(codeID wasmvm.Checksum, env wasmvmtypes.Env, queryMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) ([]byte, uint64, error)
	MigrateFn           func(codeID wasmvm.Checksum, env wasmvmtypes.Env, migrateMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error)
	SudoFn              func(codeID wasmvm.Checksum, env wasmvmtypes.Env, sudoMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error)
	ReplyFn             func(codeID wasmvm.Checksum, env wasmvmtypes.Env, reply wasmvmtypes.Reply, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error)
	GetCodeFn           func(codeID wasmvm.Checksum) (wasmvm.WasmCode, error)
	CleanupFn           func()
	IBCChannelOpenFn    func(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCChannelOpenMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBC3ChannelOpenResponse, uint64, error)
	IBCChannelConnectFn func(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCChannelConnectMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error)
	IBCChannelCloseFn   func(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCChannelCloseMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error)
	IBCPacketReceiveFn  func(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCPacketReceiveMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCReceiveResult, uint64, error)
	IBCPacketAckFn      func(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCPacketAckMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error)
	IBCPacketTimeoutFn  func(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCPacketTimeoutMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error)
	PinFn               func(checksum wasmvm.Checksum) error
	UnpinFn             func(checksum wasmvm.Checksum) error
	GetMetricsFn        func() (*wasmvmtypes.Metrics, error)
}

func (m *MockWasmer) IBCChannelOpen(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCChannelOpenMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBC3ChannelOpenResponse, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) IBCChannelConnect(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCChannelConnectMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) IBCChannelClose(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCChannelCloseMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) IBCPacketReceive(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCPacketReceiveMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCReceiveResult, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) IBCPacketAck(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCPacketAckMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) IBCPacketTimeout(codeID wasmvm.Checksum, env wasmvmtypes.Env, msg wasmvmtypes.IBCPacketTimeoutMsg, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.IBCBasicResponse, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) Create(codeID wasmvm.WasmCode) (wasmvm.Checksum, error) {
	_ = "STUB: not implemented"
	return *new(wasmvm.Checksum), nil
}

func (m *MockWasmer) AnalyzeCode(codeID wasmvm.Checksum) (*wasmvmtypes.AnalysisReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockWasmer) Instantiate(codeID wasmvm.Checksum, env wasmvmtypes.Env, info wasmvmtypes.MessageInfo, initMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) Execute(codeID wasmvm.Checksum, env wasmvmtypes.Env, info wasmvmtypes.MessageInfo, executeMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) Query(codeID wasmvm.Checksum, env wasmvmtypes.Env, queryMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) Migrate(codeID wasmvm.Checksum, env wasmvmtypes.Env, migrateMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) Sudo(codeID wasmvm.Checksum, env wasmvmtypes.Env, sudoMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) Reply(codeID wasmvm.Checksum, env wasmvmtypes.Env, reply wasmvmtypes.Reply, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MockWasmer) GetCode(codeID wasmvm.Checksum) (wasmvm.WasmCode, error) {
	_ = "STUB: not implemented"
	return *new(wasmvm.WasmCode), nil
}

func (m *MockWasmer) Cleanup() { _ = "STUB: not implemented"; return }

func (m *MockWasmer) Pin(checksum wasmvm.Checksum) error { _ = "STUB: not implemented"; return nil }

func (m *MockWasmer) Unpin(checksum wasmvm.Checksum) error { _ = "STUB: not implemented"; return nil }

func (m *MockWasmer) GetMetrics() (*wasmvmtypes.Metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var AlwaysPanicMockWasmer = &MockWasmer{}

// SelfCallingInstMockWasmer prepares a Wasmer mock that calls itself on instantiation.
func SelfCallingInstMockWasmer(executeCalled *bool) *MockWasmer {
	_ = "STUB: not implemented"
	return nil
}

// IBCContractCallbacks defines the methods from wasmvm to interact with the wasm contract.
// A mock contract would implement the interface to fully simulate a wasm contract's behaviour.
type IBCContractCallbacks interface {
	IBCChannelOpen(
		codeID wasmvm.Checksum,
		env wasmvmtypes.Env,
		channel wasmvmtypes.IBCChannelOpenMsg,
		store wasmvm.KVStore,
		goapi wasmvm.GoAPI,
		querier wasmvm.Querier,
		gasMeter wasmvm.GasMeter,
		gasLimit uint64,
		deserCost wasmvmtypes.UFraction,
	) (*wasmvmtypes.IBC3ChannelOpenResponse, uint64, error)

	IBCChannelConnect(
		codeID wasmvm.Checksum,
		env wasmvmtypes.Env,
		channel wasmvmtypes.IBCChannelConnectMsg,
		store wasmvm.KVStore,
		goapi wasmvm.GoAPI,
		querier wasmvm.Querier,
		gasMeter wasmvm.GasMeter,
		gasLimit uint64,
		deserCost wasmvmtypes.UFraction,
	) (*wasmvmtypes.IBCBasicResponse, uint64, error)

	IBCChannelClose(
		codeID wasmvm.Checksum,
		env wasmvmtypes.Env,
		channel wasmvmtypes.IBCChannelCloseMsg,
		store wasmvm.KVStore,
		goapi wasmvm.GoAPI,
		querier wasmvm.Querier,
		gasMeter wasmvm.GasMeter,
		gasLimit uint64,
		deserCost wasmvmtypes.UFraction,
	) (*wasmvmtypes.IBCBasicResponse, uint64, error)

	IBCPacketReceive(
		codeID wasmvm.Checksum,
		env wasmvmtypes.Env,
		packet wasmvmtypes.IBCPacketReceiveMsg,
		store wasmvm.KVStore,
		goapi wasmvm.GoAPI,
		querier wasmvm.Querier,
		gasMeter wasmvm.GasMeter,
		gasLimit uint64,
		deserCost wasmvmtypes.UFraction,
	) (*wasmvmtypes.IBCReceiveResult, uint64, error)

	IBCPacketAck(
		codeID wasmvm.Checksum,
		env wasmvmtypes.Env,
		ack wasmvmtypes.IBCPacketAckMsg,
		store wasmvm.KVStore,
		goapi wasmvm.GoAPI,
		querier wasmvm.Querier,
		gasMeter wasmvm.GasMeter,
		gasLimit uint64,
		deserCost wasmvmtypes.UFraction,
	) (*wasmvmtypes.IBCBasicResponse, uint64, error)

	IBCPacketTimeout(
		codeID wasmvm.Checksum,
		env wasmvmtypes.Env,
		packet wasmvmtypes.IBCPacketTimeoutMsg,
		store wasmvm.KVStore,
		goapi wasmvm.GoAPI,
		querier wasmvm.Querier,
		gasMeter wasmvm.GasMeter,
		gasLimit uint64,
		deserCost wasmvmtypes.UFraction,
	) (*wasmvmtypes.IBCBasicResponse, uint64, error)
}

type contractExecutable interface {
	Execute(
		codeID wasmvm.Checksum,
		env wasmvmtypes.Env,
		info wasmvmtypes.MessageInfo,
		executeMsg []byte,
		store wasmvm.KVStore,
		goapi wasmvm.GoAPI,
		querier wasmvm.Querier,
		gasMeter wasmvm.GasMeter,
		gasLimit uint64,
		deserCost wasmvmtypes.UFraction,
	) (*wasmvmtypes.Response, uint64, error)
}

// MakeInstantiable adds some noop functions to not fail when contract is used for instantiation
func MakeInstantiable(m *MockWasmer) { _ = "STUB: not implemented"; return }

// MakeIBCInstantiable adds some noop functions to not fail when contract is used for instantiation
func MakeIBCInstantiable(m *MockWasmer) { _ = "STUB: not implemented"; return }

// NewIBCContractMockWasmer prepares a mocked wasm_engine for testing with an IBC contract test type.
// It is safe to use the mock with store code and instantiate functions in keeper as is also prepared
// with stubs. Execute is optional. When implemented by the Go test contract then it can be used with
// the mock.
func NewIBCContractMockWasmer(c IBCContractCallbacks) *MockWasmer {
	_ = "STUB: not implemented"
	return nil
}

// optional function

func HashOnlyCreateFn(code wasmvm.WasmCode) (wasmvm.Checksum, error) {
	_ = "STUB: not implemented"
	return *new(wasmvm.Checksum), nil
}

func NoOpInstantiateFn(wasmvm.Checksum, wasmvmtypes.Env, wasmvmtypes.MessageInfo, []byte, wasmvm.KVStore, wasmvm.GoAPI, wasmvm.Querier, wasmvm.GasMeter, uint64, wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func NoOpCreateFn(_ wasmvm.WasmCode) (wasmvm.Checksum, error) {
	_ = "STUB: not implemented"
	return *new(wasmvm.Checksum), nil
}

func HasIBCAnalyzeFn(wasmvm.Checksum) (*wasmvmtypes.AnalysisReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithoutIBCAnalyzeFn(wasmvm.Checksum) (*wasmvmtypes.AnalysisReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
