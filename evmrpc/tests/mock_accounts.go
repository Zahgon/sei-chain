package tests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sei-protocol/sei-chain/app"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/config"
)

var chainId = big.NewInt(config.DefaultChainID)
var mnemonic1 = "fish mention unlock february marble dove vintage sand hub ordinary fade found inject room embark supply fabric improve spike stem give current similar glimpse"

func signTxWithMnemonic(txData ethtypes.TxData, mnemonic string) *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

func signCosmosTxWithMnemonic(msg sdk.Msg, mnemonic string, accountNumber uint64, sequenceNumber uint64) (sdk.Tx, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Tx), nil
}

func getAddrWithMnemonic(mnemonic string) common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func getSeiAddrWithMnemonic(mnemonic string) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func mnemonicInitializer(mnemonic string) func(ctx sdk.Context, a *app.App) {
	_ = "STUB: not implemented"
	return nil
}

func fundSeiAddr(ctx sdk.Context, a *app.App, addr sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func multiCoinInitializer(mnemonic string) func(ctx sdk.Context, a *app.App) {
	_ = "STUB: not implemented"
	return nil
}
