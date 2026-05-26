package upgrades

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("app", "upgrades")

// Chain-ID constants for use in hard fork handlers.
const (
	ChainIDSeiHardForkTest = "sei-hard-fork-test"
)

type HardForkHandler interface {
	// a unique identifying name to ensure no duplicate handlers are registered
	GetName() string
	// The target chain ID for which to run this handler (which chain to run on)
	GetTargetChainID() string
	// The target height at which the handler should be executed
	GetTargetHeight() int64
	// An execution function used to process a hard fork handler
	ExecuteHandler(ctx sdk.Context) error
}

type HardForkManager struct {
	chainID       string
	handlerMap    map[int64][]HardForkHandler
	uniqueNameMap map[string]struct{}
}

// Create a new hard fork manager for a given chain ID.
// This will filter out any handlers for other chain IDs,
// and create a map that maps between target height and handlers to run for the given heights.
func NewHardForkManager(chainID string) *HardForkManager { _ = "STUB: not implemented"; return nil }

// This tries to register a handler with the hard fork manager.
// If the handler's target chain ID doesn't match the chain ID of the manager,
// this is a no-op and the handler is ignored.
func (hfm *HardForkManager) RegisterHandler(handler HardForkHandler) {
	_ = "STUB: not implemented"
	return
}

// we already have a migration with this name - panic

// register name for uniqueness assertion

// This returns a boolean indicating whether or not the current height is a
// target height for which there are hard fork handlers to run.
func (hfm *HardForkManager) TargetHeightReached(ctx sdk.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// This executes the hard fork handlers for the current height. This will function will panic upon receiving an error during hard fork handler execution
func (hfm *HardForkManager) ExecuteForTargetHeight(ctx sdk.Context) {
	_ = "STUB: not implemented"
	return
}

// TODO: do we want to emit any events to the context event manager (for use in beginBlockResponse)
