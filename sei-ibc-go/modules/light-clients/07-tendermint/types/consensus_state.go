package types

import (
	"time"

	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"

	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// SentinelRoot is used as a stand-in root value for the consensus state set at the upgrade height
const SentinelRoot = "sentinel_root"

// NewConsensusState creates a new ConsensusState instance.
func NewConsensusState(
	timestamp time.Time, root commitmenttypes.MerkleRoot, nextValsHash tmbytes.HexBytes,
) *ConsensusState {
	_ = "STUB: not implemented"
	return nil
}

// ClientType returns Tendermint
func (ConsensusState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetRoot returns the commitment Root for the specific
func (cs ConsensusState) GetRoot() exported.Root {
	_ = "STUB: not implemented"

	// GetTimestamp returns block time in nanoseconds of the header that created consensus state
	return *new(exported.Root)
}

func (cs ConsensusState) GetTimestamp() uint64 {
	_ = "STUB: not implemented"
	// #nosec G115 --- checked at state population.
	return 0
}

// ValidateBasic defines a basic validation for the tendermint consensus state.
// NOTE: ProcessedTimestamp may be zero if this is an initial consensus state passed in by relayer
// as opposed to a consensus state constructed by the chain.
func (cs ConsensusState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
