package ibctesting

import (
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/stretchr/testify/assert"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// ParseClientIDFromEvents parses events emitted from a MsgCreateClient and returns the
// client identifier.
func ParseClientIDFromEvents(events sdk.Events) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseConnectionIDFromEvents parses events emitted from a MsgConnectionOpenInit or
// MsgConnectionOpenTry and returns the connection identifier.
func ParseConnectionIDFromEvents(events sdk.Events) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseChannelIDFromEvents parses events emitted from a MsgChannelOpenInit or
// MsgChannelOpenTry and returns the channel identifier.
func ParseChannelIDFromEvents(events sdk.Events) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParsePacketFromEvents parses events emitted from a MsgRecvPacket and returns the
// acknowledgement.
func ParsePacketFromEvents(events sdk.Events) (channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet), nil
}

// ParseAckFromEvents parses events emitted from a MsgRecvPacket and returns the
// acknowledgement.
func ParseAckFromEvents(events sdk.Events) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AssertEvents asserts that expected events are present in the actual events.
func AssertEvents(
	t assert.TestingT,
	expected []abci.Event,
	actual []abci.Event,
) {
	_ = "STUB: not implemented"
	return
}

// any expected attributes that are not contained in the actual events will cause this event
// not to match

// shouldProcessEvent returns true if the given expected event should be processed based on event type.
func shouldProcessEvent(expectedEvent abci.Event, actualEvent abci.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// the actual event will have an extra attribute added automatically
// by Cosmos SDK since v0.50, that's why we subtract 1 when comparing
// with the number of attributes in the expected event.

// containsAttribute returns true if the given key/value pair is contained in the given attributes.
// NOTE: this ignores the indexed field, which can be set or unset depending on how the events are retrieved.
func containsAttribute(attrs []abci.EventAttribute, key, value string) bool {
	_ = "STUB: not implemented"
	return false
}

// containsAttributeKey returns true if the given key is contained in the given attributes.
func containsAttributeKey(attrs []abci.EventAttribute, key string) bool {
	_ = "STUB: not implemented"
	return false
}

// attributeByKey returns the event attribute's value keyed by the given key and a boolean indicating its presence in the given attributes.
func attributeByKey(attributes []abci.EventAttribute, key string) (abci.EventAttribute, bool) {
	_ = "STUB: not implemented"
	return *new(abci.EventAttribute), false
}
