package v614

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v614"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	DelegateMethod        = "delegate"
	RedelegateMethod      = "redelegate"
	UndelegateMethod      = "undelegate"
	DelegationMethod      = "delegation"
	CreateValidatorMethod = "createValidator"
	EditValidatorMethod   = "editValidator"
)

const (
	StakingAddress = "0x0000000000000000000000000000000000001005"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	stakingKeeper  utils.StakingKeeper
	stakingQuerier utils.StakingQuerier
	evmKeeper      utils.EVMKeeper
	bankKeeper     utils.BankKeeper
	address        common.Address

	DelegateID        []byte
	RedelegateID      []byte
	UndelegateID      []byte
	DelegationID      []byte
	CreateValidatorID []byte
	EditValidatorID   []byte
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// This should never happen since this is going to fail during Run

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) delegate(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if delegator is associated, then it must have Account set already
// if delegator is not associated, then it can't delegate anyway (since
// there is no good way to merge delegations if it becomes associated)

func (p PrecompileExecutor) redelegate(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) undelegate(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Delegation struct {
	Balance    Balance
	Delegation DelegationDetails
}

type Balance struct {
	Amount *big.Int
	Denom  string
}

type DelegationDetails struct {
	DelegatorAddress string
	Shares           *big.Int
	Decimals         *big.Int
	ValidatorAddress string
}

func (p PrecompileExecutor) delegation(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) createValidator(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract arguments

// Get validator address (caller's associated Sei address)

// Parse public key from hex

// Create ed25519 public key

// Parse commission rates

// Call the staking keeper

func (p PrecompileExecutor) editValidator(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract arguments

// Get validator address (caller's associated Sei address)

// Parse commission rate - if empty string, don't change commission

// If commissionRateStr is empty, commissionRate remains nil

// Parse minSelfDelegation - if 0, don't change it

// If minSelfDelegation is 0, minSelfDelegationPtr remains nil
