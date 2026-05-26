package v630

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/state"
)

const UnknownMethodCallGas uint64 = 3000

type Contexter interface {
	Ctx() sdk.Context
}

type StateEVMKeeperGetter interface {
	EVMKeeper() state.EVMKeeper
}

type PrecompileExecutor interface {
	RequiredGas([]byte, *abi.Method) uint64
	Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, hooks *tracing.Hooks) ([]byte, error)
}

type Precompile struct {
	abi.ABI
	address  common.Address
	name     string
	executor PrecompileExecutor
}

var _ vm.PrecompiledContract = &Precompile{}

func NewPrecompile(a abi.ABI, executor PrecompileExecutor, address common.Address, name string) *Precompile {
	_ = "STUB: not implemented"
	return nil
}

func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// This should never happen since this method is going to fail during Run

func (p Precompile) Run(evm *vm.EVM, caller common.Address, callingContract common.Address, input []byte, value *big.Int, readOnly bool, isFromDelegateCall bool, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HandlePrecompileError(err error, evm *vm.EVM, operation string) {
	_ = "STUB: not implemented"
	return
}

func (p Precompile) Prepare(evm *vm.EVM, input []byte) (sdk.Context, *abi.Method, []interface{}, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil, nil, nil
}

func (p Precompile) GetABI() abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) GetExecutor() PrecompileExecutor {
	_ = "STUB: not implemented"
	return *new(PrecompileExecutor)
}

type DynamicGasPrecompileExecutor interface {
	Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, hooks *tracing.Hooks) (ret []byte, remainingGas uint64, err error)
	EVMKeeper() putils.EVMKeeper
}

type DynamicGasPrecompile struct {
	*Precompile
	executor DynamicGasPrecompileExecutor
}

var _ vm.DynamicGasPrecompiledContract = &DynamicGasPrecompile{}

func NewDynamicGasPrecompile(a abi.ABI, executor DynamicGasPrecompileExecutor, address common.Address, name string) *DynamicGasPrecompile {
	_ = "STUB: not implemented"
	return nil
}

func (d DynamicGasPrecompile) RunAndCalculateGas(evm *vm.EVM, caller common.Address, callingContract common.Address, input []byte, suppliedGas uint64, value *big.Int, hooks *tracing.Hooks, readOnly bool, isFromDelegateCall bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (d DynamicGasPrecompile) GetExecutor() DynamicGasPrecompileExecutor {
	_ = "STUB: not implemented"
	return *new(DynamicGasPrecompileExecutor)
}

func ValidateArgsLength(args []interface{}, length int) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateNonPayable(value *big.Int) error { _ = "STUB: not implemented"; return nil }

func HandlePaymentUsei(ctx sdk.Context, precompileAddr sdk.AccAddress, payer sdk.AccAddress, value *big.Int, bankKeeper putils.BankKeeper, evmKeeper putils.EVMKeeper, hooks *tracing.Hooks, depth int) (sdk.Coin, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coin), nil
}

// refund payer because the following precompile logic will debit the payments from payer's account
// this creates a new event manager to avoid surfacing these as cosmos events

func HandlePaymentUseiWei(ctx sdk.Context, precompileAddr sdk.AccAddress, payer sdk.AccAddress, value *big.Int, bankKeeper putils.BankKeeper, evmKeeper putils.EVMKeeper, hooks *tracing.Hooks, depth int) (sdk.Int, sdk.Int, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), *new(sdk.Int), nil
}

// refund payer because the following precompile logic will debit the payments from payer's account
// this creates a new event manager to avoid surfacing these as cosmos events

/*
*
sei gas = evm gas * multiplier
sei gas price = fee / sei gas = fee / (evm gas * multiplier) = evm gas / multiplier
*/
func GetRemainingGas(ctx sdk.Context, evmKeeper putils.EVMKeeper) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func ExtractMethodID(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Check if the input has at least the length needed for methodID
	return nil, nil
}

func DefaultGasCost(input []byte, isTransaction bool) uint64 { _ = "STUB: not implemented"; return 0 }

func MustGetABI(f embed.FS, filename string) abi.ABI {
	_ = "STUB: not implemented"
	return *new(abi.ABI)
}

func GetSeiAddressByEvmAddress(ctx sdk.Context, evmAddress common.Address, evmKeeper putils.EVMKeeper) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

func GetSeiAddressFromArg(ctx sdk.Context, arg interface{}, evmKeeper putils.EVMKeeper) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}
