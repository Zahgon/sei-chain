package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// LaneProposal .
type LaneProposal struct {
	utils.ReadOnly
	block *Block
}

// NewLaneProposal constructs a new LaneProposal.
func NewLaneProposal(block *Block) *LaneProposal { _ = "STUB: not implemented"; return nil }

// Block .
func (m *LaneProposal) Block() *Block {
	_ = "STUB: not implemented"

	// Verify verifies that the LaneProposal is consistent with the Committee.
	return nil
}

func (m *LaneProposal) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// LaneProposalConv is a protobuf converter for LaneProposal.
var LaneProposalConv = protoutils.Conv[*LaneProposal, *pb.Block]{
	Encode: func(m *LaneProposal) *pb.Block {
		return BlockConv.Encode(m.block)
	},
	Decode: func(m *pb.Block) (*LaneProposal, error) {
		block, err := BlockConv.Decode(m)
		if err != nil {
			return nil, fmt.Errorf("block: %w", err)
		}
		return &LaneProposal{block: block}, nil
	},
}
