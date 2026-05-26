package v555

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/core/tracing"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v555"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	SetWithdrawAddressMethod                = "setWithdrawAddress"
	WithdrawDelegationRewardsMethod         = "withdrawDelegationRewards"
	WithdrawMultipleDelegationRewardsMethod = "withdrawMultipleDelegationRewards"
)

const (
	DistrAddress = "0x0000000000000000000000000000000000001007"
)

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

func GetABI() abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }

type Precompile struct {
	pcommon.Precompile
	distrKeeper putils.DistributionKeeper
	evmKeeper   putils.EVMKeeper
	address     common.Address

	SetWithdrawAddrID                   []byte
	WithdrawDelegationRewardsID         []byte
	WithdrawMultipleDelegationRewardsID []byte
}

func NewPrecompile(keepers putils.Keepers) (*Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// This should never happen since this method is going to fail during Run

func (Precompile) IsTransaction(method string) bool { _ = "STUB: not implemented"; return false }

func (p Precompile) RunAndCalculateGas(evm *vm.EVM, caller common.Address, _ common.Address, input []byte, suppliedGas uint64, value *big.Int, _ *tracing.Hooks, _ bool, _ bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) Run(evm *vm.EVM, caller common.Address, callingContract common.Address, input []byte, value *big.Int, readOnly bool, _ bool, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) setWithdrawAddress(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p Precompile) withdrawDelegationRewards(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p Precompile) validateInput(value *big.Int, args []interface{}, expectedArgsLength int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p Precompile) withdraw(ctx sdk.Context, delegator sdk.AccAddress, validatorAddress string) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

func (p Precompile) getDelegator(ctx sdk.Context, caller common.Address) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

func (p Precompile) withdrawMultipleDelegationRewards(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p Precompile) accAddressFromArg(ctx sdk.Context, arg interface{}) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}
