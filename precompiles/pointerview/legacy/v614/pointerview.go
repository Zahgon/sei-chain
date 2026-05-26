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
	GetNativePointer = "getNativePointer"
	GetCW20Pointer   = "getCW20Pointer"
	GetCW721Pointer  = "getCW721Pointer"
	GetCW1155Pointer = "getCW1155Pointer"
)

const PointerViewAddress = "0x000000000000000000000000000000000000100A"

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	evmKeeper utils.EVMKeeper

	GetNativePointerID []byte
	GetCW20PointerID   []byte
	GetCW721PointerID  []byte
	GetCW1155PointerID []byte
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas([]byte, *abi.Method) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, hooks *tracing.Hooks) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) GetNative(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) GetCW20(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) GetCW721(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) GetCW1155(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
