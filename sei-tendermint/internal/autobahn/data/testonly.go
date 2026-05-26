package data

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

func TestAppQC(keys []types.SecretKey, proposal *types.AppProposal) *types.AppQC {
	_ = "STUB: not implemented"
	return nil
}

func TestLaneQC(keys []types.SecretKey, header *types.BlockHeader) *types.LaneQC {
	_ = "STUB: not implemented"
	return nil
}

func TestCommitQC(
	rng utils.Rng,
	committee *types.Committee,
	keys []types.SecretKey,
	prev utils.Option[*types.CommitQC],
) (*types.FullCommitQC, []*types.Block) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make some blocks

// Construct a proposal.

var _ StateAPI = (*MockState)(nil)

type innerMockState struct {
	blocks map[types.GlobalBlockNumber]*types.GlobalBlock // [first,next)
	first  types.GlobalBlockNumber
	next   types.GlobalBlockNumber
}

// MockState is a mock implementation of the StateAPI interface.
// Allows for pushing global blocks directly (without going through consensus).
type MockState struct {
	capacity uint64
	inner    utils.Watch[*innerMockState]
}

// NewMockState creates a new MockState with the given block capacity.
func NewMockState(capacity uint64) *MockState { _ = "STUB: not implemented"; return nil }

// GlobalBlock returns the global block with the given number.
func (s *MockState) GlobalBlock(ctx context.Context, n types.GlobalBlockNumber) (*types.GlobalBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProduceBlock appends a new global block with the given payload.
func (s *MockState) ProduceBlock(ctx context.Context, payload *types.Payload) error {
	_ = "STUB: not implemented"
	return nil
}

// PushAppHash marks all blocks up to n as executed.
func (s *MockState) PushAppHash(_ context.Context, n types.GlobalBlockNumber, appHash types.AppHash) error {
	_ = "STUB: not implemented"
	return nil
}

// Describe from prometheus.Collector.
func (s *MockState) Describe(chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"

	// Collect from prometheus.Collector.
	return
}

func (s *MockState) Collect(chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }
