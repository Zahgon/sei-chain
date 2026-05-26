package service

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/sei-protocol/seilog"
)

var (
	logger = seilog.NewLogger("tendermint", "internal", "libs", "service")

	_ Service = (*BaseService)(nil)
)

// Service defines a service that can be started, stopped, and reset.
type Service interface {
	// Start is called to start the service, which should run until
	// the context terminates. If the service is already running, Start
	// must report an error.
	Start(context.Context) error

	// Manually terminates the service
	Stop()

	// Return true if the service is running
	IsRunning() bool

	// Wait blocks until the service is stopped.
	Wait()
}

// Implementation describes the implementation that the
// BaseService implementation wraps.
type Implementation interface {
	// Called by the Services Start Method
	OnStart(context.Context) error

	// Called when the service's context is canceled.
	OnStop()
}

type baseService struct {
	// This is the context that (structured concurrency) service tasks will be executed with.
	// It is canceled when outer context is canceled or when the service is stopped.
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	done   chan struct{}
}

/*
Classical-inheritance-style service declarations. Services can be started, then
stopped, but cannot be restarted.

Users must implement OnStart/OnStop methods. In the absence of errors, these
methods are guaranteed to be called at most once. If OnStart returns an error,
service won't be marked as started, so the user can call Start again.

The BaseService implementation ensures that the OnStop method is
called after the context passed to Start is canceled.

Typical usage:

	type FooService struct {
		BaseService
		// private fields
	}

	func NewFooService() *FooService {
		fs := &FooService{
			// init
		}
		fs.BaseService = *NewBaseService( "FooService", fs)
		return fs
	}

	func (fs *FooService) OnStart(ctx context.Context) error {
		// initialize private fields
		// start subroutines, etc.
	}

	func (fs *FooService) OnStop() {
		// close/destroy private fields and releases resources
	}
*/
type BaseService struct {
	name string
	// The "subclass" of BaseService
	impl  Implementation
	inner atomic.Pointer[baseService]
}

// NewBaseService creates a new BaseService.
func NewBaseService(name string, impl Implementation) *BaseService {
	_ = "STUB: not implemented"
	return nil
}

// Start starts the Service and calls its OnStart method. An error
// will be returned if the service is stopped, but not if it is
// already running.
func (bs *BaseService) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// free the context.

// Currently sei-tendermint services (and tests) rely on the fact that OnStart is called with
// exactly the same context as Start.

// free the context.

// free the context.

// wait for all spawned tasks to finish

// Stop manually terminates the service by calling OnStop method from
// the implementation and releases all resources related to the
// service.
func (bs *BaseService) Stop() { _ = "STUB: not implemented"; return }

// Spawn spawns a new goroutine executing the task, which will be cancelled
// when outer context is cancelled or when the service is stopped.
// Error (other than ctx.Canceled) is logged after the task finishes.
// Both Wait and Stop calls will block until the spawned task is finished.
// It should be called ONLY from within OnStart().
// Note that the task is provided with a narrower context than the context
// provided to OnStart(). This is intentional.
// Panics if the service has not been started yet.
func (bs *BaseService) Spawn(name string, task func(ctx context.Context) error) {
	_ = "STUB: not implemented"
	return
}

// Spawns a critical task which should run until success OR as long as the service is running.
// It panics in any of the following cases:
// * task returns context.Canceled BEFORE the service is canceled.
// * task returns an error other than context.Canceled.
func (bs *BaseService) SpawnCritical(name string, task func(ctx context.Context) error) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck // QF1001: linter wants to apply De Morgan's law, as if outer negation was objectively worse.

// IsRunning implements Service by returning true or false depending on the
// service's state.
func (bs *BaseService) IsRunning() bool { _ = "STUB: not implemented"; return false }

// Wait blocks until the service is stopped.
func (bs *BaseService) Wait() { _ = "STUB: not implemented"; return }

// String provides a human-friendly representation of the service.
func (bs *BaseService) String() string { _ = "STUB: not implemented"; return "" }
