package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// Validators gets the validator set at the given block height.
//
// If no height is provided, it will fetch the latest validator set. Note the
// validators are sorted by their voting power - this is the canonical order
// for the validators in the set as used in computing their Merkle root.
//
// More: https://docs.tendermint.com/master/rpc/#/Info/validators
//
// Under Autobahn the CometBFT StateStore is not populated, but the committee
// is fixed at genesis (no validator-updates path under Autobahn), so any
// retained height returns the genesis committee. Pagination + error shape
// mirror the CometBFT path so external tools see consistent responses.
func (env *Environment) Validators(ctx context.Context, req *coretypes.RequestValidators) (*coretypes.ResultValidators, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProposerPriority left at 0; Autobahn uses round-robin
// per-block leader selection (Committee.Leader), not
// Tendermint's proposer-priority accumulator.

// The latest validator that we know is the NextValidator of the last block.

// DumpConsensusState dumps consensus state.
// UNSTABLE
// More: https://docs.tendermint.com/master/rpc/#/Info/dump_consensus_state
func (env *Environment) DumpConsensusState(ctx context.Context) (*coretypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	// Get Peer consensus states.
	return nil, nil
}

// Peer basic info.

// Peer consensus state.

// Get self round state.

// ConsensusState returns a concise summary of the consensus state.
// UNSTABLE
// More: https://docs.tendermint.com/master/rpc/#/Info/consensus_state
func (env *Environment) GetConsensusState(ctx context.Context) (*coretypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	// Get self round state.
	return nil, nil
}

// ConsensusParams gets the consensus parameters at the given block height.
// If no height is provided, it will fetch the latest consensus params.
// More: https://docs.tendermint.com/master/rpc/#/Info/consensus_params
func (env *Environment) ConsensusParams(ctx context.Context, req *coretypes.RequestConsensusParams) (*coretypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	// The latest consensus params that we know is the consensus params after
	// the last block.
	return nil, nil
}
