package utils

import (
	"context"
	"encoding"
	"sync/atomic"
	"time"
)

// IgnoreCancel returns nil if the error is context.Canceled, err otherwise.
func IgnoreCancel(err error) error { _ = "STUB: not implemented"; return nil }

// WithDeadline executes a function with a deadline.
// If deadline is none, it executes the function without a deadline.
func WithDeadline(ctx context.Context, md Option[time.Time], f func(ctx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

// WithTimeout executes a function with a timeout.
func WithTimeout(ctx context.Context, d time.Duration, f func(ctx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

// WithOptTimeout executes a function with a timeout.
func WithOptTimeout(ctx context.Context, d Option[time.Duration], f func(ctx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

// WithTimeout1 executes a function with a timeout.
func WithTimeout1[R any](ctx context.Context, d time.Duration, f func(ctx context.Context) (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// WithOptTimeout1 executes a function with a timeout.
func WithOptTimeout1[R any](ctx context.Context, d Option[time.Duration], f func(ctx context.Context) (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// Sleep sleeps for a duration or until the context is canceled.
func Sleep(ctx context.Context, d time.Duration) error { _ = "STUB: not implemented"; return nil }

// SleepUntil sleeps until deadline t or until the context is canceled.
func SleepUntil(ctx context.Context, t time.Time) error { _ = "STUB: not implemented"; return nil }

// WaitFor polls a check function until it returns true or the context is canceled.
func WaitFor(ctx context.Context, interval time.Duration, check func() bool) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForWithTimeout polls a check function until it returns true, the context is canceled, or the timeout is reached.
func WaitForWithTimeout(ctx context.Context, interval, timeout time.Duration, check func() bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Duration is a wrapper type around time.Duration that supports JSON marshaling/unmarshaling.
// nolint:recvcheck
type Duration time.Duration

// MarshalText implements json.TextMarshaler interface to convert Duration to JSON string.
func (d Duration) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements json.TextUnmarshaler.
func (d *Duration) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }

var _ encoding.TextMarshaler = Zero[Duration]()
var _ encoding.TextUnmarshaler = (*Duration)(nil)

// Duration returns the underlying time.Duration value.
func (d Duration) Duration() time.Duration {
	_ = "STUB: not implemented"
	return *

	// Seconds returns the underlying time.Duration value in seconds.
	new(time.Duration)
}

func (d Duration) Seconds() float64 { _ = "STUB: not implemented"; return 0 }

// Once is an idempotent signal.
type Once struct {
	_    NoCopy
	ch   chan struct{}
	done atomic.Bool
}

func NewOnce() (o Once) { _ = "STUB: not implemented"; return *new(Once) }

func (o *Once) Send() { _ = "STUB: not implemented"; return }

func (o *Once) Recv(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
