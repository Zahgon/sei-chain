package cli

import (
	"time"

	"github.com/spf13/cobra"
)

// flag for feegrant module
const (
	FlagExpiration  = "expiration"
	FlagPeriod      = "period"
	FlagPeriodLimit = "period-limit"
	FlagSpendLimit  = "spend-limit"
	FlagAllowedMsgs = "allowed-messages"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCmdFeeGrant returns a CLI command handler for creating a MsgGrantAllowance transaction.
func NewCmdFeeGrant() *cobra.Command { _ = "STUB: not implemented"; return nil }

// if `FlagSpendLimit` isn't set, limit will be nil

// Check any of period or periodLimit flags set, If set consider it as periodic fee allowance.

// NewCmdRevokeFeegrant returns a CLI command handler for creating a MsgRevokeAllowance transaction.
func NewCmdRevokeFeegrant() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getPeriodReset(duration int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func getPeriod(duration int64) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }
