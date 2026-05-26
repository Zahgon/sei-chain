package v552

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v552"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	ExtractAsBytesMethod     = "extractAsBytes"
	ExtractAsBytesListMethod = "extractAsBytesList"
	ExtractAsUint256Method   = "extractAsUint256"
)

const JSONAddress = "0x0000000000000000000000000000000000001003"
const GasCostPerByte = 100 // TODO: parameterize

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type Precompile struct {
	pcommon.Precompile
	address common.Address

	ExtractAsBytesID     []byte
	ExtractAsBytesListID []byte
	ExtractAsUint256ID   []byte
}

func NewPrecompile(keepers putils.Keepers) (*Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec

func (Precompile) IsTransaction(string) bool { _ = "STUB: not implemented"; return false }

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) Run(evm *vm.EVM, _ common.Address, _ common.Address, input []byte, value *big.Int, _ bool, _ bool, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) extractAsBytes(_ sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// in the case of a string value, remove the quotes

func (p Precompile) extractAsBytesList(_ sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

func (p Precompile) ExtractAsUint256(_ sdk.Context, _ *abi.Method, args []interface{}, value *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type assertion will always succeed because it's already validated in p.Prepare call in Run()

// Assuming result is your byte slice
// Convert byte slice to string and trim quotation marks

// Convert the string to big.Int
