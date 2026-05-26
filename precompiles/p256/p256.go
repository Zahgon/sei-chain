package p256

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/core/tracing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	VerifyMethod = "verify"
)

const (
	P256VerifyAddress = "0x0000000000000000000000000000000000001011"
	GasCostPerByte    = 300
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	VerifyID []byte
}

func NewPrecompile(utils.Keepers) (*pcommon.Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, hooks *tracing.Hooks) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verify verifies the secp256r1 signature
// Implements https://github.com/ethereum/RIPs/blob/master/RIPS/rip-7212.md
func (p PrecompileExecutor) verify(ctx sdk.Context, method *abi.Method, args []interface{}, caller common.Address) (ret []byte, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Required input length is 160 bytes

// Check the input length

// Input length is invalid

// Extract the hash, r, s, x, y from the input

// Verify the secp256r1 signature

// Signature is valid

// Signature is invalid
