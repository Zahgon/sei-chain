package cmd

import (
	"crypto/ecdsa"

	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
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

// this args[0] is for the key name so "admin"

// create concrete account type based on input parameters

// associate the eth address with the sei address through the genesis file

// Add the new account to the set of genesis accounts and sanitize the
// accounts afterwards.

func getPrivateKeyOfAddr(kb keyring.Keyring, addr sdk.Address) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// will only show local key

// Need to use private key to convert to sei address here
