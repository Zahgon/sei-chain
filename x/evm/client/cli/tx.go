package cli

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	FlagGasFeeCap = "gas-fee-cap"
	FlagGas       = "gas-limit"
	FlagValue     = "value"
	FlagRPC       = "evm-rpc"
	FlagNonce     = "nonce"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdAssociateAddress() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdSend() *cobra.Command { _ = "STUB: not implemented"; return nil }

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
}

func CmdDeployContract() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdCallContract() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdERC20Send() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdCallPrecompile() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdDeployWSEI() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getPrivateKey(cmd *cobra.Command) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNonce(rpc string, key ecdsa.PublicKey) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getChainId(rpc string) (*big.Int, error) { _ = "STUB: not implemented"; return nil, nil }

func getTxData(cmd *cobra.Command) (*ethtypes.DynamicFeeTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sendTx(txData *ethtypes.DynamicFeeTx, rpcUrl string, key *ecdsa.PrivateKey) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}
