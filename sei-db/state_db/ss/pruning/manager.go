package pruning

import (
	"sync"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("db", "state-db", "ss", "pruning")

type Manager struct {
	stateStore    types.StateStore
	keepRecent    int64
	pruneInterval int64

	// Lifecycle management
	startOnce sync.Once
	stopCh    chan struct{}
	stopOnce  sync.Once
	wg        sync.WaitGroup
}

// NewPruningManager creates a new pruning manager for state store
// Pruning Manager will periodically prune state store based on keep-recent and prune-interval configs.
func NewPruningManager(
	stateStore types.StateStore,
	keepRecent int64,
	pruneInterval int64,
) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) Start() { _ = "STUB: not implemented"; return }

// Stop gracefully stops the pruning goroutine and waits for it to exit.
// Safe to call multiple times (idempotent).
func (m *Manager) Stop() { _ = "STUB: not implemented"; return }

// Safe: WaitGroup.Wait() is idempotent when counter is 0

func (m *Manager) pruneLoop() { _ = "STUB: not implemented"; return }

// Check for stop signal before pruning

// prune all versions up to and including the pruneVersion

// Generate a random percentage (between 0% and 100%) of the fixed interval as a delay

// Wait with stop signal check

// Continue to next iteration
