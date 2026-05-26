package v65

import (
	"embed"
	"errors"
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v65"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

const (
	GetExchangeRatesMethod = "getExchangeRates"
	GetOracleTwapsMethod   = "getOracleTwaps"
)

const (
	OracleAddress = "0x0000000000000000000000000000000000001008"
)

var ErrOraclePrecompileRetired = errors.New("oracle precompile is retired; oracle data queries are disabled")

var oracleRetiredRevertData = mustEncodeRevertReason(ErrOraclePrecompileRetired.Error())

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	evmKeeper utils.EVMKeeper

	GetExchangeRatesId []byte
	GetOracleTwapsId   []byte
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.DynamicGasPrecompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, hooks *tracing.Hooks) (bz []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	// Needed to catch gas meter panics
	return nil, 0, nil
}

func (p PrecompileExecutor) getExchangeRates(ctx sdk.Context, _ *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) getOracleTwaps(ctx sdk.Context, _ *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) EVMKeeper() utils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(utils.EVMKeeper)
}

func (PrecompileExecutor) IsTransaction(string) bool { _ = "STUB: not implemented"; return false }

func mustEncodeRevertReason(reason string) []byte { _ = "STUB: not implemented"; return nil }
