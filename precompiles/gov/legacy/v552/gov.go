package v552

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v552"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	VoteMethod    = "vote"
	DepositMethod = "deposit"
)

const (
	GovAddress = "0x0000000000000000000000000000000000001006"
)

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

func GetABI() abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }

type Precompile struct {
	pcommon.Precompile
	govKeeper  utils.GovKeeper
	evmKeeper  utils.EVMKeeper
	bankKeeper utils.BankKeeper
	address    common.Address

	VoteID    []byte
	DepositID []byte
}

func NewPrecompile(keepers utils.Keepers) (*Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// This should never happen since this is going to fail during Run

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) Run(evm *vm.EVM, caller common.Address, callingContract common.Address, input []byte, value *big.Int, readOnly bool, _ bool, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) vote(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) deposit(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
