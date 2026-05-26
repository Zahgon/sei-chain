package giga

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/consensus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/rpc"
)

type Service struct {
	getBlockReqs chan req
	state        *consensus.State
}

func NewService(state *consensus.State) *Service { _ = "STUB: not implemented"; return nil }

func (x *Service) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (x *Service) RunServer(ctx context.Context, server rpc.Server[API]) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *Service) RunClient(ctx context.Context, client rpc.Client[API]) error {
	_ = "STUB: not implemented"
	return nil
}
