package v575

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v575"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

const (
	InstantiateMethod  = "instantiate"
	ExecuteMethod      = "execute"
	ExecuteBatchMethod = "execute_batch"
	QueryMethod        = "query"
)

const WasmdAddress = "0x0000000000000000000000000000000000001002"

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	evmKeeper       utils.EVMKeeper
	bankKeeper      utils.BankKeeper
	wasmdKeeper     utils.WasmdKeeper
	wasmdViewKeeper utils.WasmdViewKeeper
	address         common.Address

	InstantiateID  []byte
	ExecuteID      []byte
	ExecuteBatchID []byte
	QueryID        []byte
}

type ExecuteMsg struct {
	ContractAddress string `json:"contractAddress"`
	Msg             []byte `json:"msg"`
	Coins           []byte `json:"coins"`
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.DynamicGasPrecompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, hooks *tracing.Hooks) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) EVMKeeper() utils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(utils.EVMKeeper)
}

func (p PrecompileExecutor) instantiate(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, hooks *tracing.Hooks, evm *vm.EVM) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// Run basic validation, can also just expose validateLabel and validate validateWasmCode in sei-wasmd

// sanity check coin amounts match

func (p PrecompileExecutor) executeBatch(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, hooks *tracing.Hooks, evm *vm.EVM) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// validate coins add up to value

// if validateValue is greater than zero, then value must be provided, and they must be equal

// Copy to avoid modifying the original value

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// process coin amount from the value provided

// sanity check coin amounts match

// Run basic validation, can also just expose validateLabel and validate validateWasmCode in sei-wasmd

func (p PrecompileExecutor) execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, hooks *tracing.Hooks, evm *vm.EVM) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// addresses will be sent in Sei format

// Run basic validation, can also just expose validateLabel and validate validateWasmCode in sei-wasmd

// sanity check coin amounts match

func (p PrecompileExecutor) query(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// addresses will be sent in Sei format
