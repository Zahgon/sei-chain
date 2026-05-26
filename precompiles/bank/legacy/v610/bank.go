package v610

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v606"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	SendMethod        = "send"
	SendNativeMethod  = "sendNative"
	BalanceMethod     = "balance"
	AllBalancesMethod = "all_balances"
	NameMethod        = "name"
	SymbolMethod      = "symbol"
	DecimalsMethod    = "decimals"
	SupplyMethod      = "supply"
)

const (
	BankAddress = "0x0000000000000000000000000000000000001001"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	accountKeeper putils.AccountKeeper
	bankKeeper    putils.BankKeeper
	bankMsgServer putils.BankMsgServer
	evmKeeper     putils.EVMKeeper
	address       common.Address

	SendID        []byte
	SendNativeID  []byte
	BalanceID     []byte
	AllBalancesID []byte
	NameID        []byte
	SymbolID      []byte
	DecimalsID    []byte
	SupplyID      []byte
}

type CoinBalance struct {
	Amount *big.Int
	Denom  string
}

func GetABI() abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }

func NewPrecompile(keepers putils.Keepers) (*pcommon.DynamicGasPrecompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, hooks *tracing.Hooks) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	// Needed to catch gas meter panics
	return nil, 0, nil
}

func (p PrecompileExecutor) send(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// short circuit

func (p PrecompileExecutor) sendNative(ctx sdk.Context, method *abi.Method, args []interface{}, caller common.Address, callingContract common.Address, value *big.Int, readOnly bool, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) balance(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) all_balances(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// convert to coin balance structs

func (p PrecompileExecutor) name(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) symbol(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) decimals(ctx sdk.Context, method *abi.Method, _ []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// all native tokens are integer-based, returns decimals for microdenom (usei)

func (p PrecompileExecutor) totalSupply(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) accAddressFromArg(ctx sdk.Context, arg interface{}) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

// return the casted version instead

func (PrecompileExecutor) IsTransaction(method string) bool {
	_ = "STUB: not implemented"
	return false
}

func (p PrecompileExecutor) EVMKeeper() putils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(putils.EVMKeeper)
}
