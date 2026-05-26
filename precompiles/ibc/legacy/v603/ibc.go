package v603

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"

	pcommon "github.com/sei-protocol/sei-chain/precompiles/common"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
)

const (
	TransferMethod                   = "transfer"
	TransferWithDefaultTimeoutMethod = "transferWithDefaultTimeout"
)

const (
	IBCAddress = "0x0000000000000000000000000000000000001009"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	transferKeeper   utils.TransferKeeper
	evmKeeper        utils.EVMKeeper
	clientKeeper     utils.ClientKeeper
	connectionKeeper utils.ConnectionKeeper
	channelKeeper    utils.ChannelKeeper

	TransferID                   []byte
	TransferWithDefaultTimeoutID []byte
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.DynamicGasPrecompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, _ *tracing.Hooks) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) EVMKeeper() utils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(utils.EVMKeeper)
}

func (p PrecompileExecutor) transfer(ctx sdk.Context, method *abi.Method, args []interface{}, caller common.Address) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// short circuit

func (p PrecompileExecutor) transferWithDefaultTimeout(ctx sdk.Context, method *abi.Method, args []interface{}, caller common.Address) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// short circuit

func (p PrecompileExecutor) accAddressFromArg(ctx sdk.Context, arg interface{}) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

func (p PrecompileExecutor) getChannelConnection(ctx sdk.Context, port string, channelID string) (*connectiontypes.ConnectionEnd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) getConsensusLatestHeight(ctx sdk.Context, connection connectiontypes.ConnectionEnd) (*clienttypes.Height, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetAdjustedHeight(latestConsensusHeight clienttypes.Height) (clienttypes.Height, error) {
	_ = "STUB: not implemented"
	return *new(clienttypes.Height), nil
}

func (p PrecompileExecutor) GetAdjustedTimestamp(ctx sdk.Context, clientId string, height clienttypes.Height) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type ValidatedArgs struct {
	senderSeiAddr         sdk.AccAddress
	receiverAddressString string
	port                  string
	channelID             string
	denom                 string
	amount                *big.Int
}

func (p PrecompileExecutor) validateCommonArgs(ctx sdk.Context, args []interface{}, caller common.Address) (*ValidatedArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addMemo(memoArg interface{}, transferMsg types.MsgTransfer) types.MsgTransfer {
	_ = "STUB: not implemented"
	return *new(types.MsgTransfer)
}
