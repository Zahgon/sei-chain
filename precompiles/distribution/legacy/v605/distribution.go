package v605

import (
	"embed"
	"math/big"

	distrtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v605"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

const (
	SetWithdrawAddressMethod                = "setWithdrawAddress"
	WithdrawDelegationRewardsMethod         = "withdrawDelegationRewards"
	WithdrawMultipleDelegationRewardsMethod = "withdrawMultipleDelegationRewards"
	RewardsMethod                           = "rewards"
)

const (
	DistrAddress = "0x0000000000000000000000000000000000001007"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	distrKeeper utils.DistributionKeeper
	evmKeeper   utils.EVMKeeper
	address     common.Address

	SetWithdrawAddrID                   []byte
	WithdrawDelegationRewardsID         []byte
	WithdrawMultipleDelegationRewardsID []byte
	RewardsID                           []byte
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.DynamicGasPrecompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, hooks *tracing.Hooks) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) EVMKeeper() utils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(utils.EVMKeeper)
}

func (p PrecompileExecutor) setWithdrawAddress(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) withdrawDelegationRewards(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) validateInput(value *big.Int, args []interface{}, expectedArgsLength int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PrecompileExecutor) withdraw(ctx sdk.Context, delegator sdk.AccAddress, validatorAddress string) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

func (p PrecompileExecutor) getDelegator(ctx sdk.Context, caller common.Address) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

func (p PrecompileExecutor) withdrawMultipleDelegationRewards(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) accAddressFromArg(ctx sdk.Context, arg interface{}) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

type Coin struct {
	Amount   *big.Int
	Denom    string
	Decimals *big.Int
}

type Reward struct {
	ValidatorAddress string
	Coins            []Coin
}

type Rewards struct {
	Rewards []Reward
	Total   []Coin
}

func (p PrecompileExecutor) rewards(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func getResponseOutput(response *distrtypes.QueryDelegationTotalRewardsResponse) Rewards {
	_ = "STUB: not implemented"
	return *new(Rewards)
}
