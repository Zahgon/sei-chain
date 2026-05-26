package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type EpochHooks interface {
	// AfterEpochEnd defines the first block whose timestamp is after the duration
	// is counted as the end of the epoch.
	AfterEpochEnd(ctx sdk.Context, epoch Epoch)
	// BeforeEpochStart defines the new epoch is next block of epoch EndBlock.
	BeforeEpochStart(ctx sdk.Context, epoch Epoch)
}

var _ EpochHooks = MultiEpochHooks{}

type MultiEpochHooks []EpochHooks

func NewMultiEpochHooks(hooks ...EpochHooks) MultiEpochHooks {
	_ = "STUB: not implemented"

	// AfterEpochEnd is called when epoch is going to be ended, epochNumber is the
	// number of epoch that is ending.
	return *new(MultiEpochHooks)
}

func (h MultiEpochHooks) AfterEpochEnd(ctx sdk.Context, epoch Epoch) {
	_ = "STUB: not implemented"
	return
}

// BeforeEpochStart is called when epoch is going to be started, epochNumber is
// the number of epoch that is starting.
func (h MultiEpochHooks) BeforeEpochStart(ctx sdk.Context, epoch Epoch) {
	_ = "STUB: not implemented"
	return
}

func panicCatchingEpochHook(ctx sdk.Context, hookFn func(sdk.Context, Epoch), epoch Epoch) {
	_ = "STUB: not implemented"
	return
}

// cache the context and only write if no panic (which is caught above)
