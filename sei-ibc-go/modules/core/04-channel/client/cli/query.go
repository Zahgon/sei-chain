package cli

import (
	"github.com/spf13/cobra"
)

const (
	flagSequences = "sequences"
)

// GetCmdQueryChannels defines the command to query all the channels ends
// that this chain mantains.
func GetCmdQueryChannels() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryChannel defines the command to query a channel end
func GetCmdQueryChannel() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryConnectionChannels defines the command to query all the channels associated with a
// connection
func GetCmdQueryConnectionChannels() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryChannelClientState defines the command to query a client state from a channel
func GetCmdQueryChannelClientState() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryPacketCommitments defines the command to query all packet commitments associated with
// a channel
func GetCmdQueryPacketCommitments() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryPacketCommitment defines the command to query a packet commitment
func GetCmdQueryPacketCommitment() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryPacketReceipt defines the command to query a packet receipt
func GetCmdQueryPacketReceipt() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryPacketAcknowledgement defines the command to query a packet acknowledgement
func GetCmdQueryPacketAcknowledgement() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryUnreceivedPackets defines the command to query all the unreceived
// packets on the receiving chain
func GetCmdQueryUnreceivedPackets() *cobra.Command { _ = "STUB: not implemented"; return nil }

// #nosec G115 -- will hit cli limits first.

// GetCmdQueryUnreceivedAcks defines the command to query all the unreceived acks on the original sending chain
func GetCmdQueryUnreceivedAcks() *cobra.Command { _ = "STUB: not implemented"; return nil }

// #nosec G115 -- will hit cli limits first.

// GetCmdQueryNextSequenceReceive defines the command to query a next receive sequence for a given channel
func GetCmdQueryNextSequenceReceive() *cobra.Command { _ = "STUB: not implemented"; return nil }

// #nosec G115 -- revision height is bounds checked above
