package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the gov MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SubmitProposal(goCtx context.Context, msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k msgServer) VoteWeighted(goCtx context.Context, msg *types.MsgVoteWeighted) (*types.MsgVoteWeightedResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
