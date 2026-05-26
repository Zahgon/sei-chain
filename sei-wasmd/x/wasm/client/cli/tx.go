package cli

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

const (
	flagAmount                 = "amount"
	flagLabel                  = "label"
	flagAdmin                  = "admin"
	flagNoAdmin                = "no-admin"
	flagRunAs                  = "run-as"
	flagInstantiateByEverybody = "instantiate-everybody"
	flagInstantiateNobody      = "instantiate-nobody"
	flagInstantiateByAddress   = "instantiate-only-address"
	flagProposalType           = "type"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// StoreCodeCmd will upload code to be reused.
func StoreCodeCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func parseStoreCodeArgs(file string, sender sdk.AccAddress, flags *flag.FlagSet) (types.MsgStoreCode, error) {
	_ = "STUB: not implemented"
	return *new(types.MsgStoreCode), nil
}

// gzip the wasm file

// InstantiateContractCmd will instantiate a contract from previously uploaded code.
func InstantiateContractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func parseInstantiateArgs(rawCodeID, initMsg string, sender sdk.AccAddress, flags *flag.FlagSet) (types.MsgInstantiateContract, error) {
	_ = "STUB: not implemented"
	// get the id of the code to instantiate
	return *new(types.MsgInstantiateContract), nil
}

// ensure sensible admin is set (or explicitly immutable)

// build and sign the transaction, then broadcast to Tendermint

// ExecuteContractCmd will instantiate a contract from previously uploaded code.
func ExecuteContractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func parseExecuteArgs(contractAddr string, execMsg string, sender sdk.AccAddress, flags *flag.FlagSet) (types.MsgExecuteContract, error) {
	_ = "STUB: not implemented"
	return *new(types.MsgExecuteContract), nil
}
