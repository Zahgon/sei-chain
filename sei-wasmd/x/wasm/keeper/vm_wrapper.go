package keeper

import (
	"sync"
	"time"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	wasmvm "github.com/sei-protocol/sei-chain/sei-wasmvm"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

const CreateTimeout time.Duration = 15 * time.Second

type VMWrapper struct {
	types.WasmerEngine

	mu *sync.Mutex
}

func NewVMWrapper(inner types.WasmerEngine) types.WasmerEngine {
	_ = "STUB: not implemented"
	return *new(types.WasmerEngine)
}

func (w *VMWrapper) Create(code wasmvm.WasmCode) (checksum wasmvm.Checksum, err error) {
	_ = "STUB: not implemented"
	return *new(wasmvm.Checksum), nil
}

func (w *VMWrapper) Instantiate(
	checksum wasmvm.Checksum,
	env wasmvmtypes.Env,
	info wasmvmtypes.MessageInfo,
	initMsg []byte,
	store wasmvm.KVStore,
	goapi wasmvm.GoAPI,
	querier wasmvm.Querier,
	gasMeter wasmvm.GasMeter,
	gasLimit uint64,
	deserCost wasmvmtypes.UFraction,
) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (w *VMWrapper) Execute(
	code wasmvm.Checksum,
	env wasmvmtypes.Env,
	info wasmvmtypes.MessageInfo,
	executeMsg []byte,
	store wasmvm.KVStore,
	goapi wasmvm.GoAPI,
	querier wasmvm.Querier,
	gasMeter wasmvm.GasMeter,
	gasLimit uint64,
	deserCost wasmvmtypes.UFraction,
) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (w *VMWrapper) Migrate(
	checksum wasmvm.Checksum,
	env wasmvmtypes.Env,
	migrateMsg []byte,
	store wasmvm.KVStore,
	goapi wasmvm.GoAPI,
	querier wasmvm.Querier,
	gasMeter wasmvm.GasMeter,
	gasLimit uint64,
	deserCost wasmvmtypes.UFraction,
) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (w *VMWrapper) Sudo(
	checksum wasmvm.Checksum,
	env wasmvmtypes.Env,
	sudoMsg []byte,
	store wasmvm.KVStore,
	goapi wasmvm.GoAPI,
	querier wasmvm.Querier,
	gasMeter wasmvm.GasMeter,
	gasLimit uint64,
	deserCost wasmvmtypes.UFraction,
) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (w *VMWrapper) Reply(
	checksum wasmvm.Checksum,
	env wasmvmtypes.Env,
	reply wasmvmtypes.Reply,
	store wasmvm.KVStore,
	goapi wasmvm.GoAPI,
	querier wasmvm.Querier,
	gasMeter wasmvm.GasMeter,
	gasLimit uint64,
	deserCost wasmvmtypes.UFraction,
) (*wasmvmtypes.Response, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (w *VMWrapper) Unpin(checksum wasmvm.Checksum) error { _ = "STUB: not implemented"; return nil }

func (w *VMWrapper) Pin(checksum wasmvm.Checksum) error { _ = "STUB: not implemented"; return nil }
