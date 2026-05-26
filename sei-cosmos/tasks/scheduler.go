package tasks

import (
	"context"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/multiversion"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/occ"
	"github.com/sei-protocol/sei-chain/sei-cosmos/utils/tracing"
	"github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/seilog"
	"go.opentelemetry.io/otel/trace"
)

var logger = seilog.NewLogger("cosmos", "tasks")

type status string

const (
	// statusPending tasks are ready for execution
	// all executing tasks are in pending state
	statusPending status = "pending"
	// statusExecuted tasks are ready for validation
	// these tasks did not abort during execution
	statusExecuted status = "executed"
	// statusAborted means the task has been aborted
	// these tasks transition to pending upon next execution
	statusAborted status = "aborted"
	// statusValidated means the task has been validated
	// tasks in this status can be reset if an earlier task fails validation
	statusValidated status = "validated"
	// statusWaiting tasks are waiting for another tx to complete
	statusWaiting status = "waiting"
	// maximumIterations before we revert to sequential (for high conflict rates)
	maximumIterations = 10
)

type deliverTxTask struct {
	Ctx     sdk.Context
	AbortCh chan occ.Abort

	mx            sync.RWMutex
	Status        status
	Dependencies  map[int]struct{}
	Abort         *occ.Abort
	Incarnation   int
	Request       types.RequestDeliverTxV2
	SdkTx         sdk.Tx
	Checksum      [32]byte
	AbsoluteIndex int
	Response      *types.ResponseDeliverTx
	VersionStores map[sdk.StoreKey]*multiversion.VersionIndexedStore
	TxTracer      sdk.TxTracer
}

// AppendDependencies appends the given indexes to the task's dependencies
func (dt *deliverTxTask) AppendDependencies(deps []int) { _ = "STUB: not implemented"; return }

func (dt *deliverTxTask) IsStatus(s status) bool { _ = "STUB: not implemented"; return false }

func (dt *deliverTxTask) SetStatus(s status) { _ = "STUB: not implemented"; return }

func (dt *deliverTxTask) Reset() { _ = "STUB: not implemented"; return }

func (dt *deliverTxTask) Increment() {
	_ = "STUB: not implemented"

	// Scheduler processes tasks concurrently
	return
}

type Scheduler interface {
	ProcessAll(ctx sdk.Context, reqs []*sdk.DeliverTxEntry) ([]types.ResponseDeliverTx, error)
}

type scheduler struct {
	deliverTx          func(ctx sdk.Context, req types.RequestDeliverTxV2, tx sdk.Tx, checksum [32]byte) (res types.ResponseDeliverTx)
	workers            int
	multiVersionStores map[sdk.StoreKey]multiversion.MultiVersionStore
	tracingInfo        *tracing.Info
	allTasksMap        map[int]*deliverTxTask
	allTasks           []*deliverTxTask
	executeCh          chan func()
	validateCh         chan func()
	metrics            *schedulerMetrics
	synchronous        bool           // true if maxIncarnation exceeds threshold
	maxIncarnation     int            // current highest incarnation
	conflictKeyCounts  map[string]int // per-key conflict counts accumulated over the block
	conflictKeyMu      sync.Mutex
}

// NewScheduler creates a new scheduler
func NewScheduler(workers int, tracingInfo *tracing.Info, deliverTxFunc func(ctx sdk.Context, req types.RequestDeliverTxV2, tx sdk.Tx, checksum [32]byte) (res types.ResponseDeliverTx)) Scheduler {
	_ = "STUB: not implemented"
	return *new(Scheduler)
}

func (s *scheduler) invalidateTask(task *deliverTxTask) { _ = "STUB: not implemented"; return }

func start(ctx context.Context, ch chan func(), workers int) { _ = "STUB: not implemented"; return }

func (s *scheduler) DoValidate(work func()) { _ = "STUB: not implemented"; return }

func (s *scheduler) DoExecute(work func()) { _ = "STUB: not implemented"; return }

func (s *scheduler) findConflicts(task *deliverTxTask) (bool, []int, []string) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// any non-ok value makes valid false

func toTasks(reqs []*sdk.DeliverTxEntry) ([]*deliverTxTask, map[int]*deliverTxTask) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *scheduler) collectResponses(tasks []*deliverTxTask) []types.ResponseDeliverTx {
	_ = "STUB: not implemented"
	return nil
}

func (s *scheduler) tryInitMultiVersionStore(ctx sdk.Context) { _ = "STUB: not implemented"; return }

func dependenciesValidated(tasksMap map[int]*deliverTxTask, deps map[int]struct{}) bool {
	_ = "STUB: not implemented"
	return false

	// because idx contains absoluteIndices, we need to fetch from map
}

func allValidated(tasks []*deliverTxTask) bool { _ = "STUB: not implemented"; return false }

// schedulerMetrics contains metrics for the scheduler
type schedulerMetrics struct {
	// maxIncarnation is the highest incarnation seen in this set
	maxIncarnation int
	// retries is the number of tx attempts beyond the first attempt
	retries int
}

func (s *scheduler) emitMetrics() { _ = "STUB: not implemented"; return }

func (s *scheduler) ProcessAll(ctx sdk.Context, reqs []*sdk.DeliverTxEntry) ([]types.ResponseDeliverTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize mutli-version stores if they haven't been initialized yet

// default to number of tasks if workers is negative or 0 by this point

// execution tasks are limited by workers

// validation tasks uses length of tasks to avoid blocking on validation

// if the max incarnation >= x, we should revert to synchronous

// process synchronously

// execute sets statuses of tasks to either executed or aborted

// validate returns any that should be re-executed
// note this processes ALL tasks, not just those recently executed

// these are retries which apply to metrics

func (s *scheduler) shouldRerun(task *deliverTxTask) bool { _ = "STUB: not implemented"; return false }

// validated tasks can become unvalidated if an earlier re-run task now conflicts

// With the current scheduler, we won't actually get to this step if a previous task has already been determined to be invalid,
// since we choose to fail fast and mark the subsequent tasks as invalid as well.
// TODO: in a future async scheduler that no longer exhaustively validates in order, we may need to carefully handle the `valid=true` with conflicts case

// if the conflicts are now validated, then rerun this task

// otherwise, wait for completion

// mark as validated, which will avoid re-validating unless a lower-index re-validates

// conflicts and valid, so it'll validate next time

// if conflicts are done, then this task is ready to run again

func (s *scheduler) validateTask(ctx sdk.Context, task *deliverTxTask) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *scheduler) findFirstNonValidated() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (s *scheduler) validateAll(ctx sdk.Context, tasks []*deliverTxTask) ([]*deliverTxTask, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update max incarnation for scheduler

// ExecuteAll executes all tasks concurrently
func (s *scheduler) executeAll(ctx sdk.Context, tasks []*deliverTxTask) error {
	_ = "STUB: not implemented"
	return nil
}

// validationWg waits for all validations to complete
// validations happen in separate goroutines in order to wait on previous index

func (s *scheduler) prepareAndRunTask(wg *sync.WaitGroup, ctx sdk.Context, task *deliverTxTask) {
	_ = "STUB: not implemented"
	// Must be deferred to prevent deadlock on panic
	return
}

func (s *scheduler) traceSpan(ctx sdk.Context, name string, task *deliverTxTask) (sdk.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), *new(trace.Span)
}

// prepareTask initializes the context and version stores for a task
func (s *scheduler) prepareTask(task *deliverTxTask) { _ = "STUB: not implemented"; return }

// initialize the context

// if there are no stores, don't try to wrap, because there's nothing to wrap

// non-blocking

// init version stores by store key

// save off version store so we can ask it things later

func (s *scheduler) executeTask(task *deliverTxTask) { _ = "STUB: not implemented"; return }

// in the synchronous case, we only want to re-execute tasks that need re-executing

// even if already validated, it could become invalid again due to preceding
// reruns. Make sure previous writes are invalidated before rerunning.

// waiting transactions may not yet have been reset
// this ensures a task has been reset and incremented

// close the abort channel

// if there is an abort item that means we need to wait on the dependent tx

// write from version store to multiversion stores

// write from version store to multiversion stores
