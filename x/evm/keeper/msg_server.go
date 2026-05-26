package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/x/evm/state"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

var logger = seilog.NewLogger("x", "evm", "keeper")

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

var _ types.MsgServer = msgServer{}

func (k *Keeper) PrepareCtxForEVMTransaction(ctx sdk.Context, tx *ethtypes.Transaction) (sdk.Context, sdk.GasMeter) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), *new(sdk.GasMeter)
}

// EVM has a special case here, mainly because for an EVM transaction the gas limit is set on EVM payload level, not on top-level GasWanted field
// as normal transactions (because existing eth client can't). As a result EVM has its own dedicated ante handler chain. The full sequence is:

// 	1. At the beginning of the ante handler chain, gas meter is set to infinite so that the ante processing itself won't run out of gas (EVM ante is pretty light but it does read a parameter or two)
// 	2. At the end of the ante handler chain, gas meter is set based on the gas limit specified in the EVM payload; this is only to provide a GasWanted return value to tendermint mempool when CheckTx returns, and not used for anything else.
// 	3. At the beginning of message server (here), gas meter is set to infinite again, because EVM internal logic will then take over and manage out-of-gas scenarios.
// 	4. At the end of message server, gas consumed by EVM is adjusted to Sei's unit and counted in the original gas meter, because that original gas meter will be used to count towards block gas after message server returns

func (server msgServer) EVMTransaction(goCtx context.Context, msg *types.MsgEVMTransaction) (serverRes *types.MsgEVMTransactionResponse, err error) {
	_ = "STUB: not implemented"
	return nil,

		// no-op in msg server for associate tx; all the work have been done in ante handler
		nil
}

// TODO(PLT-330): remove once evm_panics_total verified

// TODO(PLT-330): remove once evm_errors_total verified

// TODO(PLT-330): remove once evm_errors_total verified

//nolint:gosec

//nolint:gosec

// TODO(PLT-330): remove once evm_errors_total verified

// Add metrics for receipt status

// TODO(PLT-330): remove once evm_receipt_status_total verified

// TODO(PLT-330): remove once evm_receipt_status_total verified

// GasUsed in serverRes is in EVM's gas unit, not Sei's gas unit.
// PriorityNormalizer is the coefficient that's used to adjust EVM
// transactions' priority, which is based on gas limit in EVM unit,
// to Sei transactions' priority, which is based on gas limit in
// Sei unit, so we use the same coefficient to convert gas unit here.
//nolint:gosec

// This should not happen, as anything that could cause applyErr is supposed to
// be checked in CheckTx first

// TODO(PLT-330): remove once evm_errors_total verified

// if applyErr is nil then res must be non-nil

// TODO(PLT-330): remove once evm_errors_total verified

func (k *Keeper) GetGasPool() core.GasPool { _ = "STUB: not implemented"; return *new(core.GasPool) }

func (k *Keeper) GetEVMMessage(ctx sdk.Context, tx *ethtypes.Transaction, sender common.Address) *core.Message {
	_ = "STUB: not implemented"
	return nil
}

// If baseFee provided, set gasPrice to effectiveGasPrice.

func (k Keeper) applyEVMMessage(ctx sdk.Context, msg *core.Message, stateDB *state.DBImpl, gp core.GasPool, shouldIncrementNonce bool) (*core.ExecutionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fee already charged in ante handler

func (server msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (server msgServer) RegisterPointer(goCtx context.Context, msg *types.MsgRegisterPointer) (*types.MsgRegisterPointerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (server msgServer) AssociateContractAddress(goCtx context.Context, msg *types.MsgAssociateContractAddress) (*types.MsgAssociateContractAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// already validated
// check if address is for a contract

func (server msgServer) Associate(context.Context, *types.MsgAssociate) (*types.MsgAssociateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
