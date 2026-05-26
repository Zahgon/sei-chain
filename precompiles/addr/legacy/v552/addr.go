package v552

import (
	"embed"
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v552"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

const (
	GetSeiAddressMethod = "getSeiAddr"
	GetEvmAddressMethod = "getEvmAddr"
)

const (
	AddrAddress = "0x0000000000000000000000000000000000001004"
)

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type Precompile struct {
	pcommon.Precompile
	evmKeeper utils.EVMKeeper
	address   common.Address

	GetSeiAddressID []byte
	GetEvmAddressID []byte
}

func NewPrecompile(keepers utils.Keepers) (*Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// This should never happen since this method is going to fail during Run

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) Run(evm *vm.EVM, _ common.Address, _ common.Address, input []byte, value *big.Int, _ bool, _ bool, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) getSeiAddr(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) getEvmAddr(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (Precompile) IsTransaction(string) bool { _ = "STUB: not implemented"; return false }
