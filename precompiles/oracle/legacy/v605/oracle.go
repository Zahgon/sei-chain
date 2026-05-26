package v605

import (
	"embed"
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v605"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

const (
	GetExchangeRatesMethod = "getExchangeRates"
	GetOracleTwapsMethod   = "getOracleTwaps"
)

const (
	OracleAddress = "0x0000000000000000000000000000000000001008"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	evmKeeper    utils.EVMKeeper
	oracleKeeper utils.OracleKeeper

	GetExchangeRatesId []byte
	GetOracleTwapsId   []byte
}

// Define types which deviate slightly from cosmos types (ExchangeRate string vs sdk.Dec)
type OracleExchangeRate struct {
	ExchangeRate        string `json:"exchangeRate"`
	LastUpdate          string `json:"lastUpdate"`
	LastUpdateTimestamp int64  `json:"lastUpdateTimestamp"`
}

type DenomOracleExchangeRatePair struct {
	Denom                 string             `json:"denom"`
	OracleExchangeRateVal OracleExchangeRate `json:"oracleExchangeRateVal"`
}

type OracleTwap struct {
	Denom           string `json:"denom"`
	Twap            string `json:"twap"`
	LookbackSeconds int64  `json:"lookbackSeconds"`
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

func (p PrecompileExecutor) getExchangeRates(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) getOracleTwaps(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Convert twap to string

func (p PrecompileExecutor) EVMKeeper() utils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(utils.EVMKeeper)
}

func (PrecompileExecutor) IsTransaction(string) bool { _ = "STUB: not implemented"; return false }
