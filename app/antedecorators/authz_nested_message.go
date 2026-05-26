package antedecorators

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

// maxNestedMsgs defines a cap for the number of nested messages on a MsgExec message
const maxNestedMsgs = 5

type AuthzNestedMessageDecorator struct{}

func NewAuthzNestedMessageDecorator() AuthzNestedMessageDecorator {
	_ = "STUB: not implemented"
	return *new(AuthzNestedMessageDecorator)
}

func (ad AuthzNestedMessageDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// find nested evm messages

func (ad AuthzNestedMessageDecorator) CheckAuthzContainsEvm(ctx sdk.Context, authzMsg *authz.MsgExec, nestedLvl int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check if message type is authz exec or evm

// find nested to check for evm
