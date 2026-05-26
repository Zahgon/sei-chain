package v575

import (
	"embed"

	"math/big"

	putils "github.com/sei-protocol/sei-chain/precompiles/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v575"
)

const (
	GetSeiAddressMethod = "getSeiAddr"
	GetEvmAddressMethod = "getEvmAddr"
	Associate           = "associate"
)

const (
	AddrAddress = "0x0000000000000000000000000000000000001004"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	evmKeeper     putils.EVMKeeper
	bankKeeper    putils.BankKeeper
	accountKeeper putils.AccountKeeper

	GetSeiAddressID []byte
	GetEvmAddressID []byte
	AssociateID     []byte
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

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, _ common.Address, _ common.Address, args []interface{}, value *big.Int, _ bool, _ *vm.EVM, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) getSeiAddr(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) getEvmAddr(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) associate(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// v, r and s are components of a signature over the customMessage sent.
// We use the signature to construct the user's pubkey to obtain their addresses.

// Derive addresses

// Check that address is not already associated

// Associate Addresses:

func (PrecompileExecutor) IsTransaction(method string) bool {
	_ = "STUB: not implemented"
	return false
}

func decodeHexString(hexString string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
