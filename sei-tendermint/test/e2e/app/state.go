// nolint: gosec
package app

import (
	"sync"
)

const stateFileName = "app_state.json"
const prevStateFileName = "prev_app_state.json"

// State is the application state.
type State struct {
	sync.RWMutex
	Height uint64
	Values map[string]string
	Hash   []byte

	// private fields aren't marshaled to disk.
	currentFile string
	// app saves current and previous state for rollback functionality
	previousFile    string
	persistInterval uint64
	initialHeight   uint64
}

// NewState creates a new state.
func NewState(dir string, persistInterval uint64) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// load loads state from disk. It does not take out a lock, since it is called
// during construction.
func (s *State) load() error { _ = "STUB: not implemented"; return nil }

// if the current state doesn't exist then we try recover from the previous state

// save saves the state to disk. It does not take out a lock since it is called
// internally by Commit which does lock.
func (s *State) save() error { _ = "STUB: not implemented"; return nil }

// We write the state to a separate file and move it to the destination, to
// make it atomic.

// We take the current state and move it to the previous state, replacing it

// Finally, we take the new state and replace the current state.

// Export exports key/value pairs as JSON, used for state sync snapshots.
func (s *State) Export() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Import imports key/value pairs from JSON bytes, used for InitChain.AppStateBytes and
// state sync snapshots. It also saves the state once imported.
func (s *State) Import(height uint64, jsonBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Get fetches a value. A missing value is returned as an empty string.
func (s *State) Get(key string) string { _ = "STUB: not implemented"; return "" }

// Set sets a value. Setting an empty value is equivalent to deleting it.
func (s *State) Set(key, value string) { _ = "STUB: not implemented"; return }

// Finalize is called after applying a block, updating the height and returning the new app_hash
func (s *State) Finalize() []byte { _ = "STUB: not implemented"; return nil }

// Commit commits the current state.
func (s *State) Commit() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Checks if rollback will succeed.
func (s *State) CanRollback() bool { _ = "STUB: not implemented"; return false }

func (s *State) Rollback() error { _ = "STUB: not implemented"; return nil }

// hashItems hashes a set of key/value items.
func hashItems(items map[string]string, height uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}
