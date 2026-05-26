package cli

import (
	"github.com/spf13/cobra"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Flag names and values
const (
	FlagSpendLimit        = "spend-limit"
	FlagMsgType           = "msg-type"
	FlagExpiration        = "expiration"
	FlagAllowedValidators = "allowed-validators"
	FlagDenyValidators    = "deny-validators"
	delegate              = "delegate"
	redelegate            = "redelegate"
	unbond                = "unbond"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewCmdGrantAuthorization() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewCmdRevokeAuthorization() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewCmdExecAuthorization() *cobra.Command { _ = "STUB: not implemented"; return nil }

func bech32toValidatorAddresses(validators []string) ([]sdk.ValAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
