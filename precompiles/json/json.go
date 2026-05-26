package json

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	ExtractAsBytesMethod          = "extractAsBytes"
	ExtractAsBytesListMethod      = "extractAsBytesList"
	ExtractAsUint256Method        = "extractAsUint256"
	ExtractAsBytesFromArrayMethod = "extractAsBytesFromArray"
)

const JSONAddress = "0x0000000000000000000000000000000000001003"
const GasCostPerByte = 100 // TODO: parameterize

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	ExtractAsBytesID          []byte
	ExtractAsBytesListID      []byte
	ExtractAsUint256ID        []byte
	ExtractAsBytesFromArrayID []byte
}

func NewPrecompile(keepers putils.Keepers) (*pcommon.Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) extractAsBytes(_ sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// in the case of a string value, remove the quotes

func (p PrecompileExecutor) extractAsBytesList(_ sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

func (p PrecompileExecutor) ExtractAsUint256(_ sdk.Context, _ *abi.Method, args []interface{}, value *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// Assuming result is your byte slice
// Convert byte slice to string and trim quotation marks

// Convert the string to big.Int

func (p PrecompileExecutor) extractAsBytesFromArray(_ sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// in the case of a string value, remove the quotes
