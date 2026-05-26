package cli

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil/types"
)

// GenTxCmd builds the application's gentx command.
func GenTxCmd(mbm module.BasicManager, txEncCfg client.TxEncodingConfig, genBalIterator types.GenesisBalancesIterator, defaultNodeHome string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// read --nodeID, if empty take it from priv_validator.json

// read --pubkey, if empty take it from priv_validator.json

// set flags for creating a gentx

// The following line comes from a discrepancy between the `gentx`
// and `create-validator` commands:
// - `gentx` expects amount as an arg,
// - `create-validator` expects amount as a required flag.
// ref: https://github.com/cosmos/cosmos-sdk/issues/8251
// Since gentx doesn't set the amount flag (which `create-validator`
// reads from), we copy the amount arg into the valCfg directly.
//
// Ideally, the `create-validator` command should take a validator
// config file instead of so many flags.
// ref: https://github.com/cosmos/cosmos-sdk/issues/8177

// create a 'create-validator' message

// write the unsigned transaction to the buffer

// read the transaction

// sign the transaction and write it to the output file

func makeOutputFilepath(rootDir, nodeID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readUnsignedGenTxFile(clientCtx client.Context, r io.Reader) (sdk.Tx, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Tx), nil
}

func writeSignedGenTx(clientCtx client.Context, outputDocument string, tx sdk.Tx) error {
	_ = "STUB: not implemented"
	return nil
}
