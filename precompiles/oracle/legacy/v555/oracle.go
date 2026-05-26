package v555

import (
	"embed"
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v555"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

const (
	GetExchangeRatesMethod = "getExchangeRates"
	GetOracleTwapsMethod   = "getOracleTwaps"
)

const (
	OracleAddress = "0x0000000000000000000000000000000000001008"
)

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

func GetABI() abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }

type Precompile struct {
	pcommon.Precompile
	evmKeeper    utils.EVMKeeper
	oracleKeeper utils.OracleKeeper
	address      common.Address

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

func (p Precompile) getExchangeRates(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) getOracleTwaps(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert twap to string

func (Precompile) IsTransaction(string) bool { _ = "STUB: not implemented"; return false }
