package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

func (ms msgServer) AggregateExchangeRateVote(goCtx context.Context, msg *types.MsgAggregateExchangeRateVote) (*types.MsgAggregateExchangeRateVoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check all denoms are in the vote target

func (ms msgServer) DelegateFeedConsent(goCtx context.Context, msg *types.MsgDelegateFeedConsent) (*types.MsgDelegateFeedConsentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check the delegator is a validator

// Set the delegation
