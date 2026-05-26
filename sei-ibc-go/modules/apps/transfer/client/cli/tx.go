package cli

import (
	"github.com/spf13/cobra"
)

const (
	flagPacketTimeoutHeight    = "packet-timeout-height"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagAbsoluteTimeouts       = "absolute-timeouts"
	flagMemo                   = "memo"
)

// NewTransferTxCmd returns the command to create a NewMsgTransfer transaction
func NewTransferTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// if the timeouts are not absolute, retrieve latest block height and block timestamp
// for the consensus state connected to the destination port/channel

// use local clock time as reference time if it is later than the
// consensus state timestamp of the counter party chain, otherwise
// still use consensus state timestamp as reference
