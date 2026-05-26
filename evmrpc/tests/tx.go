package tests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func send(nonce uint64) ethtypes.TxData { _ = "STUB: not implemented"; return *new(ethtypes.TxData) }

// sendInsufficientFunds builds a correct-nonce tx whose intrinsic fee
// (GasFeeCap * Gas) exceeds mnemonicInitializer's funded balance, triggering
// an "insufficient funds" ante failure. Unlike send(N) for high N, this
// passes the nonce check, so SetNonceBumped fires and the chain writes a
// deferred-info stub receipt — the case the trace-side filter needs to
// catch.
func sendInsufficientFunds(nonce uint64) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

// mnemonicInitializer funds 10_000_000_000 usei = 1e22 wei. A fee of
// 2.1e34 wei is comfortably over budget.

func sendAmount(nonce uint64, amount *big.Int) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func sendErc20(nonce uint64) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func depositErc20(nonce uint64) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func mixedLogTesterTransfer(nonce uint64, recipient common.Address) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func registerCW20Pointer(nonce uint64, cw20Addr string) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func transferCW20Msg(mnemonic string, cw20Addr string) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

func transferCW20MsgTo(mnemonic string, cw20Addr string, recipient sdk.AccAddress) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

func jsonExtractAsBytesFromArray(nonce uint64) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}

func bankSendMsg(mnemonic string) sdk.Msg { _ = "STUB: not implemented"; return *new(sdk.Msg) }

func callWasmIter(nonce uint64, contractAddr string) ethtypes.TxData {
	_ = "STUB: not implemented"
	return *new(ethtypes.TxData)
}
