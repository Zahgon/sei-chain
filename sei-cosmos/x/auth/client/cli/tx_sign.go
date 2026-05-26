package cli

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

const (
	flagMultisig        = "multisig"
	flagOverwrite       = "overwrite"
	flagSigOnly         = "signature-only"
	flagAmino           = "amino"
	flagNoAutoIncrement = "no-auto-increment"
)

// GetSignBatchCommand returns the transaction sign-batch command.
func GetSignBatchCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func makeSignBatchCmd() func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// prepare output document

func setOutputFile(cmd *cobra.Command) (func(), error) { _ = "STUB: not implemented"; return nil, nil }

// GetSignCommand returns the transaction sign command.
func GetSignCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func preSignCmd(cmd *cobra.Command, _ []string) {
	_ = "STUB: not implemented"
	// Conditionally mark the account and sequence numbers required as no RPC
	// query will be done.
	return
}

func makeSignCmd() func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Bech32 decode error, maybe it's a name, we try to fetch from keyring

func marshalSignatureJSON(txConfig client.TxConfig, txBldr client.TxBuilder, signatureOnly bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMultisigRecord(clientCtx client.Context, name string) (keyring.Info, error) {
	_ = "STUB: not implemented"
	return *new(keyring.Info), nil
}
