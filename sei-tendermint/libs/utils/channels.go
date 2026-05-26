package utils

import (
	"context"
)

// Recv receives a value from a channel or returns an error if the context is canceled.
func Recv[T any](ctx context.Context, ch <-chan T) (zero T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// We are not interested in channel closing,
// patiently wait for the context to be done instead.

// RecvOrClosed receives a value from a channel, returns false if channel got closed,
// or returns an error if the context is canceled.
func RecvOrClosed[T any](ctx context.Context, ch <-chan T) (T, bool, error) {
	_ = "STUB: not implemented"
	return *new(T), false, nil
}

// Send a value to channel or returns an error if the context is canceled.
func Send[T any](ctx context.Context, ch chan<- T, v T) error {
	_ = "STUB: not implemented"
	return nil
}

// SendOrDrop send a value to channel if not full or drop the item if the channel is full.
func SendOrDrop[T any](ch chan<- T, v T) error { _ = "STUB: not implemented"; return nil }

// drop the item

// ForEach is a helper function that reads from a channel and calls a handler for each item.
// this avoids needing a lot of for/select boilerplate everywhere.
func ForEach[T any](ctx context.Context, ch <-chan T, handler func(T) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Channel closed

// Stop on error
