package simulation

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
)

// entry kinds for use within OperationEntry
const (
	BeginBlockEntryKind = "begin_block"
	EndBlockEntryKind   = "end_block"
	MsgEntryKind        = "msg"
	QueuedMsgEntryKind  = "queued_msg"
)

// OperationEntry - an operation entry for logging (ex. BeginBlock, EndBlock, XxxMsg, etc)
type OperationEntry struct {
	EntryKind string          `json:"entry_kind" yaml:"entry_kind"`
	Height    int64           `json:"height" yaml:"height"`
	Order     int64           `json:"order" yaml:"order"`
	Operation json.RawMessage `json:"operation" yaml:"operation"`
}

// NewOperationEntry creates a new OperationEntry instance
func NewOperationEntry(entry string, height, order int64, op json.RawMessage) OperationEntry {
	_ = "STUB: not implemented"
	return *new(OperationEntry)
}

// BeginBlockEntry - operation entry for begin block
func BeginBlockEntry(height int64) OperationEntry {
	_ = "STUB: not implemented"
	return *new(OperationEntry)
}

// EndBlockEntry - operation entry for end block
func EndBlockEntry(height int64) OperationEntry {
	_ = "STUB: not implemented"
	return *new(OperationEntry)
}

// MsgEntry - operation entry for standard msg
func MsgEntry(height, order int64, opMsg simulation.OperationMsg) OperationEntry {
	_ = "STUB: not implemented"
	return *new(OperationEntry)
}

// QueuedMsgEntry creates an operation entry for a given queued message.
func QueuedMsgEntry(height int64, opMsg simulation.OperationMsg) OperationEntry {
	_ = "STUB: not implemented"
	return *new(OperationEntry)
}

// MustMarshal marshals the operation entry, panic on error.
func (oe OperationEntry) MustMarshal() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// OperationQueue defines an object for a queue of operations
type OperationQueue map[int][]simulation.Operation

// NewOperationQueue creates a new OperationQueue instance.
func NewOperationQueue() OperationQueue { _ = "STUB: not implemented"; return *new(OperationQueue) }

// queueOperations adds all future operations into the operation queue.
func queueOperations(queuedOps OperationQueue, queuedTimeOps []simulation.FutureOperation, futureOps []simulation.FutureOperation) {
	_ = "STUB: not implemented"
	return
}

// TODO: Replace with proper sorted data structure, so don't have the
// copy entire slice

// WeightedOperation is an operation with associated weight.
// This is used to bias the selection operation within the simulator.
type WeightedOperation struct {
	weight int
	op     simulation.Operation
}

func (w WeightedOperation) Weight() int { _ = "STUB: not implemented"; return 0 }

func (w WeightedOperation) Op() simulation.Operation {
	_ = "STUB: not implemented"

	// NewWeightedOperation creates a new WeightedOperation instance
	return *new(simulation.Operation)
}

func NewWeightedOperation(weight int, op simulation.Operation) WeightedOperation {
	_ = "STUB: not implemented"
	return *new(WeightedOperation)
}

// WeightedOperations is the group of all weighted operations to simulate.
type WeightedOperations []simulation.WeightedOperation

func (ops WeightedOperations) totalWeight() int { _ = "STUB: not implemented"; return 0 }

func (ops WeightedOperations) getSelectOpFn() simulation.SelectOpFn {
	_ = "STUB: not implemented"
	return *new(simulation.SelectOpFn)
}

// shouldn't happen
