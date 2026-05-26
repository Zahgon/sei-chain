package v614

import (
	"embed"
	"math/big"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v614"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	"github.com/sei-protocol/sei-chain/utils"
)

const (
	ClaimMethod         = "claim"
	ClaimSpecificMethod = "claimSpecific"
)

const SoloAddress = "0x000000000000000000000000000000000000100C"

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var F embed.FS

type PrecompileExecutor struct {
	evmKeeper      putils.EVMKeeper
	bankKeeper     putils.BankKeeper
	accountKeeper  putils.AccountKeeper
	wasmKeeper     putils.WasmdKeeper
	wasmViewKeeper putils.WasmdViewKeeper

	txConfig client.TxConfig

	ClaimMethodID         []byte
	ClaimSpecificMethodID []byte
}

func NewPrecompile(
	keepers putils.Keepers,
) (*pcommon.DynamicGasPrecompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewExecutor(
	a abi.ABI,
	evmKeeper putils.EVMKeeper,
	bankKeeper putils.BankKeeper,
	accountKeeper putils.AccountKeeper,
	wasmKeeper putils.WasmdKeeper,
	wasmViewKeeper putils.WasmdViewKeeper,
	txConfig client.TxConfig,
) *PrecompileExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, _ *tracing.Hooks) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	// Needed to catch gas meter panics
	return nil, 0, nil
}

// depth is incremented upon entering the precompile call so it's
// expected to be 1.

func (p PrecompileExecutor) EVMKeeper() putils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(putils.EVMKeeper)
}

type claimMsg interface {
	GetClaimer() string
	GetSender() string
}

type claimSpecificMsg interface {
	claimMsg
	GetIAssets() []utils.IAsset
}

func (p PrecompileExecutor) Claim(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) ClaimSpecific(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) validate(ctx sdk.Context, caller common.Address, args []interface{}, readOnly bool) (claimMsg, sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(claimMsg), *new(sdk.AccAddress), nil
}

func (p PrecompileExecutor) sigverify(ctx sdk.Context, tx sdk.Tx, claimMsg claimMsg, sender sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// try find pubkey from storage

// increment sequence

func CW20BalanceQueryPayload(addr sdk.AccAddress) []byte { _ = "STUB: not implemented"; return nil }

// should be impossible

func ParseCW20BalanceQueryResponse(res []byte) (sdk.Int, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), nil
}

func CW20TransferPayload(recipient sdk.AccAddress, amount sdk.Int) []byte {
	_ = "STUB: not implemented"
	return nil
}

// should be impossible

func CW721TokensQueryPayload(addr sdk.AccAddress, startAfter string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// should be impossible

func ParseCW721TokensQueryResponse(res []byte) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CW721TransferPayload(recipient sdk.AccAddress, token string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// should be impossible
