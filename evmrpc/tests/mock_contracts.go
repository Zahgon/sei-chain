package tests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/app"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func cw20Initializer(mnemonic string, pointer bool) func(ctx sdk.Context, a *app.App) {
	_ = "STUB: not implemented"
	return nil
}

func cwIterInitializer(mnemonic string) func(ctx sdk.Context, a *app.App) {
	_ = "STUB: not implemented"
	return nil
}

var erc20DeployerMnemonics = "number friend tray advice become blame morning glow final under unlock core employ side mimic local load flag birth hire doctor immense guess net"
var erc20Addr = common.HexToAddress("0x8bFEF0785c95Cb3D4a64202AB283c45ae6c50436") // deterministic with the mnemonic above as the deployer
var mixedLogTesterDeployerMnemonics = "area level during surge alley leader clock hard teach feel evidence tattoo snack betray scare six industry winner false improve various never silent protect"
var mixedLogTesterAddr = common.HexToAddress("0x9023C8C1dB86337278f64c79bDf0aD8402B9b17c") // deterministic with the mnemonic above as the deployer

func erc20Initializer() func(ctx sdk.Context, a *app.App) { _ = "STUB: not implemented"; return nil }

func mixedLogTesterInitializer() func(ctx sdk.Context, a *app.App) {
	_ = "STUB: not implemented"
	return nil
}

func GetBin(name string) []byte { _ = "STUB: not implemented"; return nil }

func GetABI(name string) []byte { _ = "STUB: not implemented"; return nil }
