package commands

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
)

// LoadTendermintState loads the tendermint state from the database.
// Returns the state, or an error if loading fails.
func LoadTendermintState(config *config.Config) (state.State, error) {
	_ = "STUB: not implemented"
	return *new(state.State), nil
}

// RollbackStateToTargetHeight rolls back the tendermint state to the target height.
// It repeatedly calls state.Rollback until the target height is reached.
func RollbackStateToTargetHeight(config *config.Config, removeBlock bool, targetHeight int64) (int64, []byte, error) {
	_ = "STUB: not implemented"
	// use the parsed config to load the block and state store
	return 0, nil, nil
}

// Get initial state to verify we are above target height

// Since state.Rollback modifies the store, we can just call it in a loop.
