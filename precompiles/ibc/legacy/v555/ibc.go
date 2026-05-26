package v555

import (
	"embed"
	"math/big"

	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v555"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
)

const (
	TransferMethod                   = "transfer"
	TransferWithDefaultTimeoutMethod = "transferWithDefaultTimeout"
)

const (
	IBCAddress = "0x0000000000000000000000000000000000001009"
)

var _ vm.PrecompiledContract = &Precompile{}
var _ vm.DynamicGasPrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

func GetABI() abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }

type Precompile struct {
	pcommon.Precompile
	address          common.Address
	transferKeeper   putils.TransferKeeper
	evmKeeper        putils.EVMKeeper
	clientKeeper     putils.ClientKeeper
	connectionKeeper putils.ConnectionKeeper
	channelKeeper    putils.ChannelKeeper

	TransferID                   []byte
	TransferWithDefaultTimeoutID []byte
}

func NewPrecompile(keepers putils.Keepers) (*Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// This should never happen since this method is going to fail during Run

func (p Precompile) RunAndCalculateGas(evm *vm.EVM, caller common.Address, callingContract common.Address, input []byte, suppliedGas uint64, _ *big.Int, _ *tracing.Hooks, readOnly bool, _ bool) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p Precompile) Run(*vm.EVM, common.Address, common.Address, []byte, *big.Int, bool, bool, *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) transfer(ctx sdk.Context, method *abi.Method, args []interface{}, caller common.Address) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// short circuit

func (p Precompile) transferWithDefaultTimeout(ctx sdk.Context, method *abi.Method, args []interface{}, caller common.Address) (ret []byte, remainingGas uint64, rerr error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// short circuit

func (Precompile) IsTransaction(method string) bool { _ = "STUB: not implemented"; return false }

func (p Precompile) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (p Precompile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p Precompile) accAddressFromArg(ctx sdk.Context, arg interface{}) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

func (p Precompile) getChannelConnection(ctx sdk.Context, port string, channelID string) (*connectiontypes.ConnectionEnd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Precompile) getConsensusLatestHeight(ctx sdk.Context, connection connectiontypes.ConnectionEnd) (*clienttypes.Height, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetAdjustedHeight(latestConsensusHeight clienttypes.Height) (clienttypes.Height, error) {
	_ = "STUB: not implemented"
	return *new(clienttypes.Height), nil
}

func (p Precompile) GetAdjustedTimestamp(ctx sdk.Context, clientId string, height clienttypes.Height) (uint64, error) {
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

func (p Precompile) validateCommonArgs(ctx sdk.Context, args []interface{}, caller common.Address) (*ValidatedArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addMemo(memoArg interface{}, transferMsg types.MsgTransfer) types.MsgTransfer {
	_ = "STUB: not implemented"
	return *new(types.MsgTransfer)
}
