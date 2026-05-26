package avail

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
)

func (s *State) SubscribeLaneProposals(first types.BlockNumber) *LaneProposalsRecv {
	_ = "STUB: not implemented"
	return nil
}

type LaneProposalsRecv struct {
	state *State
	lane  types.LaneID
	next  types.BlockNumber
}

func (r *LaneProposalsRecv) Recv(ctx context.Context) (*types.Signed[*types.LaneProposal], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) SubscribeLaneVotes() *LaneVotesRecv { _ = "STUB: not implemented"; return nil }

type LaneVotesRecv struct {
	state *State
	next  map[types.LaneID]types.BlockNumber
}

func (r *LaneVotesRecv) RecvBatch(ctx context.Context) ([]*types.Signed[*types.LaneVote], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(gprusak): we sign the votes per VotesRecv instance, which is suboptimal.
// We should sign votes as soon as blocks arrive and cache them (in PushBlock and ProduceBlock).

type AppVotesRecv struct {
	state *State
	next  types.GlobalBlockNumber
}

func (s *State) SubscribeAppVotes() *AppVotesRecv { _ = "STUB: not implemented"; return nil }

func (r *AppVotesRecv) Recv(ctx context.Context) (*types.Signed[*types.AppVote], error) {
	_ = "STUB: not implemented"

	// If needed, fast forward to the first global number without known AppQC.
	return nil, nil
}

// Fetch the proposal.

// AppProposal currently might return a proposal with a higher global number than the one we requested.
// Correct the n in such a case.
// TODO(gprusak): perhaps it would be possible to require AppHash at every block from the execution engine.
// This would simplify the data state.
