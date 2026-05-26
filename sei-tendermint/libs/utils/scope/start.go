package scope

import (
	"context"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type scope struct {
	// scope is a concurrecy primitive, so no-ctx-in-struct rule does not apply
	// nolint:containedctx
	ctx     context.Context
	cancel  context.CancelFunc
	all     sync.WaitGroup
	main    sync.WaitGroup
	errOnce sync.Once
	err     error
}

// Scope of concurrenct tasks.
type Scope struct{ *scope }

// SpawnBg spawns a background task.
// Background tasks get canceled when all the main tasks return.
func (s Scope) SpawnBg(t func() error) { _ = "STUB: not implemented"; return }

// Spawn spawns a main task.
// Scope gets automatically canceled when all the main tasks return.
func (s Scope) Spawn(t func() error) { _ = "STUB: not implemented"; return }

// Cancels the scope.
// If err is not nil and no error was set before,
// sets err as the scope error.
func (s Scope) Cancel(err error) { _ = "STUB: not implemented"; return }

// JoinHandle is a handle to an awaitable task.
type JoinHandle[R any] struct {
	result utils.AtomicRecv[*R]
}

// Spawn1 is the same as Scope.Spawn, but allows awaiting completion of a task and getting its result.
func Spawn1[R any](s Scope, t func() (R, error)) JoinHandle[R] {
	_ = "STUB: not implemented"
	return nil
}

// Join awaits completion of a task and returns its result.
// WARNING: it does NOT return the error of the task - error is returned from the Run() command.
// Join() can only fail when context is canceled.
func (h JoinHandle[R]) Join(ctx context.Context) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// If true, tasks that do not respect context cancellation will be logged.
// This is useful for debugging, but causes unnecessary overhead.
// Since this is a constant, debug guard should be optimized out by the compiler.
const enableDebugGuard = false

func (s Scope) debugGuard(name string, done chan struct{}) { _ = "STUB: not implemented"; return }

// SpawnNamed spawns a named main task.
func (s Scope) SpawnNamed(name string, t func() error) { _ = "STUB: not implemented"; return }

// SpawnBgNamed spawns a named background task.
func (s Scope) SpawnBgNamed(name string, t func() error) { _ = "STUB: not implemented"; return }

// Run runs a scope capable of spawning tasks.
// It is guaranteed that all the spawned tasks will be executed (even if spawned after the context is cancelled),
// and that `Run` will return only after all the tasks have completed.
// Context of the tasks will be automatically cancelled as soon as ANY task returns an error.
// Returns the first error returned by any task (main or background).
func Run(ctx context.Context, main func(context.Context, Scope) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Run1 is the same as Run, but returns the result of the main task.
func Run1[R any](ctx context.Context, main func(context.Context, Scope) (R, error)) (res R, err error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

//nolint:nakedret
