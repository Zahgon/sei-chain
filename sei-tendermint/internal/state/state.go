package state

import (
	"time"

	tmstate "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/version"
)

//-----------------------------------------------------------------------------

type Version struct {
	Consensus version.Consensus ` json:"consensus"`
	Software  string            ` json:"software"`
}

// InitStateVersion sets the Consensus.Block and Software versions,
// but leaves the Consensus.App version blank.
// The Consensus.App version will be set during the Handshake, once
// we hear from the app what protocol version it is running.
var InitStateVersion = Version{
	Consensus: version.Consensus{
		Block: version.BlockProtocol,
		App:   0,
	},
	Software: version.TMVersion,
}

func (v *Version) ToProto() tmstate.Version {
	_ = "STUB: not implemented"
	return *new(tmstate.Version)
}

func VersionFromProto(v tmstate.Version) Version { _ = "STUB: not implemented"; return *new(Version) }

//-----------------------------------------------------------------------------

// State is a short description of the latest committed block of the Tendermint consensus.
// It keeps all information necessary to validate new blocks,
// including the last validator set and the consensus params.
// All fields are exposed so the struct can be easily serialized,
// but none of them should be mutated directly.
// Instead, use state.Copy() or updateState(...).
// NOTE: not goroutine-safe.
type State struct {
	// FIXME: This can be removed as TMVersion is a constant, and version.Consensus should
	// eventually be replaced by VersionParams in ConsensusParams
	Version Version

	// immutable
	ChainID       string
	InitialHeight int64 // should be 1, not 0, when starting from height 1

	// LastBlockHeight=0 at genesis (ie. block(H=0) does not exist)
	LastBlockHeight int64
	LastBlockID     types.BlockID
	LastBlockTime   time.Time

	// LastValidators is used to validate block.LastCommit.
	// Validators are persisted to the database separately every time they change,
	// so we can query for historical validator sets.
	// Note that if s.LastBlockHeight causes a valset change,
	// we set s.LastHeightValidatorsChanged = s.LastBlockHeight + 1 + 1
	// Extra +1 due to nextValSet delay.
	NextValidators              *types.ValidatorSet
	Validators                  *types.ValidatorSet
	LastValidators              *types.ValidatorSet
	LastHeightValidatorsChanged int64

	// Consensus parameters used for validating blocks.
	// Changes returned by FinalizeBlock and updated after Commit.
	ConsensusParams                  types.ConsensusParams
	LastHeightConsensusParamsChanged int64

	// Merkle root of the results from executing prev block
	LastResultsHash []byte

	// the latest AppHash we've received from calling abci.Commit()
	AppHash []byte
}

// Copy makes a copy of the State for mutating.
func (state State) Copy() State { _ = "STUB: not implemented"; return *new(State) }

// Equals returns true if the States are identical.
func (state State) Equals(state2 State) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Bytes serializes the State using protobuf, propagating marshaling
// errors
func (state State) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsEmpty returns true if the State is equal to the empty State.
func (state State) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// XXX can't compare to Empty

// ToProto takes the local state type and returns the equivalent proto type
func (state *State) ToProto() (*tmstate.State, error) { _ = "STUB: not implemented"; return nil, nil }

// At Block 1 LastValidators is nil

// FromProto takes a state proto message & returns the local state type
func FromProto(pb *tmstate.State) (*State, error) { _ = "STUB: not implemented"; return nil, nil }

// At Block 1 LastValidators is nil

//------------------------------------------------------------------------
// Create a block from the latest state

// MakeBlock builds a block from the current state with the given txs, commit,
// and evidence. Note it also takes a proposerAddress because the state does not
// track rounds, and hence does not know the correct proposer. TODO: fix this!
func (state State) MakeBlock(
	height int64,
	txs []types.Tx,
	commit *types.Commit,
	evidence []types.Evidence,
	proposerAddress []byte,
) *types.Block {
	_ = "STUB: not implemented"

	// Build base block with block data.
	return nil
}

// Fill rest of header with state data.

//------------------------------------------------------------------------
// Genesis

// MakeGenesisStateFromFile reads and unmarshals state from the given
// file.
//
// Used during replay and in tests.
func MakeGenesisStateFromFile(genDocFile string) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// MakeGenesisDocFromFile reads and unmarshals genesis doc from the given file.
func MakeGenesisDocFromFile(genDocFile string) (*types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MakeGenesisState creates state from types.GenesisDoc.
func MakeGenesisState(genDoc *types.GenesisDoc) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}
