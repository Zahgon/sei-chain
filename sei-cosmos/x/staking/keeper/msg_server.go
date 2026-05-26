package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

var _ types.MsgServer = msgServer{}

// CreateValidator defines a method for creating a new validator
func (k msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check to see if the pubkey or sender has been registered before

// call the after-creation hook

// move coins from the msg.Address account to a (self-delegation) delegator account
// the validator account and global shares are updated within here
// NOTE source will always be from a wallet which are unbonded

// EditValidator defines a method for editing an existing validator
func (k msgServer) EditValidator(goCtx context.Context, msg *types.MsgEditValidator) (*types.MsgEditValidatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validator must already be registered

// replace all editable fields (clients should autofill existing values)

// call the before-modification hook since we're about to update the commission

// Delegate defines a method for performing a delegation of coins from a delegator to a validator
func (k msgServer) Delegate(goCtx context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: source funds are always unbonded

// BeginRedelegate defines a method for performing a redelegation of coins from a delegator and source validator to a destination validator
func (k msgServer) BeginRedelegate(goCtx context.Context, msg *types.MsgBeginRedelegate) (*types.MsgBeginRedelegateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Undelegate defines a method for performing an undelegation from a delegate and a validator
func (k msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
