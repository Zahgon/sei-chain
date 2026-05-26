package v555

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v555"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	GetNativePointer = "getNativePointer"
	GetCW20Pointer   = "getCW20Pointer"
	GetCW721Pointer  = "getCW721Pointer"
)

const PointerViewAddress = "0x000000000000000000000000000000000000100A"

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type Precompile struct {
	pcommon.Precompile
	evmKeeper utils.EVMKeeper
	address   common.Address

	GetNativePointerID []byte
	GetCW20PointerID   []byte
	GetCW721PointerID  []byte
}

func NewPrecompile(keepers utils.Keepers) (*Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) Run(evm *vm.EVM, _ common.Address, _ common.Address, input []byte, _ *big.Int, _ bool, _ bool, hooks *tracing.Hooks) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) GetNative(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) GetCW20(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) GetCW721(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
