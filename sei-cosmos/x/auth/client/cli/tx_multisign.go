package cli

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	signingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
)

// BroadcastReq defines a tx broadcasting request.
type BroadcastReq struct {
	Tx   legacytx.StdTx `json:"tx" yaml:"tx"`
	Mode string         `json:"mode" yaml:"mode"`
}

// GetSignCommand returns the sign command
func GetMultiSignCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func makeMultiSignCmd() func(cmd *cobra.Command, args []string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// read each signature and add it to the multisig if valid

func GetMultiSignBatchCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func makeBatchMultisignCmd() func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// prepare output document

func unmarshalSignatureJSON(clientCtx client.Context, filename string) (sigs []signingtypes.SignatureV2, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readSignaturesFromFile(ctx client.Context, filename string) (sigs []signingtypes.SignatureV2, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMultisigInfo(clientCtx client.Context, name string) (keyring.Info, error) {
	_ = "STUB: not implemented"
	return *new(keyring.Info), nil
}
