package ibctesting

import (
	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

func getSendPackets(evts []abci.Event) []channeltypes.Packet { _ = "STUB: not implemented"; return nil }

func getAckPackets(evts []abci.Event) []PacketAck { _ = "STUB: not implemented"; return nil }

// Used for various debug statements above when needed... do not remove
// func showEvent(evt abci.Event) {
//	fmt.Printf("evt.Type: %s\n", evt.Type)
//	for _, attr := range evt.Attributes {
//		fmt.Printf("  %s = %s\n", string(attr.Key), string(attr.Value))
//	}
//}

func parsePacketFromEvent(evt abci.Event) channeltypes.Packet {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet)
}

// return the value for the attribute with the given name
func getField(evt abci.Event, key string) string { _ = "STUB: not implemented"; return "" }

func getUintField(evt abci.Event, key string) uint64 { _ = "STUB: not implemented"; return 0 }

func toUint64(raw string) uint64 { _ = "STUB: not implemented"; return 0 }

func parseTimeoutHeight(raw string) clienttypes.Height {
	_ = "STUB: not implemented"
	return *new(clienttypes.Height)
}
