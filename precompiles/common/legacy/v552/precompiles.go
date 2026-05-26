package v552

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const UnknownMethodCallGas uint64 = 3000

type Contexter interface {
	Ctx() sdk.Context
}

type Precompile struct {
	abi.ABI
}

func (p Precompile) RequiredGas(input []byte, isTransaction bool) uint64 {
	_ = "STUB: not implemented"
	// first four bytes are method ID
	return 0
}

func (p Precompile) Prepare(evm *vm.EVM, input []byte) (sdk.Context, *abi.Method, []interface{}, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil, nil, nil
}

func (p Precompile) GetABI() abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }

func ValidateArgsLength(args []interface{}, length int) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateNonPayable(value *big.Int) error { _ = "STUB: not implemented"; return nil }

func HandlePaymentUsei(ctx sdk.Context, precompileAddr sdk.AccAddress, payer sdk.AccAddress, value *big.Int, bankKeeper utils.BankKeeper, evmKeeper utils.EVMKeeper, hooks *tracing.Hooks, depth int) (sdk.Coin, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coin), nil
}

// refund payer because the following precompile logic will debit the payments from payer's account
// this creates a new event manager to avoid surfacing these as cosmos events

func HandlePaymentUseiWei(ctx sdk.Context, precompileAddr sdk.AccAddress, payer sdk.AccAddress, value *big.Int, bankKeeper utils.BankKeeper, evmKeeper utils.EVMKeeper, hooks *tracing.Hooks, depth int) (sdk.Int, sdk.Int, error) {
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
func GetRemainingGas(ctx sdk.Context, evmKeeper utils.EVMKeeper) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func ExtractMethodID(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Check if the input has at least the length needed for methodID
	return nil, nil
}
