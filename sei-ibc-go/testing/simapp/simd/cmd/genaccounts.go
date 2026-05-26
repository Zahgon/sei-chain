package cmd

import (
	"github.com/spf13/cobra"
)

const (
	flagVestingStart = "vesting-start-time"
	flagVestingEnd   = "vesting-end-time"
	flagVestingAmt   = "vesting-amount"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
func AddGenesisAccountCmd(defaultNodeHome string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// attempt to lookup address from Keybase if no address was provided

// create concrete account type based on input parameters

// Add the new account to the set of genesis accounts and sanitize the
// accounts afterwards.
