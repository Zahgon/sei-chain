package v620

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v620"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	VoteMethod           = "vote"
	VoteWeightedMethod   = "voteWeighted"
	DepositMethod        = "deposit"
	SubmitProposalMethod = "submitProposal"
)

const (
	GovAddress = "0x0000000000000000000000000000000000001006"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	govMsgServer     utils.GovMsgServer
	evmKeeper        utils.EVMKeeper
	bankKeeper       utils.BankKeeper
	address          common.Address
	proposalHandlers map[string]ProposalHandler

	VoteID           []byte
	VoteWeightedID   []byte
	DepositID        []byte
	SubmitProposalID []byte
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.Precompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register proposal handlers

// Register method IDs

// Create the precompile

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// This should never happen since this is going to fail during Run

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, hooks *tracing.Hooks) (bz []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) vote(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) voteWeighted(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// args[1] is the struct array for weighted vote options
// The ABI decoder gives us the actual struct slice

// Convert to WeightedVoteOptions

// Parse weight as decimal

func (p PrecompileExecutor) deposit(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) submitProposal(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse the proposal JSON

// Create the proposal content using the handler system

// Create the MsgSubmitProposal

// Validate the Msg

// Create a MsgServer context

// Submit the proposal using the MsgServer

// Return the proposal ID
