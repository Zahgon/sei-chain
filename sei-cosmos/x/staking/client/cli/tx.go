package cli

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client/tx"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// default values
var (
	DefaultTokens                  = sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)
	defaultAmount                  = DefaultTokens.String() + sdk.DefaultBondDenom
	defaultCommissionRate          = "0.1"
	defaultCommissionMaxRate       = "0.2"
	defaultCommissionMaxChangeRate = "0.01"
	defaultMinSelfDelegation       = "1"
)

// NewTxCmd returns a root CLI command handler for all x/staking transaction commands.
func NewTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewCreateValidatorCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewEditValidatorCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewDelegateCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewRedelegateCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewUnbondCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func newBuildCreateValidatorMsg(clientCtx client.Context, txf tx.Factory, fs *flag.FlagSet) (tx.Factory, *types.MsgCreateValidator, error) {
	_ = "STUB: not implemented"
	return *new(tx.Factory), nil, nil
}

// get the initial validator commission parameters

// get the initial validator min self delegation

// Return the flagset, particular flags, and a description of defaults
// this is anticipated to be used with the gen-tx
func CreateValidatorMsgFlagSet(ipDefault string) (fs *flag.FlagSet, defaultsDesc string) {
	_ = "STUB: not implemented"
	return nil, ""
}

type TxCreateValidatorConfig struct {
	ChainID string
	NodeID  string
	Moniker string

	Amount string

	CommissionRate          string
	CommissionMaxRate       string
	CommissionMaxChangeRate string
	MinSelfDelegation       string

	PubKey cryptotypes.PubKey

	IP              string
	P2PPort         string
	Website         string
	SecurityContact string
	Details         string
	Identity        string
}

func PrepareConfigForTxCreateValidator(flagSet *flag.FlagSet, moniker, nodeID, chainID string, valPubKey cryptotypes.PubKey) (TxCreateValidatorConfig, error) {
	_ = "STUB: not implemented"
	return *new(TxCreateValidatorConfig), nil
}

// BuildCreateValidatorMsg makes a new MsgCreateValidator.
func BuildCreateValidatorMsg(clientCtx client.Context, config TxCreateValidatorConfig, txBldr tx.Factory, generateOnly bool) (tx.Factory, sdk.Msg, error) {
	_ = "STUB: not implemented"
	return *new(tx.Factory), *new(sdk.Msg), nil
}

// get the initial validator commission parameters

// get the initial validator min self delegation
