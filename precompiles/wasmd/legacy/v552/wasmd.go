package v552

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v552"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	InstantiateMethod  = "instantiate"
	ExecuteMethod      = "execute"
	ExecuteBatchMethod = "execute_batch"
	QueryMethod        = "query"
)

const WasmdAddress = "0x0000000000000000000000000000000000001002"

var _ vm.PrecompiledContract = &Precompile{}
var _ vm.DynamicGasPrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type Precompile struct {
	pcommon.Precompile
	evmKeeper       putils.EVMKeeper
	bankKeeper      putils.BankKeeper
	wasmdKeeper     putils.WasmdKeeper
	wasmdViewKeeper putils.WasmdViewKeeper
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

func NewPrecompile(keepers putils.Keepers) (*Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// This should never happen since this method is going to fail during Run

func (Precompile) IsTransaction(method string) bool { _ = "STUB: not implemented"; return false }

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) RunAndCalculateGas(evm *vm.EVM, caller common.Address, callingContract common.Address, input []byte, suppliedGas uint64, value *big.Int, hooks *tracing.Hooks, readOnly bool, _ bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p Precompile) Run(*vm.EVM, common.Address, common.Address, []byte, *big.Int, bool, bool, *tracing.Hooks) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) instantiate(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, hooks *tracing.Hooks, evm *vm.EVM) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// Run basic validation, can also just expose validateLabel and validate validateWasmCode in sei-wasmd

// sanity check coin amounts match

func (p Precompile) executeBatch(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, hooks *tracing.Hooks, evm *vm.EVM) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// validate coins add up to value

// if validateValue is greater than zero, then value must be provided, and they must be equal

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// process coin amount from the value provided

// sanity check coin amounts match

// Run basic validation, can also just expose validateLabel and validate validateWasmCode in sei-wasmd

func (p Precompile) execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, hooks *tracing.Hooks, evm *vm.EVM) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// addresses will be sent in Sei format

// Run basic validation, can also just expose validateLabel and validate validateWasmCode in sei-wasmd

// sanity check coin amounts match

func (p Precompile) query(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// addresses will be sent in Sei format
