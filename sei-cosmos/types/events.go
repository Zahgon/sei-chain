package types

import (
	"sync"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	proto "github.com/gogo/protobuf/proto"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

// ----------------------------------------------------------------------------
// Event Manager
// ----------------------------------------------------------------------------

// EventManager implements a simple wrapper around a slice of Event objects that
// can be emitted from.
type EventManager struct {
	events Events

	mtx sync.RWMutex
}

// Common Event Types and Attributes
const (
	EventTypeResourceAccess = "resource_access"

	AttributeKeyStoreKey      = "store_key"
	AttributeKeyResourceValue = "value"
	AttributeKeyResourceKey   = "key"
	AttributeKeyOperation     = "operation"

	AttributeKeyAccessType      = "access_type"
	AttributeKeyAccessTypeWrite = "write"
	AttributeKeyAccessTypeRead  = "read"
)

func NewEventManager() *EventManager { _ = "STUB: not implemented"; return nil }

func (em *EventManager) Events() Events {
	_ = "STUB: not implemented"

	// EmitEvent stores a single Event object.
	// Deprecated: Use EmitTypedEvent
	return *new(Events)
}

func (em *EventManager) EmitEvent(event Event) { _ = "STUB: not implemented"; return }

// EmitEvents stores a series of Event objects.
// Deprecated: Use EmitTypedEvents
func (em *EventManager) EmitEvents(events Events) { _ = "STUB: not implemented"; return }

// ABCIEvents returns all stored Event objects as abci.Event objects.
func (em *EventManager) ABCIEvents() []abci.Event { _ = "STUB: not implemented"; return nil }

// EmitTypedEvent takes typed event and emits converting it into Event
func (em *EventManager) EmitTypedEvent(tev proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// EmitTypedEvents takes series of typed events and emit
func (em *EventManager) EmitTypedEvents(tevs ...proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// TypedEventToEvent takes typed event and converts to Event object
func TypedEventToEvent(tev proto.Message) (Event, error) {
	_ = "STUB: not implemented"
	return *new(Event), nil
}

// sort the keys to ensure the order is always the same

// ParseTypedEvent converts abci.Event back to typed event
func ParseTypedEvent(event abci.Event) (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}

// ----------------------------------------------------------------------------
// Events
// ----------------------------------------------------------------------------

type (
	// Event is a type alias for an ABCI Event
	Event abci.Event

	// Events defines a slice of Event objects
	Events []Event
)

// NewEvent creates a new Event object with a given type and slice of one or more
// attributes.
func NewEvent(ty string, attrs ...Attribute) Event { _ = "STUB: not implemented"; return *new(Event) }

// NewAttribute returns a new key/value Attribute object.
func NewAttribute(k, v string) Attribute {
	_ = "STUB: not implemented"
	return *

	// EmptyEvents returns an empty slice of events.
	new(Attribute)
}

func EmptyEvents() Events { _ = "STUB: not implemented"; return *new(Events) }

func (a Attribute) String() string { _ = "STUB: not implemented"; return "" }

// ToKVPair converts an Attribute object into a Tendermint key/value pair.
func (a Attribute) ToKVPair() abci.EventAttribute {
	_ = "STUB: not implemented"
	return *new(abci.EventAttribute)
}

// AppendAttributes adds one or more attributes to an Event.
func (e Event) AppendAttributes(attrs ...Attribute) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

// AppendEvent adds an Event to a slice of events.
func (e Events) AppendEvent(event Event) Events {
	_ = "STUB: not implemented"
	return *

	// AppendEvents adds a slice of Event objects to an exist slice of Event objects.
	new(Events)
}

func (e Events) AppendEvents(events Events) Events { _ = "STUB: not implemented"; return *new(Events) }

// ToABCIEvents converts a slice of Event objects to a slice of abci.Event
// objects.
func (e Events) ToABCIEvents() []abci.Event { _ = "STUB: not implemented"; return nil }

// Common event types and attribute keys
var (
	EventTypeTx = "tx"

	AttributeKeyAccountSequence = "acc_seq"
	AttributeKeySignature       = "signature"
	AttributeKeyFee             = "fee"
	AttributeKeyFeePayer        = "fee_payer"

	EventTypeMessage = "message"

	AttributeKeyAction = "action"
	AttributeKeyModule = "module"
	AttributeKeySender = "sender"
	AttributeKeyAmount = "amount"
)

type (
	// StringAttributes defines a slice of StringEvents objects.
	StringEvents []StringEvent
)

func (se StringEvents) String() string { _ = "STUB: not implemented"; return "" }

// Flatten returns a flattened version of StringEvents by grouping all attributes
// per unique event type.
func (se StringEvents) Flatten() StringEvents { _ = "STUB: not implemented"; return *new(StringEvents) }

// appeneded to keys, same length of what is allocated to keys

// StringifyEvent converts an Event object to a StringEvent object.
func StringifyEvent(e abci.Event) StringEvent { _ = "STUB: not implemented"; return *new(StringEvent) }

// StringifyEvents converts a slice of Event objects into a slice of StringEvent
// objects.
func StringifyEvents(events []abci.Event) StringEvents {
	_ = "STUB: not implemented"
	return *new(StringEvents)
}

// MarkEventsToIndex returns the set of ABCI events, where each event's attribute
// has it's index value marked based on the provided set of events to index.
func MarkEventsToIndex(events []abci.Event, indexSet map[string]struct{}) []abci.Event {
	_ = "STUB: not implemented"
	return nil
}

type EVMEventManager struct {
	events []*ethtypes.Log
}

func NewEVMEventManager() *EVMEventManager { _ = "STUB: not implemented"; return nil }

func (eem *EVMEventManager) Events() []*ethtypes.Log { _ = "STUB: not implemented"; return nil }

func (eem *EVMEventManager) EmitEvents(events []*ethtypes.Log) { _ = "STUB: not implemented"; return }
