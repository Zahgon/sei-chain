package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type contextKey int

const (
	// private type creates an interface key for Context that cannot be accessed by any other package
	contextKeyTXCount contextKey = iota
)

// WithTXCounter stores a transaction counter value in the context
func WithTXCounter(ctx sdk.Context, counter uint32) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// TXCounter returns the tx counter value and found bool from the context.
// The result will be (0, false) for external queries or simulations where no counter available.
func TXCounter(ctx sdk.Context) (uint32, bool) { _ = "STUB: not implemented"; return 0, false }
