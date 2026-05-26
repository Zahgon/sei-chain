package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagAddressValidator    = "validator"
	FlagAddressValidatorSrc = "addr-validator-source"
	FlagAddressValidatorDst = "addr-validator-dest"
	FlagPubKey              = "pubkey"
	FlagAmount              = "amount"
	FlagSharesAmount        = "shares-amount"
	FlagSharesFraction      = "shares-fraction"

	FlagMoniker         = "moniker"
	FlagEditMoniker     = "new-moniker"
	FlagIdentity        = "identity"
	FlagWebsite         = "website"
	FlagSecurityContact = "security-contact"
	FlagDetails         = "details"

	FlagCommissionRate          = "commission-rate"
	FlagCommissionMaxRate       = "commission-max-rate"
	FlagCommissionMaxChangeRate = "commission-max-change-rate"

	FlagMinSelfDelegation = "min-self-delegation"

	FlagGenesisFormat = "genesis-format"
	FlagNodeID        = "node-id"
	FlagIP            = "ip"
	FlagP2PPort       = "p2p-port"
)

// common flagsets to add to various functions
var (
	fsShares       = flag.NewFlagSet("", flag.ContinueOnError)
	fsValidator    = flag.NewFlagSet("", flag.ContinueOnError)
	fsRedelegation = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsShares.String(FlagSharesAmount, "", "Amount of source-shares to either unbond or redelegate as a positive integer or decimal")
	fsShares.String(FlagSharesFraction, "", "Fraction of source-shares to either unbond or redelegate as a positive integer or decimal >0 and <=1")
	fsValidator.String(FlagAddressValidator, "", "The Bech32 address of the validator")
	fsRedelegation.String(FlagAddressValidatorSrc, "", "The Bech32 address of the source validator")
	fsRedelegation.String(FlagAddressValidatorDst, "", "The Bech32 address of the destination validator")
}

// FlagSetCommissionCreate Returns the FlagSet used for commission create.
func FlagSetCommissionCreate() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

// FlagSetMinSelfDelegation Returns the FlagSet used for minimum set delegation.
func FlagSetMinSelfDelegation() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

// FlagSetAmount Returns the FlagSet for amount related operations.
func FlagSetAmount() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

// FlagSetPublicKey Returns the flagset for Public Key related operations.
func FlagSetPublicKey() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func flagSetDescriptionEdit() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func flagSetCommissionUpdate() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func flagSetDescriptionCreate() *flag.FlagSet { _ = "STUB: not implemented"; return nil }
