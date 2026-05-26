package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// AppHash represents EVM state hash.
// We don't even know here how long the hash can be.
type AppHash []byte

// AppProposal .
type AppProposal struct {
	utils.ReadOnly
	globalNumber GlobalBlockNumber
	roadIndex    RoadIndex
	appHash      AppHash
}

// NewAppProposal creates a new AppProposal.
func NewAppProposal(globalNumber GlobalBlockNumber, roadIndex RoadIndex, appHash AppHash) *AppProposal {
	_ = "STUB: not implemented"
	return nil
}

// GlobalNumber .
func (m *AppProposal) GlobalNumber() GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *

	// RoadIndex returns the road index of the proposal.
	new(GlobalBlockNumber)
}

func (m *AppProposal) RoadIndex() RoadIndex {
	_ = "STUB: not implemented"

	// AppHash .
	return *new(RoadIndex)
}

func (m *AppProposal) AppHash() AppHash {
	_ = "STUB: not implemented"

	// Next is the next global block number to compute AppHash for.
	return *new(AppHash)
}

func (m *AppProposal) Next() RoadIndex {
	_ = "STUB: not implemented"
	return *

	// Verify verifies that the AppProposal is consistent with the CommitQC.
	new(RoadIndex)
}

func (m *AppProposal) Verify(c *Committee, qc *CommitQC) error {
	_ = "STUB: not implemented"
	return nil
}

// AppProposalConv is a protobuf converter for AppProposal.
var AppProposalConv = protoutils.Conv[*AppProposal, *pb.AppProposal]{
	Encode: func(m *AppProposal) *pb.AppProposal {
		return &pb.AppProposal{
			GlobalNumber: utils.Alloc(uint64(m.globalNumber)),
			RoadIndex:    utils.Alloc(uint64(m.roadIndex)),
			AppHash:      m.appHash,
		}
	},
	Decode: func(m *pb.AppProposal) (*AppProposal, error) {
		if m.GlobalNumber == nil {
			return nil, fmt.Errorf("GlobalNumber: missing")
		}
		if m.RoadIndex == nil {
			return nil, fmt.Errorf("RoadIndex: missing")
		}
		return &AppProposal{
			globalNumber: GlobalBlockNumber(*m.GlobalNumber),
			roadIndex:    RoadIndex(*m.RoadIndex),
			appHash:      AppHash(m.AppHash),
		}, nil
	},
}
