package threading

import (
	"context"
)

// Push to a channel, returning an error if the context is cancelled before the value is pushed.
func InterruptiblePush[T any](ctx context.Context, ch chan T, value T) error {
	_ = "STUB: not implemented"
	return nil
}

// Pull from a channel, returning an error if the context is cancelled before the value is pulled.
func InterruptiblePull[T any](ctx context.Context, ch <-chan T) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
